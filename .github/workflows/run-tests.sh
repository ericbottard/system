#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

source ${FATS_DIR}/.configure.sh

# setup namespace
kubectl create namespace ${NAMESPACE}
fats_create_push_credentials ${NAMESPACE}
source ${FATS_DIR}/macros/create-riff-dev-pod.sh

if [ $RUNTIME = "core" ] || [ $RUNTIME = "knative" ]; then
  for location in cluster local; do
    for test in function application; do
      name=system-${RUNTIME}-${location}-uppercase-node-${test}
      image=$(fats_image_repo ${name})

      echo "##[group]Run function $name"

      if [ $location = 'cluster' ] ; then
        riff $test create $name --image $image --namespace $NAMESPACE --tail \
          --git-repo https://github.com/${FATS_REPO}.git --git-revision ${FATS_REFSPEC} --sub-path ${test}s/uppercase/node &
      else
        riff $test create $name --image $image --namespace $NAMESPACE --tail \
          --local-path ${FATS_DIR}/${test}s/uppercase/node
      fi

      riff $RUNTIME deployer create $name --${test}-ref $name --namespace $NAMESPACE --ingress-policy External --tail
      if [ $test = 'function' ] ; then
        curl_opts="-H Content-Type:text/plain -H Accept:text/plain -d system"
        expected_data="SYSTEM"
      else
        curl_opts="--get --data-urlencode input=system"
        expected_data="SYSTEM"
      fi
      # invoke ClusterLocal
      source ${FATS_DIR}/macros/invoke_incluster.sh \
        "$(kubectl get deployers.${RUNTIME}.projectriff.io ${name} --namespace ${NAMESPACE} -ojsonpath='{.status.address.url}')" \
        "${curl_opts}" \
        "${expected_data}"
      # invoke External
      # TODO also test ingress for the core runtime
      if [ $RUNTIME = "knative" ]; then
        source ${FATS_DIR}/macros/invoke_${RUNTIME}_deployer.sh \
          "${name}" \
          "${curl_opts}" \
          "${expected_data}"
      fi
      riff $RUNTIME deployer delete $name --namespace $NAMESPACE

      riff $test delete $name --namespace $NAMESPACE
      fats_delete_image $image

      echo "##[endgroup]"
    done
  done

elif [ $RUNTIME = "streaming" ]; then
  riff streaming kafka-provider create franz --bootstrap-servers kafka.kafka.svc.cluster.local:9092 --namespace $NAMESPACE

  for test in node ; do
    name=system-${RUNTIME}-fn-uppercase-${test}
    image=$(fats_image_repo ${name})

    echo "##[group]Run function ${name}"

    riff function create ${name} --image ${image} --namespace ${NAMESPACE} --tail \
      --git-repo https://github.com/${FATS_REPO} --git-revision ${FATS_REFSPEC} --sub-path functions/uppercase/${test}

    lower_stream=${name}-lower
    upper_stream=${name}-upper

    provider=$(kubectl get kafkaproviders.streaming.projectriff.io franz --namespace ${NAMESPACE} -ojsonpath='{.status.provisionerServiceRef.name}')
    riff streaming stream create ${lower_stream} --namespace $NAMESPACE --provider ${provider} --content-type 'text/plain' --tail
    riff streaming stream create ${upper_stream} --namespace $NAMESPACE --provider ${provider} --content-type 'text/plain' --tail

    riff streaming processor create $name --function-ref $name --namespace $NAMESPACE --input ${lower_stream} --output ${upper_stream} --tail

    kubectl exec riff-dev -n $NAMESPACE -- subscribe ${upper_stream} -n $NAMESPACE --payload-as-string | tee result.txt &
    sleep 10
    kubectl exec riff-dev -n $NAMESPACE -- publish ${lower_stream} -n $NAMESPACE --payload "system" --content-type "text/plain"

    actual_data=""
    expected_data="SYSTEM"
    cnt=1
    while [ $cnt -lt 60 ]; do
      echo -n "."
      cnt=$((cnt+1))

      actual_data=`cat result.txt | jq -r .payload`
      if [ "$actual_data" == "$expected_data" ]; then
        break
      fi

      sleep 1
    done
    fats_assert "$expected_data" "$actual_data"

    kubectl exec riff-dev -n $NAMESPACE -- sh -c 'kill $(pidof subscribe)'

    riff streaming stream delete ${lower_stream} --namespace $NAMESPACE
    riff streaming stream delete ${upper_stream} --namespace $NAMESPACE
    riff streaming processor delete $name --namespace $NAMESPACE

    riff function delete ${name} --namespace ${NAMESPACE}
    fats_delete_image ${image}

    echo "##[endgroup]"
  done

  riff streaming kafka-provider delete franz --namespace $NAMESPACE

fi

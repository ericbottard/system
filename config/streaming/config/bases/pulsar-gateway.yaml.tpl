apiVersion: v1
kind: ConfigMap
metadata:
  name: pulsar-gateway
data:
  gatewayImage: ctr.run/github.com/bsideup/liiklus:12037c2c6c326e3c1f15af370c4a6d212d58084b
  provisionerImage: {{ gcloud container images describe gcr.io/projectriff/pulsar-provisioner/provisioner:0.6.0-snapshot --format="value(image_summary.fully_qualified_digest)" }}

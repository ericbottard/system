/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"github.com/knative/pkg/configmap"
	"github.com/knative/pkg/logging/logkey"
	"github.com/knative/pkg/signals"
	"github.com/knative/pkg/system"
	"github.com/knative/pkg/version"
	projectriffclientset "github.com/projectriff/system/pkg/client/clientset/versioned"
	projectriffinformers "github.com/projectriff/system/pkg/client/informers/externalversions"
	"github.com/projectriff/system/pkg/gateway"
	"github.com/projectriff/system/pkg/logging"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

const (
	component = "gateway"
)

var (
	masterURL  = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()
	cm, err := configmap.Load("/etc/config-logging")
	if err != nil {
		log.Fatalf("Error loading logging configuration: %v", err)
	}
	config, err := logging.NewConfigFromMap(cm)
	if err != nil {
		log.Fatalf("Error parsing logging configuration: %v", err)
	}
	logger, atomicLevel := logging.NewLoggerFromConfig(config, component)
	defer logger.Sync()
	logger = logger.With(zap.String(logkey.ControllerType, component))

	logger.Info("Starting the Http -> Streaming Gateway")

	// Set up signals so we handle the first shutdown signal gracefully.
	stopCh := signals.SetupSignalHandler()

	clusterConfig, err := clientcmd.BuildConfigFromFlags(*masterURL, *kubeconfig)
	if err != nil {
		logger.Fatalw("Failed to get cluster config", zap.Error(err))
	}

	kubeClient, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		logger.Fatalw("Failed to get the client set", zap.Error(err))
	}

	if err := version.CheckMinimumVersion(kubeClient.Discovery()); err != nil {
		logger.Fatalf("Version check failed: %v", err)
	}

	// Watch the logging config map and dynamically update logging levels.
	configMapWatcher := configmap.NewInformedWatcher(kubeClient, system.Namespace())
	configMapWatcher.Watch(logging.ConfigName, logging.UpdateLevelFromConfigMap(logger, atomicLevel, component))
	if err = configMapWatcher.Start(stopCh); err != nil {
		logger.Fatalw("Failed to start the ConfigMap watcher", zap.Error(err))
	}

	projectriffClient, err := projectriffclientset.NewForConfig(clusterConfig)
	if err != nil {
		logger.Fatalw("Error building projectriff clientset", zap.Error(err))
	}

	projectriffInformerFactory := projectriffinformers.NewSharedInformerFactory(projectriffClient, 10*time.Hour)
	streamInformer := projectriffInformerFactory.Streaming().V1alpha1().Streams()
	_ = streamInformer.Informer()

	projectriffInformerFactory.Start(stopCh)

	logger.Info("Waiting for informer caches to sync")
	informersSynced := []cache.InformerSynced{
		streamInformer.Informer().HasSynced,
	}
	for i, synced := range informersSynced {
		if ok := cache.WaitForCacheSync(stopCh, synced); !ok {
			logger.Fatalf("Failed to wait for cache at index %d to sync", i)
		}
	}

	log.Printf("Synched")

	gateway := gateway.NewGateway(streamInformer)

	if err = gateway.Run(stopCh); err != nil {
		logger.Fatalw("Failed to start the http -> streaming gateway", zap.Error(err))
	}
}

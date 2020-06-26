/*
Copyright 2020 the Velero contributors.

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
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/backupdriver"
	"github.com/vmware-tanzu/velero/pkg/util/logging"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

var (
	master             = flag.String("master", "", "Master URL to build a client config from. Either this or kubeconfig needs to be set if the backupdriver is being run out of cluster.")
	kubeConfig         = flag.String("kubeconfig", "", "Absolute path to the kubeconfig")
	resyncPeriod       = flag.Duration("resync-period", time.Minute*10, "Resync period for cache")
	workers            = flag.Int("workers", 10, "Concurrency to process multiple backup requests")
	retryIntervalStart = flag.Duration("retry-interval-start", time.Second, "Initial retry interval of failed backup request. It exponentially increases with each failure, up to retry-interval-max.")
	retryIntervalMax   = flag.Duration("retry-interval-max", 5*time.Minute, "Maximum retry interval of failed backup request.")

	showVersion = flag.Bool("version", false, "Show version")

	version = "unknown"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	if *showVersion {
		fmt.Println(os.Args[0], version)
		os.Exit(0)
	}

	logLevelFlag := logging.LogLevelFlag(logrus.InfoLevel)
	formatFlag := logging.NewFormatFlag()
	// go-plugin uses log.Println to log when it's waiting for all plugin processes to complete so we need to
	// set its output to stdout.
	log.SetOutput(os.Stdout)

	logLevel := logLevelFlag.Parse()
	format := formatFlag.Parse()

	// Make sure we log to stdout so cloud log dashboards don't show this as an error.
	logrus.SetOutput(os.Stdout)

	// Velero's DefaultLogger logs to stdout, so all is good there.
	logger := logging.DefaultLogger(logLevel, format)

	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = time.RFC3339
	formatter.FullTimestamp = true
	logger.SetFormatter(formatter)

	logger.Debugf("setting log-level to %s", strings.ToUpper(logLevel.String()))

	// kubeClient is the client to the current cluster
	// (vanilla, guest, or supervisor)
	kubeClient, err := newK8sClient(*master, *kubeConfig)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	informerFactory := informers.NewSharedInformerFactory(kubeClient, *resyncPeriod)

	// TODO: If CLUSTER_FLAVOR is GUEST_CLUSTER, set up svcKubeClient to communicate with the Supervisor Cluster
	// If CLUSTER_FLAVOR is WORKLOAD, it is a Supervisor Cluster.
	// By default we are in the Vanilla Cluster
	rc := backupdriver.NewBackupDriverController("BackupDriverController", logger, kubeClient, *resyncPeriod, informerFactory, workqueue.NewItemExponentialFailureRateLimiter(*retryIntervalStart, *retryIntervalMax) /*svcKubeClient,*/)
	run := func(ctx context.Context) {
		informerFactory.Start(wait.NeverStop)
		rc.Run(ctx, *workers)

	}

	run(context.TODO())
}

// newK8sClient is an utility function used to create a kubernetes sdk client.
func newK8sClient(master, kubeConfig string) (kubernetes.Interface, error) {
	var config *rest.Config
	var err error
	if master != "" || kubeConfig != "" {
		config, err = clientcmd.BuildConfigFromFlags(master, kubeConfig)
	} else {
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create config: %v", err)
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return kubeClient, nil
}
package main

import (
	"context"
	"flag"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

var (
	kubeconfig *string
	interval   *int
)

func init() {
	// Default kubeconfig path
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// Default interval in seconds
	interval = flag.Int("interval", 30, "interval in seconds to check pod status")
	flag.Parse()
}

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Error creating clientset: %s", err.Error())
	}

	ctx := context.TODO()

	// Periodically check pod liveness probe status
	for {
		checkPodLivenessProbe(ctx, clientset)
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}

func checkPodLivenessProbe(ctx context.Context, clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Error listing pods: %v", err)
		return
	}

	for _, pod := range pods.Items {
		klog.Infof("Checking liveness probe status for pod %s/%s", pod.Namespace, pod.Name)

		// Check the liveness probe status of the pod
		if isPodReady(pod) {
			klog.Infof("Pod %s/%s is ready", pod.Namespace, pod.Name)
		} else {
			klog.Infof("Pod %s/%s is not ready", pod.Namespace, pod.Name)
		}
	}
}

func isPodReady(pod corev1.Pod) bool {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady {
			return condition.Status == corev1.ConditionTrue
		}
	}
	return false
}

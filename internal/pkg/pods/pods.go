package pods

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// This is important for kubeconfigs that use auth plugins.
	// Examples: OIDC, GCP, Azure, exec-based auth, etc.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type GetPodsOptions struct {
	Namespace           string
	Timeout             time.Duration
	GlobalTimeout       time.Duration
	Verbose             bool
	Filter              string
	SortBy              string
	Phase               string
	Node                string
	RestartsGreaterThan uint
	// These fields should be handled at a higher level between the CLI and the
	// internal package, as they are more about output formatting and behavior
	// rather than pod-specific options
	//
	// Watch  bool
	// Wide   bool
}

func GetPods(ctx context.Context, options GetPodsOptions) (*corev1.PodList, error) {
	// TODO: This is just a PoC! Will be refactored and functionality pulled out
	// 1. Resolve kubeconfig path.
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("couldn't get user home dir: %w", err)
	}

	kubeconfig := filepath.Join(home, ".kube", "config")

	// 2. Load kubeconfig, optionally overriding context.
	loadingRules := &clientcmd.ClientConfigLoadingRules{
		ExplicitPath: kubeconfig,
	}

	overrides := &clientcmd.ConfigOverrides{}

	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		overrides,
	)

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("build kube client config: %w", err)
	}

	// 3. Create Kubernetes clientset.
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("create kubernetes clientset: %w", err)
	}

	// 4. Decide namespace.
	namespace := options.Namespace
	if namespace == "" {
		namespace = corev1.NamespaceAll // same as ""
	}

	fmt.Printf("Namespace: %s\n", namespace)

	// 5. Add timeout so the CLI does not hang forever.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 6. List pods through Kubernetes API.
	podList, err := clientset.CoreV1().
		Pods(namespace).
		List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, fmt.Errorf("list pods: %w", err)
	}

	return podList, nil
}

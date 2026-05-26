package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type K8sDemo struct {
	clientset  *kubernetes.Clientset
	namespace  string
	serverReady chan bool
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Try to use in-cluster config first, fall back to kubeconfig
	config, err := rest.InClusterConfig()
	if err != nil {
		// Not running in cluster, use kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			log.Fatalf("Error building kubeconfig: %v", err)
		}
		log.Println("Using kubeconfig from:", *kubeconfig)
	} else {
		log.Println("Using in-cluster configuration")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		namespace = "default"
	}

	demo := &K8sDemo{
		clientset:   clientset,
		namespace:   namespace,
		serverReady: make(chan bool, 1),
	}

	// Start HTTP server for health checks in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		if err := demo.startHTTPServer(); err != nil {
			serverErrors <- err
		}
	}()

	// Wait for server to be ready or timeout
	select {
	case <-demo.serverReady:
		log.Println("HTTP server is ready to accept connections")
	case err := <-serverErrors:
		log.Fatalf("HTTP server failed to start: %v", err)
	case <-time.After(10 * time.Second):
		log.Fatalf("HTTP server failed to start within timeout period")
	}

	// Run demonstrations after server is ready
	demo.runDemonstrations()

	// Keep the application running
	log.Println("Application is running. Press Ctrl+C to exit.")
	select {}
}

func (d *K8sDemo) startHTTPServer() error {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("READY"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Kubernetes Demo Application\n"))
	})

	// Create a listener to check if port binding succeeds
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("failed to bind to port 8080: %w", err)
	}

	log.Println("HTTP server successfully bound to :8080")
	
	// Signal that server is ready
	d.serverReady <- true
	
	// Start serving
	log.Println("HTTP server is now accepting connections")
	server := &http.Server{}
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server error: %w", err)
	}
	return nil
}

func (d *K8sDemo) runDemonstrations() {
	ctx := context.Background()

	log.Println("=== Kubernetes Go Client Demonstrations ===")

	// 1. List all namespaces
	d.listNamespaces(ctx)

	// 2. List pods in current namespace
	d.listPods(ctx)

	// 3. List nodes
	d.listNodes(ctx)

	// 4. Create a ConfigMap
	d.createConfigMap(ctx)

	// 5. List services
	d.listServices(ctx)

	// 6. Get deployment information
	d.getDeploymentInfo(ctx)

	// 7. Watch pods continuously
	go d.watchPods(ctx)
}

func (d *K8sDemo) listNamespaces(ctx context.Context) {
	log.Println("\n--- Listing Namespaces ---")
	namespaces, err := d.clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing namespaces: %v", err)
		return
	}

	fmt.Printf("Found %d namespaces:\n", len(namespaces.Items))
	for _, ns := range namespaces.Items {
		fmt.Printf("  - %s (Status: %s)\n", ns.Name, ns.Status.Phase)
	}
}

func (d *K8sDemo) listPods(ctx context.Context) {
	log.Printf("\n--- Listing Pods in namespace: %s ---\n", d.namespace)
	pods, err := d.clientset.CoreV1().Pods(d.namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing pods: %v", err)
		return
	}

	fmt.Printf("Found %d pods:\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("  - %s (Status: %s, IP: %s)\n", pod.Name, pod.Status.Phase, pod.Status.PodIP)
		for _, container := range pod.Spec.Containers {
			fmt.Printf("    Container: %s (Image: %s)\n", container.Name, container.Image)
		}
	}
}

func (d *K8sDemo) listNodes(ctx context.Context) {
	log.Println("\n--- Listing Nodes ---")
	nodes, err := d.clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing nodes: %v", err)
		return
	}

	fmt.Printf("Found %d nodes:\n", len(nodes.Items))
	for _, node := range nodes.Items {
		fmt.Printf("  - %s\n", node.Name)
		fmt.Printf("    OS: %s\n", node.Status.NodeInfo.OSImage)
		fmt.Printf("    Kubernetes Version: %s\n", node.Status.NodeInfo.KubeletVersion)
		for _, addr := range node.Status.Addresses {
			if addr.Type == corev1.NodeInternalIP {
				fmt.Printf("    Internal IP: %s\n", addr.Address)
			}
		}
	}
}

func (d *K8sDemo) createConfigMap(ctx context.Context) {
	log.Println("\n--- Creating/Updating ConfigMap ---")
	
	configMapName := "k8s-demo-config"
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: d.namespace,
		},
		Data: map[string]string{
			"app.name":        "k8s-demo",
			"app.version":     "1.0.0",
			"demo.timestamp":  time.Now().Format(time.RFC3339),
			"demo.message":    "This ConfigMap was created by the Go Kubernetes client!",
		},
	}

	// Try to get existing ConfigMap
	existingCM, err := d.clientset.CoreV1().ConfigMaps(d.namespace).Get(ctx, configMapName, metav1.GetOptions{})
	if err != nil {
		// Create new ConfigMap
		result, err := d.clientset.CoreV1().ConfigMaps(d.namespace).Create(ctx, configMap, metav1.CreateOptions{})
		if err != nil {
			log.Printf("Error creating ConfigMap: %v", err)
			return
		}
		fmt.Printf("ConfigMap '%s' created successfully\n", result.Name)
	} else {
		// Update existing ConfigMap
		existingCM.Data = configMap.Data
		result, err := d.clientset.CoreV1().ConfigMaps(d.namespace).Update(ctx, existingCM, metav1.UpdateOptions{})
		if err != nil {
			log.Printf("Error updating ConfigMap: %v", err)
			return
		}
		fmt.Printf("ConfigMap '%s' updated successfully\n", result.Name)
	}

	// Read back the ConfigMap
	cm, err := d.clientset.CoreV1().ConfigMaps(d.namespace).Get(ctx, configMapName, metav1.GetOptions{})
	if err != nil {
		log.Printf("Error reading ConfigMap: %v", err)
		return
	}

	fmt.Println("ConfigMap data:")
	for key, value := range cm.Data {
		fmt.Printf("  %s: %s\n", key, value)
	}
}

func (d *K8sDemo) listServices(ctx context.Context) {
	log.Printf("\n--- Listing Services in namespace: %s ---\n", d.namespace)
	services, err := d.clientset.CoreV1().Services(d.namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing services: %v", err)
		return
	}

	fmt.Printf("Found %d services:\n", len(services.Items))
	for _, svc := range services.Items {
		fmt.Printf("  - %s (Type: %s, ClusterIP: %s)\n", svc.Name, svc.Spec.Type, svc.Spec.ClusterIP)
		if len(svc.Spec.Ports) > 0 {
			fmt.Printf("    Ports: ")
			for _, port := range svc.Spec.Ports {
				fmt.Printf("%d/%s ", port.Port, port.Protocol)
			}
			fmt.Println()
		}
	}
}

func (d *K8sDemo) getDeploymentInfo(ctx context.Context) {
	log.Printf("\n--- Getting Deployment Information in namespace: %s ---\n", d.namespace)
	deployments, err := d.clientset.AppsV1().Deployments(d.namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing deployments: %v", err)
		return
	}

	fmt.Printf("Found %d deployments:\n", len(deployments.Items))
	for _, deploy := range deployments.Items {
		fmt.Printf("  - %s\n", deploy.Name)
		replicas := int32(0)
		if deploy.Spec.Replicas != nil {
			replicas = *deploy.Spec.Replicas
		}
		fmt.Printf("    Replicas: %d desired, %d ready, %d available\n",
			replicas,
			deploy.Status.ReadyReplicas,
			deploy.Status.AvailableReplicas)
		if len(deploy.Spec.Template.Spec.Containers) > 0 {
			fmt.Printf("    Image: %s\n", deploy.Spec.Template.Spec.Containers[0].Image)
		}
	}
}

func (d *K8sDemo) watchPods(ctx context.Context) {
	log.Printf("\n--- Watching Pods in namespace: %s ---\n", d.namespace)
	
	for {
		watcher, err := d.clientset.CoreV1().Pods(d.namespace).Watch(ctx, metav1.ListOptions{})
		if err != nil {
			log.Printf("Error creating watcher: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for event := range watcher.ResultChan() {
			pod, ok := event.Object.(*corev1.Pod)
			if !ok {
				continue
			}
			
			log.Printf("Pod Event [%s]: %s (Phase: %s)\n", event.Type, pod.Name, pod.Status.Phase)
		}

		// If the watch channel closes, restart it
		time.Sleep(5 * time.Second)
	}
}


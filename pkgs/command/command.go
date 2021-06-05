package command

import (
	"context"
	"fmt"
	"github.com/sheikh1309/k8s-exec/pkgs/k8s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"strings"
)

func RunCommandInDeployment(contextName string, ns string, cmd string, deployment string) {
	fmt.Printf("Start Exec Command In Deployment: %s, Context: %s \n", deployment, contextName)
	client, config := k8s.GetClient(contextName)
	pods, _ := client.CoreV1().Pods(ns).List(context.TODO(), v1.ListOptions{})
	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, deployment) {
			err := k8s.ExecCmdInPod(client, config, pod.Name, ns, cmd, os.Stdin, os.Stdout, os.Stderr)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Printf("Finish Exec Command In Deployment: %s, Context: %s \n", deployment, contextName)
}

func RunCommandInPod(contextName string, ns string, cmd string, pod string) {
	fmt.Printf("Start Exec Command In Pod: %s, Context: %s \n", pod, contextName)
	client, config := k8s.GetClient(contextName)
	err := k8s.ExecCmdInPod(client, config, pod, ns, cmd, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Finish Exec Command In Pod: %s, Context: %s \n", pod, contextName)
}

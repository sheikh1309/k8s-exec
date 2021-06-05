package main

import "github.com/sheikh1309/k8s-exec-pod/pkgs/command"

func main()  {
	command.SetExecutable("k8s-exec", "1.0.1", "This tool if for exec commands in deployment pods or in specific pod")
}

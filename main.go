package main

import "github.com/sheikh1309/k8s-exec/pkgs/command"

func main()  {
	command.SetExecutable("k8s-exec", "1.0.3", "This tool if for exec commands in deployment pods or in specific pod")
}

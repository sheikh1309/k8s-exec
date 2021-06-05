package command

import (
	"github.com/thatisuday/commando"
)

func SetExecutable(name string, version string, desc string)  {
	commando.
		SetExecutableName(name).
		SetVersion(version).
		SetDescription(desc)

	commando.
		Register("deployment").
		AddFlag("context", "k8s context", commando.String, nil).
		AddFlag("ns", "k8s namespace", commando.String, nil).
		AddFlag("cmd", "command to exec in pod", commando.String, nil).
		AddFlag("deployment", "deployment name", commando.String, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			contextName, _ := flags["context"].GetString()
			ns, _ := flags["ns"].GetString()
			cmd, _ := flags["cmd"].GetString()
			deployment, _ := flags["deployment"].GetString()
			RunCommandInDeployment(contextName, ns, cmd, deployment)
		})

	commando.
		Register("pod").
		AddFlag("context", "k8s context", commando.String, nil).
		AddFlag("ns", "k8s namespace", commando.String, nil).
		AddFlag("cmd", "command to exec in pod", commando.String, nil).
		AddFlag("pod", "pod name", commando.String, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			contextName, _ := flags["context"].GetString()
			ns, _ := flags["ns"].GetString()
			cmd, _ := flags["cmd"].GetString()
			pod, _ := flags["pod"].GetString()
			RunCommandInPod(contextName, ns, cmd, pod)
		})

	commando.Parse(nil)
}

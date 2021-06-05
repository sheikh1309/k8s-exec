# tree
![go-version](https://img.shields.io/badge/Go%20Version-v1.15-blue) &nbsp;

A command-line tool to exec commands in k8s pods

![Alt text](./assets/ScreenShot1.png?raw=true)
> This CLI application is made using [Commando](https://github.com/thatisuday/commando).

## Installation
```
$ GO111MODULE=on go get github.com/sheikh1309/k8s-exec
```

## Usage
```
$  k8s-exec --help

This tool is for exec commands in deployment pods or in specific pod

Usage:
   k8s-exec {flags}
   k8s-exec <command> {flags}

Commands:
   deployment
   help                          displays usage informationn
   pod
   version                       displays version number

Flags:
   -h, --help                    displays usage information of the application or a command (default: false)
   -v, --version                 displays version number (default: false)
```

## Example
```
$ k8s-exec deployment --context docker-desktop --ns default --cmd 'echo 123' --deployment nginx-deployment
Start Exec Command In Deployment: nginx-deployment, Context: docker-desktop
123

123

123

Finish Exec Command In Deployment: nginx-deployment, Context: docker-desktop
```
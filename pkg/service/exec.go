package service

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	"k8s.agent/pkg/config"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func ExecInPod(k8sCli *kubernetes.Clientset, namespace, podName, command, containerName string) {

	cmd := []string{
		"sh",
		"-c",
		command,
	}
	const tty = false
	req := k8sCli.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).SubResource("exec").Param("container", containerName)
	req.VersionedParams(
		&v1.PodExecOptions{
			Command: cmd,
			Stdin:   false,
			Stdout:  true,
			Stderr:  true,
			TTY:     tty,
		},
		scheme.ParameterCodec,
	)

	exec, err := remotecommand.NewSPDYExecutor(config.GetK8sConfig(), "POST", req.URL())
	if err != nil {
		panic(err)
	}

	// 检查是不是终端
	if !terminal.IsTerminal(0) || !terminal.IsTerminal(1) {
		fmt.Errorf("stdin/stdout should be terminal")
	}
	// 这个应该是处理Ctrl + C 这种特殊键位
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		fmt.Println(err)
	}
	defer terminal.Restore(0, oldState)

	// 用IO读写替换 os stdout
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  screen,
		Stdout: screen,
		Stderr: screen,
		Tty:    false,
	})
	if err != nil {
		panic(err)
	}

}

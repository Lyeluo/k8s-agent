package model

type PodExecl struct {
	Command string `JSON:"command"`

	ContainerName string `JSON:"containerName"`
}

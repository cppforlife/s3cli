package cmd

import (
	"fmt"

	s3cliclient "s3cli/client"
)

type Factory interface {
	Create(cmdName string) (cmd Cmd, err error)
}

type concreteFactory struct {
	cmds map[string]Cmd
}

func NewFactory(s3Client s3cliclient.Client) (factory Factory) {
	return concreteFactory{
		cmds: map[string]Cmd{
			"get": newGet(s3Client),
			"put": newPut(s3Client),
		},
	}
}

func (factory concreteFactory) Create(cmdName string) (cmd Cmd, err error) {
	cmd, found := factory.cmds[cmdName]
	if !found {
		err = fmt.Errorf("Command not found: %s", cmdName)
	}
	return
}

package docker

import (
	"github.com/docker/libcompose/project"
	"github.com/samalba/dockerclient"
)

type Container struct {
	project.EmptyService

	name          string
	serviceConfig *project.ServiceConfig
	context       *Context
}

func (c *Container) findExisting() (*dockerclient.Container, error) {
	containers, err := c.context.Client.ListContainers(true, false, "label=io.docker.compose.name="+c.name)
	if len(containers) > 0 {
		return &containers[0], err
	} else {
		return nil, err
	}
}

func (c *Container) Create() (*dockerclient.Container, error) {
	container, err := c.findExisting()
	if err != nil {
		return nil, err
	}

	if container == nil {
		container, err = c.createContainer()
		if err != nil {
			return nil, err
		}
	}

	return container, err
}

func (c *Container) Up() error {
	var err error

	defer func() {
		if err == nil && c.context.Log {
			go c.Log()
		}
	}()

	container, err := c.Create()
	if err != nil {
		return err
	}

	info, err := c.context.Client.InspectContainer(container.Id)
	if err != nil {
		return err
	}

	if !info.State.Running {
		return c.context.Client.StartContainer(container.Id, info.HostConfig)
	}

	return nil
}

func (c *Container) createContainer() (*dockerclient.Container, error) {
	config, err := ConvertToApi(c.serviceConfig)
	if err != nil {
		return nil, err
	}

	if config.Labels == nil {
		config.Labels = map[string]string{}
	}

	config.Labels["io.docker.compose.name"] = c.name
	config.Labels["io.docker.compose.project"] = c.context.Project.Name

	_, err = c.context.Client.CreateContainer(config, c.name)
	if err != nil {
		return nil, err
	}

	return c.findExisting()
}

package docker

import "github.com/docker/libcompose/project"

type Service struct {
	project.EmptyService

	tty           bool
	name          string
	serviceConfig *project.ServiceConfig
	context       *Context
}

func (s *Service) Create() error {
	_, err := s.doCreate()
	return err
}

func (s *Service) doCreate() (Container, error) {
	c := Container{
		name:          s.name,
		serviceConfig: s.serviceConfig,
		context:       s.context,
	}

	_, err := c.Create()
	return c, err
}

func (s *Service) Up() error {
	c, err := s.doCreate()
	if err != nil {
		return err
	}

	return c.Up()
}

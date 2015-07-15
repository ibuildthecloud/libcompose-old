package integration

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

func (s *RunSuite) TestHelloWorld(c *C) {
	p := s.CreateProjectFromText(c, `
	hello:
	  image: tianon/true
	`)

	name := fmt.Sprintf("%s_%s_1", p, "hello")
	cn := s.GetContainerByName(c, name)
	c.Assert(cn, NotNil)

	c.Assert(cn.Name, Equals, "/"+name)
}

func (s *RunSuite) TestUp(c *C) {
	p := s.ProjectFromText(c, "up", `
	hello:
	  image: busybox
	  stdin_open: true
	  tty: true
	`)

	name := fmt.Sprintf("%s_%s_1", p, "hello")
	cn := s.GetContainerByName(c, name)
	c.Assert(cn, NotNil)

	c.Assert(cn.State.Running, Equals, true)
}

func (s *RunSuite) TestStop(c *C) {
	p := s.ProjectFromText(c, "up", `
	hello:
	  image: busybox
	  stdin_open: true
	  tty: true
	`)

	name := fmt.Sprintf("%s_%s_1", p, "hello")

	cn := s.GetContainerByName(c, name)
	c.Assert(cn, NotNil)
	c.Assert(cn.State.Running, Equals, true)

	s.FromText(c, p, "stop", `
	hello:
	  image: busybox
	  stdin_open: true
	  tty: true
	`)

	cn = s.GetContainerByName(c, name)
	c.Assert(cn, NotNil)
	c.Assert(cn.State.Running, Equals, false)
}

func (s *RunSuite) TestDelete(c *C) {
	p := s.ProjectFromText(c, "up", `
	hello:
	  image: busybox
	  stdin_open: true
	  tty: true
	`)

	name := fmt.Sprintf("%s_%s_1", p, "hello")

	cn := s.GetContainerByName(c, name)
	c.Assert(cn, NotNil)
	c.Assert(cn.State.Running, Equals, true)

	s.FromText(c, p, "rm", "--force", `
	hello:
	  image: busybox
	  stdin_open: true
	  tty: true
	`)

	cn = s.GetContainerByName(c, name)
	c.Assert(cn, IsNil)
}

func (s *RunSuite) TestDown(c *C) {
	p := s.ProjectFromText(c, "up", `
	hello:
	  image: busybox
	  stdin_open: true
	  tty: true
	`)

	name := fmt.Sprintf("%s_%s_1", p, "hello")

	cn := s.GetContainerByName(c, name)
	c.Assert(cn, NotNil)
	c.Assert(cn.State.Running, Equals, true)

	s.FromText(c, p, "down", `
	hello:
	  image: busybox
	  stdin_open: true
	  tty: true
	`)

	cn = s.GetContainerByName(c, name)
	c.Assert(cn, NotNil)
	c.Assert(cn.State.Running, Equals, false)
}

func (s *RunSuite) TestLink(c *C) {
	p := s.ProjectFromText(c, "up", `
	server:
	  image: busybox
	  command: cat
	  stdin_open: true
	  expose:
	  - 80
	client:
	  image: busybox
	  links:
	  - server:foo
	  - server
	`)

	serverName := fmt.Sprintf("%s_%s_1", p, "server")

	cn := s.GetContainerByName(c, serverName)
	c.Assert(cn, NotNil)
	c.Assert(cn.Config.ExposedPorts, DeepEquals, map[string]struct{}{
		"80/tcp": {},
	})

	clientName := fmt.Sprintf("%s_%s_1", p, "client")
	cn = s.GetContainerByName(c, clientName)
	c.Assert(cn, NotNil)
	c.Assert(cn.HostConfig.Links, DeepEquals, []string{
		fmt.Sprintf("/%s:/%s/%s", serverName, clientName, "foo"),
		fmt.Sprintf("/%s:/%s/%s", serverName, clientName, "server"),
		fmt.Sprintf("/%s:/%s/%s", serverName, clientName, serverName),
	})
}

func (s *RunSuite) TestPull(c *C) {
	//TODO: This doesn't test much
	s.ProjectFromText(c, "pull", `
	hello:
	  image: busybox
	  stdin_open: true
	  tty: true
	`)
}

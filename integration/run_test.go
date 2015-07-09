package integration

import (
	"bytes"
	"math/rand"
	"os/exec"
	"testing"

	"github.com/Sirupsen/logrus"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

func RandStr(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type RunSuite struct {
	command  string
	projects []string
}

var _ = Suite(&RunSuite{
	command: "/home/darren/src/libcompose/test",
})

func (s *RunSuite) CreateProjectFromText(c *C, input string) string {
	projectName := "test-project-" + RandStr(7)
	args := []string{"--verbose", "-p", projectName, "-f", "-", "create"}
	cmd := exec.Command(s.command, args...)
	cmd.Stdin = bytes.NewBufferString(input)

	err := cmd.Run()
	if err != nil {
		logrus.Errorf("Failed to run %s %v: %v\n with input:\n%s", s.command, err, args, input)
	}

	c.Assert(err, NotNil)

	return projectName
}

func (s *RunSuite) TestHelloWorld(c *C) {
	s.CreateProjectFromText(c, `
	hello:
	  image: busybox
	  command: echo hi
	`)
}

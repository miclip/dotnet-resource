package dotnetresource

import (
	"os/exec"
)
// DotnetClient ...
type DotnetClient interface {
	Build() ([]byte, error)
}

type dotnetclient struct {
	path      string
	framework string
	runtime   string
}

// NewDotnetClient ...
func NewDotnetClient(
	path string,
	framework string,
	runtime string,
) DotnetClient {
	projectPath := path
	targetFramework := framework
	targetRuntime := runtime

	return &dotnetclient{
		path:      projectPath,
		framework: targetFramework,
		runtime:   targetRuntime,
	}
}

var ExecCommand = exec.Command

func (client *dotnetclient) Build() ([]byte, error) {
	cmd := ExecCommand("dotnet", "build", client.path, "-f", client.framework, "-r", client.runtime)
	out, err := cmd.CombinedOutput()
	return out, err
}
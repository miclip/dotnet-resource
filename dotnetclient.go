package dotnetresource

import (
	"bufio"
	"os/exec"
	"path"
	"strings"
)

// DotnetClient ...
type DotnetClient interface {
	Build() ([]byte, error)
	Test(testfilter string) ([]byte, error)
}

type dotnetclient struct {
	path      string
	framework string
	runtime   string
	sourceDir string
}

// NewDotnetClient ...
func NewDotnetClient(
	path string,
	framework string,
	runtime string,
	sourceDir string,
) DotnetClient {
	projectPath := path
	targetFramework := framework
	targetRuntime := runtime
	root := sourceDir
	return &dotnetclient{
		path:      projectPath,
		framework: targetFramework,
		runtime:   targetRuntime,
		sourceDir: root,
	}
}

var ExecCommand = exec.Command

func (client *dotnetclient) Build() ([]byte, error) {
	cmd := ExecCommand("dotnet", "build", path.Join(client.sourceDir, client.path), "-f", client.framework, "-r", client.runtime)
	out, err := cmd.CombinedOutput()
	return out, err
}
func (client *dotnetclient) Test(testfilter string) ([]byte, error) {
	output := []byte{}
	cmd := ExecCommand("find", client.sourceDir, "-type", "f", "-name", testfilter)
	out, err := cmd.CombinedOutput()
	if(err!= nil){
		Fatal("error searching for test projects: \n"+string(out), err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		cmd = ExecCommand("dotnet", "test", "-f", client.framework, "--no-build", "--no-restore", scanner.Text(),
			"-p:RuntimeIdentifier="+client.runtime)
		out, _ := cmd.CombinedOutput()
		output = append(output, out...)
	}
	return output, nil
}

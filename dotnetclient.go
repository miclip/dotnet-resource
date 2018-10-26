package dotnetresource

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

// DotnetClient ...
type DotnetClient interface {
	Build() ([]byte, error)
	Test(testfilter string) ([]byte, error)
	Pack(projectPath string, version string) ([]byte, error)
	Push(sourceURL string, apiKey string, timeout int) ([]byte, error)
	Publish(projectPath string, packageID string) ([]byte, error)
	ManualPack(packageID string, version string) ([]byte, error)
	AddFileToPackage(packageID string, version string,filename string, reader io.Reader) error
	ManualUnpack(packageID string, version string) ([]byte, error)
}

type dotnetclient struct {
	path       string
	framework  string
	runtime    string
	sourceDir  string
	packageDir string
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
		path:       projectPath,
		framework:  targetFramework,
		runtime:    targetRuntime,
		sourceDir:  root,
		packageDir: root + "/packages",
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
	if err != nil {
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

func (client *dotnetclient) Pack(projectPath string, version string) ([]byte, error) {
	cmd := ExecCommand("dotnet", "pack", projectPath, "--no-build", "--no-restore", "--output", client.packageDir, "--runtime", client.runtime, "--include-symbols", "-p:PackageVersion="+version)
	out, err := cmd.CombinedOutput()
	return out, err
}

func (client *dotnetclient) Push(sourceURL string, apiKey string, timeout int) ([]byte, error) {
	cmd := ExecCommand("dotnet", "nuget", "push", client.packageDir+"/*.*", "--api-key", apiKey, "--source", sourceURL, "--timeout", strconv.Itoa(timeout))
	out, err := cmd.CombinedOutput()
	return out, err
}

func (client *dotnetclient) Publish(projectPath string, packageID string) ([]byte, error) {
	cmd := ExecCommand("dotnet", "publish", projectPath, "--no-build", "--no-restore", "-f", client.framework, "-r", client.runtime, "-o", client.packageDir+"/"+packageID)
	out, err := cmd.CombinedOutput()
	return out, err
}

func (client *dotnetclient) ManualPack(packageID string, version string) ([]byte, error) {
	out := []byte{}
	packageName := client.packageDir + "/" + packageID + "." + version
	cmd := ExecCommand("7z", "a", "-r", packageName+".zip", client.packageDir+"/"+packageID+"/*")
	zipOut, err := cmd.CombinedOutput()
	out = append(out, zipOut...)
	if err != nil {
		return out, err
	}
	cmd = ExecCommand("/bin/sh", "-c", "mv -v "+packageName+".zip"+" "+packageName+".nupkg")
	mvOut, err := cmd.CombinedOutput()
	out = append(out, mvOut...)
	if err != nil {
		return out, err
	}
	return out, err
}

func (client *dotnetclient) ManualUnpack(packageID string, version string) ([]byte, error) {
	out := []byte{}
	packageName := client.packageDir + "/" + packageID + "." + version
	cmd := ExecCommand("/bin/sh", "-c", "mv -v "+packageName+".nupkg"+" "+packageName+".zip")
	mvOut, err := cmd.CombinedOutput()
	out = append(out, mvOut...)
	if err != nil {
		return out, err
	}	
	cmd = ExecCommand("7z", "x", packageName+".zip", "-o"+client.sourceDir+"/"+packageID, "-r")
	zipOut, err := cmd.CombinedOutput()
	out = append(out, zipOut...)
	if err != nil {
		return out, err
	}	
	return out, err
}

func (client *dotnetclient) AddFileToPackage(packageID string, version string, filename string, reader io.Reader) error {

	fo, err := os.Create(client.packageDir + "/" + packageID + "/" + filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}

		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
	return nil
}

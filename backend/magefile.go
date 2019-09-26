// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

func Lint() error {
	return sh.RunV("golint", "./src/...")
}

func Build() error {
	ldflags := map[string]string{
		"main.version":   getVersionString(),
		"main.buildTime": time.Now().UTC().Format(time.RFC3339),
		"main.builder":   getBuilderString(),
		"main.goversion": runtime.Version()[2:],
	}

	env := map[string]string{
		"CGO_ENABLED": "0",
	}

	_, err := sh.Exec(env, os.Stdout, os.Stderr,
		"go", "build",
		"-o", "bin/koala-pos",
		"-v",
		"-ldflags",
		formatLDFlags(ldflags),
		"./cmd/koala-pos",
	)
	return err
}

func Generate() error {
	return sh.RunV("go", "generate", "./...")
}

func BuildInDocker() error {
	pwd, _ := os.Getwd()
	gopath, _ := os.LookupEnv("GOPATH")
	if gopath == "" {
		home, _ := os.UserHomeDir()
		gopath = filepath.Join(home, "go")
	}

	env := map[string]string{
		"BUILDTIME": time.Now().UTC().Format(time.RFC3339),
		"BUILDER":   getBuilderString(),
	}

	_, err := sh.Exec(env, os.Stdout, os.Stderr,
		"docker", "run", "--rm",
		"-e", "GOPATH=/go",
		"-e", "CGO_ENABLED=0",
		"-e", "BUILDTIME",
		"-e", "BUILDER",
		"-v", fmt.Sprintf("%s/..:/usr/src/koala-pos", pwd),
		"-v", fmt.Sprintf("%s/pkg/mod:/go/pkg/mod", gopath),
		"-w", "/usr/src/koala-pos/backend",
		"golang:1.13-alpine", "go", "run", "mage.go",
	)
	return err
}

func RunDev() error {
	mg.Deps(BuildInDocker)
	os.Chdir("docker")
	return sh.RunV("docker-compose", "up", "-d")
}

func RunDevLogs() error {
	mg.Deps(BuildInDocker, RunDev)
	return sh.RunV("docker-compose", "logs", "-f")
}

func StopDev() error {
	os.Chdir("docker")
	return sh.RunV("docker-compose", "down")
}

func StopDevClean() error {
	os.Chdir("docker")
	return sh.RunV("docker-compose", "down", "-v")
}

func Clean() {
	sh.Rm("./bin/*")
	sh.Rm("./logs/*")
	sh.Rm("./sessions/*")
	mg.Deps(StopDevClean)
}

func Test() error {
	mg.Deps(Lint)
	return sh.RunV("go", "test", "./src/...")
}

func Format() error {
	return sh.RunV("gofmt", "-s", "-l", "-d", "./src/*")
}

func formatLDFlags(flags map[string]string) string {
	s := strings.Builder{}

	for tag, value := range flags {
		s.WriteString("-X '")
		s.WriteString(tag)
		s.WriteByte('=')
		s.WriteString(value)
		s.WriteString("' ")
	}

	return s.String()
}

func getVersionString() string {
	out, _ := sh.Output("git", "describe", "--tags", "--always", "--dirty")
	return out
}

func getBuilderString() string {
	gitConfigCmd := sh.OutCmd("git", "config")
	name, _ := gitConfigCmd("user.name")
	email, _ := gitConfigCmd("user.email")
	return fmt.Sprintf("%s <%s>", name, email)
}

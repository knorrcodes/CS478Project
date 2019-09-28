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

var Default = BuildInDocker

// Will build the koala-pos binary, this target should only be used inside
// the Docker container building the project.
func Build() error {
	ldflags := map[string]string{
		"main.version":   lookupEnv("BUILD_VERSION", getVersionString()),
		"main.builder":   lookupEnv("BUILDER", getBuilderString()),
		"main.buildTime": time.Now().UTC().Format(time.RFC3339),
		"main.goversion": runtime.Version()[2:], // Strip off "go"
	}

	env := map[string]string{
		"CGO_ENABLED": lookupEnv("CGO_ENABLED", "0"),
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

// Generate GraphQL files
func Generate() error {
	return sh.RunV("go", "generate", "./...")
}

// Build the project inside a Docker container using caches from the host
func BuildInDocker() error {
	return buildInDocker(false)
}

// Build the project inside a Docker container with a clean environment
func BuildInDockerClean() error {
	return buildInDocker(true)
}

func buildInDocker(clean bool) error {
	pwd, _ := os.Getwd()
	gopath := lookupEnv("GOPATH", "")
	home, _ := os.UserHomeDir()
	if gopath == "" {
		// Default Go path used by go toolchain
		gopath = filepath.Join(home, "go")
	}

	env := map[string]string{
		"BUILDER":       getBuilderString(),
		"BUILD_VERSION": getVersionString(),
	}

	args := []string{
		"run", "--rm", // Remove container after execution
		"-e", "CGO_ENABLED=0", // Ensure cgo is disabled
		"-e", "BUILDER", // Pass in builder
		"-e", "BUILD_VERSION", // Pass in build version
		"-w", "/usr/src/koala-pos", // Set working directory
		"-v", fmt.Sprintf("%s:/usr/src/koala-pos", pwd), // Mount code
	}

	if !clean {
		args = append(args,
			// Mount package cache
			"-v", fmt.Sprintf("%s/pkg/mod:/go/pkg/mod", gopath),
			// Mount build cache
			"-v", fmt.Sprintf("%s/.cache/go-build:/root/.cache/go-build", home),
		)
	}

	args = append(args,
		// Image and command to run
		"golang:1.13-alpine", "go", "run", "mage.go", "build",
	)

	_, err := sh.Exec(env, os.Stdout, os.Stderr, "docker", args...)
	return err
}

// Start the backend server and database
func RunDev() error {
	mg.Deps(BuildInDocker)
	os.Chdir("docker")
	return sh.RunV("docker-compose", "up", "-d")
}

// Start the backend server and database and show logs
func RunDevLogs() error {
	mg.Deps(BuildInDocker, RunDev)
	return sh.RunV("docker-compose", "logs", "-f")
}

// Restart the backend server, useful after a rebuild
func RestartDev() error {
	os.Chdir("docker")
	return sh.RunV("docker-compose", "restart", "pos-backend")
}

// Rebuild and restart the backend server with logs
func RestartDevLogs() error {
	// Rebuild application
	mg.Deps(BuildInDocker)

	// Restart container
	os.Chdir("docker")
	if err := sh.RunV("docker-compose", "restart", "pos-backend"); err != nil {
		return err
	}

	// Start log stream
	return sh.RunV("docker-compose", "logs", "-f")
}

// Stop the backend server and database
func StopDev() error {
	os.Chdir("docker")
	return sh.RunV("docker-compose", "down")
}

// Stop the backend server and database and delete database data
func StopDevClean() error {
	os.Chdir("docker")
	return sh.RunV("docker-compose", "down", "-v")
}

// Remove build artifacts, logs, and Docker containers
func Clean() {
	sh.Rm("./bin/*")
	sh.Rm("./logs/*")
	sh.Rm("./sessions/*")
	mg.Deps(StopDevClean)
}

// Run tests
func Test() error {
	mg.Deps(Lint)
	return sh.RunV("go", "test", "./src/...")
}

// Ensure Go formatting
func Format() error {
	return sh.RunV("gofmt", "-s", "-l", "-d", "./src/*")
}

// Run golint
func Lint() error {
	return sh.RunV("golint", "./src/...")
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

// Generate a version string using git tags or commit hash.
func getVersionString() string {
	out, _ := sh.Output("git", "describe", "--tags", "--always", "--dirty")
	if out == "" {
		return "Unknow version"
	}
	return out
}

// Generate builder ID using name and email from git config.
func getBuilderString() string {
	gitConfigCmd := sh.OutCmd("git", "config")

	name, _ := gitConfigCmd("user.name")
	if name == "" {
		name = "Unknown Builder"
	}

	email, _ := gitConfigCmd("user.email")

	if email == "" {
		return fmt.Sprintf("%s", name)
	}
	return fmt.Sprintf("%s <%s>", name, email)
}

// lookupEnv checks the environment for key and returns its value if
// it exists otherwise def is returned.
func lookupEnv(key string, def string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return def
}

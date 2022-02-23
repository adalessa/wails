package system

import (
	"github.com/wailsapp/wails/v2/internal/system/operatingsystem"
	"github.com/wailsapp/wails/v2/internal/system/packagemanager"
	"os/exec"
	"strings"
)

var (
	IsAppleSilicon bool
)

// Info holds information about the current operating system,
// package manager and required dependancies
type Info struct {
	OS           *operatingsystem.OS
	PM           packagemanager.PackageManager
	Dependencies packagemanager.DependencyList
}

// GetInfo scans the system for operating system details,
// the system package manager and the status of required
// dependancies.
func GetInfo() (*Info, error) {
	var result Info
	err := result.discover()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func checkNPM() *packagemanager.Dependancy {

	// Check for npm
	output, err := exec.Command("npm", "-version").Output()
	installed := true
	version := ""
	if err != nil {
		installed = false
	} else {
		version = strings.TrimSpace(strings.Split(string(output), "\n")[0])
	}
	return &packagemanager.Dependancy{
		Name:           "npm ",
		PackageName:    "N/A",
		Installed:      installed,
		InstallCommand: "Available at https://nodejs.org/en/download/",
		Version:        version,
		Optional:       false,
		External:       false,
	}
}

func checkUPX() *packagemanager.Dependancy {

	// Check for npm
	output, err := exec.Command("upx", "-V").Output()
	installed := true
	version := ""
	if err != nil {
		installed = false
	} else {
		version = strings.TrimSpace(strings.Split(string(output), "\n")[0])
	}
	return &packagemanager.Dependancy{
		Name:           "upx ",
		PackageName:    "N/A",
		Installed:      installed,
		InstallCommand: "Available at https://upx.github.io/",
		Version:        version,
		Optional:       true,
		External:       false,
	}
}

func checkNSIS() *packagemanager.Dependancy {

	// Check for nsis installer
	output, err := exec.Command("makensis", "-VERSION").Output()
	installed := true
	version := ""
	if err != nil {
		installed = false
	} else {
		version = strings.TrimSpace(strings.Split(string(output), "\n")[0])
	}
	return &packagemanager.Dependancy{
		Name:           "nsis ",
		PackageName:    "N/A",
		Installed:      installed,
		InstallCommand: "Available at https://nsis.sourceforge.io/Download",
		Version:        version,
		Optional:       true,
		External:       false,
	}
}

func checkGCC() *packagemanager.Dependancy {

	version := packagemanager.AppVersion("gcc")

	return &packagemanager.Dependancy{
		Name:           "gcc ",
		PackageName:    "N/A",
		Installed:      version != "",
		InstallCommand: "Install via your package manager",
		Version:        version,
		Optional:       false,
		External:       false,
	}
}

func checkPkgConfig() *packagemanager.Dependancy {

	version := packagemanager.AppVersion("pkg-config")

	return &packagemanager.Dependancy{
		Name:           "pkg-config ",
		PackageName:    "N/A",
		Installed:      version != "",
		InstallCommand: "Install via your package manager",
		Version:        version,
		Optional:       false,
		External:       false,
	}
}

func checkLibrary(name string) func() *packagemanager.Dependancy {
	return func() *packagemanager.Dependancy {
		output, _, _ := shell.RunCommand(".", "pkg-config", "--cflags", name)
		installed := len(strings.TrimSpace(output)) > 0

		return &packagemanager.Dependancy{
			Name:           "lib" + name + " ",
			PackageName:    "N/A",
			Installed:      installed,
			InstallCommand: "Install via your package manager",
			Version:        "N/A",
			Optional:       false,
			External:       false,
		}
	}
}

func checkDocker() *packagemanager.Dependancy {

	// Check for npm
	output, err := exec.Command("docker", "version").Output()
	installed := true
	version := ""

	// Docker errors if it is not running so check for that
	if len(output) == 0 && err != nil {
		installed = false
	} else {
		// Version is in a line like: " Version:           20.10.5"
		versionOutput := strings.Split(string(output), "\n")
		for _, line := range versionOutput[1:] {
			splitLine := strings.Split(line, ":")
			if len(splitLine) > 1 {
				key := strings.TrimSpace(splitLine[0])
				if key == "Version" {
					version = strings.TrimSpace(splitLine[1])
					break
				}
			}
		}
	}
	return &packagemanager.Dependancy{
		Name:           "docker ",
		PackageName:    "N/A",
		Installed:      installed,
		InstallCommand: "Available at https://www.docker.com/products/docker-desktop",
		Version:        version,
		Optional:       true,
		External:       false,
	}
}

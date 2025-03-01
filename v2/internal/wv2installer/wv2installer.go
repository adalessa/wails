//go:build windows

package wv2installer

import (
	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/go-webview2/webviewloader"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

const MinimumRuntimeVersion string = "91.0.992.28"

type installationStatus int

const (
	needsInstalling installationStatus = iota
	needsUpdating
	installed
)

func Process(appoptions *options.App) (string, error) {

	messages := windows.DefaultMessages()
	if appoptions.Windows != nil && appoptions.Windows.Messages != nil {
		messages = appoptions.Windows.Messages
	}
	installStatus := needsInstalling
	installedVersion, err := webviewloader.GetInstalledVersion()
	if err != nil {
		return "", err
	}
	if installedVersion != "" {
		installStatus = installed
		compareResult, err := webviewloader.CompareBrowserVersions(installedVersion, MinimumRuntimeVersion)
		if err != nil {
			return "", err
		}
		updateRequired := compareResult == -1
		// Installed and does not require updating
		if !updateRequired {
			return installedVersion, nil
		}

	}
	return installedVersion, doInstallationStrategy(installStatus, messages)
}

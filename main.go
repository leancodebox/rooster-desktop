package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/leancodebox/rooster-desktop/resource"
	"github.com/leancodebox/rooster/jobmanager"
	"github.com/leancodebox/rooster/jobmanagerserver"
	"log"
	"os/exec"
	"runtime"
)

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		stop()
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func main() {
	url := "http://localhost:9090/actor/"
	a := app.New()
	logLifecycle(a)
	a.SetIcon(resource.GetLogo())
	serverErr := startRoosterServer()
	// 桌面系统设置托盘
	if desk, ok := a.(desktop.App); ok {
		var list []*fyne.MenuItem
		open := fyne.NewMenuItem("打开管理", func() {
			err := openURL(url)
			if err != nil {
				fmt.Println(err)
			}
		})
		list = append(list, open)
		if serverErr != nil {
			list = append(list, fyne.NewMenuItem(serverErr.Error(), func() {
				err := openURL(url)
				if err != nil {
					fmt.Println(err)
				}
			}))
		} else {
			err := openURL(url)
			if err != nil {
				fmt.Println(err)
			}
		}

		m := fyne.NewMenu("cock-desktop",
			list...,
		)
		desk.SetSystemTrayMenu(m)
	}
	a.Run()
}

func startRoosterServer() error {
	err := jobmanager.RegByUserConfig()
	if err != nil {
		return err
	}
	jobmanagerserver.ServeRun()
	return nil
}

func stop() {
	jobmanagerserver.ServeStop()
}

func openURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return fmt.Errorf("unsupported platform")
	}
	runCmd := exec.Command(cmd, args...)
	jobmanager.HideWindows(runCmd)
	return runCmd.Start()
}

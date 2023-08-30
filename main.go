package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/caseymrm/menuet"

	"github.com/kyleshepherd/mr-menubar/internal/gitlab"
)

var username string

func main() {
	go runMenuApp()
	menuet.App().Label = "com.github.kyleshepherd.mrbar"
	menuet.App().Children = menuItems
	menuet.App().RunApplication()
}

func runMenuApp() {
	for {
		token := menuet.Defaults().String("gitlab_token")
		if token == "" {
			updateState(nil)
		} else {
			mrs, err := gitlab.GetMRs(token)
			if err != nil {
				handleError(err)
				updateState(nil)
				return
			}
			username = mrs.Username
			updateState(mrs)
		}
		time.Sleep(time.Second * 60)
	}
}

func updateState(mrs *gitlab.MRs) {
	if mrs != nil {
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: strconv.Itoa(mrs.Assigned.Count) + " | " + strconv.Itoa(mrs.Review.Count),
			Image: "gitlab",
		})
		return
	}
	menuet.App().SetMenuState(&menuet.MenuState{
		Image: "gitlab",
	})
}

func menuItems() []menuet.MenuItem {
	browserItem := menuet.MenuItem{
		Text: "Open in Browser",
		Clicked: func() {
			err := openBrowser("https://gitlab.com/dashboard/merge_requests?assignee_username=" + username)
			if err != nil {
				handleError(err)
				return
			}
			err = openBrowser("https://gitlab.com/dashboard/merge_requests?reviewer_username=" + username)
			if err != nil {
				handleError(err)
				return
			}
		},
	}

	tokenItem := menuet.MenuItem{
		Text: "Set Token",
		Clicked: func() {
			tokenRes := menuet.App().Alert(menuet.Alert{
				MessageText:     "Set Token",
				InformativeText: "Enter your Gitlab Token below",
				Inputs:          []string{"Token"},
				Buttons:         []string{"Save", "Cancel"},
			})
			if tokenRes.Button != 0 {
				return
			}
			menuet.Defaults().SetString("gitlab_token", tokenRes.Inputs[0])
		},
	}

	if username == "" {
		return []menuet.MenuItem{tokenItem}
	}

	return []menuet.MenuItem{
		browserItem,
		tokenItem,
	}
}

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		return err
	}
	return nil
}

func handleError(err error) {
	fmt.Println(err)
	menuet.App().Alert(menuet.Alert{MessageText: "Error", InformativeText: err.Error()})
}

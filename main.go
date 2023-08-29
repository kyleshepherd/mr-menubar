package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/caseymrm/menuet"

	"github.com/kyleshepherd/mr-menubar/internal/gitlab"
)

// Application entry point
func main() {
	go runMenuApp()
	menuet.App().Label = "com.github.kyleshepherd.mrbar"
	menuet.App().RunApplication()
}

func runMenuApp() {
	for {
		mrs, err := gitlab.GetMRs()
		if err != nil {
			fmt.Println(err)
		}
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: strconv.Itoa(mrs.Assigned.Count) + " / " + strconv.Itoa(mrs.Review.Count),
			Image: "gitlab",
		})
		time.Sleep(time.Second * 60)
	}
}

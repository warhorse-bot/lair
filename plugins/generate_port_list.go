package main

import (
	"fmt"
	"sort"
	"time"

	lair "github.com/lair-framework/go-lair"
)

//Update identifies version
func Update(p *lair.Project) {
	ports := make(map[int]bool)
	for _, host := range p.Hosts {
		for _, s := range host.Services {
			if ok, _ := ports[s.Port]; ok == false {
				ports[s.Port] = true
			}
		}
	}
	if len(ports) == 0 {
		return
	}
	var portList []int
	for port := range ports {
		portList = append(portList, port)
	}
	sort.Ints(portList)
	ourNote := lair.Note{}
	ourNote.Title = fmt.Sprintf("Ports from Generate Port List Plugin %s",
		time.Now().Format("2006_01_02_15_04_05"))
	var noteContent string
	for _, port := range portList {

		noteContent += fmt.Sprintf("%d,", port)
	}
	ourNote.Content = noteContent
	ourNote.LastModifiedBy = "Generate Ports List Plugin"
	p.Notes = append(p.Notes, ourNote)
}

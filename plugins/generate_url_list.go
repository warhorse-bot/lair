package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	lair "github.com/lair-framework/go-lair"
)

type urlInfo struct {
	port  int
	host  string
	proto string
}

func formatURL(h *urlInfo) string {
	return fmt.Sprintf("%s%s:%d", h.proto, h.host, h.port)
}

//Update identifies version
func Update(p *lair.Project) {
	var urls []string
	mainReg, err := regexp.Compile(`web|www|ssl|http|https`)
	if err != nil {
		log.Printf("Generate URL List Transform error: %s\n", err.Error())
		return
	}
	httpsReg, err := regexp.Compile(`ssl|https`)
	if err != nil {
		log.Printf("Generate URL List Transform error: %s\n", err.Error())
		return
	}
	for _, host := range p.Hosts {
		for _, s := range host.Services {
			var u urlInfo
			u.proto = "http://"
			service := strings.ToLower(s.Service)
			if mainReg.MatchString(service) {
				if httpsReg.MatchString(service) {
					u.proto = "https://"
				}
				u.host = host.IPv4
				u.port = s.Port
				urls = append(urls, formatURL(&u))
				for _, n := range host.Hostnames {
					u.host = n
					urls = append(urls, formatURL(&u))
				}
			}
		}
	}
	if len(urls) == 0 {
		return
	}
	ourNote := lair.Note{}
	ourNote.Title = fmt.Sprintf("URLs from Generate URL List Plugin %s",
		time.Now().Format("2006_01_02_15_04_05"))
	ourNote.Content = strings.Join(urls, "\n")
	ourNote.LastModifiedBy = "Generate URL List Plugin"
	p.Notes = append(p.Notes, ourNote)
}

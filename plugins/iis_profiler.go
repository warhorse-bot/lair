package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	lair "github.com/lair-framework/go-lair"
)

// Update identifies versions of IIS and modifies the host OS.
func Update(p *lair.Project) {
	reg, err := regexp.Compile(`\d+\.\d+`)
	if err != nil {
		log.Printf("IIS Transform error: %s\n", err.Error())
		return
	}
	for i, h := range p.Hosts {
		for _, s := range h.Services {
			product := s.Product

			if !strings.Contains(product, "IIS") {
				continue
			}

			match := reg.FindString(product)
			if match == "" {
				continue
			}

			version, err := strconv.ParseFloat(match, 64)
			if err != nil {
				log.Printf("IIS Transform error: %s\n", err.Error())
				continue
			}
			var os string
			switch {
			case version < 6:
				os = "Microsoft Windows Server 2000"
			case version < 7:
				os = "Microsoft Windows Server 2003"
			case version < 8:
				os = "Microsoft Windows Server 2008"
			case version < 9:
				os = "Microsoft Windows Server 2012"
			case version < 10:
				os = "Microsoft Windows Server 2016"
			}
			if os == "" {
				continue
			}
			p.Hosts[i].OS.Fingerprint = os
			p.Hosts[i].OS.Tool = "IIS Profiler"
			p.Hosts[i].OS.Weight = 90
		}
	}
}

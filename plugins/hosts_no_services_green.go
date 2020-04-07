package main

import lair "github.com/lair-framework/go-lair"

// Update and mark host green with no services.
func Update(p *lair.Project) {
	for i, h := range p.Hosts {
		if len(h.Services) == 0 || (len(h.Services) == 1 && h.Services[0].Port == 0) {
			p.Hosts[i].Status = lair.StatusGreen
			p.Hosts[i].OS.Tool = "Hosts No Services Green"
		}
	}
}

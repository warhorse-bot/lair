package main

import (
	"strings"

	lair "github.com/lair-framework/go-lair"
)

// Update services.
func Update(p *lair.Project) {
	mappedServices := genServiceMap()
	for i, h := range p.Hosts {
		for j, s := range h.Services {
			service := strings.ToUpper(s.Service)

			if value, ok := mappedServices[s.Port]; ok {
				service = value
			}

			if service == "" {
				service = "UNKNOWN"
			}

			service = strings.Replace(service, "WWW", "HTTP", -1)
			service = strings.Replace(service, "HTTP-ALT", "HTTP", -1)
			service = strings.Replace(service, "HTTPS-ALT", "HTTPS", -1)
			service = strings.Replace(service, "?", "", -1)

			if service == s.Service {
				continue
			}
			p.Hosts[i].Services[j].Service = service
		}
		p.Hosts[i].OS.Tool = "Normalize Services"
	}
}

func genServiceMap() map[int]string {
	services := map[int]string{
		22:    "SSH",
		21:    "FTP",
		23:    "TELNET",
		25:    "SMTP",
		53:    "DNS",
		79:    "FINGER",
		80:    "HTTP",
		81:    "HTTP",
		111:   "RPC",
		123:   "NTP",
		135:   "MS-RPC",
		137:   "NETBIOS",
		139:   "CIFS",
		161:   "SNMP",
		443:   "HTTPS",
		445:   "CIFS",
		500:   "ISAKMP",
		1433:  "MS-SQL-TDS",
		1434:  "MS-SQL-MONITOR",
		2222:  "SSH",
		2638:  "SYBASE",
		3389:  "MS RDP",
		4786:  "SMARTINSTALL",
		5060:  "SIP",
		5222:  "XMPPCLIENT",
		7777:  "HTTP",
		8000:  "HTTP",
		8080:  "HTTP",
		8081:  "HTTP",
		8443:  "HTTPS",
		9090:  "HTTP",
		49316: "MS-SQL-TDS",
	}
	return services
}

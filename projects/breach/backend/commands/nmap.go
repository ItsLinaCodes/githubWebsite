package commands

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"

	"breach/backend/state"
	"breach/backend/world"
)

func Nmap(args []string, callback js.Value) {
	if len(args) != 1 {
		callback.Invoke("Usage: nmap <target>", "red")
		return
	}

	var ips []string
	if strings.Contains(args[0], "/") {
		ips = world.SubnetHosts()
	} else {
		ips = []string{args[0]}
	}

	callback.Invoke("Starting Nmap scan...", "dim")
	time.Sleep(500 * time.Millisecond)

	up := 0
	for _, ip := range ips {
		host, found := world.Network[ip]
		if !found {
			callback.Invoke(fmt.Sprintf("%-16s — no response (host down)", ip), "dim")
			time.Sleep(300 * time.Millisecond)
			continue
		}
		up++
		callback.Invoke(fmt.Sprintf("%-16s (%s) — Host is up", ip, host.Hostname), "green")
		for _, p := range host.Ports {
			time.Sleep(80 * time.Millisecond)
			color := "green"
			if p.State == "filtered" {
				color = "dim"
			}
			callback.Invoke(fmt.Sprintf("  %-6d/tcp  %-12s %s", p.Number, p.State, p.Service), color)
		}
		state.Current.ScannedHosts = append(state.Current.ScannedHosts, ip)
		time.Sleep(300 * time.Millisecond)
	}

	callback.Invoke(fmt.Sprintf("Scan complete. %d host(s) up.", up), "white")
}

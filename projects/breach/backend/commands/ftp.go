package commands

import (
	"breach/backend/state"
	"breach/backend/utils"
	"breach/backend/world"
	"syscall/js"
)

func FTP(args []string, callback js.Value) {
	if state.Current.ScannedHosts == nil {
		callback.Invoke("No scanned hosts available.", "red")
		return
	}

	if !utils.Contains(state.Current.ScannedHosts, "10.0.0.5") {
		callback.Invoke("Host not found. Have you scanned the network?", "red")
		return
	}

	host := world.Network["10.0.0.5"]
	ftpOpen := false
	for _, p := range host.Ports {
		if p.Number == 21 && p.State == "open" {
			ftpOpen = true
			break
		}
	}
	if !ftpOpen {
		callback.Invoke("Port 21 is not open on this host.", "red")
		return
	}

	state.Current.ActiveSession = "ftp"
	state.Current.CurrentDirectory = "/"
	callback.Invoke("Connected to 10.0.0.5 (corp-ftp).", "green")
	callback.Invoke("FTP session open. Type 'ls' to list files.", "dim")
}

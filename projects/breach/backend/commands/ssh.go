package commands

import (
	"breach/backend/state"
	"breach/backend/utils"
	"fmt"
	"strings"
	"syscall/js"
	"time"
)

func SSH(args []string, callback js.Value) {
	if len(args) != 1 {
		callback.Invoke("Usage: ssh <user>@<host>", "red")
		return
	}

	parts := strings.Split(args[0], "@")
	if len(parts) != 2 {
		callback.Invoke("Usage: ssh <user>@<host>", "red")
		return
	}

	user := parts[0]
	host := parts[1]

	if host != "10.0.0.12" {
		callback.Invoke(fmt.Sprintf("ssh: connect to host %s port 22: Connection refused", host), "red")
		return
	}

	if !utils.Contains(state.Current.ScannedHosts, "10.0.0.12") {
		callback.Invoke("ssh: connect to host 10.0.0.12 port 22: No route to host", "red")
		return
	}

	_, found := state.Current.FoundCredentials[user]
	if !found {
		callback.Invoke("Permission denied (publickey,password).", "red")
		return
	}

	callback.Invoke("Connecting to 10.0.0.12...", "dim")
	time.Sleep(600 * time.Millisecond)
	callback.Invoke("Warning: Permanently added '10.0.0.12' to the list of known hosts.", "dim")
	time.Sleep(400 * time.Millisecond)
	callback.Invoke(fmt.Sprintf("%s@10.0.0.12's password:", user), "password-prompt")

	typed := <-state.InputChan
	if typed != state.Current.FoundCredentials[user] {
		callback.Invoke("Permission denied (publickey,password).", "red")
		return
	}

	time.Sleep(500 * time.Millisecond)
	callback.Invoke("Last login: Sat Jan  4 02:11:43 2026 from 10.0.0.99", "dim")
	callback.Invoke("Welcome to corp-internal.", "green")

	state.Current.ActiveSession = "ssh"
	state.Current.CurrentDirectory = "/home/admin"
}

package main

import (
	"breach/backend/commands"
	"breach/backend/state"
	"fmt"
	"strings"
	"syscall/js"
)

type command struct {
	name string
	desc string
}

var cmds = []command{
	{"nmap <target>", "scan a host or subnet (e.g. nmap 10.0.0.0/24)"},
	{"ftp <host>", "connect to an FTP server"},
	{"ls", "list files in current directory"},
	{"cd <dir>", "change directory"},
	{"cat <file>", "read a file"},
	{"get <file>", "download a file"},
	{"crack <file>", "crack a password hash file"},
	{"ssh <user>@<host>", "connect via SSH"},
	{"exfil <file>", "exfiltrate a file (win condition)"},
	{"clear", "clear the terminal"},
	{"help", "show this message"},
}

func main() {
	js.Global().Set("runCommand", js.FuncOf(runCommand))
	js.Global().Set("submitInput", js.FuncOf(func(this js.Value, args []js.Value) any {
		state.InputChan <- args[0].String()
		return nil
	}))
	js.Global().Set("getPrompt", js.FuncOf(func(this js.Value, args []js.Value) any {
		switch state.Current.ActiveSession {
		case "ssh":
			return "admin@corp-internal:~$"
		case "ftp":
			return "ftp>"
		default:
			return "guest@corp-net:~$"
		}
	}))
	select {}
}

func runCommand(this js.Value, args []js.Value) any {
	cmd := args[0].String()
	callback := args[1]

	go func() {
		handleCommand(cmd, callback)
	}()

	return nil
}

func handleCommand(cmd string, callback js.Value) {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return
	}

	state.Current.CommandCount++

	switch parts[0] {
	case "help":
		callback.Invoke("Available commands:", "white")
		for _, cmd := range cmds {
			callback.Invoke(fmt.Sprintf("  %-24s — %s", cmd.name, cmd.desc), "dim")
		}
	case "hint":
		commands.Hint(parts[1:], callback)
	case "nmap":
		commands.Nmap(parts[1:], callback)
	case "ftp":
		commands.FTP(parts[1:], callback)
	case "ls":
		commands.Ls(parts[1:], callback)
	case "cd":
		commands.Cd(parts[1:], callback)
	case "cat":
		commands.Cat(parts[1:], callback)
	case "get":
		commands.Get(parts[1:], callback)
	case "crack":
		commands.Crack(parts[1:], callback)
	case "ssh":
		commands.SSH(parts[1:], callback)
	case "exfil":
		commands.Exfil(parts[1:], callback)
	case "clear":
		callback.Invoke("CLEAR", "clear")
	default:
		callback.Invoke("command not found: "+parts[0], "red")
	}
}

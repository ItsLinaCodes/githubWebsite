package commands

import (
	"breach/backend/state"
	"breach/backend/utils"
	"syscall/js"
)

type hint struct {
	check func() bool
	lines []string
}

var hints = []hint{
	{
		check: func() bool {
			return !utils.Contains(state.Current.ScannedHosts, "10.0.0.5")
		},
		lines: []string{
			"Tip: Start by scanning the network.",
			"  Try: nmap 10.0.0.0/24",
		},
	},
	{
		check: func() bool {
			return state.Current.ActiveSession == ""
		},
		lines: []string{
			"Tip: You found an FTP server. Connect to it.",
			"  Try: ftp 10.0.0.5",
		},
	},
	{
		check: func() bool {
			return !utils.Contains(state.Current.DownloadedFiles, "passwords.hash")
		},
		lines: []string{
			"Tip: Browse the FTP server and grab anything interesting.",
			"  Try: ls, cd private, get passwords.hash",
		},
	},
	{
		check: func() bool {
			return len(state.Current.FoundCredentials) == 0
		},
		lines: []string{
			"Tip: You have a hash file. Crack it.",
			"  Try: crack passwords.hash",
		},
	},
	{
		check: func() bool {
			return state.Current.ActiveSession != "ssh"
		},
		lines: []string{
			"Tip: You have credentials. Use them.",
			"  Try: ssh admin@10.0.0.12",
		},
	},
	{
		check: func() bool {
			return !utils.Contains(state.Current.DownloadedFiles, "confidential.db")
		},
		lines: []string{
			"Tip: You're in. Find the target file.",
			"  Try: ls, cd /data, cat confidential.db",
		},
	},
	{
		check: func() bool { return true },
		lines: []string{
			"Tip: You found the target. Exfiltrate it.",
			"  Try: exfil confidential.db",
		},
	},
}

func Hint(args []string, callback js.Value) {
	for _, h := range hints {
		if h.check() {
			for _, line := range h.lines {
				callback.Invoke(line, "dim")
			}
			return
		}
	}
}

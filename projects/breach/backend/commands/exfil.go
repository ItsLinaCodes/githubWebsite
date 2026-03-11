package commands

import (
	"breach/backend/state"
	"fmt"
	"strings"
	"syscall/js"
	"time"
)

func Exfil(args []string, callback js.Value) {
	if state.Current.ActiveSession != "ssh" {
		callback.Invoke("exfil: no active SSH session.", "red")
		return
	}
	if len(args) == 0 {
		callback.Invoke("Usage: exfil <file>", "red")
		return
	}
	if args[0] != "confidential.db" {
		callback.Invoke(fmt.Sprintf("exfil: %s: file not found or not exfiltrable.", args[0]), "red")
		return
	}
	if state.Current.CurrentDirectory != "/data" {
		callback.Invoke("exfil: confidential.db: no such file in current directory.", "red")
		return
	}

	callback.Invoke("Initiating secure exfiltration channel...", "dim")
	time.Sleep(600 * time.Millisecond)
	callback.Invoke("Encrypting payload...", "dim")
	time.Sleep(400 * time.Millisecond)

	steps := []struct {
		pct   int
		delay time.Duration
	}{
		{3, 150}, {12, 400}, {13, 80}, {27, 350}, {40, 200},
		{41, 60}, {55, 450}, {68, 280}, {81, 320}, {94, 190}, {100, 250},
	}
	for i, s := range steps {
		filled := s.pct / 5
		bar := fmt.Sprintf("[%s%s] %d%%", strings.Repeat("█", filled), strings.Repeat("░", 20-filled), s.pct)
		if i == 0 {
			callback.Invoke(bar, "dim")
		} else {
			callback.Invoke(bar, "update")
		}
		time.Sleep(s.delay * time.Millisecond)
	}

	elapsed := time.Since(state.Current.StartTime).Round(time.Second)
	state.Current.GameWon = true

	callback.Invoke("", "dim")
	callback.Invoke("████████████████████████████████████████", "green")
	callback.Invoke("  MISSION COMPLETE", "green")
	callback.Invoke("████████████████████████████████████████", "green")
	callback.Invoke("", "dim")
	callback.Invoke(fmt.Sprintf("  File exfiltrated : confidential.db"), "white")
	callback.Invoke(fmt.Sprintf("  Commands used    : %d", state.Current.CommandCount), "white")
	callback.Invoke(fmt.Sprintf("  Time elapsed     : %s", elapsed), "white")
	callback.Invoke("", "dim")
	callback.Invoke("  Target: CorpNet Q4 financials. Mission logged.", "dim")
}

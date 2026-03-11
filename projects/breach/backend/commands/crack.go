package commands

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"

	"breach/backend/state"
	"breach/backend/utils"
)

type crackStep struct {
	pct   int
	delay time.Duration
}

func Crack(args []string, callback js.Value) {
	if len(args) != 1 {
		callback.Invoke("Usage: crack <file>", "red")
		return
	}

	filename := args[0]

	if !utils.Contains(state.Current.DownloadedFiles, filename) {
		callback.Invoke(fmt.Sprintf("crack: %s: file not found. Use 'get' to download it first.", filename), "red")
		return
	}

	if filename != "passwords.hash" {
		callback.Invoke(fmt.Sprintf("crack: %s: unrecognized hash format.", filename), "red")
		return
	}

	callback.Invoke(fmt.Sprintf("Loading %s...", filename), "dim")
	time.Sleep(500 * time.Millisecond)
	callback.Invoke("Detected: bcrypt $2b$10 — launching dictionary attack", "dim")
	time.Sleep(400 * time.Millisecond)

	steps := []crackStep{
		{2, 120 * time.Millisecond},
		{5, 340 * time.Millisecond},
		{6, 90 * time.Millisecond},
		{11, 210 * time.Millisecond},
		{19, 480 * time.Millisecond},
		{20, 60 * time.Millisecond},
		{31, 390 * time.Millisecond},
		{44, 270 * time.Millisecond},
		{45, 150 * time.Millisecond},
		{58, 520 * time.Millisecond},
		{63, 180 * time.Millisecond},
		{71, 310 * time.Millisecond},
		{79, 440 * time.Millisecond},
		{80, 70 * time.Millisecond},
		{88, 290 * time.Millisecond},
		{95, 410 * time.Millisecond},
		{98, 130 * time.Millisecond},
		{100, 200 * time.Millisecond},
	}

	for i, step := range steps {
		filled := step.pct / 5
		bar := fmt.Sprintf("[%s%s] %d%%", strings.Repeat("█", filled), strings.Repeat("░", 20-filled), step.pct)
		if i == 0 {
			callback.Invoke(bar, "dim")
		} else {
			callback.Invoke(bar, "update")
		}
		time.Sleep(step.delay)
	}

	callback.Invoke("", "dim")
	state.Current.FoundCredentials["admin"] = "Summer2024!"
	callback.Invoke("[+] admin : Summer2024!", "green")
	callback.Invoke("[+] guest : (uncrackable — hash mismatch)", "dim")
	callback.Invoke("Crack complete. Credentials saved.", "white")
}

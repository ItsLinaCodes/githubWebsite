package state

import "time"

type State struct {
	ScannedHosts     []string
	ActiveSession    string
	CurrentDirectory string
	DownloadedFiles  []string
	FoundCredentials map[string]string
	GameWon          bool
	StartTime        time.Time
	CommandCount     int
}

var Current = &State{
	StartTime:        time.Now(),
	FoundCredentials: make(map[string]string),
}

var InputChan = make(chan string, 1)

package commands

import (
	"fmt"
	"sort"
	"strings"
	"syscall/js"

	"breach/backend/state"
	"breach/backend/world"
)

func currentFS() *world.Node {
	if state.Current.ActiveSession == "ssh" {
		return world.InternalFilesystem
	}
	return world.FTPFilesystem
}

func currentNode() (*world.Node, bool) {
	parts := strings.Split(strings.Trim(state.Current.CurrentDirectory, "/"), "/")
	return world.GetNode(currentFS(), parts)
}

func Ls(args []string, callback js.Value) {
	if state.Current.ActiveSession == "" {
		callback.Invoke("No active session.", "red")
		return
	}
	node, ok := currentNode()
	if !ok || !node.IsDir {
		callback.Invoke("Invalid directory.", "red")
		return
	}
	names := make([]string, 0, len(node.Children))
	for name := range node.Children {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		child := node.Children[name]
		if child.IsDir {
			callback.Invoke(name+"/", "green")
		} else {
			callback.Invoke(name, "white")
		}
	}
}

// resolvePath resolves a target path against the current directory,
// handling absolute paths, "..", and trailing slashes.
func resolvePath(target string) string {
	var base string
	if strings.HasPrefix(target, "/") {
		base = target
	} else {
		base = strings.TrimRight(state.Current.CurrentDirectory, "/") + "/" + target
	}
	// process each segment, resolving ".."
	segments := strings.Split(base, "/")
	var resolved []string
	for _, seg := range segments {
		switch seg {
		case "", ".":
			// skip empty and current-dir segments
		case "..":
			if len(resolved) > 0 {
				resolved = resolved[:len(resolved)-1]
			}
		default:
			resolved = append(resolved, seg)
		}
	}
	return "/" + strings.Join(resolved, "/")
}

func Cd(args []string, callback js.Value) {
	if state.Current.ActiveSession == "" {
		callback.Invoke("No active session.", "red")
		return
	}
	if len(args) == 0 {
		callback.Invoke("Usage: cd <dir>", "red")
		return
	}
	newPath := resolvePath(args[0])
	parts := strings.Split(strings.Trim(newPath, "/"), "/")
	node, ok := world.GetNode(currentFS(), parts)
	if !ok || !node.IsDir {
		callback.Invoke(fmt.Sprintf("cd: %s: no such directory", args[0]), "red")
		return
	}
	state.Current.CurrentDirectory = newPath
	callback.Invoke(state.Current.CurrentDirectory, "dim")
}

func Cat(args []string, callback js.Value) {
	if state.Current.ActiveSession == "" {
		callback.Invoke("No active session.", "red")
		return
	}
	if len(args) == 0 {
		callback.Invoke("Usage: cat <file>", "red")
		return
	}
	node, ok := resolveFile(args[0])
	if !ok {
		callback.Invoke(fmt.Sprintf("cat: %s: no such file", args[0]), "red")
		return
	}
	for _, line := range strings.Split(node.Content, "\n") {
		callback.Invoke(line, "white")
	}
}

func Get(args []string, callback js.Value) {
	if state.Current.ActiveSession == "" {
		callback.Invoke("No active session.", "red")
		return
	}
	if len(args) == 0 {
		callback.Invoke("Usage: get <file>", "red")
		return
	}
	node, ok := resolveFile(args[0])
	if !ok {
		callback.Invoke(fmt.Sprintf("get: %s: no such file", args[0]), "red")
		return
	}
	if node.IsDir {
		callback.Invoke(fmt.Sprintf("get: %s: is a directory", args[0]), "red")
		return
	}
	state.Current.DownloadedFiles = append(state.Current.DownloadedFiles, args[0])
	callback.Invoke(fmt.Sprintf("Downloaded: %s", args[0]), "green")
}

func resolveFile(name string) (*world.Node, bool) {
	node, ok := currentNode()
	if !ok {
		return nil, false
	}
	child, ok := node.Children[name]
	return child, ok
}

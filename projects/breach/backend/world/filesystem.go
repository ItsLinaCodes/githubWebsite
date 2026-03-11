package world

type Node struct {
	IsDir    bool
	Content  string
	Children map[string]*Node
}

var FTPFilesystem = &Node{
	IsDir: true,
	Children: map[string]*Node{
		"pub": {
			IsDir: true,
			Children: map[string]*Node{
				"readme.txt": {
					Content: "Welcome to CorpNet FTP server.\nAuthorized access only.\nContact sysadmin@corp.internal for issues.",
				},
				"backup.tar.gz": {
					Content: "[binary archive — use 'get backup.tar.gz' to download]",
				},
			},
		},
		"private": {
			IsDir: true,
			Children: map[string]*Node{
				"passwords.hash": {
					Content: "admin:$2b$12$KIXoB1234abcdEFGHijklOZpqrstUVWXyz567890ABCDEFghijklmno\nguest:$2b$12$aaaabbbbccccddddeeee00ZZZZyyyyxxxx1111222233334444555566",
				},
			},
		},
	},
}

var InternalFilesystem = &Node{
	IsDir: true,
	Children: map[string]*Node{
		"home": {
			IsDir: true,
			Children: map[string]*Node{
				"admin": {
					IsDir: true,
					Children: map[string]*Node{
						".bash_history": {
							Content: "ls /data\ncat /data/confidential.db\nmysql -u root -p\nexit",
						},
					},
				},
			},
		},
		"data": {
			IsDir: true,
			Children: map[string]*Node{
				"confidential.db": {
					Content: "CONFIDENTIAL — CorpNet Q4 financials, merger targets, and employee SSNs.\n[EXFIL TARGET]",
				},
				"logs.txt": {
					Content: "2026-01-04 02:11 admin logged in from 10.0.0.99\n2026-01-04 02:13 file access: /data/confidential.db\n2026-01-04 02:15 admin logged out",
				},
			},
		},
		"etc": {
			IsDir: true,
			Children: map[string]*Node{
				"passwd": {
					Content: "root:x:0:0:root:/root:/bin/bash\nadmin:x:1000:1000::/home/admin:/bin/bash\nguest:x:1001:1001::/home/guest:/bin/sh",
				},
			},
		},
	},
}

func GetNode(root *Node, parts []string) (*Node, bool) {
	current := root
	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}
		child, ok := current.Children[part]
		if !ok {
			return nil, false
		}
		current = child
	}
	return current, true
}

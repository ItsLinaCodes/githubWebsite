package world

type Port struct {
	Number  int
	Service string
	State   string
}

type Host struct {
	Hostname string
	OS       string
	Ports    []Port
}

var Network = map[string]Host{
	"10.0.0.1": {
		Hostname: "corp-gateway",
		OS:       "Linux 4.15",
		Ports: []Port{
			{22, "ssh", "open"},
			{80, "http", "open"},
			{443, "https", "open"},
		},
	},
	"10.0.0.5": {
		Hostname: "corp-ftp",
		OS:       "Linux 3.10",
		Ports: []Port{
			{21, "ftp", "open"},
			{22, "ssh", "filtered"},
		},
	},
	"10.0.0.12": {
		Hostname: "corp-internal",
		OS:       "Linux 5.4",
		Ports: []Port{
			{22, "ssh", "open"},
			{3306, "mysql", "filtered"},
			{8080, "http-alt", "open"},
		},
	},
}

func SubnetHosts() []string {
	return []string{"10.0.0.1", "10.0.0.5", "10.0.0.12"}
}

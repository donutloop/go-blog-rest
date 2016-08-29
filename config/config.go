package config

type Configuration struct {
	DebugMode bool
	Server server `toml:"server"`
	Database   database `toml:"database"`
}

type database struct {
	Hostname  string
	Port   int
}

type server struct {
	Hostname string
	Port int
}

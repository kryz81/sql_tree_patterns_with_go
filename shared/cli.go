package shared

import "flag"

type Config struct {
	Pattern string
	Cmd     string
	DbName  string
	DbUser  string
	DbPass  string
	DbSrv   string
	DbPort  int
}

func ParseFlags() Config {
	cfg := Config{}

	flag.StringVar(&cfg.Pattern, "pattern", "path", "Pattern to use")
	flag.StringVar(&cfg.Cmd, "cmd", "list", "Command to run")
	flag.StringVar(&cfg.DbName, "db", "go_tree", "Database name")
	flag.StringVar(&cfg.DbUser, "user", "postgres", "Username")
	flag.StringVar(&cfg.DbPass, "pass", "root", "Password")
	flag.StringVar(&cfg.DbSrv, "srv", "localhost", "DB server")
	flag.IntVar(&cfg.DbPort, "port", 54321, "DB port")
	flag.Parse()

	return cfg
}

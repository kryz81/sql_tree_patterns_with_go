package main

import (
	"fmt"
	"github.com/kryz81/sql_tree_patterns_with_go/path_enumeration"
	"github.com/kryz81/sql_tree_patterns_with_go/shared"
	"log"
)

func main() {
	cfg := shared.ParseFlags()

	db, err := shared.Connect(cfg.DbSrv, cfg.DbPort, cfg.DbName, cfg.DbUser, cfg.DbPass)

	if err != nil {
		_ = fmt.Errorf("error when connecting: %s", err)
	}

	var e error

	switch cfg.Pattern {
	case "path":
		e = path_enumeration.Run(db, cfg.Cmd)
	default:
		log.Fatalln("provided pattern is not yet supported")
	}

	if e != nil {
		_ = fmt.Errorf("error: %s", e)
	}
}

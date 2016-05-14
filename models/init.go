package models

import (
    _ "github.com/mattn/go-sqlite3"
    "github.com/go-xorm/xorm"
    "github.com/go-xorm/core"
    "log"
    "flag"
)

var (
    Engine *xorm.Engine
    db string
)

func init() {
    flag.StringVar(&db, "db", "./test.db", "db path")
    var err error
    Engine, err = xorm.NewEngine("sqlite3", db)
    if err != nil {
        log.Fatal("Cannot initialize db:", db)
    }
    Engine.SetMapper(core.SameMapper{})
}

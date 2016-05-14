package models

import (
    _ "github.com/mattn/go-sqlite3"
    "github.com/go-xorm/xorm"
    "github.com/go-xorm/core"
    "log"
    "github.com/tuxlinuxien/gootstrap/config"
)

var (
    Engine *xorm.Engine
)

func init() {
    var err error
    db := config.Get("db").(string)
    Engine, err = xorm.NewEngine("sqlite3", db)
    if err != nil {
        log.Fatal("Cannot initialize db:", db)
    }
    Engine.SetMapper(core.SameMapper{})
}

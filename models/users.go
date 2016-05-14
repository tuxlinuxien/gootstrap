package models

import (
    "time"
    "log"
)

type User struct {
    Id int64
    Email string `xorm:"varchar(50) not null unique"`
    Password string `xorm:"TEXT"`
    Type string `xorm:"varchar(20) not null"`
    Created time.Time `xorm:"created"`
}

func init() {
    log.Println("Create user table...")
    err := Engine.Sync2(new(User))
    if err != nil {
        log.Fatal("Create user table error:", err)
    }
}

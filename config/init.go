package config

import (
    "github.com/spf13/viper"
    "log"
)

func init() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.SetDefault("port", "8080")
    viper.SetDefault("db", "./test.db")
    err := viper.ReadInConfig() // Find and read the config file
    if err != nil { // Handle errors reading the config file
        log.Fatal("Fatal error config file:", err)
    }
}

func Get(name string) interface{} {
    return viper.Get(name)
}

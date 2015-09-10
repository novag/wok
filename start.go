package main

import (
    "github.com/novag/wok/bot"
    _ "github.com/novag/wok/modules"
)

func main() {
    bot.Run(&bot.Config {
        Server:   "irc.hackint.org:6667",
        UseTLS:   false,

        Channel:  "#wok-test",
        User:     "",
        Nick:     "",
        Password: "",

        Admin:    "",
        LGServer: "172.22.232.1:4747",

        Verbose:  false,
    })
}
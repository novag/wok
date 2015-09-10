package modules

import (
    "fmt"
    "net"
    "io"
    "bufio"

    "github.com/novag/wok/bot"
)

func TCPConnect(channel string, sender string, arg string) {
    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Notice(sender, "Error: Please try again later")
        return
    }
    fmt.Fprintf(lg, "tcpconnect %s\n", arg)

    reader := bufio.NewReader(lg)
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            lg.Close()
            break
        }

        bot.Notice(sender, line)
    }
}

func CTCPConnect(channel string, sender string, arg string) {
    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Notice(channel, "Error: Please try again later")
        return
    }
    fmt.Fprintf(lg, "tcpconnect %s\n", arg)

    reader := bufio.NewReader(lg)
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            lg.Close()
            break
        }

        bot.Notice(channel, line)
    }
}

func init() {
    bot.RegisterPrivate(&bot.ModInfo {
        Name:        "TCPConnect",
        Description: "Checks if the given TCP port is reachable",

        Help:        true,

        Command:     "tcpconnect",
        Callback:    TCPConnect,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Name:        "TCPConnect",
        Description: "Checks if the given TCP port is reachable",

        Help:        true,

        Command:     "tcpconnect",
        Callback:    CTCPConnect,
    })
}
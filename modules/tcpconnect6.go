package modules

import (
    "fmt"
    "net"
    "io"
    "bufio"

    "github.com/novag/wok/bot"
)

func TCPConnect6(channel string, sender string, arg string) {
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

func CTCPConnect6(channel string, sender string, arg string) {
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

        bot.Notice(channel, line)
    }
}

func init() {
    bot.RegisterPrivate(&bot.ModInfo {
        Name:        "TCPConnect6",
        Description: "Checks if the given TCP port is reachable",

        Help:        true,

        Command:     "tcpconnect6",
        Callback:    TCPConnect6,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Name:        "TCPConnect6",
        Description: "Checks if the given TCP port is reachable",

        Help:        true,

        Command:     "tcpconnect6",
        Callback:    CTCPConnect6,
    })
}
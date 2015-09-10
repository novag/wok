package modules

import (
    "fmt"
    "strings"
    "net"
    "io"
    "bufio"

    "github.com/novag/wok/bot"
)

func Route(channel string, sender string, arg string) {
    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Notice(sender, "Error: Please try again later")
        return
    }
    fmt.Fprintf(lg, "route %s\n", arg)

    reader := bufio.NewReader(lg)
    for i := 0; ; i++ {
        line, err := reader.ReadString('\n')
        if err == io.EOF || (strings.Contains(line, "via") && i > 0) {
            lg.Close()
            break
        }

        bot.Notice(sender, line)
    }
}

func CRoute(channel string, sender string, arg string) {
    status := 0

    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Noticef(channel, "route %s: Error", arg)
        return
    }
    fmt.Fprintf(lg, "route %s\n", arg)

    for {
        res, err := bufio.NewReader(lg).ReadString('\n')
        if err == io.EOF {
            lg.Close()
            break
        }
        if strings.Contains(res, "via") {
            status = 1
            arg = res[:strings.Index(res, " ")]
            break
        }
    }
    if status == 1 {
        bot.Noticef(channel, "route %s: Network in table", arg)
    } else {
        bot.Noticef(channel, "route %s: Network not in table", arg)
    }
}

func init() {
    bot.RegisterPrivate(&bot.ModInfo {
        Name:        "Route",
        Description: "Preferred peer and additional information for given IPv4 network",

        Help:        true,

        Command:     "route",
        Callback:    Route,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Name:        "Route",
        Description: "Wheter given IPv4 network is in IPv4 table and it's CIDR notation",

        Help:        true,

        Command:     "route",
        Callback:    CRoute,
    })
}
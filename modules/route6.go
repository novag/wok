package modules

import (
    "fmt"
    "strings"
    "net"
    "io"
    "bufio"

    "github.com/novag/wok/bot"
)

func Route6(channel string, sender string, arg string) {
    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Notice(sender, "Error: Please try again later")
        return
    }
    fmt.Fprintf(lg, "route6 %s\n", arg)

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

func CRoute6(channel string, sender string, arg string) {
    status := 0

    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Noticef(channel, "route6 %s: Error", arg)
        return
    }
    fmt.Fprintf(lg, "route6 %s\n", arg)
    
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
        bot.Noticef(channel, "route6 %s: Network in table", arg)
    } else {
        bot.Noticef(channel, "route6 %s: Network not in table", arg)
    }
}

func init() {
    bot.RegisterPrivate(&bot.ModInfo {
        Name:        "Route6",
        Description: "Preferred peer and additional information for given IPv6 network",

        Help:        true,

        Command:     "route6",
        Callback:    Route6,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Name:        "Route6",
        Description: "Wheter given IPv4 network is in IPv6 table and it's CIDR notation",

        Help:        true,

        Command:     "route6",
        Callback:    CRoute6,
    })
}
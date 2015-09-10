package modules

import (
    "fmt"
    "strings"
    "net"
    "io"
    "bufio"

    "github.com/novag/wok/bot"
)

func Trace6(channel string, sender string, arg string) {
    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Notice(sender, "Error: Please try again later")
        return
    }
    fmt.Fprintf(lg, "traceroute6 %s\n", arg)

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

func CTrace6(channel string, sender string, arg string) {
    status := 0
    var hops int

    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Noticef(channel, "traceroute6 %s: Error", arg)
        return
    }
    fmt.Fprintf(lg, "traceroute6 %s\n", arg)

    buf := make([]byte, 1024)
    n, err := io.ReadFull(lg, buf)
    if err == io.EOF {
        lg.Close()
        return
    }
    strbuf := string(buf[:n])

    if strings.Contains(strbuf, "unreachable") {
        status = 1
    } else if strings.Contains(strbuf, "Name or service") {
        status = 2
    } else if strings.Contains(strbuf, "traceroute to") {
        if strings.Contains(strbuf, "*") {
            status = 3
        }
        hops = len(strings.Split(strbuf, "\n")) - 2
    } else {
        status = 4
    }
    if status == 1 {
        bot.Noticef(channel, "traceroute6 %s: Unreachable", arg)
    } else if status == 2 {
        bot.Noticef(channel, "traceroute6 %s: Name or service not known", arg)
    } else if status == 4 {
        bot.Noticef(channel, "traceroute6 %s: Unknown error occurred", arg)
    } else {
        if hops == 1 {
            bot.Noticef(channel, "traceroute6 %s: 1 hop", arg)
        } else if hops == 30 {
            bot.Noticef(channel, "traceroute6 %s: 30+ hops", arg)
        } else if status == 3 {
            bot.Noticef(channel, "traceroute6 %s: %d hops (Timeouts occured)", arg, hops)
        } else {
            bot.Noticef(channel, "traceroute6 %s: %d hops", arg, hops)
        }
    }
}

func init() {
    bot.RegisterPrivate(&bot.ModInfo {
        Name:        "Trace6",
        Description: "Result of traceroute to given IPv6 destination",

        Help:        true,

        Command:     "trace6",
        Callback:    Trace6,
    })

    bot.RegisterPrivate(&bot.ModInfo {
        Help:        false,

        Command:     "traceroute6",
        Callback:    Trace6,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Name:        "Trace6",
        Description: "Number of hops to given IPv6 destination",

        Help:        true,

        Command:     "trace6",
        Callback:    CTrace6,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Help:        false,

        Command:     "traceroute6",
        Callback:    CTrace6,
    })
}
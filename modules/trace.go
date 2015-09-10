package modules

import (
    "fmt"
    "strings"
    "net"
    "io"
    "bufio"

    "github.com/novag/wok/bot"
)

func Trace(channel string, sender string, arg string) {
    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Notice(sender, "Error: Please try again later")
        return
    }
    fmt.Fprintf(lg, "traceroute %s\n", arg)

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

func CTrace(channel string, sender string, arg string) {
    status := 0
    var hops int

    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Noticef(channel, "traceroute %s: Error", arg)
        return
    }
    fmt.Fprintf(lg, "traceroute %s\n", arg)

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
        bot.Noticef(channel, "traceroute %s: Unreachable", arg)
    } else if status == 2 {
        bot.Noticef(channel, "traceroute %s: Name or service not known", arg)
    } else if status == 4 {
        bot.Noticef(channel, "traceroute %s: Unknown error occurred", arg)
    } else {
        if hops == 1 {
            bot.Noticef(channel, "traceroute %s: 1 hop", arg)
        } else if hops == 30 {
            bot.Noticef(channel, "traceroute %s: 30+ hops", arg)
        } else if status == 3 {
            bot.Noticef(channel, "traceroute %s: %d hops (Timeouts occured)", arg, hops)
        } else {
            bot.Noticef(channel, "traceroute %s: %d hops", arg, hops)
        }
    }
}

func init() {
    bot.RegisterPrivate(&bot.ModInfo {
        Name:        "Trace",
        Description: "Result of traceroute to given IPv4 destination",

        Help:        true,

        Command:     "trace",
        Callback:    Trace,
    })

    bot.RegisterPrivate(&bot.ModInfo {
        Help:        false,

        Command:     "traceroute",
        Callback:    Trace,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Name:        "Trace",
        Description: "Number of hops to given IPv4 destination",

        Help:        true,

        Command:     "trace",
        Callback:    CTrace,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Help:        false,

        Command:     "traceroute",
        Callback:    CTrace,
    })
}
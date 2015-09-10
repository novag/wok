package modules

import (
    "fmt"
    "regexp"
    "strings"
    "net"
    "io"
    "bufio"
    "strconv"

    "github.com/novag/wok/bot"
)

func Ping(channel string, sender string, arg string) {
    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Notice(sender, "Error: Please try again later")
        return
    }
    fmt.Fprintf(lg, "ping %s\n", arg)

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

func CPing(channel string, sender string, arg string) {
    status := 0
    var average float64

    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Noticef(channel, "ping %s: Error", arg)
        return
    }
    fmt.Fprintf(lg, "ping %s\n", arg)

    buf := make([]byte, 1024)
    n, err := io.ReadFull(lg, buf)
    if err == io.EOF {
        lg.Close()
        return
    }
    strbuf := string(buf[:n])
    
    if strings.Contains(strbuf, "unreachable") {
        status = 1
    } else if strings.Contains(strbuf, "Time to live exceeded") {
        status = 2
    } else if strings.Contains(strbuf, "statistics") {
        re := regexp.MustCompile(`time=(\d+\.?\d*)`)
        matches := re.FindAllStringSubmatch(strbuf, -1)
        if matches != nil {
            for _, ms := range matches[1:] {
                fms, _ := strconv.ParseFloat(ms[1], 64)
                average += fms
            }
            average = average / float64(len(matches[1:]))
        } else {
            status = 3
        }
    } else {
        status = 4
    }
    if status == 1 {
        bot.Noticef(channel, "ping %s: Unreachable", arg)
    } else if status == 2 {
        bot.Noticef(channel, "ping %s: Time to live exceeded", arg)
    } else if status == 3 {
        bot.Noticef(channel, "ping %s: 100%% packet loss", arg)
    } else if status == 4 {
        bot.Noticef(channel, "ping %s: Unknown error occurred", arg)
    } else {
        bot.Noticef(channel, "ping %s: %.1f ms", arg, average)
    }
}

func init() {
    bot.RegisterPrivate(&bot.ModInfo {
        Name:        "Ping",
        Description: "Result of 5 ECHO_REQUEST packets to given IPv4 destination",

        Help:        true,

        Command:     "ping",
        Callback:    Ping,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Name:        "Ping",
        Description: "Average time of 5 ECHO_REQUEST packets to given IPv4 destination",

        Help:        true,

        Command:     "ping",
        Callback:    CPing,
    })
}
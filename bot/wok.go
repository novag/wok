package bot

import (
    "fmt"
    "regexp"
    "time"
    "log"

    "github.com/thoj/go-ircevent"
)

var (
    irccon *irc.Connection

    blockedNicks map[string]int64
)

func handleMessage(e *irc.Event) {
    _channel := e.Arguments[0]
    _sender := e.Nick
    _line := e.Message()
    isPrivate := 0
    if _channel == config.Nick { isPrivate = 1 }

    if _, ok := blockedNicks[_sender]; ok {
        return
    }

    // Validate message
    var re *regexp.Regexp
    if isPrivate == 1 {
        re = regexp.MustCompile(`^!?([a-zA-Z1-9]+)(?:$|\s)(.*)$`)
    } else {
        re = regexp.MustCompile(`^!([a-zA-Z1-9]+)(?:$|\s)(.*)$`)
    }
    match := re.FindStringSubmatch(_line)
    if match == nil { return }

    _cmd := match[1]
    _args := match[2]

    if info := registry[isPrivate][_cmd]; info != nil {
        fmt.Printf("Received message: channel=\"%s\" sender=\"%s\" command=\"%s\" arguments=\"%s\"\n", _channel, _sender, _cmd, _args)
        info.Callback(_channel, _sender, _args)
    }
}

func Run(c *Config) {
    config = c

    blockedNicks = make(map[string]int64)

    irccon = irc.IRC(config.User, config.Nick)
    irccon.UseTLS = config.UseTLS
    irccon.VerboseCallbackHandler = config.Verbose
    err := irccon.Connect(config.Server)
    if err != nil {
        log.Fatal(err)
        return
    }

    irccon.AddCallback("001", func(e *irc.Event) {
        irccon.Privmsgf("nickserv", "identify %s", config.Password)
        time.Sleep(5 * time.Second)
        irccon.Join(config.Channel)
    })

    setupMessageQueue()

    irccon.AddCallback("PRIVMSG", handleMessage)

    irccon.Loop()
}
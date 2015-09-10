# wok - IRC bot

## Overview

TODO

Version: v1.0.0

## Implemented commands

TODO

## Contribute

```go
package modules

import (
    "github.com/novag/wok/bot"
)

func Example(channel string, sender string, arg string) {
    bot.Notice(sender, "Ohai!")
    
    bot.Noticef(sender, "Ohai, %s!", sender)
}

func CExample(channel string, sender string, arg string) {
    bot.Notice(channel, "Ohai!")
    
    bot.Noticef(channel, "Ohai, %s!", sender)
}

func init() {
    bot.RegisterPrivate(&bot.ModInfo {
        Name:        "Example",
        Description: "This is an example query command",

        Help:        true,

        Command:     "example",
        Callback:    Example,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Name:        "Example",
        Description: "This is an example channel command",

        Help:        true,

        Command:     "example",
        Callback:    CExample,
    })
}
```

## License

Code is licensed under GNU General Public License, version 3 (GPL-3.0).

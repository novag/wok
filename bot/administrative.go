package bot

import (
    "fmt"
    "time"
)

func ABlock(channel string, sender string, arg string) {
    if sender != config.Admin { return }

    blockedNicks[arg] = time.Now().Unix()
    fmt.Printf("User blocked: admin=\"%s\" nick=\"%s\"\n", sender, arg)
}

func AUnblock(channel string, sender string, arg string) {
    if sender != config.Admin { return }

    delete(blockedNicks, arg)
    fmt.Printf("User unblocked: admin=\"%s\" nick=\"%s\"\n", sender, arg)
}

func init() {
    RegisterPrivate(&ModInfo {
        Name:        "Nickblock",
        Description: "Blocks specified nickname from executing commands",

        Help:        false,

        Command:     "block",
        Callback:    ABlock,
    })

    RegisterPrivate(&ModInfo {
        Name:        "Nickblock",
        Description: "Unblocks specified nickname from executing commands",

        Help:        false,

        Command:     "unblock",
        Callback:    AUnblock,
    })
}
package bot

import (
    "strings"
    "sort"
)

func showHelp(sender string) {
    var channel_commands []string
    var private_commands []string

    for _, info := range registry[0] {
        if info.Help {
            channel_commands = append(channel_commands, info.Command)
        }
    }
    for _, info := range registry[1] {
        if info.Help {
            private_commands = append(private_commands, info.Command)
        }
    }
    
    // Alphabetical order
    sort.Strings(channel_commands)
    sort.Strings(private_commands)

    Notice(sender, "Type: 'help <command>' to get details about a specific command")
    Noticef(sender, "Channel commands: %v", strings.Join(channel_commands, ", "))
    Noticef(sender, "Query commands: %v", strings.Join(private_commands, ", "))
}

func showCommandHelp(sender string, arg string) {
    available := false
    if info, ok := registry[0][arg]; ok && info.Help {
        Noticef(sender, "Channel description: %s", info.Description)

        available = true
    }
    if info, ok := registry[1][arg]; ok && info.Help {
        Noticef(sender, "Query description: %s", info.Description)

        available = true
    }
    if !available {
        Notice(sender, "No description available")
    }
}

func Help(channel string, sender string, arg string) {
    if arg == "" {
        showHelp(sender)
    } else {
        showCommandHelp(sender, arg)
    }
}

func init() {
    RegisterPrivate(&ModInfo {
        Name:        "Help",
        Description: "Available commands",

        Help:        true,

        Command:     "help",
        Callback:    Help,
    })
}

package bot

import (
    "fmt"
)

type CallbackFunc func(room string, sender string, args string)

type ModInfo struct {
    Name string
    Description string

    Help bool

    Command string
    Callback CallbackFunc
}

var registry = make(map[int]map[string]*ModInfo)

func register(isPrivate int, info *ModInfo) {
    if _, ok := registry[isPrivate]; !ok {
        registry[isPrivate] = make(map[string]*ModInfo)
    }

    registry[isPrivate][info.Command] = info
    
    fmt.Printf("Module '%s' registered: private=%d command=\"%s\"\n", info.Name, isPrivate, info.Command)
}

func RegisterChannel(info *ModInfo) {
    register(0, info)
}

func RegisterPrivate(info *ModInfo) {
    register(1, info)
}
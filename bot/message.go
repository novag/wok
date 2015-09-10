package bot

type Message struct {
    sender  string
    data    string
    args    []interface{}

    msgType int
}

var queue chan Message

func Notice(sender string, data string) {
    queue <- Message {
        sender:  sender,
        data:    data,
        args:    nil,

        msgType: 0,
    }
}

func Noticef(sender string, formatData string, args ...interface{}) {
    queue <- Message {
        sender: sender,
        data:   formatData,
        args:   args,

        msgType: 0,
    }
}

func setupMessageQueue() {
    queue = make(chan Message)

    go func() {
        for {
            message := <- queue

            switch message.msgType {
            case 0:
                if message.args == nil {
                    irccon.Notice(message.sender, message.data)
                } else {
                    irccon.Noticef(message.sender, message.data, message.args...)
                }
            }
        }
    }()
}
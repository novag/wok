package bot

type Config struct {
    Server        string   // IRC server:port
    UseTLS        bool     // Enable TLS

    Channel       string   // Channel to connect
    User          string   // IRC username
    Nick          string   // IRC nick
    Password      string   // NickServ password

    Admin         string   // Admin
    LGServer      string   // LG daemon server:port

    Verbose       bool     // Verbose IRC communication
}

var config *Config

func GetLGServer() string {
    return config.LGServer
}
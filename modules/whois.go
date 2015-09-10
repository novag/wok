package modules

import (
    "fmt"
    "strings"
    "regexp"
    "net"
    "io"
    "time"

    "github.com/novag/wok/bot"
)

func Whois(channel string, sender string, arg string) {
    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Notice(sender, "Error: Please try again later")
        return
    }
    fmt.Fprintf(lg, "whois %s\n", arg)

    buf := make([]byte, 3072)
    n, err := io.ReadFull(lg, buf)
    if err == io.EOF {
        lg.Close()
        return
    }
    strbuf := string(buf[:n])
    ilb := strings.LastIndex(strbuf, "% Information")
    if ilb == -1 { ilb = 0 }
    for _, line := range strings.Split(strbuf[ilb:], "\n") {
        bot.Notice(sender, line)
        time.Sleep(1 * time.Second)
    }
}

func CWhois(channel string, sender string, arg string) {
    status := 0

    lg, err := net.Dial("tcp", bot.GetLGServer())
    if err != nil {
        bot.Noticef(channel, "whois %s: Error", arg)
        return
    }
    fmt.Fprintf(lg, "whois %s\n", arg)

    buf := make([]byte, 3072)
    n, err := io.ReadFull(lg, buf)
    if err == io.EOF {
        lg.Close()
        return
    }
    strbuf := string(buf[:n])
    
    if strings.Contains(strbuf, "No match found") {
        status = 1
    } else if strings.Contains(strbuf, "dns/") {
        re_admin := regexp.MustCompile(`admin-c:\s+([\w-]+)`)
        re_status := regexp.MustCompile(`status:\s+([\w\s#:']+)`)
        admin_matches := re_admin.FindAllStringSubmatch(strbuf, -1)
        status_matches := re_status.FindAllStringSubmatch(strbuf, -1)
        if admin_matches != nil && status_matches != nil {
            bot.Noticef(channel, "whois %s: Admin-C: %s, Status: %s", arg, admin_matches[len(admin_matches)-1][1], status_matches[len(status_matches)-1][1])
            return
        } else {
            status = 2
        }
    } else if strings.Contains(strbuf, "inetnum/") {
        re_admin := regexp.MustCompile(`admin-c:\s+([\w-]+)`)
        re_status := regexp.MustCompile(`bgp-status:\s+([\w\s#:']+)`)
        admin_matches := re_admin.FindAllStringSubmatch(strbuf, -1)
        status_matches := re_status.FindAllStringSubmatch(strbuf, -1)
        if admin_matches != nil && status_matches != nil {
            bot.Noticef(channel, "whois %s: Admin-C: %s, BGP: %s", arg, admin_matches[len(admin_matches)-1][1], status_matches[len(status_matches)-1][1])
            return
        } else {
            status = 2
        }
    } else if strings.Contains(strbuf, "person/") {
        re_email := regexp.MustCompile(`e-mail:\s+([\w@\.-]+)`)
        re_contact := regexp.MustCompile(`contact:\s+(.+)`)
        email_matches := re_email.FindAllStringSubmatch(strbuf, -1)
        contact_matches := re_contact.FindAllStringSubmatch(strbuf, -1)
        if email_matches != nil && contact_matches == nil {
            bot.Noticef(channel, "whois %s: E-Mail: %s", arg, email_matches[len(email_matches)-1][1])
            return
        } else if contact_matches != nil && email_matches == nil {
            bot.Noticef(channel, "whois %s: Contact: %s", arg, contact_matches[len(contact_matches)-1][1])
            return
        } else if contact_matches != nil && email_matches != nil {
            bot.Noticef(channel, "whois %s: E-Mail: %s, Contact: %s", arg, email_matches[len(email_matches)-1][1], contact_matches[len(contact_matches)-1][1])
            return
        } else {
            status = 2
        }
    } else if strings.Contains(strbuf, "aut-num/") {
        re := regexp.MustCompile(`admin-c:\s+([\w-]+)`)
        matches := re.FindAllStringSubmatch(strbuf, -1)
        if matches != nil {
            bot.Noticef(channel, "whois %s: Admin-C: %s", arg, matches[len(matches)-1][1])
            return
        } else {
            status = 2
        }
    } else {
        status = 2
    }
    if status == 1 {
        bot.Noticef(channel, "whois %s: No match found", arg)
    } else if status == 2 {
        bot.Noticef(channel, "whois %s: No channel preview available, please query me", arg)
    }
}

func init() {
    bot.RegisterPrivate(&bot.ModInfo {
        Name:        "Whois",
        Description: "Last whois block for query",

        Help:        true,

        Command:     "whois",
        Callback:    Whois,
    })

    bot.RegisterChannel(&bot.ModInfo {
        Name:        "Whois",
        Description: "Admin-C, status, contact, ...",

        Help:        true,

        Command:     "whois",
        Callback:    CWhois,
    })
}
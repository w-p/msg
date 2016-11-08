package main

import (
    "os"
    "fmt"
    "bytes"
    "strings"
    "net/http"
    "io/ioutil"
    "github.com/urfave/cli"
)

func main () {
    var endpoint string
    var title string
    var level string
    var message string
    var blob string
    var debug bool

    app := cli.NewApp()

    app.Name = "msg"
    app.Usage = "Send a message to a slack webhook"
    app.Version = "1.0.0"
    app.Authors = []cli.Author{
      cli.Author{
        Name:  "Will Palmer",
        Email: "will@steelhive.com",
      },
    }

    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "endpoint, e",
            Usage: "slack incoming webhook `url`",
            Destination: &endpoint,
        },
        cli.StringFlag{
            Name: "title, t",
            Usage: "`title` for the message",
            Destination: &title,
        },
        cli.StringFlag{
            Name: "level, l",
            Usage: "log `level` (info, warn, error)",
            Destination: &level,
        },
        cli.StringFlag{
            Name: "message, m",
            Usage: "`text` to send",
            Destination: &message,
        },
        cli.StringFlag{
            Name: "json, j",
            Usage: "json `blob` to send (overrides messsage)",
            Destination: &blob,
        },
        cli.BoolFlag{
            Name: "debug, d",
            Usage: "print debug info",
            Destination: &debug,
        },
    }

    app.Action = func (c *cli.Context) error {
        if (debug) {
            fmt.Println("*** debugging enabled - this might be noisy ***")
        }

        icon := ""
        color := "#777777";
        switch strings.ToLower(level) {
        case "info":
            icon = ":no_mouth:"
            color = "#22EE22"
        case "warn":
            icon = ":open_mouth:"
            color = "#FF9922"
        case "error":
            icon = ":angry:"
            color = "#FF0000"
        default:
            panic("Error: log level must be one of info, warn, or error.")
        }

        payload := fmt.Sprintf(
            `{
                "attachments": [
                    {
                        "title": "%s %s",
                        "text": "%s",
                        "mrkdwn_in": ["text"],
                        "color": "%s"
                    }
                ]
            }`,
            icon, title, message, color)

        if (endpoint == "" || message == "") {
            panic("Error: Both and endpoint and a message are required.")
        }

        data := bytes.NewBuffer([]byte(payload))
        client := &http.Client{}
        req, err := http.NewRequest("POST", endpoint, data)
        req.Header.Set("Content-Type", "application/json")

        resp, err := client.Do(req)
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        if (strings.Contains(resp.Status, "200")) {
            fmt.Println("send: ok")
        } else {
            fmt.Println("send: fail")
        }
        if (debug) {
            fmt.Println(" code:     ", resp.Status)
            fmt.Println(" body:     ", string(body))
            fmt.Println(" endpoint: ", endpoint)
            fmt.Println(" message:  ", message)
            fmt.Println(" payload:  ", payload)
        }
        return nil
    }

    app.Run(os.Args)
}

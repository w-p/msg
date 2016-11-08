
# msg - Send a message to Slack
A simple command-line tool to post a message to an inbound webhook url.

`msg` provides a simple mechanism for sending notifications to a webhooked Slack channel. Use this to provide visibility into your automation pipeline. `msg` is written in go so you only need the binary.

## Usage
Create an 'Incoming WebHook' in Slack. Copy the 'Webhook URL'. Call `msg` like so:

```
msg -e https://hooks.slack.com/services/T000/B000/xxxx \
    -t 'Cluster Notification' \
    -m 'My notification message.' \
    -l info
```

## Installation
```
sudo wget -qO /usr/local/bin/msg https://raw.githubusercontent.com/w-p/msg/master/msg
sudo chmod +x /usr/local/bin/msg
```

## License
MIT

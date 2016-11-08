
# msg - Send a message to Slack
A simple command-line tool to post a message to Slack.

`msg` provides a simple mechanism for sending notifications to a webhooked Slack channel. Use this to provide visibility into your automation pipeline. `msg` is written in go so you only need the binary.

## Usage
Create an 'Incoming WebHook' in Slack.
Copy the 'Webhook URL'.
Call `msg` like so:
```
msg -e https://hooks.slack.com/services/T000/B000/xxxx \
    -t 'Cluster Notification' \
    -m 'My notification message.' \
    -l info
```
In Slack, you should see something like:
![alt text](readme-screenshot.png "Description goes here")

## Installation
Install the binary directly:
```
sudo wget -qO /usr/local/bin/msg https://raw.githubusercontent.com/w-p/msg/master/bin/msg
sudo chmod +x /usr/local/bin/msg
```
Or build it:
```
git clone https://github.com/w-p/msg.git
cd msg
make && make install
```

## License
MIT

# Kubebot

Kubebot is a Kubernetes chatbot for Slack. 

This project is in active development and it's __not ready__ for production yet.

## Setup
To run Kubebot on Slack, first you need to [create a new bot](https://my.slack.com/services/new/bot) user integration on Slack and get the `token`.

Then you need to know the channel ids where you want to run the Kubebot. You can get them on `https://slack.com/api/channels.list?token={REPLACE WITH YOUR TOKEN}`

## How to run it

### Using Kubernetes charts

The fastest way to run Kubebot in your Kubernetes cluster is using the [Kubebot chart](https://github.com/harbur/kubebot-chart) for Kubernetes.


### Running the binary

It is possible to run the binary locally or in a server. First you need to download and compile this project using the Go compiler:

```
mkdir -p $GOPATH/github.com/harbur/kubebot
cd $_
git clone git@github.com:harbur/kubebot.git
go install -v github.com/harbur/kubebot
```


Then set up the following environment variables:

```
# use the token you generated in the setup
KUBEBOT_SLACK_TOKEN="replacewithyourtoken" 

# use as many channels ids you want; use a space as a separator
KUBEBOT_SLACK_CHANNELS_IDS="1234 4321" 

# use as many admin nicknames as you want; use a space as separator
KUBEBOT_SLACK_ADMINS_NICKNAMES="nickname1 nickname2" 

# set which kubectl commands the admins will be able to run
KUBEBOT_SLACK_VALID_COMMANDS="get describe logs explain"
```


After the setup, you can run the binary:

```
kubebot
```
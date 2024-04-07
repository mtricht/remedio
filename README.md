# remedio
[![Go Report Card](https://goreportcard.com/badge/github.com/mtricht/remedio)](https://goreportcard.com/report/github.com/mtricht/remedio)

A simple self-hosted medication helper/reminder.

Features:
- Keep track of current supply
- Get reminded through a notification (ntfy, Gotify, etc.)

Screenshot:

![screenshot](screenshot.png)

## Run as docker
TODO

## Docker compose
TODO

## Configuration
| Configuration name                                                                                         | Required | Description                                                                                                  |
| ---------------------------------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------ |
| `REMEDIO_NOTIFICATION_URL`                                                                                 | Yes      | A [shoutrrr](https://containrrr.dev/shoutrrr/v0.8/services/overview/) notification URL to send reminders to.<br>Example: `ntfy://username:password@ntfy.sh/remedio?actions=view, Open Remedio, https://remedio.my.website` |
| `REMEDIO_PORT`                                                                                             | No       | The port to listen on for HTTP                                                                               |
| `REMEDIO_USERNAME`                                                                                         | No       | Username to use for basic authentication. When omitted, no authentication is required.                       |
| `REMEDIO_PASSWORD`                                                                                         | No       | Password to use for basic authentication. When omitted, no authentication is required.                       |

## TODO

- [ ] Github actions
- [ ] Support multiple users
- [ ] Use a different authentication mechanism
- [ ] Multiple times a day / every other day / weekly

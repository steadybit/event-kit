# Event API

This document explains the event API, control flow and the contracts behind it. It starts with a high-level overview and then explains every API in detail.

## Overview

Event Extension are implemented with the help of EventKit and the event API through an implementation of an extension. Extensions are HTTP servers
implementing the event API to describe which events are supported and how to consume these. The following diagram illustrates who is issuing calls and in what
phases.

![UML sequence diagram showing in what order the APIs are called](img/event-flow.svg)

As can be seen above, the extension is called by the Steadybit agent in two phases:

- In the event listeners registration phase, Steadybit learns about the supported event listeners. Once this phase is completed,
- Event listeners will be automaticlly called by Steadybit, e.g.,
  during the experiment execution.

The following sections explain the various API endpoints, their responsibilities and structures in more detail.

## Event Listeners List

As the name implies, the action list returns a list of supported event listeners. Or, more specifically, HTTP endpoints that the agent should call to learn more
about
the event listeners.

This endpoint needs to be [registered with Steadybit agents](./event-registration.md).

### Example

```json
// Request: GET /events

// Response: 200
{
  "listeners": [
    {
      "method": "GET",
      "path": "/events/mylistener"
    }
  ]
}
```

### References

- [Go API](https://github.com/steadybit/event-kit/tree/main/go/event_kit_api): `EventListResponse`
- [TypeScript API](https://github.com/steadybit/event-kit/tree/main/typescript/event_kit_api): `EventListResponse`
- [OpenAPI Schema](https://github.com/steadybit/event-kit/tree/main/openapi): `EventListResponse`

## Event Listener Description

An event listener description is required for each event listener. The HTTP endpoint serving the description is discovered through the event listener list
endpoint.

Event listener descriptions expose information about the presentation, configuration and behavior of the listener. For example:

- How should the event listener be called?

### Example

```json
// Request: GET /events/mylistener

// Response: 200
{
  "id": "com.steadybit.example.events.listener.datadog.publisher",
  "label": "Steadybit Event Listener which forwards Events to Datadog",
  "description": "Execute a rollout restart for a Kubernetes deployment",
  "version": "1.0.0",
  "icon": "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4NCjwhLS0gR2VuZXJhdG9yOiBBZG9iZSBJbGx1c3RyYXRvciAxNi4wLjQsIFNWRyBFeHBvcnQgUGx1Zy1JbiAuIFNWRyBWZXJzaW9uOiA2LjAwIEJ1aWxkIDApICAtLT4NCjwhRE9DVFlQRSBzdmcgUFVCTElDICItLy9XM0MvL0RURCBTVkcgMS4xLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL0dyYXBoaWNzL1NWRy8xLjEvRFREL3N2ZzExLmR0ZCI+DQo8c3ZnIHZlcnNpb249IjEuMSIgaWQ9IkxheWVyXzEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHg9IjBweCIgeT0iMHB4Ig0KCSB3aWR0aD0iNjEycHgiIGhlaWdodD0iNTAyLjE3NHB4IiB2aWV3Qm94PSIwIDY1LjMyNiA2MTIgNTAyLjE3NCIgZW5hYmxlLWJhY2tncm91bmQ9Im5ldyAwIDY1LjMyNiA2MTIgNTAyLjE3NCINCgkgeG1sOnNwYWNlPSJwcmVzZXJ2ZSI+DQo8ZWxsaXBzZSBmaWxsPSIjQzZDNkM2IiBjeD0iMjgzLjUiIGN5PSI0ODcuNSIgcng9IjI1OSIgcnk9IjgwIi8+DQo8cGF0aCBpZD0iYmlyZCIgZD0iTTIxMC4zMzMsNjUuMzMxQzEwNC4zNjcsNjYuMTA1LTEyLjM0OSwxNTAuNjM3LDEuMDU2LDI3Ni40NDljNC4zMDMsNDAuMzkzLDE4LjUzMyw2My43MDQsNTIuMTcxLDc5LjAzDQoJYzM2LjMwNywxNi41NDQsNTcuMDIyLDU0LjU1Niw1MC40MDYsMTEyLjk1NGMtOS45MzUsNC44OC0xNy40MDUsMTEuMDMxLTE5LjEzMiwyMC4wMTVjNy41MzEtMC4xNywxNC45NDMtMC4zMTIsMjIuNTksNC4zNDENCgljMjAuMzMzLDEyLjM3NSwzMS4yOTYsMjcuMzYzLDQyLjk3OSw1MS43MmMxLjcxNCwzLjU3Miw4LjE5MiwyLjg0OSw4LjMxMi0zLjA3OGMwLjE3LTguNDY3LTEuODU2LTE3LjQ1NC01LjIyNi0yNi45MzMNCgljLTIuOTU1LTguMzEzLDMuMDU5LTcuOTg1LDYuOTE3LTYuMTA2YzYuMzk5LDMuMTE1LDE2LjMzNCw5LjQzLDMwLjM5LDEzLjA5OGM1LjM5MiwxLjQwNyw1Ljk5NS0zLjg3Nyw1LjIyNC02Ljk5MQ0KCWMtMS44NjQtNy41MjItMTEuMDA5LTEwLjg2Mi0yNC41MTktMTkuMjI5Yy00LjgyLTIuOTg0LTAuOTI3LTkuNzM2LDUuMTY4LTguMzUxbDIwLjIzNCwyLjQxNWMzLjM1OSwwLjc2Myw0LjU1NS02LjExNCwwLjg4Mi03Ljg3NQ0KCWMtMTQuMTk4LTYuODA0LTI4Ljg5Ny0xMC4wOTgtNTMuODY0LTcuNzk5Yy0xMS42MTctMjkuMjY1LTI5LjgxMS02MS42MTctMTUuNjc0LTgxLjY4MWMxMi42MzktMTcuOTM4LDMxLjIxNi0yMC43NCwzOS4xNDcsNDMuNDg5DQoJYy01LjAwMiwzLjEwNy0xMS4yMTUsNS4wMzEtMTEuMzMyLDEzLjAyNGM3LjIwMS0yLjg0NSwxMS4yMDctMS4zOTksMTQuNzkxLDBjMTcuOTEyLDYuOTk4LDM1LjQ2MiwyMS44MjYsNTIuOTgyLDM3LjMwOQ0KCWMzLjczOSwzLjMwMyw4LjQxMy0xLjcxOCw2Ljk5MS02LjAzNGMtMi4xMzgtNi40OTQtOC4wNTMtMTAuNjU5LTE0Ljc5MS0yMC4wMTZjLTMuMjM5LTQuNDk1LDUuMDMtNy4wNDUsMTAuODg2LTYuODc2DQoJYzEzLjg0OSwwLjM5NiwyMi44ODYsOC4yNjgsMzUuMTc3LDExLjIxOGM0LjQ4MywxLjA3Niw5Ljc0MS0xLjk2NCw2LjkxNy02LjkxN2MtMy40NzItNi4wODUtMTMuMDE1LTkuMTI0LTE5LjE4LTEzLjQxMw0KCWMtNC4zNTctMy4wMjktMy4wMjUtNy4xMzIsMi42OTctNi42MDJjMy45MDUsMC4zNjEsOC40NzgsMi4yNzEsMTMuOTA4LDEuNzY3YzkuOTQ2LTAuOTI1LDcuNzE3LTcuMTY5LTAuODgzLTkuNTY2DQoJYy0xOS4wMzYtNS4zMDQtMzkuODkxLTYuMzExLTYxLjY2NS01LjIyNWMtNDMuODM3LTguMzU4LTMxLjU1NC04NC44ODcsMC05MC4zNjNjMjkuNTcxLTUuMTMyLDYyLjk2Ni0xMy4zMzksOTkuOTI4LTMyLjE1Ng0KCWMzMi42NjgtNS40MjksNjQuODM1LTEyLjQ0Niw5Mi45MzktMzMuODVjNDguMTA2LTE0LjQ2OSwxMTEuOTAzLDE2LjExMywyMDQuMjQxLDE0OS42OTVjMy45MjYsNS42ODEsMTUuODE5LDkuOTQsOS41MjQtNi4zNTENCgljLTE1Ljg5My00MS4xMjUtNjguMTc2LTkzLjMyOC05Mi4xMy0xMzIuMDg1Yy0yNC41ODEtMzkuNzc0LTE0LjM0LTYxLjI0My0zOS45NTctOTEuMjQ3DQoJYy0yMS4zMjYtMjQuOTc4LTQ3LjUwMi0yNS44MDMtNzcuMzM5LTE3LjM2NWMtMjMuNDYxLDYuNjM0LTM5LjIzNC03LjExNy01Mi45OC0zMS4yNzNDMzE4LjQyLDg3LjUyNSwyNjUuODM4LDY0LjkyNywyMTAuMzMzLDY1LjMzMQ0KCXogTTQ0NS43MzEsMjAzLjAxYzYuMTIsMCwxMS4xMTIsNC45MTksMTEuMTEyLDExLjAzOGMwLDYuMTE5LTQuOTk0LDExLjExMS0xMS4xMTIsMTEuMTExcy0xMS4wMzgtNC45OTQtMTEuMDM4LTExLjExMQ0KCUM0MzQuNjkzLDIwNy45MjksNDM5LjYxMywyMDMuMDEsNDQ1LjczMSwyMDMuMDF6Ii8+DQo8L3N2Zz4NCg==",
  "listenTo": [
    "experiment.started",
    "experiment.ended"
  ],
  "includeEnvironmentVariables": [
    "env",
    "endpoint",
    "datadog_tag"
  ],
  "listen": {
    "method": "POST",
    "path": "/events/mylistener/consume"
  }
}
```

### References

- [Go API](https://github.com/steadybit/event-kit/tree/main/go/event_kit_api): `DescribeActionResponse`
- [TypeScript API](https://github.com/steadybit/event-kit/tree/main/typescript/event_kit_api): `DescribeActionResponse`
- [OpenAPI Schema](https://github.com/steadybit/event-kit/tree/main/openapi): `DescribeActionResponse`

### Versioning

Extensions are versioned strictly, and Steadybit will ignore definition changes for the same version. Remember to update the version every time you update the
action description.

## Event Consumption

Each time a Steadybit event occurs that matches the `listenTo` field in the action description, Steadybit will send a request to the `listen` endpoint. The
request will contain the event data and the filtered environment variables.

#### Example

```json
// Request: POST /events/mylistener/consume
{
  "event": {
    "name": "Experiment DEMO-0815 Rolling Restart started",
    "type": "com.steadybit.experiment.started",
    "timestamp": "2022-11-03_15:37:00",
    "payload": {
      "experiment": {
        "experimentKey": "DEMO-0815",
        "experimentName": "Rolling Restart",
        "experimentExecutionId": 123
      },
      "team": {
        "teamId": "team1",
        "teamName": "Administrator"
      },
      "environment": {
        "environmentName": "Global",
        "environmentVariables": {
          "envVar1": "envValue1",
          "envVar2": "envValue2"
        }
      }
      "user": {
        "userId": "admin"
      }
    }
  }
}

// Response: 200
{
  "ok"
}
```

#### References

- [Go API](https://github.com/steadybit/event-kit/tree/main/go/event_kit_api): `ConsumeEventRequestBody`, `ConsumeEventResponse`
- [TypeScript API](https://github.com/steadybit/event-kit/tree/main/typescript/event_kit_api): `ConsumeEventRequestBody`, `ConsumeEventResponse`
- [OpenAPI Schema](https://github.com/steadybit/event-kit/tree/main/openapi): `ConsumeEventRequestBody`, `ConsumeEventResponse`

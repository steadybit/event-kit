# Event API

This document explains the event API, control flow and the contracts behind it. It starts with a high-level overview and then explains every API in detail.

## Overview

Event extensions are implemented with the help of EventKit and the event API through an implementation of an extension. Extensions are HTTP servers
implementing the event API to describe which events are desired and how to consume these. The following diagram illustrates who is issuing calls and in what phases.

![UML sequence diagram showing in what order the APIs are called](img/event-flow.svg)

As can be seen above, the extension is called by the Steadybit agent in two phases:

- In the event listeners registration phase, Steadybit learns about the supported event listeners. Once this phase is completed,
- Event listeners will be automatically called by Steadybit, e.g., during the experiment execution.

The following sections explain the various API endpoints, their responsibilities and structures in more detail.

## Event Listeners List

As the name implies, the action list returns a list of supported event listeners. Or, more specifically, HTTP endpoints that the agent should call to learn more
about
the event listeners.

This endpoint needs to be [registered with Steadybit agents](./event-registration.md).

### Example

```json
// Request: GET /

// Response: 200
{
  "eventListeners": [
    {
      "method": "POST",
      "path": "/events/all",
      "listenTo": ["*"],
      "restrictTo": "leader"
    }
  ]
}
```

### listenTo

Defines the event names that you want to subscribe to. The special case `["*"]` means subscribe to all. At the moment we support:

 - `*`
 - `user.invited`
 - `killswitch.engaged`
 - `killswitch.disengaged`
 - `experiment.execution.created`
 - `experiment.execution.failed`
 - `experiment.execution.errored`
 - `experiment.execution.canceled`
 - `experiment.execution.completed`
 - `experiment.execution.step-started`
 - `experiment.execution.step-errored`
 - `experiment.execution.step-canceled`
 - `experiment.execution.step-skipped`
 - `experiment.execution.step-failed`
 - `experiment.execution.step-completed`

### restrictTo

Defines which agents should call this event listener (defaults to `leader`). This is useful to simplify extension and agent deployment, i.e., you can configure the extension on all agents in a cluster while ensuring that you only get a single call.

### References

- [Go API](https://github.com/steadybit/event-kit/tree/main/go/event_kit_api): `EventListenerList`
- [OpenAPI Schema](https://github.com/steadybit/event-kit/tree/main/openapi): `EventListenerList`


## Event Consumption

Each time a Steadybit event occurs that matches the `listenTo` and `restrictTo` configuration, Steadybit will send a request to the endpoint. The request will contain the event data.

Please note that event data is not provided in real time. Outpost agents are polling the platform for new events within the configured registration interval (default: 5 seconds). The amount of returned events per registration interval is limited to 50, following events will be submitted within the next intervals.

#### Example

```json
// Request: POST /events/all
{
  "id": "da059724-a8ae-4b4b-b4f0-ee01898232d2",
  "eventName": "experiment.execution.created",
  "eventTime": "2021-09-01T12:00:00Z",
  "tenant": {
    "key": "exmpl",
    "name": "Example Inc."
  },
  "principal": {
    "principalType": "user",
    "username": "tom.mason",
    "name": "Tom Mason",
    "email": "tom.mason@example.com"
  },
  "environment": {
    "id": "STG",
    "name": "Staging"
  },
  "team": {
    "key": "ADM",
    "name": "Administrators"
  },
  "experimentExecution": {
    "experimentKey": "ADM-4",
    "executionId": 34,
    "name": "Rollout restart does not impact service availability",
    "state": "COMPLETED",
    "preparedTime": "2022-11-08T16:42:32.303762Z",
    "startedTime": "2022-11-08T16:42:32.329718Z",
    "endedTime": "2022-11-08T16:42:42.636157Z"
  }
}

// Response: 200
```

### References

- [Go API](https://github.com/steadybit/event-kit/tree/main/go/event_kit_api): `EventRequestBody`
- [OpenAPI Schema](https://github.com/steadybit/event-kit/tree/main/openapi): `EventRequestBody`

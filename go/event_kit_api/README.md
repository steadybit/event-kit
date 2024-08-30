# EventKit Go API

This module exposes Go types that you will find helpful when implementing an EventKit extension.

The types are generated automatically from the EventKit [OpenAPI specification](https://github.com/steadybit/event-kit/tree/main/openapi).

## Installation

Add the following to your `go.mod` file:

```
go get github.com/steadybit/event-kit/go/event_kit_api
```

## Usage

```go
import (
	"github.com/steadybit/event-kit/go/event_kit_api"
    "github.com/steadybit/extension-kit/extutil"
)

eventList := event_kit_api.EventListenerList{
    EventListeners: []event_kit_api.EventListener{
        {
            Method:     "POST",
            Path:       "/events/experiment-started",
            ListenTo:   []string{"experiment.execution.created"},
        },
        {
            Method:     "POST",
            Path:       "/events/experiment-completed",
            ListenTo:   []string{"experiment.execution.completed"},
        },
    },
}
```
// Package event_kit_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.1-0.20220629212257-2cf7fcf5b26d DO NOT EDIT.
package event_kit_api

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Defines values for AccessTokenType.
const (
	AccessTokenTypeAdmin AccessTokenType = "admin"
	AccessTokenTypeTeam  AccessTokenType = "team"
)

// Defines values for DescribingEndpointReferenceMethod.
const (
	Get DescribingEndpointReferenceMethod = "get"
)

// Defines values for EventListenerRestrictTo.
const (
	Any    EventListenerRestrictTo = "any"
	Aws    EventListenerRestrictTo = "aws"
	Leader EventListenerRestrictTo = "leader"
)

// Defines values for ExperimentExecutionState.
const (
	Canceled  ExperimentExecutionState = "canceled"
	Completed ExperimentExecutionState = "completed"
	Created   ExperimentExecutionState = "created"
	Errored   ExperimentExecutionState = "errored"
	Failed    ExperimentExecutionState = "failed"
	Prepared  ExperimentExecutionState = "prepared"
	Running   ExperimentExecutionState = "running"
	Skipped   ExperimentExecutionState = "skipped"
)

// Defines values for MutatingHttpMethod.
const (
	Delete MutatingHttpMethod = "delete"
	Post   MutatingHttpMethod = "post"
	Put    MutatingHttpMethod = "put"
)

// Defines values for PrincipalType.
const (
	AccessToken PrincipalType = "access_token"
	BatchJob    PrincipalType = "batch_job"
	User        PrincipalType = "user"
)

// AccessTokenPrincipal defines model for AccessTokenPrincipal.
type AccessTokenPrincipal struct {
	Id string `json:"id"`

	// A human-readable name for the user.
	Name          string          `json:"name"`
	PrincipalType string          `json:"principalType"`
	TokenType     AccessTokenType `json:"tokenType"`
}

// AccessTokenType defines model for AccessTokenType.
type AccessTokenType string

// BatchPrincipal defines model for BatchPrincipal.
type BatchPrincipal struct {
	PrincipalType string `json:"principalType"`

	// This is a unique identifier for the user.
	Username string `json:"username"`
}

// HTTP endpoint which the Steadybit platform/agent could communicate with.
type DescribingEndpointReference struct {
	// HTTP method to use when calling the HTTP endpoint.
	Method DescribingEndpointReferenceMethod `json:"method"`

	// Absolute path of the HTTP endpoint.
	Path string `json:"path"`
}

// HTTP method to use when calling the HTTP endpoint.
type DescribingEndpointReferenceMethod string

// Environment defines model for Environment.
type Environment struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// EventListener defines model for EventListener.
type EventListener struct {
	// By default, event listeners do not receive any environment variables. To receive environment variables you need to explicitly list the environment variables keys here.
	IncludeEnvironmentVariables *[]string `json:"includeEnvironmentVariables,omitempty"`

	// List of event names that the event listener want to listen to. You may optionally define the special `*` event name to listen to all events.
	ListenTo []string           `json:"listenTo"`
	Method   MutatingHttpMethod `json:"method"`

	// Absolute path of the HTTP endpoint.
	Path string `json:"path"`

	// If the agent is deployed as a daemonset in Kubernetes, should the discovery only be called from the leader agent? This can be helpful to avoid duplicate event listener calls every running agent. You may alternatively define that the listener should run only for Steadybit agents operating within the AWS agent mode (as defined by the `STEADYBIT_AGENT_MODE` Steadybit agent environment variable / the `agent.mode` Steadybit agent Helm chart value).
	RestrictTo *EventListenerRestrictTo `json:"restrictTo,omitempty"`
}

// If the agent is deployed as a daemonset in Kubernetes, should the discovery only be called from the leader agent? This can be helpful to avoid duplicate event listener calls every running agent. You may alternatively define that the listener should run only for Steadybit agents operating within the AWS agent mode (as defined by the `STEADYBIT_AGENT_MODE` Steadybit agent environment variable / the `agent.mode` Steadybit agent Helm chart value).
type EventListenerRestrictTo string

// Lists all listeners that the platform/agent could call.
type EventListenerList struct {
	EventListeners []EventListener `json:"eventListeners"`
}

// ExperimentExecution defines model for ExperimentExecution.
type ExperimentExecution struct {
	EndedTime            *time.Time               `json:"endedTime,omitempty"`
	ExecutionId          string                   `json:"executionId"`
	ExperimentKey        string                   `json:"experimentKey"`
	FailureReason        *string                  `json:"failureReason,omitempty"`
	FailureReasonDetails *string                  `json:"failureReasonDetails,omitempty"`
	Hypothesis           string                   `json:"hypothesis"`
	Name                 string                   `json:"name"`
	PreparedTime         time.Time                `json:"preparedTime"`
	StartedTime          time.Time                `json:"startedTime"`
	State                ExperimentExecutionState `json:"state"`
	Variables            map[string]string        `json:"variables"`
}

// ExperimentExecutionState defines model for ExperimentExecution.State.
type ExperimentExecutionState string

// ListenResult defines model for ListenResult.
type ListenResult = map[string]interface{}

// MutatingHttpMethod defines model for MutatingHttpMethod.
type MutatingHttpMethod string

// PrincipalType defines model for PrincipalType.
type PrincipalType string

// Team defines model for Team.
type Team struct {
	Id   string `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

// Tenant defines model for Tenant.
type Tenant struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

// UserPrincipal defines model for UserPrincipal.
type UserPrincipal struct {
	Email *string `json:"email,omitempty"`

	// A human-readable name for the user.
	Name          string `json:"name"`
	PrincipalType string `json:"principalType"`

	// This is a unique identifier for the user.
	Username string `json:"username"`
}

// EventListenerListResponse defines model for EventListenerListResponse.
type EventListenerListResponse struct {
	union json.RawMessage
}

// EventRequestBody defines model for EventRequestBody.
type EventRequestBody struct {
	Environment         *Environment         `json:"environment,omitempty"`
	EventName           string               `json:"eventName"`
	EventTime           time.Time            `json:"eventTime"`
	ExperimentExecution *ExperimentExecution `json:"experimentExecution,omitempty"`
	Id                  uuid.UUID            `json:"id"`

	// The principal describes through which activity the action was triggered.
	Principal Principal `json:"principal"`
	Team      *Team     `json:"team,omitempty"`
	Tenant    Tenant    `json:"tenant"`
}

func (t EventListenerListResponse) AsEventListenerList() (EventListenerList, error) {
	var body EventListenerList
	err := json.Unmarshal(t.union, &body)
	return body, err
}

func (t *EventListenerListResponse) FromEventListenerList(v EventListenerList) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

func (t EventListenerListResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *EventListenerListResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
type Principal interface {}
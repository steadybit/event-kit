// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2024 Steadybit GmbH

// Package event_kit_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package event_kit_api

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime"
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

// Defines values for ExperimentExecutionState.
const (
	ExperimentExecutionStateCanceled  ExperimentExecutionState = "canceled"
	ExperimentExecutionStateCompleted ExperimentExecutionState = "completed"
	ExperimentExecutionStateCreated   ExperimentExecutionState = "created"
	ExperimentExecutionStateErrored   ExperimentExecutionState = "errored"
	ExperimentExecutionStateFailed    ExperimentExecutionState = "failed"
	ExperimentExecutionStatePrepared  ExperimentExecutionState = "prepared"
	ExperimentExecutionStateRunning   ExperimentExecutionState = "running"
	ExperimentExecutionStateSkipped   ExperimentExecutionState = "skipped"
)

// Defines values for ExperimentStepExecutionActionKind.
const (
	Attack   ExperimentStepExecutionActionKind = "attack"
	Check    ExperimentStepExecutionActionKind = "check"
	LoadTest ExperimentStepExecutionActionKind = "load_test"
	Other    ExperimentStepExecutionActionKind = "other"
)

// Defines values for ExperimentStepExecutionState.
const (
	ExperimentStepExecutionStateCanceled  ExperimentStepExecutionState = "canceled"
	ExperimentStepExecutionStateCompleted ExperimentStepExecutionState = "completed"
	ExperimentStepExecutionStateCreated   ExperimentStepExecutionState = "created"
	ExperimentStepExecutionStateErrored   ExperimentStepExecutionState = "errored"
	ExperimentStepExecutionStateFailed    ExperimentStepExecutionState = "failed"
	ExperimentStepExecutionStatePrepared  ExperimentStepExecutionState = "prepared"
	ExperimentStepExecutionStateRunning   ExperimentStepExecutionState = "running"
	ExperimentStepExecutionStateSkipped   ExperimentStepExecutionState = "skipped"
)

// Defines values for ExperimentStepExecutionType.
const (
	Action ExperimentStepExecutionType = "action"
	Wait   ExperimentStepExecutionType = "wait"
)

// Defines values for ExperimentStepTargetExecutionState.
const (
	Canceled  ExperimentStepTargetExecutionState = "canceled"
	Completed ExperimentStepTargetExecutionState = "completed"
	Created   ExperimentStepTargetExecutionState = "created"
	Errored   ExperimentStepTargetExecutionState = "errored"
	Failed    ExperimentStepTargetExecutionState = "failed"
	Prepared  ExperimentStepTargetExecutionState = "prepared"
	Running   ExperimentStepTargetExecutionState = "running"
	Skipped   ExperimentStepTargetExecutionState = "skipped"
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

	// Name A human-readable name for the user.
	Name          string          `json:"name"`
	PrincipalType string          `json:"principalType"`
	TokenType     AccessTokenType `json:"tokenType"`
}

// AccessTokenType defines model for AccessTokenType.
type AccessTokenType string

// BatchPrincipal defines model for BatchPrincipal.
type BatchPrincipal struct {
	PrincipalType string `json:"principalType"`

	// Username This is a unique identifier for the user.
	Username string `json:"username"`
}

// DescribingEndpointReference HTTP endpoint which the Steadybit platform/agent could communicate with.
type DescribingEndpointReference struct {
	// Method HTTP method to use when calling the HTTP endpoint.
	Method DescribingEndpointReferenceMethod `json:"method"`

	// Path Absolute path of the HTTP endpoint.
	Path string `json:"path"`
}

// DescribingEndpointReferenceMethod HTTP method to use when calling the HTTP endpoint.
type DescribingEndpointReferenceMethod string

// Environment defines model for Environment.
type Environment struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// EventListener defines model for EventListener.
type EventListener struct {
	// ListenTo List of event names that the event listener want to listen to. You may optionally define the special `*` event name to listen to all events.
	ListenTo []string           `json:"listenTo"`
	Method   MutatingHttpMethod `json:"method"`

	// Path Absolute path of the HTTP endpoint.
	Path string `json:"path"`
}

// EventListenerList Lists all listeners that the platform/agent could call.
type EventListenerList struct {
	EventListeners []EventListener `json:"eventListeners"`
}

// ExperimentExecution defines model for ExperimentExecution.
type ExperimentExecution struct {
	EndedTime     *time.Time `json:"endedTime,omitempty"`
	ExecutionId   float32    `json:"executionId"`
	ExperimentKey string     `json:"experimentKey"`
	Hypothesis    string     `json:"hypothesis"`
	Name          string     `json:"name"`
	PreparedTime  time.Time  `json:"preparedTime"`

	// Reason Additional information about the reason for the state.
	Reason *string `json:"reason,omitempty"`

	// ReasonDetails Additional information about the reason for the state.
	ReasonDetails *string                  `json:"reasonDetails,omitempty"`
	StartedTime   time.Time                `json:"startedTime"`
	State         ExperimentExecutionState `json:"state"`
}

// ExperimentExecutionState defines model for ExperimentExecution.State.
type ExperimentExecutionState string

// ExperimentStepExecution defines model for ExperimentStepExecution.
type ExperimentStepExecution struct {
	ActionId      *string                            `json:"actionId,omitempty"`
	ActionKind    *ExperimentStepExecutionActionKind `json:"actionKind,omitempty"`
	ActionName    *string                            `json:"actionName,omitempty"`
	CustomLabel   *string                            `json:"customLabel,omitempty"`
	EndedTime     *time.Time                         `json:"endedTime,omitempty"`
	ExecutionId   float32                            `json:"executionId"`
	ExperimentKey string                             `json:"experimentKey"`
	Id            uuid.UUID                          `json:"id"`
	StartedTime   *time.Time                         `json:"startedTime,omitempty"`
	State         ExperimentStepExecutionState       `json:"state"`
	Type          ExperimentStepExecutionType        `json:"type"`
}

// ExperimentStepExecutionActionKind defines model for ExperimentStepExecution.ActionKind.
type ExperimentStepExecutionActionKind string

// ExperimentStepExecutionState defines model for ExperimentStepExecution.State.
type ExperimentStepExecutionState string

// ExperimentStepExecutionType defines model for ExperimentStepExecution.Type.
type ExperimentStepExecutionType string

// ExperimentStepTargetExecution defines model for ExperimentStepTargetExecution.
type ExperimentStepTargetExecution struct {
	AgentHostname string     `json:"agentHostname"`
	EndedTime     *time.Time `json:"endedTime,omitempty"`
	ExecutionId   float32    `json:"executionId"`
	ExperimentKey string     `json:"experimentKey"`
	Id            uuid.UUID  `json:"id"`

	// Reason Additional information about the reason for the state.
	Reason *string `json:"reason,omitempty"`

	// ReasonDetails Additional information about the reason for the state.
	ReasonDetails    *string                            `json:"reasonDetails,omitempty"`
	StartedTime      *time.Time                         `json:"startedTime,omitempty"`
	State            ExperimentStepTargetExecutionState `json:"state"`
	StepExecutionId  uuid.UUID                          `json:"stepExecutionId"`
	TargetAttributes map[string][]string                `json:"targetAttributes"`
	TargetName       string                             `json:"targetName"`
	TargetType       string                             `json:"targetType"`
}

// ExperimentStepTargetExecutionState defines model for ExperimentStepTargetExecution.State.
type ExperimentStepTargetExecutionState string

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

	// Name A human-readable name for the user.
	Name          string `json:"name"`
	PrincipalType string `json:"principalType"`

	// Username This is a unique identifier for the user.
	Username string `json:"username"`
}

// EventListenerListResponse defines model for EventListenerListResponse.
type EventListenerListResponse struct {
	union json.RawMessage
}

// EventRequestBody defines model for EventRequestBody.
type EventRequestBody struct {
	Environment                   *Environment                   `json:"environment,omitempty"`
	EventName                     string                         `json:"eventName"`
	EventTime                     time.Time                      `json:"eventTime"`
	ExperimentExecution           *ExperimentExecution           `json:"experimentExecution,omitempty"`
	ExperimentStepExecution       *ExperimentStepExecution       `json:"experimentStepExecution,omitempty"`
	ExperimentStepTargetExecution *ExperimentStepTargetExecution `json:"experimentStepTargetExecution,omitempty"`
	Id                            uuid.UUID                      `json:"id"`

	// Principal The principal describes through which activity the action was triggered.
	Principal Principal `json:"principal"`
	Team      *Team     `json:"team,omitempty"`
	Tenant    Tenant    `json:"tenant"`
}

// AsEventListenerList returns the union data inside the EventListenerListResponse as a EventListenerList
func (t EventListenerListResponse) AsEventListenerList() (EventListenerList, error) {
	var body EventListenerList
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventListenerList overwrites any union data inside the EventListenerListResponse as the provided EventListenerList
func (t *EventListenerListResponse) FromEventListenerList(v EventListenerList) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventListenerList performs a merge with any union data inside the EventListenerListResponse, using the provided EventListenerList
func (t *EventListenerListResponse) MergeEventListenerList(v EventListenerList) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
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

type Principal interface{}

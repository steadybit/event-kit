// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2024 Steadybit GmbH

package event_kit_api

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

// Note: These test cases only check that the code compiles as intended.
// On compilation errors, we most likely have caused a breaking change of
// the API and need to act accordingly.

func markAsUsed(t *testing.T, v any) {
	if v == nil {
		t.Fail()
	}
}

func TestEventListenerRequestBody(t *testing.T) {
	v := EventRequestBody{
		Environment: new(Environment{
			Id:   "test",
			Name: "gateway",
		}),
		EventName: "experiment.step.started",
		EventTime: time.Now(),
		ExperimentExecution: new(ExperimentExecution{
			EndedTime:     new(time.Now()),
			ExecutionId:   42,
			ExperimentKey: "ExperimentKey",
			Reason:        new("Reason"),
			ReasonDetails: new("ReasonDetails"),
			Hypothesis:    "Hypothesis",
			Name:          "Name",
			PreparedTime:  time.Time{},
			StartedTime:   time.Time{},
			State:         ExperimentExecutionStateCreated,
		}),
		ExperimentStepExecution: new(ExperimentStepExecution{
			ExecutionId:   42,
			ExperimentKey: "ExperimentKey",
			Type:          Action,
			ActionId:      new("action-id"),
			Id:            uuid.New(),
			EndedTime:     new(time.Now()),
			StartedTime:   new(time.Now()),
			State:         ExperimentStepExecutionStateCreated,
		}),
		ExperimentStepTargetExecution: new(ExperimentStepTargetExecution{
			ExecutionId:     42,
			ExperimentKey:   "ExperimentKey",
			Id:              uuid.New(),
			StepExecutionId: uuid.New(),
			AgentHostname:   "agent-123",
			TargetAttributes: map[string][]string{
				"target.name":   {"test-target"},
				"target.colors": {"red", "green", "blue"},
			},
			EndedTime:     new(time.Now()),
			StartedTime:   new(time.Now()),
			State:         "state",
			TargetName:    "test-target",
			TargetType:    "target-type-test",
			Reason:        new("Reason"),
			ReasonDetails: new("ReasonDetails"),
		}),
		Id: uuid.New(),
		Principal: UserPrincipal{
			Email:         new("email"),
			Name:          "Peter",
			Username:      "Pan",
			PrincipalType: string(User),
		},
		Team: new(Team{
			Id:   "test",
			Key:  "test",
			Name: "gateway",
		}),
		Tenant: Tenant{
			Key:  "key",
			Name: "name",
		},
	}
	markAsUsed(t, v)
}

func TestBatchPrincipal(t *testing.T) {
	v := BatchPrincipal{
		PrincipalType: string(BatchJob),
		Username:      "foo",
	}
	markAsUsed(t, v)
}

func TestAccessTokenPrincipal(t *testing.T) {
	v := AccessTokenPrincipal{
		Id:            "42",
		Name:          "Admin Token",
		PrincipalType: string(AccessToken),
		TokenType:     AccessTokenTypeAdmin,
	}
	markAsUsed(t, v)
}

func TestActionList(t *testing.T) {
	v := EventListenerList{
		EventListeners: []EventListener{},
	}
	markAsUsed(t, v)
}

//go:fix inline
func Ptr[T any](val T) *T {
	return new(val)
}

func TestEventListenerDescription(t *testing.T) {
	v := EventListener{
		Method:   Post,
		Path:     "/events/mylistener/listen",
		ListenTo: []string{"*"},
	}
	markAsUsed(t, v)
}

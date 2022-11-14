// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

import { EventListenerList, EventRequestBody } from './index';

export const eventListenerList: EventListenerList = {
	eventListeners: [
		{
			method: 'post',
			path: '/events/mylistener',
			restrictTo: 'aws',
			listenTo: ['my-event-name'],
		}
	]
};

export const eventListenerRequestBody: EventRequestBody = {
	environment: {
		id: 'my-environment-id',
		name: 'my-environment-name'
	},
	team: {
		id: 'my-team-id',
		key: 'my-team-key',
		name: 'my-team-name'
	},
	tenant: {
		key: 'my-tenant-key',
		name: 'my-tenant-name'
	},
	experimentExecution: {
		endedTime: '2021-09-01T12:00:00Z',
		executionId: 42,
		experimentKey: 'my-experiment-key',
		failureReason: 'my-failure-reason',
		hypothesis: 'my-hypothesis',
		name: 'my-execution-name',
		preparedTime: '2021-09-01T12:00:00Z',
		startedTime: '2021-09-01T12:00:00Z',
		state: 'created',
		failureReasonDetails: 'my-failure-reason-details'
	},
	eventName: 'my-event-name',
	eventTime: '2021-09-01T12:00:00Z',
	id: 'my-event-id',
	principal: {
		name: 'my-principal-name',
		username: 'my-principal-username'
	}
};

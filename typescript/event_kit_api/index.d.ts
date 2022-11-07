// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

import { components } from './schemas';

export type EventListenerList = components['schemas']['EventListenerList'];
export type EventKitError = components['schemas']['EventKitError'];
export type EventListenerDescription = components['schemas']['EventListenerDescription'];
export type ListenResult = components['schemas']['ListenResult'];
export type DescribingEndpointReference = components['schemas']['DescribingEndpointReference'];
export type MutatingHttpMethod = components['schemas']['MutatingHttpMethod'];
export type MutatingEndpointReference = components['schemas']['MutatingEndpointReference'];
export type Principal = components['schemas']['Principal'];
export type Tenant = components['schemas']['Tenant'];
export type Team = components['schemas']['Team'];
export type Environment = components['schemas']['Environment'];
export type ExperimentExecution = components['schemas']['ExperimentExecution'];

export type EventListenerListResponse =
	components['responses']['EventListenerListResponse']['content']['application/json'];
export type DescribeEventListenerResponse =
	components['responses']['DescribeEventListenerResponse']['content']['application/json'];

export type EventListenerRequestBody =
	components['requestBodies']['EventListenerRequestBody']['content']['application/json'];

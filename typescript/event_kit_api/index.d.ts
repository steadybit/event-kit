// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

import { components } from './schemas';

export type EventListenerList = components['schemas']['EventListenerList'];
export type EventListener = components['schemas']['EventListener'];
export type ListenResult = components['schemas']['ListenResult'];
export type DescribingEndpointReference = components['schemas']['DescribingEndpointReference'];
export type MutatingHttpMethod = components['schemas']['MutatingHttpMethod'];
export type Principal = components['schemas']['Principal'];
export type UserPrincipal = components['schemas']['UserPrincipal'];
export type AccessTokenPrincipal = components['schemas']['AccessTokenPrincipal'];
export type BatchPrincipal = components['schemas']['BatchPrincipal'];
export type Tenant = components['schemas']['Tenant'];
export type Team = components['schemas']['Team'];
export type Environment = components['schemas']['Environment'];
export type ExperimentExecution = components['schemas']['ExperimentExecution'];

export type EventListenerListResponse =
	components['responses']['EventListenerListResponse']['content']['application/json'];

export type EventRequestBody =
	components['requestBodies']['EventRequestBody']['content']['application/json'];

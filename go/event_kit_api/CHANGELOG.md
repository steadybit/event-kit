# Changelog

## 1.3.0

- EventRequestBody will either contain `experimentExecution`, `experimentStepExecution` or `experimentStepExecutionTarget` depending on the event type

## 1.2.3

- Use uuid type for `stepExecutionId`

## 1.2.2

- Added `stepExecutionId` for events related to steps

## 1.2.1

- Added more information to `ExperimentStepExecution` events

## 1.2.0

- Added Step-Detais to execution events
- BREAKING: Enum-Constants for `state` of `ExperimentExecution` changed 

## 1.1.0

- Initial release

# How To Write An Event Extension

This how-to article will teach you how to write an extension using EventKit that adds new event listener capabilities. We will look closely at existing extensions to learn about semantic conventions, best practices, expected behavior and necessary boilerplate.

The article assumes that you have read the [overview documentation](../event-api.md#overview) for the Event API and possibly skimmed over the expected API endpoints. We are leveraging the Go programming language within the examples, but you can use every other language as long as you adhere to the expected API.

<!-- TOC -->
  * [Necessary Boilerplate](#necessary-boilerplate)
  * [Action List](#action-list)
  * [Action Description](#action-description)
  * [Action Execution](#action-execution)
    * [Prepare](#prepare)
    * [Start](#start)
    * [Status](#status)
    * [Stop](#stop)
  * [Extension Registration](#extension-registration)
<!-- TOC -->

## Necessary Boilerplate

## Event Listener List

## Event Listener Description

## Event Consumption

## Extension Registration

Congratulations, the extension is now completed! This leaves only one last step: Announcing the extension to the Steadybit agents. Read more on this topic within our separate [event registration document](../event-registration.md).
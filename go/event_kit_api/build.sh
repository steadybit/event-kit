#!/usr/bin/env bash

#
# SPDX-License-Identifier: MIT
# SPDX-FileCopyrightText: 2022 Steadybit GmbH
#

set -eo pipefail

oapi-codegen -config generator-config.yml ../../openapi/spec.yml > event_kit_api.go

cat extras.go.txt >> event_kit_api.go
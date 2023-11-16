#!/usr/bin/env bash

#
# SPDX-License-Identifier: MIT
# SPDX-FileCopyrightText: 2022 Steadybit GmbH
#

set -eo pipefail

go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.16.2
oapi-codegen -config generator-config.yml ../../openapi/spec.yml > event_kit_api.go

cat extras.go.txt >> event_kit_api.go
#!/bin/bash
curl \
--header "Authorization: Bearer 45954192-e1f8-9526-0974-32cb5b66c235" \
--request PUT \
--data @registe-service.json \
http://127.0.0.1:8500/v1/agent/service/register\?replace-existing-checks\=true

curl \
--header "Authorization: Bearer 45954192-e1f8-9526-0974-32cb5b66c235" \
--request PUT \
http://127.0.0.1:8500/v1/agent/service/deregister/user-service
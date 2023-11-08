#!/bin/bash
set -exuo pipefail
docker build -t web5-component:latest -f .web5-component/test.Dockerfile .
cd ${GITHUB_ACTION_PATH}
docker-compose up --timeout 10 --exit-code-from test-client

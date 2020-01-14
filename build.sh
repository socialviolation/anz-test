#!/bin/bash
set -e

SHA=$(git rev-parse HEAD)
# Tag container differently for local development.
# CLOUDBUILD env variable is set in cloudbuild step:build_image,
# which triggers this build script
if [[ -n "${CLOUDBUILD}" ]]; then
        CONTAINER_TAGS="-t gcr.io/${PROJECT_ID}/anz-test:${SHA} -t gcr.io/${PROJECT_ID}/anz-test:latest"
else
        CONTAINER_TAGS="-t anz-test:${SHA} -t anz-test:latest"
fi

docker build \
        --build-arg APP_VERSION=$(cat version.txt) \
        --build-arg SHA=${SHA} \
        ${CONTAINER_TAGS} .

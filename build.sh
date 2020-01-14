#!/bin/bash
set -e

SHA=$(git rev-parse HEAD)
# Tag container differently for local development.
# CLOUDBUILD env variable is set in cloudbuild step:build_image,
# which triggers this build script
if [[ -n "${CLOUDBUILD}" ]]; then
        CONTAINER_TAG="gcr.io/${PROJECT_ID}/anz-test:${SHA}"
else
        CONTAINER_TAG="anz-test"
fi

docker build \
        --build-arg APP_VERSION=$(cat version.txt) \
        --build-arg SHA=${SHA} \
        -t ${CONTAINER_TAG} .

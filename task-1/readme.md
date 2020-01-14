# ANZ-Test
ANZ-Test is a golang webserver that dynamically returns the application's name, version, commit SHA and description.

The **APP_VERSION** is determined by the file `./version.txt`, and the **COMMIT_SHA** is determined by the commit the build is based off of. Both are passed as variables to the "Build Image" step as a part of the CloudBuild pipeline, and stored as an environment variable in the image. The web server will read the environment variable when a `/version` request is recieved.

## Requirements
To run this webserver, you will need to have installed and configured:
- [`docker`](https://www.docker.com/products/docker-desktop)
- [`gcloud`](https://cloud.google.com/sdk/install)
- [`go@1.13+`](https://golang.org/dl/)

## Local

For local development, you can run `make local-ci`, which will perform a local version of the whole CI CD pipeline. The result will be a docker image tagged `anz-test:<COMMIT_SHA>`, and a container bound on port 8080

## Remote
### CI Pipeline

This webserver has been setup to use [CloudBuild](https://cloud.google.com/cloud-build/) as a target to perform the following, any time the github repository is pushed to:

1. Download third-party deps using a `GOPROXY`
2. Run unit tests
4. Run `gosec`
5. Build and tag the docker image
6. Push the Docker image to GCR
7. Deploy the latest container to a GKE cluster,
   - This will create the kubernetes namespace, deployment and service if they do not already exist
   - Currently hardcoded to a cluster named anz-test-cluster, in the australia-southeast1-b region

You can manually trigger this build if you have:
- created the GKE cluster
- configured your gcloud sdk to to use the project that is hosting the GKE cluster

Simply run `make remote-build` to manually trigger a cloud build

## Considerations

* The third party dependencies are strictly versioned, and downloaded using a GOPROXY to avoid any malicious interception, or the possibility of the package versions becoming unavailable in the future.
* Cloudbuilds can manually be triggered, and as a result can deploy code straight to production that has not been checked in or verified. This is very very naughty.
* CloudBuild doesn't provide any nice reports, and either all succeeds, or fails. Feedback can sometimes be a bit lacking.
* CloudBuild will not currently fail if the code has not been formatted or vetted. Could be added.
*

## References:
* https://github.com/GoogleContainerTools/distroless
* https://github.com/securego/gosec
* https://github.com/gorilla/mux
* https://cloud.google.com/cloud-build/docs/deploying-builds/deploy-gke
* https://cloud.google.com/cloud-build/docs/configuring-builds/build-test-deploy-artifacts#deploying_artifacts

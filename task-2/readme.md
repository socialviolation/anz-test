# ANZ-Test - Task-2

## Prerequisites

For this part of the task, you are going to need:
- `bash`
- [`gcloud`](https://cloud.google.com/sdk/install)

It is expected that you have run at least 1 successful CloudBuild from task-1, and have an image pushed to GCR with the tag `gcr.io/$PROJECT_ID/anz-test:latest`

You can manually trigger this build if you have:
- created the GKE cluster named: `anz-test-cluster`
- configured your gcloud sdk to to use the project that is hosting the GKE cluster
- granted the cloudbuild service account the kubernetes developer role

# Deploying
To deploy the kube manifests for the webserver, we are using [CloudBuild](https://cloud.google.com/cloud-build/) as a target to interface with a GKE cluster named: `anz-test-cluster`.
The manifests will create the following kube resources:
- [namespace.yaml](kube-anz-test/namespace.yaml): namespace to isolate the anz-test resources
- [service.yaml](kube-anz-test/service.yaml): A service to publicly expose the deployment
- [deployment.yaml](kube-anz-test/deployment.yaml): A deployment to monitor and deploy pods
- [pod-autoscaler.yaml](kube-anz-test/pod-autoscaler.yaml): a horizontal pod autoscaler, to scale pods horizontally

The deploy script will submit a cloudbuild job, which will apply all the kube manifests, creating and updating the kube resources as necessary

```bash
$ ./deploy.sh
```

# Considerations
* The pod will autoscale if the cpu usage exceeds 75%
* The only restrictions possible are restricting who can execute CloudBuilds.
* Cloudbuilds can manually be triggered, and as a result can deploy manifest changes straight to production that has not been checked in or verified. This is very very naughty.

## References

* https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/

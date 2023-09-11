# Adding alerts and recording rules to the operator

In this part of the tutorial we will add alerts and recording rules to the
operator. This step assumes you have an instance of Prometheus running in your
cluster and that the `PrometheusRule` CRD is installed.

## Steps

Check the [pkg/monitoring/rules](pkg/monitoring/rules) directory in this
repository for the files that need to be created in the operator project.

It contains the following files:

- `rules.go`: contains the central logic for alerts and recording rules
registration and setup.

- `operator_alerts.go`: contains the alerts definitions related to the operator
execution. This is where you should add alerts for when the operator is down,
when it is not able to reconcile a resource, etc. Each alert subcategory should
be declared in a separate file.

- `operator_recording_rules.go`: contains the recording rules related to the
operator execution. This is where you should add recording rules for the number
of operator pods, the number of ready operator pods, etc. Each recording rule
subcategory should be declared in a separate file.

Now we need to call the `SetupRules` function from the `main.go`: 

```go
...
import (
  ...
  promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
  "github.com/machadovilaca/operator-observability/pkg/monitoring/rules"
  ...
)

func init() {
	  ...
	utilruntime.Must(promv1.AddToScheme(scheme))
  ...
}

func main() {
  ...
  rules.SetupRules(mgr.GetClient())
  ...
}
```

Let's deploy the operator

```bash
# make sure the repository used in IMAGE_NAME exists and is accessible from your cluster
export IMG="example.com/observability-operator:v0.0.1"

make docker-build docker-push

kubectl create namespace observability-operator
kubectl label namespace observability-operator openshift.io/cluster-monitoring=true # Or any other label configured in the Prometheus instance
kubectl config set-context --current --namespace=observability-operator

make deploy
```

Now, you can check the Prometheus UI to see if the alerts and recording rules are working as
expected.

## Next

=> [Generating documentation for the operator](../04-documentation/README.md)

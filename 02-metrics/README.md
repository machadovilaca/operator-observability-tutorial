# Adding metrics to the operator

In this part of the tutorial we will import and set up the
operator-observability package in the operator generated in the previous part,
and add metrics to it.

## Steps

Install the operator-observability package:

```bash
go get github.com/machadovilaca/operator-observability@latest
```

Check the [pkg/monitoring/metrics](pkg/monitoring/metrics) directory in this
repository for the files that need to be created in the operator project.

It contains the following files:

- `metrics.go`: contains the central logic for metrics setup and registration.

- `operator_metrics.go`: contains the metrics definitions for metrics specific
to the operator execution. This is where you should add metrics for the
operator's internal execution, such as the number of reconciliations, the
duration of each reconciliation, etc. Each metric subcategory should be
declared in a separate file.

- `custom_resource_collector.go`: contains the definition of a collector to 
count the number of custom resources in the cluster. This is where you should
add metrics that need to fetch information from external sources, such as the
cluster API server, cloud provider APIs, etc.


Now we need to call the `SetupMetrics` function from the `main.go`: 

```go
...
import (
  ...
  "github.com/machadovilaca/operator-observability/pkg/monitoring/metrics"
  ...
)

func main() {
  ...
  metrics.SetupMetrics(mgr.GetClient())
  ...
}
```

Run locally (only for testing) and check metrics:

```bash
make install run

# In another terminal
curl localhost:8080/metrics | grep observability_operator_
```

Also update the `Dockerfile` to include the `pkg/` directory:

```Dockerfile
...
COPY pkg/ pkg/
...
```

## Next

=> [Adding alerts and recording rules to the operator](../03-alerts-recordingrules/README.md)


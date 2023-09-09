package metrics

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	runtimemetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"
)

const metricPrefix = "observability_operator_"

var metricsLog = ctrl.Log.WithName("metrics")

func SetupMetrics(client client.Client) {
	operatormetrics.Register = runtimemetrics.Registry.Register
	metricsLog.Info("registering metrics")
	err := operatormetrics.RegisterMetrics(operatorMetrics)
	if err != nil {
		panic(err)
	}

	metricsLog.Info("registering collectors")
	SetupCustomResourceCollector(client)
	err = operatormetrics.RegisterCollector(customResourceCollector)
	if err != nil {
		panic(err)
	}
}

// ListMetrics returns a list of all metrics exposed by the operator
func ListMetrics() []operatormetrics.Metric {
	return operatormetrics.ListMetrics()
}

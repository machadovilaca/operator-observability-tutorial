package metrics

import (
	"context"
	observabilityv1alpha1 "github.com/machadovilaca/operator-observability-tutorial/api/v1alpha1"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"
)

var (
	collectorK8sClient client.Client
)

func SetupCustomResourceCollector(k8sClient client.Client) {
	collectorK8sClient = k8sClient
}

var (
	customResourceCollector = operatormetrics.Collector{
		Metrics: []operatormetrics.Metric{
			crCount,
		},
		CollectCallback: customResourceCollectorCallback,
	}

	crCount = operatormetrics.NewGaugeVec(
		operatormetrics.MetricOpts{
			Name:        metricPrefix + "cr_count",
			Help:        "Number of existing guestbook custom resources",
			ConstLabels: map[string]string{"controller": "observability"},
			ExtraFields: map[string]string{
				"StabilityLevel":    "DEPRECATED",
				"DeprecatedVersion": "1.14.0",
			},
		},
		[]string{"namespace"},
	)
)

func customResourceCollectorCallback() []operatormetrics.CollectorResult {
	if collectorK8sClient == nil {
		return nil
	}

	result := observabilityv1alpha1.TestList{}
	err := collectorK8sClient.List(context.TODO(), &result, client.InNamespace("default"))
	if err != nil {
		metricsLog.Error(err, "failed to list custom resources")
		return nil
	}

	crCountValue := float64(len(result.Items))

	return []operatormetrics.CollectorResult{
		{Metric: crCount, Value: crCountValue, Labels: []string{"default"}},
	}
}

package rules

import (
	"fmt"

	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"
	"github.com/machadovilaca/operator-observability/pkg/operatorrules"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var operatorRecordingRules = []operatorrules.RecordingRule{
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name:        recordingRulesPrefix + "number_of_pods",
			Help:        "Number of observability operator pods in the cluster",
			ConstLabels: map[string]string{"controller": "observability"},
		},
		MetricType: operatormetrics.GaugeType,
		Expr:       intstr.FromString(fmt.Sprintf("sum(up{namespace='%s', pod=~'observability-operator-.*'}) or vector(0)", namespace)),
	},
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name:        recordingRulesPrefix + "number_of_ready_pods",
			Help:        "Number of ready observability operator pods in the cluster",
			ExtraFields: map[string]string{"StabilityLevel": "ALPHA"},
			ConstLabels: map[string]string{"controller": "observability"},
		},
		MetricType: operatormetrics.GaugeType,
		Expr:       intstr.FromString(fmt.Sprintf("sum(up{namespace='%s', pod=~'observability-operator-.*', ready='true'}) or vector(0)", namespace)),
	},
}

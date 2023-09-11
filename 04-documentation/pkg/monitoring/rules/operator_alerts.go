package rules

import (
	"fmt"

	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var operatorAlerts = []promv1.Rule{
	{
		Alert: "ObservabilityOperatorDown",
		Expr:  intstr.FromString(fmt.Sprintf("%snumber_of_pods == 0", recordingRulesPrefix)),
		Annotations: map[string]string{
			"summary":     "Observability operator is down",
			"description": "Observability operator is down for more than 5 minutes.",
		},
		Labels: map[string]string{
			"severity": "critical",
		},
	},
	{
		Alert: "ObservabilityOperatorNotReady",
		Expr:  intstr.FromString(fmt.Sprintf("%snumber_of_ready_pods < %snumber_of_pods", recordingRulesPrefix, recordingRulesPrefix)),
		For:   "5m",
		Annotations: map[string]string{
			"summary":     "Observability operator is not ready",
			"description": "Observability operator is not ready for more than 5 minutes.",
		},
		Labels: map[string]string{
			"severity": "critical",
		},
	},
}

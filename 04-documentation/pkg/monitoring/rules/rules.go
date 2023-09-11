package rules

import (
	"context"
	"github.com/machadovilaca/operator-observability/pkg/operatorrules"
	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	recordingRulesPrefix = "observability_operator_"
	namespace            = "observability-operator"
)

var alertsLog = ctrl.Log.WithName("alerts")

var (
	// Add your custom recording rules here
	recordingRules = [][]operatorrules.RecordingRule{
		operatorRecordingRules,
	}

	// Add your custom alerts here
	alerts = [][]promv1.Rule{
		operatorAlerts,
	}
)

//+kubebuilder:rbac:groups=monitoring.coreos.com,resources=prometheusrules,verbs=get;list;watch;create;update;patch;delete

func SetupRules(client client.Client) {
	alertsLog.Info("registering recording rules")
	RegisterRecordingRules()

	alertsLog.Info("registering alerts")
	RegisterAlerts()

	// For simplicity, we are going to create here the PrometheusRule CR,
	// but you should use a separate controller to ensure that the CR is
	// created and updated when needed.
	pr, err := BuildPrometheusRule()
	if err != nil {
		panic(err)
	}

	err = client.Delete(context.TODO(), pr)
	if err != nil {
		if !errors.IsNotFound(err) {
			panic(err)
		}
	}

	err = client.Create(context.TODO(), pr)
	if err != nil {
		panic(err)
	}
}

func RegisterRecordingRules() {
	err := operatorrules.RegisterRecordingRules(recordingRules...)
	if err != nil {
		panic(err)
	}
}

func RegisterAlerts() {
	err := operatorrules.RegisterAlerts(alerts...)
	if err != nil {
		panic(err)
	}
}

func BuildPrometheusRule() (*promv1.PrometheusRule, error) {
	rules, err := operatorrules.BuildPrometheusRule(
		"observability-operator-prometheus-rules",
		namespace,
		map[string]string{"app": "observability-operator"},
	)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func ListRecordingRules() []operatorrules.RecordingRule {
	return operatorrules.ListRecordingRules()
}

func ListAlerts() []promv1.Rule {
	return operatorrules.ListAlerts()
}

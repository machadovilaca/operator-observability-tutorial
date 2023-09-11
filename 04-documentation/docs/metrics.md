# Observability Operator Metrics

### observability_operator_cr_count
[DEPRECATED in 1.14.0] Number of existing observability custom resources. Type: Gauge.

### observability_operator_reconcile_action_count
[ALPHA] Number of times the operator has executed the reconcile loop with a given action. Type: Counter.

### observability_operator_reconcile_count
Number of times the operator has executed the reconcile loop. Type: Counter.

### observability_operator_number_of_pods
Number of observability operator pods in the cluster. Type: Gauge.

### observability_operator_number_of_ready_pods
[ALPHA] Number of ready observability operator pods in the cluster. Type: Gauge.

## Developing new metrics

All metrics documented here are auto-generated and reflect exactly what is being
exposed. After developing new metrics or changing old ones please regenerate
this document.


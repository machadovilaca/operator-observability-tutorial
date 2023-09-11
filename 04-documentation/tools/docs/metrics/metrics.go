package main

import (
	"fmt"

	"github.com/machadovilaca/operator-observability/pkg/docs"

	"github.com/machadovilaca/operator-observability-tutorial/pkg/monitoring/metrics"
	"github.com/machadovilaca/operator-observability-tutorial/pkg/monitoring/rules"
)

const tpl = `# Observability Operator Metrics

{{- range . }}

{{ $deprecatedVersion := "" -}}
{{- with index .ExtraFields "DeprecatedVersion" -}}
    {{- $deprecatedVersion = printf " in %s" . -}}
{{- end -}}

{{- $stabilityLevel := "" -}}
{{- if and (.ExtraFields.StabilityLevel) (ne .ExtraFields.StabilityLevel "STABLE") -}}
	{{- $stabilityLevel = printf "[%s%s] " .ExtraFields.StabilityLevel $deprecatedVersion -}}
{{- end -}}

### {{ .Name }}
{{ print $stabilityLevel }}{{ .Help }}. Type: {{ .Type -}}.

{{- end }}

## Developing new metrics

All metrics documented here are auto-generated and reflect exactly what is being
exposed. After developing new metrics or changing old ones please regenerate
this document.
`

func main() {
	metrics.RegisterMetrics()
	metrics.RegisterCollectors()
	rules.RegisterRecordingRules()

	//docsString := docs.BuildMetricsDocs(metrics.ListMetrics(), rules.ListRecordingRules())
	docsString := docs.BuildMetricsDocsWithCustomTemplate(metrics.ListMetrics(), rules.ListRecordingRules(), tpl)
	fmt.Println(docsString)
}

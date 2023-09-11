package main

import (
	"fmt"

	"github.com/machadovilaca/operator-observability/pkg/docs"

	"github.com/machadovilaca/operator-observability-tutorial/pkg/monitoring/rules"
)

func main() {
	rules.RegisterAlerts()

	docsString := docs.BuildAlertsDocs(rules.ListAlerts())
	fmt.Println(docsString)
}

# Generating documentation for the operator

In this part of the tutorial we will add documentation to the operator.

## Steps

Check the [tools/docs](tools/docs) directory in this repository for the files
that need to be created in the operator project.

It contains the following files:

- `metrics/metrics.go`: where we call the functions `RegisterMetrics`,
`RegisterCollectors` and `RegisterRecordingRules` to register all the
information we need to collect for the metrics documentation.

- `alerts/alerts.go`: where we call the function `RegisterAlerts` to register all the
information we need to collect for the alerts documentation.

Now we add commands in the Makefile to generate the documentation:

```Makefile
...
.PHONY: metrics-docs
metrics-docs:
	mkdir -p docs
	go run -ldflags="-w -s" ./tools/docs/metrics.go > docs/metrics.md

.PHONY: alerts-docs
alerts-docs:
	mkdir -p docs
	go run -ldflags="-w -s" ./tools/docs/alerts.go > docs/alerts.md
...
```

And run the commands:

```bash
make metrics-docs
make alerts-docs
```

And check the generated files:

- [docs/metrics.md](docs/metrics.md)
- [docs/alerts.md](docs/alerts.md)

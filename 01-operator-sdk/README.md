# Generating an operator with operator-sdk

In this tutorial we will generate a simple operator with operator-sdk. Refer to
the [operator-sdk documentation](https://sdk.operatorframework.io/docs/building-operators/golang/quickstart/)
for more information.

## Steps

Initialize the operator with operator-sdk:

```bash
export OPERATOR_DOMAIN=observability.example.com
export OPERATOR_REPO=github.com/machadovilaca/operator-observability-tutorial

operator-sdk init --domain $OPERATOR_DOMAIN --repo $OPERATOR_REPO
```

Create a new API for the operator:

```bash
export API_GROUP=observability
export API_VERSION=v1alpha1
export KIND=Test

operator-sdk create api --group $API_GROUP --version $API_VERSION --kind $KIND --resource --controller
make manifests
```

Run locally (only for testing):

```bash
make install run

# In another terminal, make sure you have API_GROUP, API_VERSION and KIND set
# and you are in the 01-operator-sdk directory
kubectl apply -f config/samples/${API_GROUP}_${API_VERSION}_$(echo $KIND | tr 'A-Z' 'a-z').yaml
```

## Next

=> [Adding metrics to the operator](../02-metrics/README.md)

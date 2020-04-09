# cf-k8s-logging

cf-k8s-logging contains the portions of cf-for-k8s which enable logging
outcomes.

## API

### App Containers
Logs from app containers are automatically ingested and egressed from
cf-k8s-logging

### System Components
An example is located [here](examples/forwarder)

cf-k8s-logging implements the [Fluent Forward](https://docs.fluentd.org/output/forward)
API for other components to emit logs directly into the outgoing log stream on
behalf of apps.

#### Log Format
Logs emitted to cf-k8s-logging by system components must include the fields:
- app_id: id of the app for which the logs are being emitted.
- instance_id: id of the instance for which the logs are being emitted.
- log: log message to emit

```
{"log":"This is a test log from a fluent log producer","app_id":"11111111-1111-1111-1111-111111111111","instance_id":"1"}
```

### Development flow

1. make updates needed (update vendir, update k8s files, etc).
1. make local commit(should make reverting image tags easy)
1. run build ./scripts/build-images.sh, setting $REPOSITORY to a docker
   repository you can push to
1. run ./scripts/bump-cf-for-k8s.sh
1. follow cf-for-k8s deployment steps.

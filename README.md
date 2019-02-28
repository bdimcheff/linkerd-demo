This is the code for my [talk](https://www.youtube.com/watch?v=csu11mejrvU) on using [linkerd](https://linkerd.io) with third-party APIs. It simulates a service (`service.go`) that calls a third-party API represented by the cloud function.

The cloudfunction simulates a third-party service that takes ~1-4 seconds usually, but can occasionally be faster or much slower. It also 500s 10% of the time.

The top-level `service.go` is the service itself. It's just calls the cloud function via http and returns how long it took and passes through the http status code from the cloud function.

The stuff in `kubernetes/` defines a deployment and service, which are standard. `serviceprofile.yaml` is a CRD to tell linkerd how to track outgoing connections. It also enables retries, which should make the flaky 3rd party service appear to have close to 100% reliability.
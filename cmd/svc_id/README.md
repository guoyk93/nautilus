# svc_id

## Environment Variables

* `BIND`, bind address for nrpc service, default to `:3000`
* `CLUSTER_ID`, cluster id, max 5-bits length, required
* `WORKER_ID`, worker id, max 5-bits length, use `StatefulSet` sequence id automatically

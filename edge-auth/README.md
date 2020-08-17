edge-auth
=======

The edge-auth service can be used to evaluate whether access is available to a given resource

For example:

Check access to resource (r) /function/system-dashboard:

```
http://edge-auth:8080/q/?r=/function/system-dashboard
```

Responses:

* 200 - OK
* 402 - Payment required
* 403 - Forbidden

## Building

```
export TAG=0.7.3
make build push
```

## Running

All environmental variables must be set and configured for the service whether running locally as a container, via Swarm or on Kubernetes.

* `/system-dashboard` and `system-github-event` are accessible
* All pipeline functions in OpenFaaS Cloud's stack.yml are blocked by default from all ingress such as `git-tar` and `buildshiprun`
* All other (user) functions require payment

### As a local container:

```sh
docker rm -f edge-auth
export TAG=0.7.3

docker run \
 -e PORT=8080 \
 -p 8880:8080 \
 -e receipt_verifier_uri="http://receipt-verifier.openfaas:3000" \
 -e request_price="100" \
 --name edge-auth -ti openfaas/edge-auth:${TAG}
```

### On Kubernetes

Edit `yaml/core/edge-auth-dep.yml` as needed and apply that file.

### On Swarm:

```sh
export TAG=0.7.1
docker service rm edge-auth
docker service create --name edge-auth \
 -e PORT=8080 \
 -p 8085:8080 \
 -e receipt_verifier_uri="http://receipt-verifier.openfaas:3000" \
 -e request_price="100" \
 openfaas/edge-auth:$TAG
```

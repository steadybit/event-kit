# Event Example: Logging

A tiny event implementation which exposes the required HTTP APIs and logs all interactions with it to the console. If you are new to Steadybit's EventKit,
this example app will help you understand the fundamental contracts and control flows.

## Starting the example

```sh
npm install
npm start
```

## Starting the example through Kubernetes

This is the recommended approach to give this example app a try. The app is deployed within a namespace `steadybit-extension` as a deployment
called `example-nodejs-event-logging`. Several other resources are created as well to permit access to the relevant API endpoints. For more details, please
inspect `kubernetes.yml`.

```sh
kubectl apply -f kubernetes.yml
```

Once deployed in your Kubernetes cluster the example is reachable
through `http://example-nodejs-event-logging.steadybit-extension.svc.cluster.local:8085/events`. Steadybit agents can be configured to support this
event provider through the environment variable `STEADYBIT_AGENT_EVENTS_EXTENSIONS_0_URL`.

## Starting the example using Docker

```sh
docker run -it \
  --rm \
  --init \
  -p 8085:8085 \
  ghcr.io/steadybit/example-nodejs-event-logging:main
```
## Overview

Deploy a Consul datacenter (one server and one client), Grafana container, Prometheus container, and a Node-Exporter container. These resources will be used to provide metrics observability capabilities.

## Prerequisites

- Docker
- Docker Compose
- Linux or OSX

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to this directory.
3. `docker-compose up -d`

## Testing procedure

1. Navigate to [http://localhost:9090/targets](http://localhost:9090/targets).
2. Notice that Prometheus is pre-configured to scrape data from three endpoints: Consul, Node-Exporter, and itself.
3. Navigate to [http://localhost:3000/datasources](http://localhost:3000/datasources).
4. Login to Grafana with `consul/consul`.
5. Notice that Grafana is pre-configured with Prometheus as a data source.
6. Navigate to [http://localhost:3000/dashboards](http://localhost:3000/dashboards).
7. Explore the Consul Server Monitoring and Node Exporter Dashboards.
8. Explore [Monitor Datacenter Health](https://learn.hashicorp.com/tutorials/consul/monitor-datacenter-health) tutorial.

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://learn.hashicorp.com/tutorials/consul/monitor-datacenter-health](https://learn.hashicorp.com/tutorials/consul/monitor-datacenter-health)
- [https://learn.hashicorp.com/tutorials/consul/kubernetes-layer7-observability](https://learn.hashicorp.com/tutorials/consul/kubernetes-layer7-observability)
- [https://www.consul.io/docs/agent/telemetry](https://www.consul.io/docs/agent/telemetry)

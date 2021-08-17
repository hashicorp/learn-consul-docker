## Overview

Deploy a Consul datacenter containing one server and three clients. 
Client 1 and Client 2 run a counting service (data layer) while Client 3 runs a dashboard service (application layer).
Consul Intentions enables secure communication between the services in a service mesh.
Consul DNS is used to load balance between the two counting containers.

![Example Consul Service Mesh Architecture Diagram](https://learn.hashicorp.com/img/consul/connect-getting-started/consul_connect_demo_service_flow.png)

## Prerequisites

- Docker
- Docker Compose

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to this directory.
3. `docker-compose up -d`
4. `chmod 755 service-install.sh`
5. `./service-install.sh`

## Testing procedure

1. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
2. Explore services and intentions tabs of the Consul UI.
3. Navigate to [http://localhost:9002/](http://localhost:9002/) to see the Dashboard service.
4. Notice how the Dashboard service receives data from the Counting service due to connectivity provided by Consul Intentions.
5. Refresh the Dashboard UI multiple times; notice how different hostnames are shown as the browser is reloaded.

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://learn.hashicorp.com/tutorials/consul/service-mesh-with-envoy-proxy](https://learn.hashicorp.com/tutorials/consul/service-mesh-with-envoy-proxy)
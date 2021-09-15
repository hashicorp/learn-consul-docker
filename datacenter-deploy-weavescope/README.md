## Overview

Deploy a Consul datacenter containing one server and two clients. 
Client 1 runs a counting service (data layer) while Client 2 runs a dashboard service (application layer).
Consul Intentions enables secure communication between the services in a service mesh.

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

1. Navigate to [http://localhost:4040](http://localhost:4040) on your local browser.
2. Explore the network diagram and traffic flow presented by the Weavescope application.
3. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
4. Explore services and intentions tabs of the Consul UI.
5. Navigate to [http://localhost:9002/](http://localhost:9002/) to see the Dashboard service. 
6. Navigate to [http://localhost:9002/](http://localhost:9002/) to see the Counting service.
7. Notice how the Dashboard service front-end receives data from the Counting service back-end due to connectivity provided by Consul Intentions. 
8. Refresh the Dashboard and Counting service pages several times to generate traffic.
9. Navigate to [http://localhost:4040/#!/state/{%22pinnedMetricType%22:%22CPU%22,%22topologyId%22:%22processes%22}](http://localhost:4040/#!/state/{%22pinnedMetricType%22:%22CPU%22,%22topologyId%22:%22processes%22}) on your local browser.
10. Notice how traffic flows through the network.
11. Explore [Service Mesh](https://learn.hashicorp.com/tutorials/consul/service-mesh-with-envoy-proxy) tutorial.

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://learn.hashicorp.com/tutorials/consul/service-mesh-with-envoy-proxy](https://learn.hashicorp.com/tutorials/consul/service-mesh-with-envoy-proxy)
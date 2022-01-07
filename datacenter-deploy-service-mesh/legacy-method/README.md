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

1. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
2. Explore services and intentions tabs of the Consul UI.
3. Navigate to [http://localhost:9002/](http://localhost:9002/) to see the Dashboard service.
4. Notice how the Dashboard service receives data from the Counting service due to connectivity provided by Consul Intentions. In this example, this configuration was created by using the legacy `consul intention create <source_service> <destination_service>`  method.
5. Navigate to the Intentions tab in the [Consul UI](http://localhost:8500/ui/dc1/intentions/).
6. Change the Dashboard->Counting Intention's default behavior from `Allow` to `Deny`.
7. Navigate to [http://localhost:9002/](http://localhost:9002/).
8. Notice how the Dashboard service is no longer able to communicate with the Counting service due to the `Deny` default behavior setting.
9. Explore [Service Mesh](https://learn.hashicorp.com/tutorials/consul/service-mesh-with-envoy-proxy) tutorial.

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://learn.hashicorp.com/tutorials/consul/service-mesh-with-envoy-proxy](https://learn.hashicorp.com/tutorials/consul/service-mesh-with-envoy-proxy)
## Overview

Deploy a Consul datacenter containing agents with a preconfigured DNS service and health check.

## Prerequisites

- Docker
- Docker Compose

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to this directory.
3. `docker-compose up -d`

## Testing procedure

1. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
2. Explore DNS service and health check.
3. Explore [Service Discovery](https://learn.hashicorp.com/tutorials/consul/get-started-service-discovery) tutorial.

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://learn.hashicorp.com/tutorials/consul/get-started-service-discovery](https://learn.hashicorp.com/tutorials/consul/get-started-service-discovery)
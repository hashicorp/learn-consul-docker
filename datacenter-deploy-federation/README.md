## Overview

Deploy a federated Consul environment containing two secure datacenters.

## Prerequisites

- Docker
- Docker Compose

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to this directory.
3. `docker-compose up -d`

## Testing procedure

1.  Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
2.  Explore navigation between the two federated datacenters.
3.  Explore the [Federation Gossip WAN](https://learn.hashicorp.com/tutorials/consul/federation-gossip-wan) tutorial.

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://learn.hashicorp.com/tutorials/consul/tls-encryption-secure](https://learn.hashicorp.com/tutorials/consul/tls-encryption-secure)
- [https://learn.hashicorp.com/tutorials/consul/federation-gossip-wan](https://learn.hashicorp.com/tutorials/consul/federation-gossip-wan)
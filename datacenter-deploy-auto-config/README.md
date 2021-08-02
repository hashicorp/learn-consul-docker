# Currently Under Development - 08/2021

## Overview

Deploy a three-server Consul datacenter that utilizes `auto-config` for allowing clients to securely retrieve ACL tokens, TLS certificates, gossip encryption keys, and other configuration settings when joining the Consul datacenter.

## Prerequisites

- Docker
- Docker Compose

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to this directory.
3. `docker-compose up -d`

## Testing procedure

1. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
2. Login with the pre-configured master token: `e95b599e-166e-7d80-08ad-aee76e7ddf19`
3. Create additional content her

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://www.consul.io/docs/agent/options#auto_config](https://www.consul.io/docs/agent/options#auto_config)
- [https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter](https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter)
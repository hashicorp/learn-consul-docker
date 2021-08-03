# Currently Under Development - 08/2021

## Overview

Deploy a three-server Consul datacenter that utilizes `auto-config` for allowing clients to securely retrieve ACL tokens, TLS certificates, gossip encryption keys, and other configuration settings when joining the Consul datacenter.

## Prerequisites

- Docker
- Docker Compose
- golang

## Secint setup procedure

1. Clone [secint](https://github.com/banks/secint) repository.
2. Navigate to the `secint` directory.
3. `go build`
4. Copy the newly created `secint` binary to your path (for example, `cp secint /usr/bin/`)

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to the `auto_config` directory.
3. `docker-compose up -d`

## Testing procedure

1. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
2. Login with the pre-configured master token: `e95b599e-166e-7d80-08ad-aee76e7ddf19`
3. Notice that no client is present in the nodes tab.
4. Generate JWT: `secint mint -issuer auto-config-cluster -ttl 12h -node consul-client -priv-key ./certs/dc1-server-consul-0-key.pem -audience audience`
5. Copy the generated JWT to the `intro_token` line in `client.json`
6. Delete the current `consul-client` container: `docker rm consul-client --force`
7. Recreate the `consul-client` container: `docker-compose up -d`
8. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
9. Login with the pre-configured master token: `e95b599e-166e-7d80-08ad-aee76e7ddf19`
10. Notice that the client is now present in the nodes tab.

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://www.consul.io/docs/agent/options#auto_config](https://www.consul.io/docs/agent/options#auto_config)
- [https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter](https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter)
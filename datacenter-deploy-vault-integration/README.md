# Currently Under Development - 08/2021

## Overview

Deploy a three-server Consul datacenter and a Vault server.

## Prerequisites

- Docker
- Docker Compose

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to the `auto_config` directory.
3. `docker-compose up -d`

## Vault setup procedure

1. Navigate to [http://localhost:8200/ui](http://localhost:8200/ui/) on your local browser.
2. Login with the pre-configured root token: `vault-plaintext-root-token`
2. Enable the PKI engine: `vault secrets enable pki`
3. Generate a CA: `vault write pki/root/generate/internal common_name=consul-server* ttl=8760h`

## Testing procedure

1. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
2. Login with the pre-configured master token: `e95b599e-166e-7d80-08ad-aee76e7ddf19`
3. Notice that no client is present in the nodes tab.

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://www.consul.io/docs/agent/options#auto_config](https://www.consul.io/docs/agent/options#auto_config)
- [https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter](https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter)
- [https://www.vaultproject.io/docs/secrets/pki](https://www.vaultproject.io/docs/secrets/pki)
## Overview

Deploy a three-server Consul datacenter and a Vault server used to walk through the Consul + Vault integration procedure. This integration enables Vault to dynamically generate Consul ACL tokens with attributes such as specific ACL policies and lease times.

## Prerequisites

- Docker
- Docker Compose

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to this directory.
3. `docker-compose up -d`

## Vault + Consul integration setup procedure

1. Open an interactive shell to the first Consul server:
   1. `docker exec -it consul-server1 ash`
2. Install Vault (to be used as a client):
   1. `apk add vault`
3. Login to Vault server with the pre-configured root token:
   1. Initialize login: `vault login`
   2. Token: `vault-plaintext-root-token`
4. Enable Vault's consul secrets engine: 
   1. `vault secrets enable consul`
5. Configure Vault's secrets engine access:
   1. `vault write consul/config/access address=${CONSUL_HTTP_ADDR} token=${CONSUL_HTTP_TOKEN}`
6. Create a Consul ACL policy that permits server agents to register themselves:
   1. `consul acl policy create -name consul-clients -rules @consul/config/acls/consul-client-policy.hcl -token=e95b599e-166e-7d80-08ad-aee76e7ddf19`
7. Create a Vault server role that maps a name in Vault to a set of Consul ACL policies: 
   1. `vault write consul/roles/consul-client-role policies=consul-clients`
8. Obtain and test a Consul ACL token
   1. `vault read consul/creds/consul-client-role | tee consul.client`
   2. `export CONSUL_SERVER_ACCESSOR=$(cat consul.client | grep accessor | awk '{print $2}')`
   3. `consul acl token read -id ${CONSUL_SERVER_ACCESSOR}`
9. Verify token generation in Vault UI and Consul UI:
   1.  Vault: 
       1.  URL: `http://localhost:8200/ui/vault/access/leases/list/consul/creds/consul-client-role/`
       2.  Login with Token: `vault-plaintext-root-token`
       3.  Notice the lease created for the generated ACL token.
   2.  Consul:
       1.  URL: `http://localhost:8500/ui/dc1/acls/tokens`
       2.  Login with Token: `e95b599e-166e-7d80-08ad-aee76e7ddf19`
       3.  Notice that Vault created an ACL token with the `consul-clients` policy attached.
   3.  At this point, any user or process with access to Vault can now create short-lived Consul tokens in order to carry out operations, thus centralizing the management of Consul tokens.
10. Exit the terminal: `exit`

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://www.consul.io/docs/agent/options#auto_config](https://www.consul.io/docs/agent/options#auto_config)
- [https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter](https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter)
- [https://learn.hashicorp.com/tutorials/consul/vault-consul-secrets](https://learn.hashicorp.com/tutorials/consul/vault-consul-secrets)
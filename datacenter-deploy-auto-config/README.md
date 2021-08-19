# Currently Under Development - 08/2021

## Overview

Deploy a three-server Consul datacenter that utilizes `auto-config` for allowing clients to securely retrieve ACL tokens, TLS certificates, gossip encryption keys, and other configuration settings when joining the Consul datacenter.

## Prerequisites

- Docker
- Docker Compose

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to the `auto_config` directory.
3. `docker-compose up -d`

# Vault testing procedure

1. Open an interactive shell to the Vault server
   1. `docker exec -it vault-server ash`
2. Login to Vault server with the pre-configured root token
   1. Initialize login: `vault login`
   2. Token: `vault-plaintext-root-token`
3. Create a key
   1. `vault write identity/oidc/key/oidc-key-1 allowed_client_ids="*"`
4. Create a role
   1. `vault write identity/oidc/role/oidc-role-1 ttl=12h key="oidc-key-1" client_id="consul-client"`
5. Create a policy
   1. `vault policy write oidc-policy /vault/policies/policy.json`
6. Enable Vault's username/password secrets engine
   1. `vault auth enable userpass`
7. Create an example user with the oidc-policy attached
   1. `vault write auth/userpass/users/example password=password policies=oidc-policy`
8. Login as the newly created user
    1.  `vault login -method=userpass username=example`
    2.  Password: `password`
9.  Get a signed ID token
    1.  `vault read identity/oidc/token/oidc-role-1`
10. Copy the `token` value to the /tokens/jwt file in your working directory, then save the changes
11. Delete then recreate the `consul-client` container
    1.  `docker rm consul-client --force` 
    2.  `docker-compose up -d`
12. Notice that `consul-client` has now successfully joined the Consul datacenter using the `auto_config` method.



## Secint testing procedure

1. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
2. Login with the pre-configured master token: `e95b599e-166e-7d80-08ad-aee76e7ddf19`
3. Notice that no client is present in the nodes tab.
4. Generate JWT: `secint mint -issuer auto-config-cluster -ttl 12h -node consul-client -priv-key ./certs/auto-config-private-key.pem -audience audience`
5. Copy the generated JWT to the `intro_token` line in `client.json`
6. Delete the current `consul-client` container: `docker rm consul-client --force`
7. Recreate the `consul-client` container: `docker-compose up -d`
8. Navigate to [http://localhost:8500/ui](http://localhost:8500/ui/) on your local browser.
9. Login with the pre-configured master token: `e95b599e-166e-7d80-08ad-aee76e7ddf19`
10. Notice that the client is now present in the nodes tab.

## Additional information

- [secint repository](https://github.com/banks/secint)
- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://www.consul.io/docs/agent/options#auto_config](https://www.consul.io/docs/agent/options#auto_config)
- [https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter](https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter)
- [https://www.vaultproject.io/docs/secrets/identity](https://www.vaultproject.io/docs/secrets/identity)
- [https://www.vaultproject.io/api-docs/secret/identity/tokens](https://www.vaultproject.io/api-docs/secret/identity/tokens)

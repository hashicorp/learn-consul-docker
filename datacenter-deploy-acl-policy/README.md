## Overview

Deploy a one-server, one-client environment used for testing ACL policies. In this example, a policy is created that denies all requests from anonymous tokens.

## Prerequisites

- Docker
- Docker Compose

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to this directory.
3. `docker-compose up -d`

## Consul ACL Testing Procedure

1. Open an interactive shell to the Consul server:
   1. `docker exec -it consul-server1 ash`
2. Create a Consul ACL policy that denies all anonymous token requests:
   1. `consul acl policy create -name consul-anonymous-token-policy -rules @consul/config/acls/consul-anonymous-token-policy.hcl -token=e95b599e-166e-7d80-08ad-aee76e7ddf19`
3. Make a request:
   1. With the token in the request: 
      1. `curl --header "X-Consul-Token: e95b599e-166e-7d80-08ad-aee76e7ddf19" http://127.0.0.1:8500/v1/agent/members`
      2. Notice that all node data is able to be retrieved.
   2. Without the token in the request (uses anonymous token): 
      1. `curl http://consul-server1:8500/v1/catalog/nodes`
      2. Notice that the client cannot access Consul data since it uses the anonymous token for requests.
4.  Exit the terminal: `exit`

## Additional information

- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter](https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter)

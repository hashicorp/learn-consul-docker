# Currently Under Development - 07/2021

## Overview

Deploy, secure, and scale a sample microservice application with the Hashistack (Terraform, Consul, and Vault).

![](images/hashistack-cloud-operating-model.png)

## Prerequisites

- Docker
- Terraform v0.14+

## Deployment procedure

1. Clone [learn-consul-docker](https://github.com/hashicorp/learn-consul-docker) repository.
2. Navigate to this directory.
3. Edit the terraform `pwd` variable in `variables.tf` to the absolute path of your working directory
4. `terraform init`
5. `terraform apply`

## Consul: Service Discovery and Health Check setup

1. Navigate to [http://localhost:9090/ui/](http://localhost:9090/ui/).
   1. Notice the architecture of the application stack. 
2. Navigate to [http://localhost:8500/ui/dc1/services](http://localhost:8500/ui/dc1/services)
   1. Notice the services and nodes being monitored by Consul.
3. Install Consul on each application container:
   1.  `docker exec -d web apk add consul`
   2.  `docker exec -d api apk add consul`
   3.  `docker exec -d vault-server apk add consul`
4.  Start the Consul agent manually on each application container:
    1.  `docker exec -d web consul agent --config-dir=/var/consul/config/`
    2.  `docker exec -d api consul agent --config-dir=/var/consul/config/`
    3.  `docker exec -d vault-server consul agent --config-dir=/consul/config/`
5.  Navigate to [http://localhost:8500/ui/dc1/services](http://localhost:8500/ui/dc1/services)
    1.  Notice the services, health checks, and nodes now being monitored by Consul.

## Terraform: Infrastructure as Code and Scaling

1.  Append the contents of `additional-servers.txt` to your `main.tf`:
    1.  `cat additional-servers.txt >> main.tf`
2.  `terraform apply`
3.  Install Consul on each newly-created application container:
    1.  `docker exec -d web2 apk add consul`
4.  Start the Consul agent manually on the newly-created application container:
    1.  `docker exec -d web2 consul agent --config-dir=/var/consul/config/`
5.   Navigate to [http://localhost:8500/ui/dc1/services/web/instances](http://localhost:8500/ui/dc1/services/web/instances)
6.  Notice the additional instances of the `web` service now being monitored by Consul.

## Vault: (development in progress)

1. Open an interactive shell to the first Consul server:
   1. `docker exec -it vault-server ash`
2. Login to Vault server with the pre-configured root token:
   1. Initialize login: `vault login`
   2. Token: `vault-plaintext-root-token`
3. Enable Vault's ______ secrets engine: 
   1. `vault secrets enable ______`


## Consul + vault content (use-cases)

## Consul + terraform content (use-cases)

## Concepts Covered

1. Terraform
   1. Infrastructure as Code
   2. Scaling
2. Consul
   1. Service Discovery
   2. Health Checking
3. Vault
   1. Secret Management
   2. Automatic Gossip Key Rotation

## Additional information

- [https://www.hashicorp.com/cloud-operating-model](https://www.hashicorp.com/cloud-operating-model)
- [https://learn.hashicorp.com/collections/consul/docker](https://learn.hashicorp.com/collections/consul/docker)
- [https://learn.hashicorp.com/tutorials/consul/vault-kv-consul-secure-gossip](https://learn.hashicorp.com/tutorials/consul/vault-kv-consul-secure-gossip)
- [https://learn.hashicorp.com/tutorials/vault/codify-mgmt-oss](https://learn.hashicorp.com/tutorials/vault/codify-mgmt-oss)
- [https://learn.hashicorp.com/tutorials/consul/vault-consul-secrets](https://learn.hashicorp.com/tutorials/consul/vault-consul-secrets)

## Application reference

This demo consists of two example services, Web (HTTP) and API (gRPC).

```
web (HTTP)  --
                api (gRPC)
```

## Gratitude

- [Nic Jackson @ Hashicorp](https://github.com/nicholasjackson)
- [Rosemary Wang @ Hashicorp](https://github.com/joatmon08)
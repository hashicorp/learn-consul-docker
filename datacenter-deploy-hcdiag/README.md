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

1. Open an interactive shell to a Consul server:
   1. `docker exec -it consul-server1 /bin/sh`
2. Set environment variables:
   1. `export CONSUL_HTTP_ADDR=http://127.0.0.1:8500`
   2. `export CONSUL_HTTP_TOKEN=my-master-token`
   3. `export CONSUL_TOKEN=my-master-token`
3. Check to see if all clients have successfully joined the Consul datacenter:
   1. `consul members`
   2. `curl http://127.0.0.1:8500/v1/agent/members --header "X-Consul-Token: my-master-token"`
4. Update package sources;
   1. `apk update`
5. Add curl and unzup packages:
   1. `apk add curl unzip`
6. Download hcdiag:
   1. `curl --silent --remote-name https://releases.hashicorp.com/hcdiag/0.1.1/hcdiag_0.1.1_linux_amd64.zip`
7. Unzip hcdiag and remove the archive:
   1. `unzip hcdiag_0.1.1_linux_amd64.zip && rm -f hcdiag_0.1.1_linux_amd64.zip`
8. Move the hcdiag executable to your sbin directory:
   1. `mv hcdiag sbin/`
9.  Run hcdiag for consul:
   2. `hcdiag -consul`
   3. Let it run until completion
10. Look for the support package
   4. `ls -l *.gz`
11. Unpack the archive:
   5.  `tar zxvf support-2021-12-10T20:47:55Z.tar.gz`
12. Change directory into the unpacked folder:
    1.  `cd temp495511880/`
13. Examine the contents 
14. Exit the terminal: `exit`

## Additional information

- [https://github.com/hashicorp/hcdiag](https://github.com/hashicorp/hcdiag)
- [https://learn.hashicorp.com/tutorials/consul/troubleshooting](https://learn.hashicorp.com/tutorials/consul/troubleshooting)
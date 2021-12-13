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
2. Update package sources;
   1. `apk update`
3. Add curl and unzup packages:
   1. `apk add curl unzip`
4. Download hcdiag:
   1. `curl --silent --remote-name https://releases.hashicorp.com/hcdiag/0.1.1/hcdiag_0.1.1_linux_amd64.zip`
5. Unzip hcdiag and remove the archive:
   1. `unzip hcdiag_0.1.1_linux_amd64.zip && rm -f hcdiag_0.1.1_linux_amd64.zip`
6. Move the hcdiag executable to your sbin directory:
   1. `mv hcdiag sbin/`
7. Run hcdiag for consul:
   1. `hcdiag -consul`
   2. Let it run until completion
8. Look for the support package
   1. `ls -l *.gz`
9. Unpack the archive:
   1.  `tar zxvf support-2021-12-10T20:47:55Z.tar.gz`
10. Change directory into the unpacked folder:
    1.  `cd temp495511880/`
11. Examine the contents 
13. Exit the terminal: `exit`

## Additional information

- [https://github.com/hashicorp/hcdiag](https://github.com/hashicorp/hcdiag)
- [https://learn.hashicorp.com/tutorials/consul/troubleshooting](https://learn.hashicorp.com/tutorials/consul/troubleshooting)
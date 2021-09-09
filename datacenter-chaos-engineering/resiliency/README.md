## Tutorial URL

To execute Pumba and stop random containers

```shell
docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock gaiaadm/pumba \ | pumba -l info --random  --interval 60s pause --duration 59s
```
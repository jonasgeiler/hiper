# hiper

Automatically updates Hetzner Cloud Floating IP server assignments in a Docker Swarm.  
This project was inspired by the blog post ["Implementation of a Cloud HA infrastructure with Keepalived" written by Markus (@BackInBash)](https://community.hetzner.com/tutorials/configure-cloud-ha-keepalived)
and some parts of the code were taken from [hcloud-ip](https://github.com/FootprintDev/hcloud-ip).  

## How it works

`hiper` relies on the automatic container rescheduling performed by Docker Swarm. Whenever a host goes down and this container happens to run on it,
it will be started on another node, and `hiper` will then update the floating IP assigment.  
I got the idea from [this serverfault answer](https://serverfault.com/a/930938/938715), but instead of DynDNS, I rely on floating IPs.

## Usage

```shell
docker run -t --name hiper ghcr.io/zackplan/hiper
```

An example using docker compose can be found [here](https://github.com/zackplan/hiper/blob/main/docker-compose.yml).

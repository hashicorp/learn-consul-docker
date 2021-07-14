terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
    }
  }
}

provider "docker" {}

variable pwd {
  default = "/Users/eddierowe/repos/eddie-private/terraform_testing"
}

resource "docker_network" "private_network" {
  name = "bridge_network"
  attachable = true
  driver = "bridge"
}

resource "docker_image" "fake-service" {
  name         = "nicholasjackson/fake-service:v0.21.0"
  keep_locally = false
}

resource "docker_image" "consul" {
  name         = "hashicorp/consul:latest"
  keep_locally = false
}

resource "docker_image" "vault" {
  name         = "hashicorp/vault:latest"
  keep_locally = false
}

resource "docker_container" "web" {
  image = docker_image.fake-service.latest
  name  = "web"
  ports {
    internal = 9090
    external = 9090
  }
  env = ["LISTEN_ADDR=0.0.0.0:9090", "UPSTREAM_URIS=grpc://api:9091", "MESSAGE=Hello World", "NAME=web", "SERVER_TYPE=http"]
  networks_advanced {
    name = "${docker_network.private_network.name}"
  }
  volumes {
    host_path = "${var.pwd}/fake-service/web1.json"
    container_path = "/var/consul/config/web1.json"
  }
}

resource "docker_container" "api" {
  image = docker_image.fake-service.latest
  name  = "api"
  ports {
    internal = 9091
    external = 9091
  }
  env = ["LISTEN_ADDR=0.0.0.0:9091", "MESSAGE=API Response", "NAME=api", "SERVER_TYPE=grpc"]
  networks_advanced {
    name = "${docker_network.private_network.name}"
  }
  volumes {
    host_path = "${var.pwd}/fake-service/api1.json"
    container_path = "/var/consul/config/api1.json"
  }
}

resource "docker_container" "consul-server" {
  image = docker_image.consul.latest
  name  = "consul-server"
  ports {
    internal = 8500
    external = 8500
  }
  ports {
    internal = 8600
    external = 8600
  }
  volumes {
    host_path = "${var.pwd}/consul/server.json"
    container_path = "/consul/config/server.json"
  }
  networks_advanced {
    name = "${docker_network.private_network.name}"
  }
  command = ["agent"]
}

resource "docker_container" "vault" {
  image = docker_image.vault.latest
  name  = "vault"
  ports {
    internal = 8200
    external = 8200
  }
}
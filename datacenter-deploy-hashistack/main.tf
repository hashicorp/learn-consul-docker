# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
    }
  }
}

provider "docker" {}

resource "docker_network" "private_network" {
  name = "bridge_network"
  attachable = true
  driver = "bridge"
}

resource "docker_image" "consul" {
  name         = "hashicorp/consul:latest"
  keep_locally = false
}

resource "docker_image" "vault" {
  name         = "hashicorp/vault:latest"
  keep_locally = false
}

resource "docker_image" "fake-service" {
  name         = "nicholasjackson/fake-service:v0.21.0"
  keep_locally = false
}

resource "docker_container" "consul-server" {
  image = docker_image.consul.latest
  name  = "consul-server"
  hostname = "consul-server"
  env = [
    "VAULT_ADDR=http://vault-server:8200",
    "VAULT_TOKEN=vault-plaintext-root-token",
    "CONSUL_HTTP_ADDR=consul-server:8500",
    "CONSUL_HTTP_TOKEN=e95b599e-166e-7d80-08ad-aee76e7ddf19"
  ]
  capabilities {
    add = ["IPC_LOCK"]
    drop = []
  }
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

resource "docker_container" "vault-server" {
  image = docker_image.vault.latest
  name  = "vault-server"
  hostname = "vault-server"
  env = [
    "VAULT_ADDR=http://0.0.0.0:8200",
    "VAULT_DEV_ROOT_TOKEN_ID=vault-plaintext-root-token"
  ]
  capabilities {
    add = ["IPC_LOCK"]
    drop = []
  }
  ports {
    internal = 8200
    external = 8200
  }
  networks_advanced {
    name = "${docker_network.private_network.name}"
  }
  volumes {
    host_path = "${var.pwd}/vault/vault.json"
    container_path = "/consul/config/vault.json"
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


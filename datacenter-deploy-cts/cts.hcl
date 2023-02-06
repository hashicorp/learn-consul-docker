# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# Configures Consul-Terraform-Sync connection with a Consul agent to perform queries to the Consul catalog
consul {
  address = "consul-server:8500"
}

# A task is executed when any change to the defined services are detected in the Consul catalog
task {
 name        = "learn-cts-example"
 description = "Example task with two services"
 # This Terraform module creates text files on your local machine containing Consul service information
 source      = "findkim/print/cts"
 version     = "0.1.0"
 services    = ["consul", "dns"]
}

# Relays provider discovery and installation information to Terraform
driver "terraform" {
  # version = "0.14.0"
  # path = ""
  log = false
  persist_log = false
  #working_dir = ""
  backend "consul" {
    gzip = true
  }
}
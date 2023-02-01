# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data_dir = "/tmp/"

telemetry {
  prometheus_retention_time = "60s"
  disable_hostname = true
}

log_level = "TRACE"

datacenter = "dc1"

server = true

bootstrap_expect = 1
ui = true

bind_addr = "0.0.0.0"
client_addr = "0.0.0.0"

ports {
  grpc = 8502
}

connect {
  enabled = true
}

advertise_addr = "10.5.0.2"
enable_central_service_config = true

config_entries {
  bootstrap = [
    {
      kind = "proxy-defaults"
      name = "global"
      
      config {
        envoy_extra_static_clusters_json = <<EOL
          {
            "connect_timeout": "3.000s",
            "dns_lookup_family": "V4_ONLY",
            "lb_policy": "ROUND_ROBIN",
            "load_assignment": {
                "cluster_name": "zipkin",
                "endpoints": [
                    {
                        "lb_endpoints": [
                            {
                                "endpoint": {
                                    "address": {
                                        "socket_address": {
                                            "address": "10.5.0.9",
                                            "port_value": 9411,
                                            "protocol": "TCP"
                                        }
                                    }
                                }
                            }
                        ]
                    }
                ]
            },
            "name": "zipkin",
            "type": "STRICT_DNS"
          }
        EOL

        envoy_tracing_json = <<EOL
          {
              "http": {
                  "name": "envoy.zipkin",
                  "config": {
                      "collector_cluster": "zipkin",
                      "collector_endpoint": "/api/v1/spans"
                  }
              }
          }
        EOL
      }
    }
  ]
}

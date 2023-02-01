# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

node_prefix "consul-client" {
  policy = "write"
}
node_prefix "" {
  policy = "read"
}
service_prefix "" {
  policy = "read"
}
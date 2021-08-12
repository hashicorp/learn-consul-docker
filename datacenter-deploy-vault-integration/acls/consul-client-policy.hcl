node_prefix "consul-client" {
  policy = "write"
}
node_prefix "" {
  policy = "read"
}
service_prefix "" {
  policy = "read"
}
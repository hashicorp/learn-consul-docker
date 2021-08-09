# Secure Introduction Token Tool

This is an experimental tool for playing with mechanisms to generate
short-lived, single-use provisioning tokens for infrastructure components to
solve the "first secret" problem.

The idea is that the trusted platform (e.g. your infrastructure provisioning
tool or cloud provider) would create a provisioning key pair and hold the key
securely. It can then generate and inject a very short-lived and single-use
token into a new workload (VM, server instance etc) when it is provisioned which
is enough for that workload to exchange it for some arbitrary other secret data
like initial TLS certificates or longer-lived tokens.

Production applications are likely to want to use something more fully featured
like [HashiCorp Vault](https://vaultproject.io) however this is a prototype of
the simplest thing that could work without dependencies.

## Example

```
$ secint init
# Wrote secint-priv-key.pem and secint-pub-key.pem
$ secint mint -node foo
eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTE2MjcwMjQsImp0aSI6IjMxN2RiMzVhLTgxODAtMGE0Zi00NzU4LWE1ZmU3MWM2NWY0NyIsInN1YiI6ImZvbyJ9.MEUCIQCjM1Kw8dGPiZcrbHSxDcFcnPUTsZIRJC4I2KE3Wr9HkwIgUduzh855GDuUU_zzvv_9xpa4i3FB3ei2jzu_dLAY0jc
$ secint mint -node foo | secint verify -node foo
Server Allowed:     NO
JWT ID            : 5affd7ee-e784-4cb0-e3e7-308fef5d2115
Verified Node Name: foo
Server Token      : NO
VALID
$ secint mint -node foo | secint verify -node bar
ERROR: token doesn't match expected node name.
Got "foo" expect "bar"
```

## Notably Missing

This tool so far is only part of the puzzle, it assumes that there is some other
component in the infrastructure being setup capable of validating these tokens
and then providing long-lived secrets for each workload. That serving system
must also enforce the single-use aspect of the system based on stored state of
which JWT IDs have been used already.

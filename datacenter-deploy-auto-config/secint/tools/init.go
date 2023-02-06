// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/mitchellh/cli"
)

func initFactory() (cli.Command, error) {
	return &initCommand{
		ui: ui(),
	}, nil
}

type initCommand struct {
	ui cli.Ui
}

func (c *initCommand) Help() string {
	return `Create a new provisioning key pair. The key is written to
"secint-priv-key.pem" and "secint-pub-key.pem"
in the current directory.
`
}

func (c *initCommand) Synopsis() string {
	return "Creates a new provisioning key pair."
}

func (c *initCommand) Run(args []string) int {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		c.ui.Error(err.Error())
		return 1
	}

	b, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		c.ui.Error(err.Error())
		return 1
	}

	if err := writePEM("secint-priv-key.pem", "EC PRIVATE KEY", b); err != nil {
		c.ui.Error(err.Error())
		return 1
	}

	b, err = x509.MarshalPKIXPublicKey(key.Public())
	if err := writePEM("secint-pub-key.pem", "PUBLIC KEY", b); err != nil {
		c.ui.Error(err.Error())
		return 1
	}

	return 0
}

func writePEM(name string, kind string, b []byte) error {
	block := &pem.Block{Type: kind, Bytes: b}

	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := pem.Encode(file, block); err != nil {
		return err
	}
	return nil
}

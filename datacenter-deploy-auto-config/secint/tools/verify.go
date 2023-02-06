// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/cli"
	"gopkg.in/square/go-jose.v2/jwt"
)

func verifyFactory() (cli.Command, error) {
	c := &verifyCommand{
		ui: ui(),
	}

	c.init()
	return c, nil
}

type verifyCommand struct {
	ui cli.Ui

	flags *flag.FlagSet

	pubKeyFile string
	nodeName   string
	server     bool

	help string
}

const verifyHelp = `Usage: secint verify [options]

Verify a new intro token from stdin.
`

func (c *verifyCommand) Help() string {
	return Usage(c.help, nil)
}

func (c *verifyCommand) Synopsis() string {
	return `Verify an intro token from stdin.`
}

func (c *verifyCommand) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)

	c.flags.StringVar(&c.pubKeyFile, "pub-key", "secint-pub-key.pem", "public key file in PEM format")
	c.flags.BoolVar(&c.server, "server", false, "whether this token is expected to be for a server.")
	c.flags.StringVar(&c.nodeName, "node", "", "the name of the node expected.")
	c.help = Usage(verifyHelp, c.flags)
}

func (c *verifyCommand) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		return 1
	}

	// Load the public key
	key, err := pubKeyFromFile(c.pubKeyFile)
	if err != nil {
		c.ui.Error(fmt.Sprintf("ERROR: failed to read private key file: %s", err))
		return 1
	}

	tokenBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		c.ui.Error(fmt.Sprintf("ERROR: failed to read token from stdin: %s", err))
		return 1
	}

	token, err := jwt.ParseSigned(string(tokenBytes))
	if err != nil {
		c.ui.Error(fmt.Sprintf("ERROR: failed to parse JWT: %v", err))
		return 1
	}

	var stdClaims jwt.Claims
	privateClaims := make(map[string]interface{})

	// Verify token
	if err := token.Claims(key, &stdClaims, &privateClaims); err != nil {
		c.ui.Error(fmt.Sprintf("ERROR: failed to verify token: %v", err))
		return 1
	}

	if stdClaims.ID == "" {
		c.ui.Error(fmt.Sprintf("ERROR: no ID in token"))
		return 1
	}

	if stdClaims.Subject == "" {
		c.ui.Error(fmt.Sprintf("ERROR: no subject in token"))
		return 1
	}

	if c.nodeName != "" && c.nodeName != stdClaims.Subject {
		c.ui.Error(fmt.Sprintf("ERROR: token doesn't match expected node name.\n"+
			"Got %q expect %q", stdClaims.Subject, c.nodeName))
		return 1
	}

	server, ok := privateClaims["server"].(bool)
	isServer := "NO"
	if ok && server {
		isServer = "YES"
	} else {
		// Not a server token
		if c.server {
			c.ui.Error("ERROR: token is not a server token but one is expected.")
			return 1
		}
		c.ui.Output("Server Allowed:     NO")
	}

	c.ui.Output(fmt.Sprintf("JWT ID            : %s", stdClaims.ID))
	c.ui.Output(fmt.Sprintf("Verified Node Name: %s", stdClaims.Subject))
	c.ui.Output(fmt.Sprintf("Server Token      : %s", isServer))
	c.ui.Output("VALID")

	return 0
}

func pubKeyFromFile(fileName string) (*ecdsa.PublicKey, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bs)
	if block == nil {
		return nil, fmt.Errorf("No PEM block in file")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	ecpub, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Parsed public key is not an ECDSA Public Key")
	}

	return ecpub, nil
}

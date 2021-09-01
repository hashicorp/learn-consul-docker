package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"

	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/mitchellh/cli"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func mintFactory() (cli.Command, error) {
	c := &mintCommand{
		ui: ui(),
	}

	c.init()
	return c, nil
}

type mintCommand struct {
	ui cli.Ui

	flags *flag.FlagSet

	privKeyFile string
	ttl         time.Duration
	typ         string
	nodeName    string
	server      bool
	issuer      string
	audience    string

	help string
}

const mintHelp = `Usage: secint mint [options]

Create a new intro token.
`

func (c *mintCommand) Help() string {
	return Usage(c.help, nil)
}

func (c *mintCommand) Synopsis() string {
	return `Create a new intro token.`
}

func (c *mintCommand) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)

	c.flags.StringVar(&c.privKeyFile, "priv-key", "secint-priv-key.pem", "private key file in PEM format")
	c.flags.DurationVar(&c.ttl, "ttl", 30*time.Second, "duration the generated token is valid for")
	c.flags.BoolVar(&c.server, "server", false, "whether this token is for a server.")
	c.flags.StringVar(&c.nodeName, "node", "", "the name of the node")
	c.flags.StringVar(&c.issuer, "issuer", "", "the iss claim to embed in the JWT")
	c.flags.StringVar(&c.audience, "audience", "", "the aud claim to embed in the JWT")
	c.help = Usage(mintHelp, c.flags)
}

func (c *mintCommand) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		return 1
	}

	if c.nodeName == "" {
		c.ui.Error("ERROR: node name must be set")
		return 1
	}
	if c.ttl > 24*time.Hour {
		c.ui.Error("ERROR: duration can't be more than 1 hour")
		return 1
	}
	if c.ttl < time.Second {
		c.ui.Error("ERROR: duration can't be less than 1 second")
		return 1
	}

	// Load the private key
	pk, err := privKeyFromFile(c.privKeyFile)
	if err != nil {
		c.ui.Error(fmt.Sprintf("ERROR: failed to read private key file: %s", err))
		return 1
	}

	var opts jose.SignerOptions
	opts.WithType("JWT")

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.ES256, Key: pk}, &opts)
	if err != nil {
		c.ui.Error(fmt.Sprintf("Error: failed to create JWT signer: %v", err))
		return 1
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		c.ui.Error(fmt.Sprintf("ERROR: failed to generate UUID: %s", err))
		return 1
	}

	now := time.Now()

	stdClaims := jwt.Claims{
		Subject:   c.nodeName,
		Issuer:    c.issuer,
		Audience:  jwt.Audience{c.audience},
		NotBefore: jwt.NewNumericDate(now.Add(-60 * time.Second)),
		Expiry:    jwt.NewNumericDate(now.Add(c.ttl)),
		ID:        id,
	}

	privateClaims := make(map[string]interface{})

	if c.server {
		privateClaims["server"] = true
	}

	token, err := jwt.Signed(sig).
		Claims(stdClaims).
		Claims(privateClaims).
		CompactSerialize()
	if err != nil {
		c.ui.Error(fmt.Sprintf("ERROR: failed to create token: %v", err))
		return 1
	}

	c.ui.Output(string(token))

	return 0
}

func privKeyFromFile(fileName string) (*ecdsa.PrivateKey, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bs)
	if block == nil {
		return nil, fmt.Errorf("No PEM block in file")
	}

	return x509.ParseECPrivateKey(block.Bytes)
}

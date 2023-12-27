package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Networking Config defines the structure of the configuration file.
type Networking struct {
	Values struct {
		Name string `toml:"vpc_name"`
	} `toml:"networking"`
}

func cloudNetworking(ctx *pulumi.Context) error {

	// Setup config
	var config Networking

	// Read Config.toml
	if _, err := toml.DecodeFile("./Config.toml", &config); err != nil {
		return fmt.Errorf("unable to parse config.toml file: %w", err)
	}

	network, err := compute.NewNetwork(ctx, config.Values.Name, &compute.NetworkArgs{
		Name:                  pulumi.String(config.Values.Name),
		AutoCreateSubnetworks: pulumi.Bool(true), // todo!("Lets create the subnets manually next time we are here")
	})

	if err != nil {
		return err
	}

	ctx.Export("networkId", network.ID())
	return nil
}

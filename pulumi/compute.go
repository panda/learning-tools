package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Compute Config defines the structure of the configuration file.
type Compute struct {
	Values struct {
		Name        string `toml:"instance_name"`
		MachineType string `toml:"machine_type"`
		Zone        string `toml:"zone"`
	} `toml:"compute"`
}

func cloudComputeInstance(ctx *pulumi.Context) error {

	// Setup config
	var config Compute

	// Read Config.toml
	if _, err := toml.DecodeFile("./Config.toml", &config); err != nil {
		return fmt.Errorf("unable to parse config.toml file: %w", err)
	}

	// Create a new instance resource.
	instance, err := compute.NewInstance(ctx, config.Values.Name, &compute.InstanceArgs{
		MachineType: pulumi.String(config.Values.MachineType),
		Zone:        pulumi.String(config.Values.Zone),
		BootDisk: &compute.InstanceBootDiskArgs{
			InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
				Image: pulumi.String("debian-cloud/debian-11"), // panda: can probably make this configurable later on
				Labels: pulumi.Map{
					"label": pulumi.Any("steve-pulumi-dev"),
				},
			},
		},
		NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
			&compute.InstanceNetworkInterfaceArgs{
				Network: pulumi.String("default"), // panda: default for now till we figure out the vpc situation
			},
		},
	})
	if err != nil {
		return err
	}

	// Export the instance name.
	ctx.Export("instanceName", instance.Name)

	return nil
}

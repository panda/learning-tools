package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new VPC
		err := cloudNetworking(ctx)
		if err != nil {
			return err
		}

		// Create a new storage bucket
		err = cloudBucket(ctx)
		if err != nil {
			return err
		}

		// Create a new compute instance
		err = cloudComputeInstance(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}

package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Storage Config defines the structure of the configuration file.
type Storage struct {
	Values struct {
		BucketName               string `toml:"bucket_name"`
		Region                   string `toml:"region"`
		Versioning               bool   `toml:"versioning"`
		UniformBucketLevelAccess bool   `toml:"uniform_bucket_level_access"`
	} `toml:"storage"`
}

func cloudBucket(ctx *pulumi.Context) error {

	// Setup config
	var config Storage

	// Read Config.toml
	if _, err := toml.DecodeFile("./Config.toml", &config); err != nil {
		return fmt.Errorf("unable to parse config.toml file: %w", err)
	}

	bucket, err := storage.NewBucket(ctx, config.Values.BucketName, &storage.BucketArgs{
		Name:         pulumi.String(config.Values.BucketName),
		Location:     pulumi.StringInput(pulumi.String(config.Values.Region)), // panda: meh for now
		ForceDestroy: pulumi.Bool(true),
		Versioning: &storage.BucketVersioningArgs{
			Enabled: pulumi.Bool(config.Values.Versioning),
		},

		UniformBucketLevelAccess: pulumi.Bool(config.Values.UniformBucketLevelAccess),
	})

	if err != nil {
		return fmt.Errorf("unable to create storage bucket: %w", err)
	}

	_, err = storage.NewBucketObject(ctx, "files/index.html", &storage.BucketObjectArgs{
		Bucket: bucket.Name,
		Source: pulumi.NewFileAsset("files/index.html"),
	})

	if err != nil {
		return err
	}

	_, err = storage.NewBucketIAMBinding(ctx, config.Values.BucketName+"-binding", &storage.BucketIAMBindingArgs{
		Bucket: bucket.Name,
		Role:   pulumi.String("roles/storage.objectViewer"),
		Members: pulumi.StringArray{
			pulumi.String("allUsers"),
		},
	})
	if err != nil {
		return err
	}

	// lets export the name to console
	ctx.Export("bucketName", bucket.Url)
	ctx.Export("bucketURL", bucket.Url)

	return nil
}

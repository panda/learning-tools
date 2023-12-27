# Building Google Cloud infrastructure using Pulumi
## Introduction
This sub-folder contains a bare bones example of how to provision infrastructure on Google Cloud using Pulumi.
In this project we used the [Go](https://golang.org/) programming language to write our infrastructure as code.

We build out an object-store bucket, a compute instance and a new VPC network (this one not really used though).
## Prerequisites
- [Pulumi](https://www.pulumi.com/docs/get-started/install/)
- [Go](https://golang.org/doc/install)
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)

Following the prerequisites installation, you need to ensure your environment is configured properly to use your
Google Cloud project. 


## Instructions
1. Change into the `pulumi-gcp` directory.
2. Run `pulumi stack init` to initialize a new stack. 
   * I created an account on Pulumi.com and logged in to take advantage of the free state storage.
3. Run `pulumi config set gcp:project <your-project-id>` to set the GCP project ID.
4. Find the `Config.toml` file and adjust your configuration (pull-requests are welcome here)
5. Run `pulumi up` to preview and deploy the changes.
6. Run `pulumi destroy` to tear down the resources.

Keep in mind this is a bare bones example and you will need to adjust the configuration to suit your needs.
As we get more familiar with Pulumi, we can add more and more examples! Overall the Pulumi tool rocks!


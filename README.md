# paws

PAWS is an acronym for Pulumi on AWS. The idea is to create
a library of reusable Pulumi components for various popular
AWS resources.

This repository will use Go for this goal.

# Pulumi components

The Pulumi documentation describes the concept of reusable
components in its [documentation](https://www.pulumi.com/docs/iac/concepts/components/)

```Go
type MyComponent struct {
    pulumi.ResourceState
}

func NewMyComponent(ctx *pulumi.Context, name string, myComponentArgs MyComponentArgs, opts ...pulumi.ResourceOption) (*MyComponent, error) {
    myComponent := &MyComponent{}
    err := ctx.RegisterComponentResource("pkg:index:MyComponent", name, myComponent, opts...)
    if err != nil {
        return nil, err
    }
    return myComponent, nil
}
```

## Available Components

### VPC Component (`paws:vpc:Vpc`)

A reusable VPC component that provides a simplified interface for creating AWS VPCs with sensible defaults.

**Features:**
- IPv4 and IPv6 support
- Configurable DNS settings
- Instance tenancy options
- IPAM integration support
- Comprehensive output values for integration

**Basic Usage:**
```go
vpc, err := vpc.NewVpc(ctx, "my-vpc", &vpc.VpcArgs{
    CidrBlock: "10.0.0.0/16",
    Tags: map[string]string{
        "Name": "my-vpc",
        "Environment": "development",
    },
})
```

See the [VPC component documentation](vpc/README.md) for detailed usage examples.

## Additional resources

The above mentioned website suggests the following pages as additional
readings:

* [Build a Component](https://www.pulumi.com/docs/iac/using-pulumi/build-a-component/)
* [S3 folder Component](https://www.pulumi.com/registry/packages/aws/how-to-guides/s3-folder-component/)

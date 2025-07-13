# VPC Component

The VPC component (`paws:vpc:Vpc`) provides a simplified interface for creating AWS Virtual Private Clouds with sensible defaults and comprehensive configuration options.

## Features

- **IPv4 and IPv6 Support**: Configure both IPv4 CIDR blocks and optional IPv6 blocks
- **DNS Configuration**: Control DNS hostnames and support settings
- **Instance Tenancy**: Choose between default, dedicated, or host tenancy
- **IPAM Integration**: Support for AWS IP Address Manager (IPAM) pools
- **Comprehensive Outputs**: Access all VPC properties for further resource creation
- **Tagging Support**: Apply custom tags to the VPC resource

## Usage

### Basic Example

```go
package main

import (
    "github.com/Comradin/paws/vpc"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Create a basic VPC
        myVpc, err := vpc.NewVpc(ctx, "my-vpc", &vpc.VpcArgs{
            CidrBlock: "10.0.0.0/16",
            Tags: map[string]string{
                "Name":        "my-vpc",
                "Environment": "development",
            },
        })
        if err != nil {
            return err
        }

        // Export the VPC ID for use in other stacks
        ctx.Export("vpcId", myVpc.VpcId)
        return nil
    })
}
```

### Advanced Configuration

```go
func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        enableDnsHostnames := true
        instanceTenancy := "dedicated"
        
        customVpc, err := vpc.NewVpc(ctx, "custom-vpc", &vpc.VpcArgs{
            CidrBlock:          "172.16.0.0/12",
            EnableDnsHostnames: &enableDnsHostnames,
            EnableDnsSupport:   &enableDnsHostnames,
            InstanceTenancy:    &instanceTenancy,
            Tags: map[string]string{
                "Name":        "custom-vpc",
                "Environment": "production",
                "Team":        "infrastructure",
            },
        })
        if err != nil {
            return err
        }

        // Export multiple outputs
        ctx.Export("vpcId", customVpc.VpcId)
        ctx.Export("cidrBlock", customVpc.CidrBlock)
        ctx.Export("defaultRouteTableId", customVpc.DefaultRouteTableId)
        
        return nil
    })
}
```

### IPv6 Configuration

```go
func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        assignIpv6 := true
        
        ipv6Vpc, err := vpc.NewVpc(ctx, "ipv6-vpc", &vpc.VpcArgs{
            CidrBlock: "10.0.0.0/16",
            AssignGeneratedIpv6CidrBlock: &assignIpv6,
            Tags: map[string]string{
                "Name": "ipv6-enabled-vpc",
            },
        })
        if err != nil {
            return err
        }

        ctx.Export("vpcId", ipv6Vpc.VpcId)
        ctx.Export("ipv6CidrBlock", ipv6Vpc.Ipv6CidrBlock)
        
        return nil
    })
}
```

## Configuration Reference

### VpcArgs

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `CidrBlock` | `string` | Yes | The IPv4 CIDR block for the VPC |
| `EnableDnsHostnames` | `*bool` | No | Enable DNS hostnames in the VPC (default: true) |
| `EnableDnsSupport` | `*bool` | No | Enable DNS support in the VPC (default: true) |
| `InstanceTenancy` | `*string` | No | Tenancy option for instances ("default", "dedicated", "host") |
| `Ipv4IpamPoolId` | `*string` | No | ID of an IPv4 IPAM pool to allocate CIDR from |
| `Ipv4NetmaskLength` | `*int` | No | Netmask length for IPv4 CIDR allocated from IPAM pool |
| `Ipv6CidrBlock` | `*string` | No | IPv6 CIDR block for the VPC |
| `Ipv6IpamPoolId` | `*string` | No | ID of an IPv6 IPAM pool to allocate CIDR from |
| `Ipv6NetmaskLength` | `*int` | No | Netmask length for IPv6 CIDR allocated from IPAM pool |
| `AssignGeneratedIpv6CidrBlock` | `*bool` | No | Assign a generated IPv6 CIDR block |
| `Tags` | `map[string]string` | No | Tags to apply to the VPC |

### Outputs

The VPC component provides the following outputs:

| Output | Type | Description |
|--------|------|-------------|
| `VpcId` | `pulumi.StringOutput` | The ID of the VPC |
| `CidrBlock` | `pulumi.StringOutput` | The IPv4 CIDR block of the VPC |
| `Ipv6CidrBlock` | `pulumi.StringOutput` | The IPv6 CIDR block of the VPC |
| `DefaultRouteTableId` | `pulumi.StringOutput` | The ID of the default route table |
| `DefaultSecurityGroupId` | `pulumi.StringOutput` | The ID of the default security group |
| `MainRouteTableId` | `pulumi.StringOutput` | The ID of the main route table |
| `DhcpOptionsId` | `pulumi.StringOutput` | The ID of the DHCP options set |
| `Arn` | `pulumi.StringOutput` | The ARN of the VPC |
| `OwnerId` | `pulumi.StringOutput` | The ID of the AWS account that owns the VPC |

## Using VPC Outputs

The VPC component outputs can be used to create other AWS resources:

```go
// Create subnets in the VPC
subnet, err := ec2.NewSubnet(ctx, "public-subnet", &ec2.SubnetArgs{
    VpcId:     myVpc.VpcId,
    CidrBlock: pulumi.String("10.0.1.0/24"),
    Tags: pulumi.StringMap{
        "Name": pulumi.String("public-subnet"),
    },
})

// Create security groups
securityGroup, err := ec2.NewSecurityGroup(ctx, "web-sg", &ec2.SecurityGroupArgs{
    VpcId: myVpc.VpcId,
    Ingress: ec2.SecurityGroupIngressArray{
        &ec2.SecurityGroupIngressArgs{
            FromPort:   pulumi.Int(80),
            ToPort:     pulumi.Int(80),
            Protocol:   pulumi.String("tcp"),
            CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
        },
    },
})
```

## Error Handling

The component validates required arguments:

- `CidrBlock` is required and cannot be empty
- The `args` parameter cannot be nil

```go
// This will return an error
vpc, err := vpc.NewVpc(ctx, "invalid-vpc", &vpc.VpcArgs{
    // Missing required CidrBlock
})
if err != nil {
    // Handle error - "missing required argument 'CidrBlock'"
}
```
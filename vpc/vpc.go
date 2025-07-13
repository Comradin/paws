package vpc

import (
	"fmt"
	
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpcArgs struct {
	CidrBlock               string
	EnableDnsHostnames      *bool
	EnableDnsSupport        *bool
	InstanceTenancy         *string
	Ipv4IpamPoolId          *string
	Ipv4NetmaskLength       *int
	Ipv6CidrBlock           *string
	Ipv6IpamPoolId          *string
	Ipv6NetmaskLength       *int
	AssignGeneratedIpv6CidrBlock *bool
	Tags                    map[string]string
}

type Vpc struct {
	pulumi.ResourceState

	VpcId               pulumi.StringOutput `pulumi:"vpcId"`
	CidrBlock           pulumi.StringOutput `pulumi:"cidrBlock"`
	Ipv6CidrBlock       pulumi.StringOutput `pulumi:"ipv6CidrBlock"`
	DefaultRouteTableId pulumi.StringOutput `pulumi:"defaultRouteTableId"`
	DefaultSecurityGroupId pulumi.StringOutput `pulumi:"defaultSecurityGroupId"`
	MainRouteTableId    pulumi.StringOutput `pulumi:"mainRouteTableId"`
	DhcpOptionsId       pulumi.StringOutput `pulumi:"dhcpOptionsId"`
	Arn                 pulumi.StringOutput `pulumi:"arn"`
	OwnerId             pulumi.StringOutput `pulumi:"ownerId"`
}

func NewVpc(ctx *pulumi.Context, name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		return nil, fmt.Errorf("missing required argument 'args'")
	}
	if args.CidrBlock == "" {
		return nil, fmt.Errorf("missing required argument 'CidrBlock'")
	}

	vpc := &Vpc{}
	err := ctx.RegisterComponentResource("paws:vpc:Vpc", name, vpc, opts...)
	if err != nil {
		return nil, err
	}

	vpcArgs := &ec2.VpcArgs{
		CidrBlock:               pulumi.String(args.CidrBlock),
		EnableDnsHostnames:      pulumi.BoolPtr(true),
		EnableDnsSupport:        pulumi.BoolPtr(true),
		InstanceTenancy:         pulumi.StringPtr("default"),
		Tags:                    pulumi.ToStringMap(args.Tags),
	}

	if args.EnableDnsHostnames != nil {
		vpcArgs.EnableDnsHostnames = pulumi.BoolPtr(*args.EnableDnsHostnames)
	}
	if args.EnableDnsSupport != nil {
		vpcArgs.EnableDnsSupport = pulumi.BoolPtr(*args.EnableDnsSupport)
	}
	if args.InstanceTenancy != nil {
		vpcArgs.InstanceTenancy = pulumi.StringPtr(*args.InstanceTenancy)
	}
	if args.Ipv4IpamPoolId != nil {
		vpcArgs.Ipv4IpamPoolId = pulumi.StringPtr(*args.Ipv4IpamPoolId)
	}
	if args.Ipv4NetmaskLength != nil {
		vpcArgs.Ipv4NetmaskLength = pulumi.IntPtr(*args.Ipv4NetmaskLength)
	}
	if args.Ipv6CidrBlock != nil {
		vpcArgs.Ipv6CidrBlock = pulumi.StringPtr(*args.Ipv6CidrBlock)
	}
	if args.Ipv6IpamPoolId != nil {
		vpcArgs.Ipv6IpamPoolId = pulumi.StringPtr(*args.Ipv6IpamPoolId)
	}
	if args.Ipv6NetmaskLength != nil {
		vpcArgs.Ipv6NetmaskLength = pulumi.IntPtr(*args.Ipv6NetmaskLength)
	}
	if args.AssignGeneratedIpv6CidrBlock != nil {
		vpcArgs.AssignGeneratedIpv6CidrBlock = pulumi.BoolPtr(*args.AssignGeneratedIpv6CidrBlock)
	}

	awsVpc, err := ec2.NewVpc(ctx, name, vpcArgs, pulumi.Parent(vpc))
	if err != nil {
		return nil, err
	}

	vpc.VpcId = awsVpc.ID().ToStringOutput()
	vpc.CidrBlock = awsVpc.CidrBlock
	vpc.Ipv6CidrBlock = awsVpc.Ipv6CidrBlock
	vpc.DefaultRouteTableId = awsVpc.DefaultRouteTableId
	vpc.DefaultSecurityGroupId = awsVpc.DefaultSecurityGroupId
	vpc.MainRouteTableId = awsVpc.MainRouteTableId
	vpc.DhcpOptionsId = awsVpc.DhcpOptionsId
	vpc.Arn = awsVpc.Arn
	vpc.OwnerId = awsVpc.OwnerId

	ctx.RegisterResourceOutputs(vpc, pulumi.Map{
		"vpcId":                   vpc.VpcId,
		"cidrBlock":               vpc.CidrBlock,
		"ipv6CidrBlock":           vpc.Ipv6CidrBlock,
		"defaultRouteTableId":     vpc.DefaultRouteTableId,
		"defaultSecurityGroupId":  vpc.DefaultSecurityGroupId,
		"mainRouteTableId":        vpc.MainRouteTableId,
		"dhcpOptionsId":           vpc.DhcpOptionsId,
		"arn":                     vpc.Arn,
		"ownerId":                 vpc.OwnerId,
	})

	return vpc, nil
}
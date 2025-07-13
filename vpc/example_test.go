package vpc_test

import (
	"github.com/Comradin/paws/vpc"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ExampleNewVpc() {
	ctx := &pulumi.Context{}
	
	_, err := vpc.NewVpc(ctx, "example-vpc", &vpc.VpcArgs{
		CidrBlock: "10.0.0.0/16",
		Tags: map[string]string{
			"Name":        "example-vpc",
			"Environment": "development",
		},
	})
	
	if err != nil {
		panic(err)
	}
}

func ExampleNewVpc_withOptions() {
	ctx := &pulumi.Context{}
	
	enableDnsHostnames := true
	instanceTenancy := "dedicated"
	
	_, err := vpc.NewVpc(ctx, "custom-vpc", &vpc.VpcArgs{
		CidrBlock:          "172.16.0.0/12",
		EnableDnsHostnames: &enableDnsHostnames,
		InstanceTenancy:    &instanceTenancy,
		Tags: map[string]string{
			"Name":        "custom-vpc",
			"Environment": "production",
			"Team":        "infrastructure",
		},
	})
	
	if err != nil {
		panic(err)
	}
}
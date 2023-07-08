package nested_stack

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

func (s NestedStack) VpcStack() ec2.Vpc {

	vpcName := "test-cdk-go-v2-vpc"

	vpc := ec2.NewVpc(awscdk.NewNestedStack(s.Scope, s.Id, s.Env), jsii.String("Vpc"), &ec2.VpcProps{
		VpcName:                      jsii.String(vpcName),
		AvailabilityZones:            nil,
		DefaultInstanceTenancy:       "",
		EnableDnsHostnames:           nil,
		EnableDnsSupport:             nil,
		FlowLogs:                     nil,
		GatewayEndpoints:             nil,
		IpAddresses:                  nil,
		MaxAzs:                       nil,
		NatGatewayProvider:           nil,
		NatGateways:                  nil,
		NatGatewaySubnets:            nil,
		ReservedAzs:                  nil,
		RestrictDefaultSecurityGroup: nil,
		SubnetConfiguration:          nil,
		VpnConnections:               nil,
		VpnGateway:                   nil,
		VpnGatewayAsn:                nil,
		VpnRoutePropagation:          nil,
	})

	return vpc
}

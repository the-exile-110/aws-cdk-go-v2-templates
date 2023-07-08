package root_stack

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func tags() *map[string]*string {
	return &map[string]*string{
		"env": jsii.String("dev"),
	}
}

func env() *awscdk.Environment {

	return &awscdk.Environment{}

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}

type Root struct {
	Scope constructs.Construct
	Id    *string
	Env   *awscdk.StackProps
}

func CreateRootStack(scope constructs.Construct, id string) *Root {
	return &Root{
		Scope: scope,
		Id:    &id,
		Env:   &awscdk.StackProps{Env: env(), Tags: tags()},
	}
}

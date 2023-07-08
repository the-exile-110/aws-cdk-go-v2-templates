package root_stack

import (
	nestedStack "app/stacks/nested-stack"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (s Root) TestRootStack() *awscdk.Stack {
	rootStack := awscdk.NewStack(s.Scope, s.Id, s.Env)

	vpcStack := nestedStack.CreateNestedStack(rootStack, "VpcStack").VpcStack()

	fmt.Println("vpc:", vpcStack.VpcId())

	return &rootStack
}

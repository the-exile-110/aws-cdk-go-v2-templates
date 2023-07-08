package main

import (
	rootStack "app/stacks/root-stack"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func main() {
	app := awscdk.NewApp(nil)

	rootStack.CreateRootStack(app, "TestRootStack").TestRootStack()

	app.Synth(nil)
}

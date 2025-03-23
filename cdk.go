package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const CRON_EXPRESSION = "cron(15 * ? * MON-FRI *)"
const RUN_AT_TIME = "11:15"
const TARGET_TIMEZONE = "Europe/Stockholm"

type StackProps struct {
	awscdk.StackProps
}

func NewBentobotStack(scope constructs.Construct, id string, props *StackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)
	function := awslambda.NewFunction(stack, jsii.String("BentobotFunction"), &awslambda.FunctionProps{
		Architecture: awslambda.Architecture_ARM_64(),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./app/build/"), nil),
		Environment: &map[string]*string{
			"SLACK_WEBHOOK_URL": jsii.String(os.Getenv("SLACK_WEBHOOK_URL")),
			"RUN_AT_TIME":       jsii.String(RUN_AT_TIME),
			"TARGET_TIMEZONE":   jsii.String(TARGET_TIMEZONE),
		},
	})

	awsevents.NewRule(stack, jsii.String("BentobotScheduleRule"), &awsevents.RuleProps{
		Schedule: awsevents.Schedule_Expression(jsii.String(CRON_EXPRESSION)),
		Targets: &[]awsevents.IRuleTarget{
			awseventstargets.NewLambdaFunction(function, nil),
		},
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewBentobotStack(app, "BentobotStack", &StackProps{})

	app.Synth(nil)
}

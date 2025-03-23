# BentoBot
A Slack bot to automatically check what your colleagues are doing for lunch.

## Prerequisites
* A Slack workspace.
* An AWS account.

## Setup
1. Start by creating a [Slack app and setting up Incoming Webhooks](https://slack.com/intl/en-se/help/articles/115005265063-incoming-webhooks-for-slack).
2. Then, [bootstrap your AWS environment](https://docs.aws.amazon.com/cdk/v2/guide/bootstrapping-env.html).
3. [Setup a new policy](https://stackoverflow.com/questions/57118082/what-iam-permissions-are-needed-to-use-cdk-deploy/61102280#61102280) granting permission to assume CDK roles.
4. Add the following GitHub Actions secrets: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_DEFAULT_REGION`.
5. Configure a `SLACK_WEBHOOK_URL` environmental variable with your Slack Webhook URL.
5. Change any other values in `cdk.go` to fit your needs.
6. Profit!

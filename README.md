# BentoBot
A Slack bot to automatically check what your colleagues are doing for lunch.

## Prerequisites
* A Slack workspace.
* An AWS account.

## Setup
1. Start by creating a [Slack app and setting up Incoming Webhooks](https://slack.com/intl/en-se/help/articles/115005265063-incoming-webhooks-for-slack).
2. Configure your [AWS account to work with Serverless](https://serverless.com/framework/docs/providers/aws/guide/credentials/).
3. Configure a `SLACK_WEBHOOK_URL` environmental variable with your Slack Webhook URL. Change any other values in `serverless.yml` to fit your needs.
4. Profit!

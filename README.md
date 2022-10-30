# ting-task-nhk-easy [![Go](https://github.com/ting-app/ting-task-nhk-easy-lambda/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/ting-app/ting-task-nhk-easy-lambda/actions/workflows/build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/ting-app/ting-task-nhk-easy)](https://goreportcard.com/report/github.com/ting-app/ting-task-nhk-easy)
An AWS Lambda that saves [NHK NEWS WEB EASY](https://www3.nhk.or.jp/news/easy/) as ting.

## Getting Started
First create an AWS Lambda function, click `Enable VPC` under `Advanced settings`, and select the VPC of your MySQL instance. Remember to add the security group of your Lambda function as an inbound rule to your MySQL instance.

Second, follow the instructions [here](https://blog.theodo.com/2020/01/internet-access-to-lambda-in-vpc/) to enable the Lambda to have the access to Internet.

Then run `sh ./package.sh` to package the function, and upload `lambda.zip` to your Lambda function. Also add the following environment variables to the Lambda function:

 * DB_USER_NAME: user name of MySQL database
 * DB_PASSWORD: password of user
 * DB_HOST: host of MySQL database
 * DB_PORT: port of MySQL database

Finally, create an EventBridge rule to trigger the Lambda function daily:
1. Select `Schedule` as `Rule type`
2. This task relies on [nhk-easy-task](https://github.com/nhk-news-web-easy/nhk-easy-task) which runs at 10:00 AM (UTC) every day, so the cron expression should be set after that, for example `30 10 * * ? *`

## License
[MIT](LICENSE)

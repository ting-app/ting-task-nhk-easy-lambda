# ting-task-nhk-easy [![Go](https://github.com/ting-app/ting-task-nhk-easy/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/ting-app/ting-task-nhk-easy/actions/workflows/build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/ting-app/ting-task-nhk-easy)](https://goreportcard.com/report/github.com/ting-app/ting-task-nhk-easy)
An AWS Lambda that saves [NHK NEWS WEB EASY](https://www3.nhk.or.jp/news/easy/) as ting.

## Usage
Create an AWS Lambda function, and run `sh ./package.sh` to package the function, then upload `lambda.zip` to your Lambda function. Also add the following environment variables to the lambda function:

 * DB_USER_NAME: user name of MySQL database
 * DB_PASSWORD: password of user
 * DB_HOST: host of MySQL database
 * DB_PORT: port of MySQL database

## License
[MIT](LICENSE)

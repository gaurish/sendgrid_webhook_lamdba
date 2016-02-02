Sendgrid Event Webhook Lambda
=====

An AWS Lambda function written in Go which listens to Sendgrid's Event Webhook for email events. SendGrid’s Event Webhook will notify a URL of your choice via HTTP POST with information about events that occur as SendGrid processes your email but it generates a lot of events. the incoming data can easily overload your web server & bring down your production.

### Enter AWS Lambda & API Gateway
Lambda is a service by Amazon which allows you write functions which can be invoked via HTTP API using API Gateway. 

## Installation
```sh
$ go get github.com/gaurish/sendgrid_webhook_lamdba
```
The API is not yet stable. Please use a tool like [Godep](https://github.com/tools/godep) to vendor dependencies in your project.

## Features


- Serverless Architecture. Infinitly scalable. First 1 million requests per month are free
- Writes all the events from Sendgrid to AWS S3 as JSON files which can then later we loaded into database of your choice for any analytisis.
- Proxyies HTTP POSTs which have have unsubcribe in them
- Logs Events to CloudWatch, so you can search them for your operational needs(debugging missing emails etc)
- End to End Acceptance Tests which you can run locally on your development machine using `go test` command. Resulting in Faster development time. 


## Configuration 

### Environment Variables
The Following are the variable required for this application

- `PROXY_HOST_URL`: FULL URL to the HOST where you wan to post your unsubcribe events. 
- `AWS_S3_BUCKET`: the name of S3 bucket where we will store the sendgrid events as JSON files. 
- `AWS_ACCESS_KEY_ID` & `AWS_SECRET_ACCESS_KEY` Required for Authenticating with AWS & uploading Lambda.

### Depedantant Packages
Please install the following before using:-

- Apex: Install instructions given on this page https://github.com/apex/apex
- Official AWS SDK: https://github.com/aws/aws-sdk-go

### Project.json
In `project.json` please enter the role. Example

```javascript
{
  "name": "sendgrid_webhook_lamdba",
  "description": "Handle sendgrid events. two things. 1) proxy unsubcribe events to URL 2) Log all events to S3",
  "memory": 128,
  "runtime": "golang",
  "timeout": 10,
  "role": "arn:aws:iam::311249233107:role/lambda-execution-role"
}
```

## Examples

### Deploying
This is how you should deploy this lambda function.

```sh
$ apex deploy -e AWS_S3_BUCKET=myBucketName -e PROXY_HOST_URL=https://example.com/incoming_webhooks webhook
• deploying                 function=webhook
   • created build (10 MB)     function=webhook
   • updating function         function=webhook
   • updating alias            function=webhook
   • deployed                  function=webhook name=sendgrid_webhook_lamdba_webhook version=36
   • deploying config          function=webhook
```

## Testing By Running the Function
Apex provides an handy invoke features, you can use the same here. 

```sh
$ cat event.json | apex invoke --logs webhook
```

The above command will POS the contents to `event.json` file to lambda & execute the function. The logs of the execution will be available in the terminal itself. so it doesn't require you to leave the terminal. 


## Steps to contribute:

#### step 1:

Fork the repository and clone into your system - [github:help - fork a repo](https://help.github.com/articles/fork-a-repo)

#### step 2:

Do "Pull Request" - [github:help - Creating a Pull Request](https://help.github.com/articles/creating-a-pull-request)


## Running Tests
The tests written using using go's testing package & work against real-resources instead of mocks. so please be careful 

#### Test Uploading S3
```
$ cd /path/to/sendgrid_webhook_lambda
$ cd s3
$ go test
2016/02/01 17:26:06 [S3] Using File name ->  foo/2016/February/1/1454327766-738bd115-02ea-91c7-1538-09d53d163236.json
2016/02/01 17:26:09 [S3] {
  ETag: "\"ad74ac2799b6af79d108d28409f973db\""
}
PASS
ok  	github.com/gaurish/sendgrid_webhook_lambda/s3	2.267s
```

#### Test Proxing Unsubcribe Requests

```
$ cd /path/to/sendgrid_webhook_lambda
$ cd proxy
$ go test
2016/02/01 17:28:15 {unsubscribe}
PASS
ok  	github.com/gaurish/sendgrid_webhook_lambda/proxy	2.327s
```


## For any Questions/Feedback
- Please open an Issue




### License
####The MIT License (MIT)

Copyright (c) 2016 Gaurish Sharma

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


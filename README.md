# GO challenge

## API

> deploy api url ==> [https://tvwvnaqfy9.execute-api.us-west-2.amazonaws.com/api](https://tvwvnaqfy9.execute-api.us-west-2.amazonaws.com/api)

> click [here](https://tvwvnaqfy9.execute-api.us-west-2.amazonaws.com/api/swagger/index.html) to see swagger page.

## Folder structure of project

- **.github**: This folder contains of GitHub action configuration.
- **controller**: There are controller of router, and their integration tests in this file.
- **docs**: There are swagger files in this folder.
- **mocks**: There are mock interface of dependency for unit tests, that created by mockery.
- **models**: This package contains of whole models that use in project.
- **repository**: There are interface of dynamoDb in this folder.
- **routes**: There are functions that create router is their duty.
- **services**: There are services that use them in controller, and their unit tests.
- **main.go**: There is lamdad handler  in this file.
- **serverless.yml**: this file contains of configuration serverless framework that use in GitHub action for deploy project in AWS.

## environment variable
>This environment variable must be set in lambda configuration
- **TABLE_NAME** : name of dynamoDB table that data should save to it
- **ACCESS_TOKEN** : Access_token_id for access to aws service
- **SECRET_KEY** : Secret_key_access for access to aws service

:pushpin: The request and response that aws pass to lambda is different with the request and response that gorilla mux work with them. so 
we need convert aws req to golang req and use it in gorilla mux router, the end covert golang res to aws res. for this purpose 
we use [`awslabs`](https://github.com/awslabs/aws-lambda-go-api-proxy) package.

## Details
:black_nib: Lambda function defined in [`main`](https://github.com/dehghanpur10/go-challenge/blob/master/main.go) package, first gorilla mux created then set router
that we need them, then it passes as dependency to  gorillaMux adapter in [`awslabs`](https://github.com/awslabs/aws-lambda-go-api-proxy) package for convert req and res, and req within it proccessed
by gorilla mux router, finally return response.

:black_nib: In controller, first valid data input and connect to dynamoDB then as dependency pass to service for do job, finally
set status code according to different conditions, and set body of response.

:black_nib: In service, first prepare data for communication with the database. database is dependency of service because it can be tested later.

## Test

:pushpin: For controller, integration tests applied, and items that created in integration test delete at the end of test.

:pushpin: For services, unit tests applied, first we mock that dependency which use in service, and we isolated this service


## serverless
In serverless.yml file set config for apigetwaye and dynamoDB table and lambda function.github action was set to automated test, build and deploy to aws,too 

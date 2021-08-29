# GO challenge

## API

> api url ==> [https://tvwvnaqfy9.execute-api.us-west-2.amazonaws.com/api](https://tvwvnaqfy9.execute-api.us-west-2.amazonaws.com/api)

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
we use `awslabs` package.

:pushpin: For controller, integration tests applied, and items that created in integration test delete at the end of test.

:pushpin: For services, unit tests applied, first we mock that dependency which use in service, and we isolated this service


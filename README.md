# Async HTTP API Lambda Layer
This is a lambda layer package that allows you to return data from your http API before your `ASYNC_HANDLER` has finished executing.

This allows you to perform fire and forget http APIs similar to the old `X-Amz-Invocation-Type: Event` that is [enabled on v1 APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-lambda-integration-async.html).

## How to use
This repository is basically an example repo showing how to deploy and integrate with this async handler.

Each function in the `serverless.yml` file presents a different way to interact with  

## Layer details
This layer includes an external extension that is monitoring for an http event from the internal lambda function that forwards the event and context information.

Once the event and context information is received the extension will then execute the `ASYNC_HANDLER`

## Environment variables

| Name | Description | Required (Y/N) | Default |
|--|--|--|--|
| `ASYNC_HANDLER` | This is the path to your async handler. This is similar to the normal `handler` property in the function description in the `serverless.yml` file. | `Y` | N/A |
| `CUSTOM_RES` | If you want to use the built-in `proxy` handler but you would like to customize the http response you can do that via this variable. (Example on `customResponse` function 😉) | `Y` | N/A |
# go-azure-function

Simple example of an azure function written in gloang


## Prerequisites

* Existing subscription 
* Login with `az login`
* Setting the correct subscription with `az account set --subscription <SUBS-ID>`
* Install go compiler https://golang.org/doc/install
* Install azure function toolkit for Linux, Mac or Windows
  https://docs.microsoft.com/en-us/azure/azure-functions/functions-run-local?tabs=linux%2Ccsharp%2Cbash

## Initialize Azure Subscription

Run `./init.sh` for creating
* A Resource Group
* A Storage Account
* The Function App

Please Note that the function app name and the name of the storage account needs to be
globally unique. Simply use a random number as suffix and use the same name in the
script publish.sh

## Testing locally

Simple start ./testrun.sh

This compiles and start the function locally using the `func` command. 
Please note that the go build inside creates a binary for your OS!

Test the function with `curl  http://localhost:7071/api/hello`

## Deploy and test

Build and deploy the function to azure with `./publish.sh`
This compiles the code to linux-architecture and deploys the function to azure

Test the function with the url that is provided as output of the publish tool

e.g. https://functionapp<RANDOMNUMBER>.azurewebsites.net/api/hello

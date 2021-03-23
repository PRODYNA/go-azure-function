#!/bin/bash

# use $RANDOM instead of hard coded values
RESOURCE_GROUP_NAME=rgfunction
STORAGE_ACCOUNT_NAME=safunction1234
FUNCTION_APP_NAME=functionapp3328


az group create --name $RESOURCE_GROUP_NAME --location westeurope

az storage account create --name $STORAGE_ACCOUNT_NAME --resource-group $RESOURCE_GROUP_NAME --location westeurope

az functionapp create --name $FUNCTION_APP_NAME --storage-account $STORAGE_ACCOUNT_NAME --consumption-plan-location westeurope --resource-group $RESOURCE_GROUP_NAME --runtime custom --os-type Linux --functions-version 3



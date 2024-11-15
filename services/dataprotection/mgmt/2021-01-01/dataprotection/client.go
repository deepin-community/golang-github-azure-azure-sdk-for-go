// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
//
// Package dataprotection implements the Azure ARM Dataprotection service API version 2021-01-01.
//
// Open API 2.0 Specs for Azure Data Protection service
package dataprotection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Dataprotection
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Dataprotection.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckFeatureSupport sends the check feature support request.
// Parameters:
// parameters - feature support request object
func (client BaseClient) CheckFeatureSupport(ctx context.Context, location string, parameters BasicFeatureValidationRequestBase) (result FeatureValidationResponseBaseModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckFeatureSupport")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckFeatureSupportPreparer(ctx, location, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BaseClient", "CheckFeatureSupport", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckFeatureSupportSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.BaseClient", "CheckFeatureSupport", resp, "Failure sending request")
		return
	}

	result, err = client.CheckFeatureSupportResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BaseClient", "CheckFeatureSupport", resp, "Failure responding to request")
		return
	}

	return
}

// CheckFeatureSupportPreparer prepares the CheckFeatureSupport request.
func (client BaseClient) CheckFeatureSupportPreparer(ctx context.Context, location string, parameters BasicFeatureValidationRequestBase) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataProtection/locations/{location}/checkFeatureSupport", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckFeatureSupportSender sends the CheckFeatureSupport request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckFeatureSupportSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckFeatureSupportResponder handles the response to the CheckFeatureSupport request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckFeatureSupportResponder(resp *http.Response) (result FeatureValidationResponseBaseModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOperationResultPatch sends the get operation result patch request.
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
func (client BaseClient) GetOperationResultPatch(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result BackupVaultResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetOperationResultPatch")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOperationResultPatchPreparer(ctx, vaultName, resourceGroupName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BaseClient", "GetOperationResultPatch", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOperationResultPatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.BaseClient", "GetOperationResultPatch", resp, "Failure sending request")
		return
	}

	result, err = client.GetOperationResultPatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BaseClient", "GetOperationResultPatch", resp, "Failure responding to request")
		return
	}

	return
}

// GetOperationResultPatchPreparer prepares the GetOperationResultPatch request.
func (client BaseClient) GetOperationResultPatchPreparer(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/operationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOperationResultPatchSender sends the GetOperationResultPatch request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetOperationResultPatchSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetOperationResultPatchResponder handles the response to the GetOperationResultPatch request. The method always
// closes the http.Response Body.
func (client BaseClient) GetOperationResultPatchResponder(resp *http.Response) (result BackupVaultResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOperationStatus sends the get operation status request.
func (client BaseClient) GetOperationStatus(ctx context.Context, location string, operationID string) (result OperationResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetOperationStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOperationStatusPreparer(ctx, location, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BaseClient", "GetOperationStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOperationStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.BaseClient", "GetOperationStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetOperationStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BaseClient", "GetOperationStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetOperationStatusPreparer prepares the GetOperationStatus request.
func (client BaseClient) GetOperationStatusPreparer(ctx context.Context, location string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataProtection/locations/{location}/operationStatus/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOperationStatusSender sends the GetOperationStatus request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetOperationStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetOperationStatusResponder handles the response to the GetOperationStatus request. The method always
// closes the http.Response Body.
func (client BaseClient) GetOperationStatusResponder(resp *http.Response) (result OperationResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

package maintenance

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

// ApplyUpdatesClient is the azure Maintenance Management Client
type ApplyUpdatesClient struct {
	BaseClient
}

// NewApplyUpdatesClient creates an instance of the ApplyUpdatesClient client.
func NewApplyUpdatesClient(subscriptionID string) ApplyUpdatesClient {
	return NewApplyUpdatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplyUpdatesClientWithBaseURI creates an instance of the ApplyUpdatesClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewApplyUpdatesClientWithBaseURI(baseURI string, subscriptionID string) ApplyUpdatesClient {
	return ApplyUpdatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate apply maintenance updates to resource
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
func (client ApplyUpdatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (result ApplyUpdate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplyUpdatesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, providerName, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplyUpdatesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplyUpdatesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ApplyUpdatesClient) CreateOrUpdateResponder(resp *http.Response) (result ApplyUpdate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateParent apply maintenance updates to resource with parent
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceParentType - resource parent type
// resourceParentName - resource parent identifier
// resourceType - resource type
// resourceName - resource identifier
func (client ApplyUpdatesClient) CreateOrUpdateParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (result ApplyUpdate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplyUpdatesClient.CreateOrUpdateParent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdateParentPreparer(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdateParent", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateParentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdateParent", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateParentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdateParent", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdateParentPreparer prepares the CreateOrUpdateParent request.
func (client ApplyUpdatesClient) CreateOrUpdateParentPreparer(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerName":       autorest.Encode("path", providerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"resourceName":       autorest.Encode("path", resourceName),
		"resourceParentName": autorest.Encode("path", resourceParentName),
		"resourceParentType": autorest.Encode("path", resourceParentType),
		"resourceType":       autorest.Encode("path", resourceType),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateParentSender sends the CreateOrUpdateParent request. The method will close the
// http.Response Body if it receives an error.
func (client ApplyUpdatesClient) CreateOrUpdateParentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateParentResponder handles the response to the CreateOrUpdateParent request. The method always
// closes the http.Response Body.
func (client ApplyUpdatesClient) CreateOrUpdateParentResponder(resp *http.Response) (result ApplyUpdate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get track maintenance updates to resource
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
// applyUpdateName - applyUpdate Id
func (client ApplyUpdatesClient) Get(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (result ApplyUpdate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplyUpdatesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, providerName, resourceType, resourceName, applyUpdateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplyUpdatesClient) GetPreparer(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applyUpdateName":   autorest.Encode("path", applyUpdateName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/{applyUpdateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplyUpdatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplyUpdatesClient) GetResponder(resp *http.Response) (result ApplyUpdate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetParent track maintenance updates to resource with parent
// Parameters:
// resourceGroupName - resource group name
// resourceParentType - resource parent type
// resourceParentName - resource parent identifier
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
// applyUpdateName - applyUpdate Id
func (client ApplyUpdatesClient) GetParent(ctx context.Context, resourceGroupName string, resourceParentType string, resourceParentName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (result ApplyUpdate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplyUpdatesClient.GetParent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetParentPreparer(ctx, resourceGroupName, resourceParentType, resourceParentName, providerName, resourceType, resourceName, applyUpdateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "GetParent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetParentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "GetParent", resp, "Failure sending request")
		return
	}

	result, err = client.GetParentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "GetParent", resp, "Failure responding to request")
		return
	}

	return
}

// GetParentPreparer prepares the GetParent request.
func (client ApplyUpdatesClient) GetParentPreparer(ctx context.Context, resourceGroupName string, resourceParentType string, resourceParentName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applyUpdateName":    autorest.Encode("path", applyUpdateName),
		"providerName":       autorest.Encode("path", providerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"resourceName":       autorest.Encode("path", resourceName),
		"resourceParentName": autorest.Encode("path", resourceParentName),
		"resourceParentType": autorest.Encode("path", resourceParentType),
		"resourceType":       autorest.Encode("path", resourceType),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/{applyUpdateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetParentSender sends the GetParent request. The method will close the
// http.Response Body if it receives an error.
func (client ApplyUpdatesClient) GetParentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetParentResponder handles the response to the GetParent request. The method always
// closes the http.Response Body.
func (client ApplyUpdatesClient) GetParentResponder(resp *http.Response) (result ApplyUpdate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
func (client ApplyUpdatesClient) List(ctx context.Context) (result ListApplyUpdate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplyUpdatesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplyUpdatesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/applyUpdates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplyUpdatesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplyUpdatesClient) ListResponder(resp *http.Response) (result ListApplyUpdate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

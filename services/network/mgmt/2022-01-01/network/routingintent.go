package network

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

// RoutingIntentClient is the network Client
type RoutingIntentClient struct {
	BaseClient
}

// NewRoutingIntentClient creates an instance of the RoutingIntentClient client.
func NewRoutingIntentClient(subscriptionID string) RoutingIntentClient {
	return NewRoutingIntentClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoutingIntentClientWithBaseURI creates an instance of the RoutingIntentClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRoutingIntentClientWithBaseURI(baseURI string, subscriptionID string) RoutingIntentClient {
	return RoutingIntentClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a RoutingIntent resource if it doesn't exist else updates the existing RoutingIntent.
// Parameters:
// resourceGroupName - the resource group name of the RoutingIntent.
// virtualHubName - the name of the VirtualHub.
// routingIntentName - the name of the per VirtualHub singleton Routing Intent resource.
// routingIntentParameters - parameters supplied to create or update RoutingIntent.
func (client RoutingIntentClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, routingIntentParameters RoutingIntent) (result RoutingIntentCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoutingIntentClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualHubName, routingIntentName, routingIntentParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RoutingIntentClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string, routingIntentParameters RoutingIntent) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routingIntentName": autorest.Encode("path", routingIntentName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	routingIntentParameters.Etag = nil
	routingIntentParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routingIntent/{routingIntentName}", pathParameters),
		autorest.WithJSON(routingIntentParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RoutingIntentClient) CreateOrUpdateSender(req *http.Request) (future RoutingIntentCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RoutingIntentClient) CreateOrUpdateResponder(resp *http.Response) (result RoutingIntent, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a RoutingIntent.
// Parameters:
// resourceGroupName - the resource group name of the RoutingIntent.
// virtualHubName - the name of the VirtualHub.
// routingIntentName - the name of the RoutingIntent.
func (client RoutingIntentClient) Delete(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string) (result RoutingIntentDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoutingIntentClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualHubName, routingIntentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RoutingIntentClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routingIntentName": autorest.Encode("path", routingIntentName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routingIntent/{routingIntentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RoutingIntentClient) DeleteSender(req *http.Request) (future RoutingIntentDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RoutingIntentClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a RoutingIntent.
// Parameters:
// resourceGroupName - the resource group name of the RoutingIntent.
// virtualHubName - the name of the VirtualHub.
// routingIntentName - the name of the RoutingIntent.
func (client RoutingIntentClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string) (result RoutingIntent, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoutingIntentClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualHubName, routingIntentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoutingIntentClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualHubName string, routingIntentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routingIntentName": autorest.Encode("path", routingIntentName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routingIntent/{routingIntentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoutingIntentClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoutingIntentClient) GetResponder(resp *http.Response) (result RoutingIntent, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieves the details of all RoutingIntent child resources of the VirtualHub.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
func (client RoutingIntentClient) List(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListRoutingIntentResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoutingIntentClient.List")
		defer func() {
			sc := -1
			if result.lrir.Response.Response != nil {
				sc = result.lrir.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualHubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lrir.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "List", resp, "Failure sending request")
		return
	}

	result.lrir, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lrir.hasNextLink() && result.lrir.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RoutingIntentClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routingIntent", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RoutingIntentClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RoutingIntentClient) ListResponder(resp *http.Response) (result ListRoutingIntentResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RoutingIntentClient) listNextResults(ctx context.Context, lastResults ListRoutingIntentResult) (result ListRoutingIntentResult, err error) {
	req, err := lastResults.listRoutingIntentResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.RoutingIntentClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.RoutingIntentClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RoutingIntentClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoutingIntentClient) ListComplete(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListRoutingIntentResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoutingIntentClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, virtualHubName)
	return
}

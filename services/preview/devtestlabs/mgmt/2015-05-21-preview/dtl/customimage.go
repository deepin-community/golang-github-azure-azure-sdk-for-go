package dtl

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

// CustomImageClient is the azure DevTest Labs REST API version 2015-05-21-preview.
type CustomImageClient struct {
	BaseClient
}

// NewCustomImageClient creates an instance of the CustomImageClient client.
func NewCustomImageClient(subscriptionID string) CustomImageClient {
	return NewCustomImageClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCustomImageClientWithBaseURI creates an instance of the CustomImageClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCustomImageClientWithBaseURI(baseURI string, subscriptionID string) CustomImageClient {
	return CustomImageClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateResource create or replace an existing custom image. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the custom image.
func (client CustomImageClient) CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, name string, customImage CustomImage) (result CustomImageCreateOrUpdateResourceFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomImageClient.CreateOrUpdateResource")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdateResourcePreparer(ctx, resourceGroupName, labName, name, customImage)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "CreateOrUpdateResource", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateResourceSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "CreateOrUpdateResource", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateResourcePreparer prepares the CreateOrUpdateResource request.
func (client CustomImageClient) CreateOrUpdateResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string, customImage CustomImage) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", pathParameters),
		autorest.WithJSON(customImage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateResourceSender sends the CreateOrUpdateResource request. The method will close the
// http.Response Body if it receives an error.
func (client CustomImageClient) CreateOrUpdateResourceSender(req *http.Request) (future CustomImageCreateOrUpdateResourceFuture, err error) {
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

// CreateOrUpdateResourceResponder handles the response to the CreateOrUpdateResource request. The method always
// closes the http.Response Body.
func (client CustomImageClient) CreateOrUpdateResourceResponder(resp *http.Response) (result CustomImage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteResource delete custom image. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the custom image.
func (client CustomImageClient) DeleteResource(ctx context.Context, resourceGroupName string, labName string, name string) (result CustomImageDeleteResourceFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomImageClient.DeleteResource")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteResourcePreparer(ctx, resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "DeleteResource", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteResourceSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "DeleteResource", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeleteResourcePreparer prepares the DeleteResource request.
func (client CustomImageClient) DeleteResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteResourceSender sends the DeleteResource request. The method will close the
// http.Response Body if it receives an error.
func (client CustomImageClient) DeleteResourceSender(req *http.Request) (future CustomImageDeleteResourceFuture, err error) {
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

// DeleteResourceResponder handles the response to the DeleteResource request. The method always
// closes the http.Response Body.
func (client CustomImageClient) DeleteResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetResource get custom image.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the custom image.
func (client CustomImageClient) GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result CustomImage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomImageClient.GetResource")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetResourcePreparer(ctx, resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "GetResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "GetResource", resp, "Failure sending request")
		return
	}

	result, err = client.GetResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "GetResource", resp, "Failure responding to request")
		return
	}

	return
}

// GetResourcePreparer prepares the GetResource request.
func (client CustomImageClient) GetResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetResourceSender sends the GetResource request. The method will close the
// http.Response Body if it receives an error.
func (client CustomImageClient) GetResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResourceResponder handles the response to the GetResource request. The method always
// closes the http.Response Body.
func (client CustomImageClient) GetResourceResponder(resp *http.Response) (result CustomImage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list custom images.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// filter - the filter to apply on the operation.
func (client CustomImageClient) List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationCustomImagePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomImageClient.List")
		defer func() {
			sc := -1
			if result.rwcci.Response.Response != nil {
				sc = result.rwcci.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, filter, top, orderBy)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwcci.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "List", resp, "Failure sending request")
		return
	}

	result.rwcci, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rwcci.hasNextLink() && result.rwcci.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CustomImageClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderBy"] = autorest.Encode("query", orderBy)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CustomImageClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CustomImageClient) ListResponder(resp *http.Response) (result ResponseWithContinuationCustomImage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CustomImageClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationCustomImage) (result ResponseWithContinuationCustomImage, err error) {
	req, err := lastResults.responseWithContinuationCustomImagePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.CustomImageClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.CustomImageClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.CustomImageClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomImageClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationCustomImageIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomImageClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labName, filter, top, orderBy)
	return
}
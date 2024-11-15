package security

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CustomAssessmentAutomationsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type CustomAssessmentAutomationsClient struct {
	BaseClient
}

// NewCustomAssessmentAutomationsClient creates an instance of the CustomAssessmentAutomationsClient client.
func NewCustomAssessmentAutomationsClient(subscriptionID string) CustomAssessmentAutomationsClient {
	return NewCustomAssessmentAutomationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCustomAssessmentAutomationsClientWithBaseURI creates an instance of the CustomAssessmentAutomationsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewCustomAssessmentAutomationsClientWithBaseURI(baseURI string, subscriptionID string) CustomAssessmentAutomationsClient {
	return CustomAssessmentAutomationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates or updates a custom assessment automation for the provided subscription. Please note that providing
// an existing custom assessment automation will replace the existing record.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// customAssessmentAutomationName - name of the Custom Assessment Automation.
// customAssessmentAutomationBody - custom Assessment Automation body
func (client CustomAssessmentAutomationsClient) Create(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, customAssessmentAutomationBody CustomAssessmentAutomationRequest) (result CustomAssessmentAutomation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomAssessmentAutomationsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.CustomAssessmentAutomationsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, customAssessmentAutomationName, customAssessmentAutomationBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client CustomAssessmentAutomationsClient) CreatePreparer(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, customAssessmentAutomationBody CustomAssessmentAutomationRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customAssessmentAutomationName": autorest.Encode("path", customAssessmentAutomationName),
		"resourceGroupName":              autorest.Encode("path", resourceGroupName),
		"subscriptionId":                 autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customAssessmentAutomations/{customAssessmentAutomationName}", pathParameters),
		autorest.WithJSON(customAssessmentAutomationBody),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client CustomAssessmentAutomationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client CustomAssessmentAutomationsClient) CreateResponder(resp *http.Response) (result CustomAssessmentAutomation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a custom assessment automation by name for a provided subscription
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// customAssessmentAutomationName - name of the Custom Assessment Automation.
func (client CustomAssessmentAutomationsClient) Delete(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomAssessmentAutomationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.CustomAssessmentAutomationsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, customAssessmentAutomationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CustomAssessmentAutomationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customAssessmentAutomationName": autorest.Encode("path", customAssessmentAutomationName),
		"resourceGroupName":              autorest.Encode("path", resourceGroupName),
		"subscriptionId":                 autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customAssessmentAutomations/{customAssessmentAutomationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CustomAssessmentAutomationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CustomAssessmentAutomationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a single custom assessment automation by name for the provided subscription and resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// customAssessmentAutomationName - name of the Custom Assessment Automation.
func (client CustomAssessmentAutomationsClient) Get(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string) (result CustomAssessmentAutomation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomAssessmentAutomationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.CustomAssessmentAutomationsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, customAssessmentAutomationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CustomAssessmentAutomationsClient) GetPreparer(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customAssessmentAutomationName": autorest.Encode("path", customAssessmentAutomationName),
		"resourceGroupName":              autorest.Encode("path", resourceGroupName),
		"subscriptionId":                 autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customAssessmentAutomations/{customAssessmentAutomationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CustomAssessmentAutomationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CustomAssessmentAutomationsClient) GetResponder(resp *http.Response) (result CustomAssessmentAutomation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list custom assessment automations by provided subscription and resource group
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
func (client CustomAssessmentAutomationsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result CustomAssessmentAutomationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomAssessmentAutomationsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.caalr.Response.Response != nil {
				sc = result.caalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.CustomAssessmentAutomationsClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.caalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.caalr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.caalr.hasNextLink() && result.caalr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client CustomAssessmentAutomationsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customAssessmentAutomations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client CustomAssessmentAutomationsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client CustomAssessmentAutomationsClient) ListByResourceGroupResponder(resp *http.Response) (result CustomAssessmentAutomationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client CustomAssessmentAutomationsClient) listByResourceGroupNextResults(ctx context.Context, lastResults CustomAssessmentAutomationsListResult) (result CustomAssessmentAutomationsListResult, err error) {
	req, err := lastResults.customAssessmentAutomationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomAssessmentAutomationsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result CustomAssessmentAutomationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomAssessmentAutomationsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription list custom assessment automations by provided subscription
func (client CustomAssessmentAutomationsClient) ListBySubscription(ctx context.Context) (result CustomAssessmentAutomationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomAssessmentAutomationsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.caalr.Response.Response != nil {
				sc = result.caalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.CustomAssessmentAutomationsClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.caalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.caalr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.caalr.hasNextLink() && result.caalr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client CustomAssessmentAutomationsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/customAssessmentAutomations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client CustomAssessmentAutomationsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client CustomAssessmentAutomationsClient) ListBySubscriptionResponder(resp *http.Response) (result CustomAssessmentAutomationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client CustomAssessmentAutomationsClient) listBySubscriptionNextResults(ctx context.Context, lastResults CustomAssessmentAutomationsListResult) (result CustomAssessmentAutomationsListResult, err error) {
	req, err := lastResults.customAssessmentAutomationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.CustomAssessmentAutomationsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomAssessmentAutomationsClient) ListBySubscriptionComplete(ctx context.Context) (result CustomAssessmentAutomationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomAssessmentAutomationsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

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

// ConfigurationAssignmentsWithinSubscriptionClient is the azure Maintenance Management Client
type ConfigurationAssignmentsWithinSubscriptionClient struct {
	BaseClient
}

// NewConfigurationAssignmentsWithinSubscriptionClient creates an instance of the
// ConfigurationAssignmentsWithinSubscriptionClient client.
func NewConfigurationAssignmentsWithinSubscriptionClient(subscriptionID string) ConfigurationAssignmentsWithinSubscriptionClient {
	return NewConfigurationAssignmentsWithinSubscriptionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationAssignmentsWithinSubscriptionClientWithBaseURI creates an instance of the
// ConfigurationAssignmentsWithinSubscriptionClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewConfigurationAssignmentsWithinSubscriptionClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationAssignmentsWithinSubscriptionClient {
	return ConfigurationAssignmentsWithinSubscriptionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List sends the list request.
func (client ConfigurationAssignmentsWithinSubscriptionClient) List(ctx context.Context) (result ListConfigurationAssignmentsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationAssignmentsWithinSubscriptionClient.List")
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
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsWithinSubscriptionClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsWithinSubscriptionClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsWithinSubscriptionClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ConfigurationAssignmentsWithinSubscriptionClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/configurationAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationAssignmentsWithinSubscriptionClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ConfigurationAssignmentsWithinSubscriptionClient) ListResponder(resp *http.Response) (result ListConfigurationAssignmentsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

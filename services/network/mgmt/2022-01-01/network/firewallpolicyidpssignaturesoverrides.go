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

// FirewallPolicyIdpsSignaturesOverridesClient is the network Client
type FirewallPolicyIdpsSignaturesOverridesClient struct {
	BaseClient
}

// NewFirewallPolicyIdpsSignaturesOverridesClient creates an instance of the
// FirewallPolicyIdpsSignaturesOverridesClient client.
func NewFirewallPolicyIdpsSignaturesOverridesClient(subscriptionID string) FirewallPolicyIdpsSignaturesOverridesClient {
	return NewFirewallPolicyIdpsSignaturesOverridesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallPolicyIdpsSignaturesOverridesClientWithBaseURI creates an instance of the
// FirewallPolicyIdpsSignaturesOverridesClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFirewallPolicyIdpsSignaturesOverridesClientWithBaseURI(baseURI string, subscriptionID string) FirewallPolicyIdpsSignaturesOverridesClient {
	return FirewallPolicyIdpsSignaturesOverridesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get returns all signatures overrides for a specific policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// firewallPolicyName - the name of the Firewall Policy.
func (client FirewallPolicyIdpsSignaturesOverridesClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string) (result SignaturesOverrides, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyIdpsSignaturesOverridesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, firewallPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client FirewallPolicyIdpsSignaturesOverridesClient) GetPreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallPolicyName": autorest.Encode("path", firewallPolicyName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallPolicyIdpsSignaturesOverridesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FirewallPolicyIdpsSignaturesOverridesClient) GetResponder(resp *http.Response) (result SignaturesOverrides, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns all signatures overrides objects for a specific policy as a list containing a single value.
// Parameters:
// resourceGroupName - the name of the resource group.
// firewallPolicyName - the name of the Firewall Policy.
func (client FirewallPolicyIdpsSignaturesOverridesClient) List(ctx context.Context, resourceGroupName string, firewallPolicyName string) (result SignaturesOverridesList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyIdpsSignaturesOverridesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, firewallPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client FirewallPolicyIdpsSignaturesOverridesClient) ListPreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallPolicyName": autorest.Encode("path", firewallPolicyName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallPolicyIdpsSignaturesOverridesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FirewallPolicyIdpsSignaturesOverridesClient) ListResponder(resp *http.Response) (result SignaturesOverridesList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Patch will update the status of policy's signature overrides for IDPS
// Parameters:
// parameters - will contain all properties of the object to put
// resourceGroupName - the name of the resource group.
// firewallPolicyName - the name of the Firewall Policy.
func (client FirewallPolicyIdpsSignaturesOverridesClient) Patch(ctx context.Context, parameters SignaturesOverrides, resourceGroupName string, firewallPolicyName string) (result SignaturesOverrides, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyIdpsSignaturesOverridesClient.Patch")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PatchPreparer(ctx, parameters, resourceGroupName, firewallPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "Patch", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "Patch", resp, "Failure sending request")
		return
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "Patch", resp, "Failure responding to request")
		return
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client FirewallPolicyIdpsSignaturesOverridesClient) PatchPreparer(ctx context.Context, parameters SignaturesOverrides, resourceGroupName string, firewallPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallPolicyName": autorest.Encode("path", firewallPolicyName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides/default", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallPolicyIdpsSignaturesOverridesClient) PatchSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client FirewallPolicyIdpsSignaturesOverridesClient) PatchResponder(resp *http.Response) (result SignaturesOverrides, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put will override/create a new signature overrides for the policy's IDPS
// Parameters:
// parameters - will contain all properties of the object to put
// resourceGroupName - the name of the resource group.
// firewallPolicyName - the name of the Firewall Policy.
func (client FirewallPolicyIdpsSignaturesOverridesClient) Put(ctx context.Context, parameters SignaturesOverrides, resourceGroupName string, firewallPolicyName string) (result SignaturesOverrides, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyIdpsSignaturesOverridesClient.Put")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutPreparer(ctx, parameters, resourceGroupName, firewallPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "Put", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "Put", resp, "Failure sending request")
		return
	}

	result, err = client.PutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesOverridesClient", "Put", resp, "Failure responding to request")
		return
	}

	return
}

// PutPreparer prepares the Put request.
func (client FirewallPolicyIdpsSignaturesOverridesClient) PutPreparer(ctx context.Context, parameters SignaturesOverrides, resourceGroupName string, firewallPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallPolicyName": autorest.Encode("path", firewallPolicyName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides/default", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallPolicyIdpsSignaturesOverridesClient) PutSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client FirewallPolicyIdpsSignaturesOverridesClient) PutResponder(resp *http.Response) (result SignaturesOverrides, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

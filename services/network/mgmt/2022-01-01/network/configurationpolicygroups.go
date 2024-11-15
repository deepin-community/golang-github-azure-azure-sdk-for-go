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

// ConfigurationPolicyGroupsClient is the network Client
type ConfigurationPolicyGroupsClient struct {
	BaseClient
}

// NewConfigurationPolicyGroupsClient creates an instance of the ConfigurationPolicyGroupsClient client.
func NewConfigurationPolicyGroupsClient(subscriptionID string) ConfigurationPolicyGroupsClient {
	return NewConfigurationPolicyGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationPolicyGroupsClientWithBaseURI creates an instance of the ConfigurationPolicyGroupsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewConfigurationPolicyGroupsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationPolicyGroupsClient {
	return ConfigurationPolicyGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a ConfigurationPolicyGroup if it doesn't exist else updates the existing one.
// Parameters:
// resourceGroupName - the resource group name of the ConfigurationPolicyGroup.
// vpnServerConfigurationName - the name of the VpnServerConfiguration.
// configurationPolicyGroupName - the name of the ConfigurationPolicyGroup.
// vpnServerConfigurationPolicyGroupParameters - parameters supplied to create or update a
// VpnServerConfiguration PolicyGroup.
func (client ConfigurationPolicyGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, vpnServerConfigurationPolicyGroupParameters VpnServerConfigurationPolicyGroup) (result ConfigurationPolicyGroupsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationPolicyGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, vpnServerConfigurationName, configurationPolicyGroupName, vpnServerConfigurationPolicyGroupParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConfigurationPolicyGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string, vpnServerConfigurationPolicyGroupParameters VpnServerConfigurationPolicyGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationPolicyGroupName": autorest.Encode("path", configurationPolicyGroupName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
		"vpnServerConfigurationName":   autorest.Encode("path", vpnServerConfigurationName),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	vpnServerConfigurationPolicyGroupParameters.Etag = nil
	vpnServerConfigurationPolicyGroupParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}/configurationPolicyGroups/{configurationPolicyGroupName}", pathParameters),
		autorest.WithJSON(vpnServerConfigurationPolicyGroupParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationPolicyGroupsClient) CreateOrUpdateSender(req *http.Request) (future ConfigurationPolicyGroupsCreateOrUpdateFuture, err error) {
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
func (client ConfigurationPolicyGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result VpnServerConfigurationPolicyGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a ConfigurationPolicyGroup.
// Parameters:
// resourceGroupName - the resource group name of the ConfigurationPolicyGroup.
// vpnServerConfigurationName - the name of the VpnServerConfiguration.
// configurationPolicyGroupName - the name of the ConfigurationPolicyGroup.
func (client ConfigurationPolicyGroupsClient) Delete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string) (result ConfigurationPolicyGroupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationPolicyGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, vpnServerConfigurationName, configurationPolicyGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConfigurationPolicyGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationPolicyGroupName": autorest.Encode("path", configurationPolicyGroupName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
		"vpnServerConfigurationName":   autorest.Encode("path", vpnServerConfigurationName),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}/configurationPolicyGroups/{configurationPolicyGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationPolicyGroupsClient) DeleteSender(req *http.Request) (future ConfigurationPolicyGroupsDeleteFuture, err error) {
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
func (client ConfigurationPolicyGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a ConfigurationPolicyGroup.
// Parameters:
// resourceGroupName - the resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - the name of the VpnServerConfiguration.
// configurationPolicyGroupName - the name of the ConfigurationPolicyGroup being retrieved.
func (client ConfigurationPolicyGroupsClient) Get(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string) (result VpnServerConfigurationPolicyGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationPolicyGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, vpnServerConfigurationName, configurationPolicyGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigurationPolicyGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, configurationPolicyGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationPolicyGroupName": autorest.Encode("path", configurationPolicyGroupName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
		"vpnServerConfigurationName":   autorest.Encode("path", vpnServerConfigurationName),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}/configurationPolicyGroups/{configurationPolicyGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationPolicyGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigurationPolicyGroupsClient) GetResponder(resp *http.Response) (result VpnServerConfigurationPolicyGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByVpnServerConfiguration lists all the configurationPolicyGroups in a resource group for a
// vpnServerConfiguration.
// Parameters:
// resourceGroupName - the resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - the name of the VpnServerConfiguration.
func (client ConfigurationPolicyGroupsClient) ListByVpnServerConfiguration(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (result ListVpnServerConfigurationPolicyGroupsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationPolicyGroupsClient.ListByVpnServerConfiguration")
		defer func() {
			sc := -1
			if result.lvscpgr.Response.Response != nil {
				sc = result.lvscpgr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVpnServerConfigurationNextResults
	req, err := client.ListByVpnServerConfigurationPreparer(ctx, resourceGroupName, vpnServerConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "ListByVpnServerConfiguration", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVpnServerConfigurationSender(req)
	if err != nil {
		result.lvscpgr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "ListByVpnServerConfiguration", resp, "Failure sending request")
		return
	}

	result.lvscpgr, err = client.ListByVpnServerConfigurationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "ListByVpnServerConfiguration", resp, "Failure responding to request")
		return
	}
	if result.lvscpgr.hasNextLink() && result.lvscpgr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByVpnServerConfigurationPreparer prepares the ListByVpnServerConfiguration request.
func (client ConfigurationPolicyGroupsClient) ListByVpnServerConfigurationPreparer(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"vpnServerConfigurationName": autorest.Encode("path", vpnServerConfigurationName),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}/configurationPolicyGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVpnServerConfigurationSender sends the ListByVpnServerConfiguration request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationPolicyGroupsClient) ListByVpnServerConfigurationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVpnServerConfigurationResponder handles the response to the ListByVpnServerConfiguration request. The method always
// closes the http.Response Body.
func (client ConfigurationPolicyGroupsClient) ListByVpnServerConfigurationResponder(resp *http.Response) (result ListVpnServerConfigurationPolicyGroupsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVpnServerConfigurationNextResults retrieves the next set of results, if any.
func (client ConfigurationPolicyGroupsClient) listByVpnServerConfigurationNextResults(ctx context.Context, lastResults ListVpnServerConfigurationPolicyGroupsResult) (result ListVpnServerConfigurationPolicyGroupsResult, err error) {
	req, err := lastResults.listVpnServerConfigurationPolicyGroupsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "listByVpnServerConfigurationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVpnServerConfigurationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "listByVpnServerConfigurationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVpnServerConfigurationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ConfigurationPolicyGroupsClient", "listByVpnServerConfigurationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVpnServerConfigurationComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConfigurationPolicyGroupsClient) ListByVpnServerConfigurationComplete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (result ListVpnServerConfigurationPolicyGroupsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationPolicyGroupsClient.ListByVpnServerConfiguration")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVpnServerConfiguration(ctx, resourceGroupName, vpnServerConfigurationName)
	return
}

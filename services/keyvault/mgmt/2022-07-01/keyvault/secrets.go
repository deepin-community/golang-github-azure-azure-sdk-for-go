package keyvault

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

// SecretsClient is the the Azure management API provides a RESTful set of web services that interact with Azure Key
// Vault.
type SecretsClient struct {
	BaseClient
}

// NewSecretsClient creates an instance of the SecretsClient client.
func NewSecretsClient(subscriptionID string) SecretsClient {
	return NewSecretsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSecretsClientWithBaseURI creates an instance of the SecretsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSecretsClientWithBaseURI(baseURI string, subscriptionID string) SecretsClient {
	return SecretsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a secret in a key vault in the specified subscription.  NOTE: This API is intended
// for internal use in ARM deployments. Users should use the data-plane REST service for interaction with vault
// secrets.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the vault belongs.
// vaultName - name of the vault
// secretName - name of the secret
// parameters - parameters to create or update the secret
func (client SecretsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretCreateOrUpdateParameters) (result Secret, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: vaultName,
			Constraints: []validation.Constraint{{Target: "vaultName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{3,24}$`, Chain: nil}}},
		{TargetValue: secretName,
			Constraints: []validation.Constraint{{Target: "secretName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{1,127}$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("keyvault.SecretsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, vaultName, secretName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecretsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"secretName":        autorest.Encode("path", secretName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecretsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SecretsClient) CreateOrUpdateResponder(resp *http.Response) (result Secret, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the specified secret.  NOTE: This API is intended for internal use in ARM deployments. Users should use the
// data-plane REST service for interaction with vault secrets.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the vault belongs.
// vaultName - the name of the vault.
// secretName - the name of the secret.
func (client SecretsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, secretName string) (result Secret, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, vaultName, secretName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SecretsClient) GetPreparer(ctx context.Context, resourceGroupName string, vaultName string, secretName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"secretName":        autorest.Encode("path", secretName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SecretsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SecretsClient) GetResponder(resp *http.Response) (result Secret, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List operation gets information about the secrets in a vault.  NOTE: This API is intended for internal use
// in ARM deployments. Users should use the data-plane REST service for interaction with vault secrets.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the vault belongs.
// vaultName - the name of the vault.
// top - maximum number of results to return.
func (client SecretsClient) List(ctx context.Context, resourceGroupName string, vaultName string, top *int32) (result SecretListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretsClient.List")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, vaultName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "List", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.slr.hasNextLink() && result.slr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SecretsClient) ListPreparer(ctx context.Context, resourceGroupName string, vaultName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SecretsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SecretsClient) ListResponder(resp *http.Response) (result SecretListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SecretsClient) listNextResults(ctx context.Context, lastResults SecretListResult) (result SecretListResult, err error) {
	req, err := lastResults.secretListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "keyvault.SecretsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "keyvault.SecretsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SecretsClient) ListComplete(ctx context.Context, resourceGroupName string, vaultName string, top *int32) (result SecretListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, vaultName, top)
	return
}

// Update update a secret in the specified subscription.  NOTE: This API is intended for internal use in ARM
// deployments.  Users should use the data-plane REST service for interaction with vault secrets.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the vault belongs.
// vaultName - name of the vault
// secretName - name of the secret
// parameters - parameters to patch the secret
func (client SecretsClient) Update(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretPatchParameters) (result Secret, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: vaultName,
			Constraints: []validation.Constraint{{Target: "vaultName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{3,24}$`, Chain: nil}}},
		{TargetValue: secretName,
			Constraints: []validation.Constraint{{Target: "secretName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{1,127}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("keyvault.SecretsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, vaultName, secretName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.SecretsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SecretsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretPatchParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"secretName":        autorest.Encode("path", secretName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SecretsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SecretsClient) UpdateResponder(resp *http.Response) (result Secret, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

package costmanagement

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

// QueryClient is the client for the Query methods of the Costmanagement service.
type QueryClient struct {
	BaseClient
}

// NewQueryClient creates an instance of the QueryClient client.
func NewQueryClient(subscriptionID string) QueryClient {
	return NewQueryClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueryClientWithBaseURI creates an instance of the QueryClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQueryClientWithBaseURI(baseURI string, subscriptionID string) QueryClient {
	return QueryClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// UsageByBillingAccount query the usage data for billing account.
// Parameters:
// billingAccountID - billingAccount ID
// parameters - parameters supplied to the CreateOrUpdate Report Config operation.
func (client QueryClient) UsageByBillingAccount(ctx context.Context, billingAccountID string, parameters ReportConfigDefinition) (result QueryResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryClient.UsageByBillingAccount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Dataset", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
						{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
								{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
								{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("costmanagement.QueryClient", "UsageByBillingAccount", err.Error())
	}

	req, err := client.UsageByBillingAccountPreparer(ctx, billingAccountID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.UsageByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByBillingAccount", resp, "Failure sending request")
		return
	}

	result, err = client.UsageByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByBillingAccount", resp, "Failure responding to request")
		return
	}

	return
}

// UsageByBillingAccountPreparer prepares the UsageByBillingAccount request.
func (client QueryClient) UsageByBillingAccountPreparer(ctx context.Context, billingAccountID string, parameters ReportConfigDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/Query", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UsageByBillingAccountSender sends the UsageByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client QueryClient) UsageByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UsageByBillingAccountResponder handles the response to the UsageByBillingAccount request. The method always
// closes the http.Response Body.
func (client QueryClient) UsageByBillingAccountResponder(resp *http.Response) (result QueryResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UsageByDepartment query the usage data for department.
// Parameters:
// billingAccountID - billingAccount ID
// departmentID - department ID
// parameters - parameters supplied to the CreateOrUpdate Report Config operation.
func (client QueryClient) UsageByDepartment(ctx context.Context, billingAccountID string, departmentID string, parameters ReportConfigDefinition) (result QueryResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryClient.UsageByDepartment")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Dataset", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
						{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
								{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
								{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("costmanagement.QueryClient", "UsageByDepartment", err.Error())
	}

	req, err := client.UsageByDepartmentPreparer(ctx, billingAccountID, departmentID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByDepartment", nil, "Failure preparing request")
		return
	}

	resp, err := client.UsageByDepartmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByDepartment", resp, "Failure sending request")
		return
	}

	result, err = client.UsageByDepartmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByDepartment", resp, "Failure responding to request")
		return
	}

	return
}

// UsageByDepartmentPreparer prepares the UsageByDepartment request.
func (client QueryClient) UsageByDepartmentPreparer(ctx context.Context, billingAccountID string, departmentID string, parameters ReportConfigDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
		"departmentId":     autorest.Encode("path", departmentID),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}/providers/Microsoft.CostManagement/Query", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UsageByDepartmentSender sends the UsageByDepartment request. The method will close the
// http.Response Body if it receives an error.
func (client QueryClient) UsageByDepartmentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UsageByDepartmentResponder handles the response to the UsageByDepartment request. The method always
// closes the http.Response Body.
func (client QueryClient) UsageByDepartmentResponder(resp *http.Response) (result QueryResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UsageByEnrollmentAccount query the usage data for an enrollment account.
// Parameters:
// billingAccountID - billingAccount ID
// enrollmentAccountID - enrollment Account ID
// parameters - parameters supplied to the CreateOrUpdate Report Config operation.
func (client QueryClient) UsageByEnrollmentAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, parameters ReportConfigDefinition) (result QueryResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryClient.UsageByEnrollmentAccount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Dataset", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
						{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
								{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
								{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("costmanagement.QueryClient", "UsageByEnrollmentAccount", err.Error())
	}

	req, err := client.UsageByEnrollmentAccountPreparer(ctx, billingAccountID, enrollmentAccountID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByEnrollmentAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.UsageByEnrollmentAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByEnrollmentAccount", resp, "Failure sending request")
		return
	}

	result, err = client.UsageByEnrollmentAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByEnrollmentAccount", resp, "Failure responding to request")
		return
	}

	return
}

// UsageByEnrollmentAccountPreparer prepares the UsageByEnrollmentAccount request.
func (client QueryClient) UsageByEnrollmentAccountPreparer(ctx context.Context, billingAccountID string, enrollmentAccountID string, parameters ReportConfigDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId":    autorest.Encode("path", billingAccountID),
		"enrollmentAccountId": autorest.Encode("path", enrollmentAccountID),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}/providers/Microsoft.CostManagement/Query", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UsageByEnrollmentAccountSender sends the UsageByEnrollmentAccount request. The method will close the
// http.Response Body if it receives an error.
func (client QueryClient) UsageByEnrollmentAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UsageByEnrollmentAccountResponder handles the response to the UsageByEnrollmentAccount request. The method always
// closes the http.Response Body.
func (client QueryClient) UsageByEnrollmentAccountResponder(resp *http.Response) (result QueryResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UsageByManagementGroup lists the usage data for management group.
// Parameters:
// managementGroupID - managementGroup ID
// parameters - parameters supplied to the CreateOrUpdate Report Config operation.
func (client QueryClient) UsageByManagementGroup(ctx context.Context, managementGroupID string, parameters ReportConfigDefinition) (result QueryResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryClient.UsageByManagementGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Dataset", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
						{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
								{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
								{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("costmanagement.QueryClient", "UsageByManagementGroup", err.Error())
	}

	req, err := client.UsageByManagementGroupPreparer(ctx, managementGroupID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.UsageByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.UsageByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByManagementGroup", resp, "Failure responding to request")
		return
	}

	return
}

// UsageByManagementGroupPreparer prepares the UsageByManagementGroup request.
func (client QueryClient) UsageByManagementGroupPreparer(ctx context.Context, managementGroupID string, parameters ReportConfigDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.CostManagement/Query", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UsageByManagementGroupSender sends the UsageByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client QueryClient) UsageByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UsageByManagementGroupResponder handles the response to the UsageByManagementGroup request. The method always
// closes the http.Response Body.
func (client QueryClient) UsageByManagementGroupResponder(resp *http.Response) (result QueryResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UsageByResourceGroup query the usage data for subscriptionId and resource group.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// parameters - parameters supplied to the CreateOrUpdate Report Config operation.
func (client QueryClient) UsageByResourceGroup(ctx context.Context, resourceGroupName string, parameters ReportConfigDefinition) (result QueryResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryClient.UsageByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Dataset", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
						{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
								{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
								{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("costmanagement.QueryClient", "UsageByResourceGroup", err.Error())
	}

	req, err := client.UsageByResourceGroupPreparer(ctx, resourceGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.UsageByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.UsageByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// UsageByResourceGroupPreparer prepares the UsageByResourceGroup request.
func (client QueryClient) UsageByResourceGroupPreparer(ctx context.Context, resourceGroupName string, parameters ReportConfigDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.CostManagement/Query", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UsageByResourceGroupSender sends the UsageByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client QueryClient) UsageByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UsageByResourceGroupResponder handles the response to the UsageByResourceGroup request. The method always
// closes the http.Response Body.
func (client QueryClient) UsageByResourceGroupResponder(resp *http.Response) (result QueryResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UsageBySubscription query the usage data for subscriptionId.
// Parameters:
// parameters - parameters supplied to the CreateOrUpdate Report Config operation.
func (client QueryClient) UsageBySubscription(ctx context.Context, parameters ReportConfigDefinition) (result QueryResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryClient.UsageBySubscription")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil},
					}},
				{Target: "parameters.Dataset", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil}}},
						{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil}}},
								{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil},
								{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
								{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true,
											Chain: []validation.Constraint{{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil}}},
									}},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("costmanagement.QueryClient", "UsageBySubscription", err.Error())
	}

	req, err := client.UsageBySubscriptionPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.UsageBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.UsageBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.QueryClient", "UsageBySubscription", resp, "Failure responding to request")
		return
	}

	return
}

// UsageBySubscriptionPreparer prepares the UsageBySubscription request.
func (client QueryClient) UsageBySubscriptionPreparer(ctx context.Context, parameters ReportConfigDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/Query", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UsageBySubscriptionSender sends the UsageBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client QueryClient) UsageBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UsageBySubscriptionResponder handles the response to the UsageBySubscription request. The method always
// closes the http.Response Body.
func (client QueryClient) UsageBySubscriptionResponder(resp *http.Response) (result QueryResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
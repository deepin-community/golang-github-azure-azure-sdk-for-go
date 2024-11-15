// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package linksapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result links.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result links.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*links.OperationsClient)(nil)

// ResourceLinksClientAPI contains the set of methods on the ResourceLinksClient type.
type ResourceLinksClientAPI interface {
	CreateOrUpdate(ctx context.Context, linkID string, parameters links.ResourceLink) (result links.ResourceLink, err error)
	Delete(ctx context.Context, linkID string) (result autorest.Response, err error)
	Get(ctx context.Context, linkID string) (result links.ResourceLink, err error)
	ListAtSourceScope(ctx context.Context, scope string, filter links.Filter) (result links.ResourceLinkResultPage, err error)
	ListAtSourceScopeComplete(ctx context.Context, scope string, filter links.Filter) (result links.ResourceLinkResultIterator, err error)
	ListAtSubscription(ctx context.Context, filter string) (result links.ResourceLinkResultPage, err error)
	ListAtSubscriptionComplete(ctx context.Context, filter string) (result links.ResourceLinkResultIterator, err error)
}

var _ ResourceLinksClientAPI = (*links.ResourceLinksClient)(nil)

// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package dnsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2016-04-01/dns"
	"github.com/Azure/go-autorest/autorest"
)

// RecordSetsClientAPI contains the set of methods on the RecordSetsClient type.
type RecordSetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, parameters dns.RecordSet, ifMatch string, ifNoneMatch string) (result dns.RecordSet, err error)
	Delete(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType) (result dns.RecordSet, err error)
	ListByDNSZone(ctx context.Context, resourceGroupName string, zoneName string, top *int32, recordsetnamesuffix string) (result dns.RecordSetListResultPage, err error)
	ListByDNSZoneComplete(ctx context.Context, resourceGroupName string, zoneName string, top *int32, recordsetnamesuffix string) (result dns.RecordSetListResultIterator, err error)
	ListByType(ctx context.Context, resourceGroupName string, zoneName string, recordType dns.RecordType, top *int32, recordsetnamesuffix string) (result dns.RecordSetListResultPage, err error)
	ListByTypeComplete(ctx context.Context, resourceGroupName string, zoneName string, recordType dns.RecordType, top *int32, recordsetnamesuffix string) (result dns.RecordSetListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, parameters dns.RecordSet, ifMatch string) (result dns.RecordSet, err error)
}

var _ RecordSetsClientAPI = (*dns.RecordSetsClient)(nil)

// ZonesClientAPI contains the set of methods on the ZonesClient type.
type ZonesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, parameters dns.Zone, ifMatch string, ifNoneMatch string) (result dns.Zone, err error)
	Delete(ctx context.Context, resourceGroupName string, zoneName string, ifMatch string) (result dns.ZonesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, zoneName string) (result dns.Zone, err error)
	List(ctx context.Context, top *int32) (result dns.ZoneListResultPage, err error)
	ListComplete(ctx context.Context, top *int32) (result dns.ZoneListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result dns.ZoneListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result dns.ZoneListResultIterator, err error)
}

var _ ZonesClientAPI = (*dns.ZonesClient)(nil)
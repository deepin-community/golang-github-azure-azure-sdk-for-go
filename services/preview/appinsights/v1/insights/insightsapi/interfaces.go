package insightsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/appinsights/v1/insights"
)

// MetricsClientAPI contains the set of methods on the MetricsClient type.
type MetricsClientAPI interface {
	Get(ctx context.Context, appID string, metricID insights.MetricID, timespan string, interval *string, aggregation []insights.MetricsAggregation, segment []insights.MetricsSegment, top *int32, orderby string, filter string) (result insights.MetricsResult, err error)
	GetMetadata(ctx context.Context, appID string) (result insights.SetObject, err error)
	GetMultiple(ctx context.Context, appID string, body []insights.MetricsPostBodySchema) (result insights.ListMetricsResultsItem, err error)
}

var _ MetricsClientAPI = (*insights.MetricsClient)(nil)

// EventsClientAPI contains the set of methods on the EventsClient type.
type EventsClientAPI interface {
	Get(ctx context.Context, appID string, eventType insights.EventType, eventID string, timespan string) (result insights.EventsResults, err error)
	GetByType(ctx context.Context, appID string, eventType insights.EventType, timespan string, filter string, search string, orderby string, selectParameter string, skip *int32, top *int32, formatParameter string, count *bool, apply string) (result insights.EventsResults, err error)
	GetOdataMetadata(ctx context.Context, appID string) (result insights.SetObject, err error)
}

var _ EventsClientAPI = (*insights.EventsClient)(nil)

// QueryClientAPI contains the set of methods on the QueryClient type.
type QueryClientAPI interface {
	Execute(ctx context.Context, appID string, body insights.QueryBody) (result insights.QueryResults, err error)
}

var _ QueryClientAPI = (*insights.QueryClient)(nil)

// MetadataClientAPI contains the set of methods on the MetadataClient type.
type MetadataClientAPI interface {
	Get(ctx context.Context, appID string) (result insights.MetadataResults, err error)
	Post(ctx context.Context, appID string) (result insights.MetadataResults, err error)
}

var _ MetadataClientAPI = (*insights.MetadataClient)(nil)
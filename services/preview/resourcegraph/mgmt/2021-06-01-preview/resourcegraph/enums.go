package resourcegraph

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AuthorizationScopeFilter enumerates the values for authorization scope filter.
type AuthorizationScopeFilter string

const (
	// AtScopeAboveAndBelow ...
	AtScopeAboveAndBelow AuthorizationScopeFilter = "AtScopeAboveAndBelow"
	// AtScopeAndAbove ...
	AtScopeAndAbove AuthorizationScopeFilter = "AtScopeAndAbove"
	// AtScopeAndBelow ...
	AtScopeAndBelow AuthorizationScopeFilter = "AtScopeAndBelow"
	// AtScopeExact ...
	AtScopeExact AuthorizationScopeFilter = "AtScopeExact"
)

// PossibleAuthorizationScopeFilterValues returns an array of possible values for the AuthorizationScopeFilter const type.
func PossibleAuthorizationScopeFilterValues() []AuthorizationScopeFilter {
	return []AuthorizationScopeFilter{AtScopeAboveAndBelow, AtScopeAndAbove, AtScopeAndBelow, AtScopeExact}
}

// ColumnDataType enumerates the values for column data type.
type ColumnDataType string

const (
	// Boolean ...
	Boolean ColumnDataType = "boolean"
	// Datetime ...
	Datetime ColumnDataType = "datetime"
	// Integer ...
	Integer ColumnDataType = "integer"
	// Number ...
	Number ColumnDataType = "number"
	// Object ...
	Object ColumnDataType = "object"
	// String ...
	String ColumnDataType = "string"
)

// PossibleColumnDataTypeValues returns an array of possible values for the ColumnDataType const type.
func PossibleColumnDataTypeValues() []ColumnDataType {
	return []ColumnDataType{Boolean, Datetime, Integer, Number, Object, String}
}

// FacetSortOrder enumerates the values for facet sort order.
type FacetSortOrder string

const (
	// Asc ...
	Asc FacetSortOrder = "asc"
	// Desc ...
	Desc FacetSortOrder = "desc"
)

// PossibleFacetSortOrderValues returns an array of possible values for the FacetSortOrder const type.
func PossibleFacetSortOrderValues() []FacetSortOrder {
	return []FacetSortOrder{Asc, Desc}
}

// ResultFormat enumerates the values for result format.
type ResultFormat string

const (
	// ResultFormatObjectArray ...
	ResultFormatObjectArray ResultFormat = "objectArray"
	// ResultFormatTable ...
	ResultFormatTable ResultFormat = "table"
)

// PossibleResultFormatValues returns an array of possible values for the ResultFormat const type.
func PossibleResultFormatValues() []ResultFormat {
	return []ResultFormat{ResultFormatObjectArray, ResultFormatTable}
}

// ResultTruncated enumerates the values for result truncated.
type ResultTruncated string

const (
	// False ...
	False ResultTruncated = "false"
	// True ...
	True ResultTruncated = "true"
)

// PossibleResultTruncatedValues returns an array of possible values for the ResultTruncated const type.
func PossibleResultTruncatedValues() []ResultTruncated {
	return []ResultTruncated{False, True}
}

// ResultType enumerates the values for result type.
type ResultType string

const (
	// ResultTypeFacet ...
	ResultTypeFacet ResultType = "Facet"
	// ResultTypeFacetError ...
	ResultTypeFacetError ResultType = "FacetError"
	// ResultTypeFacetResult ...
	ResultTypeFacetResult ResultType = "FacetResult"
)

// PossibleResultTypeValues returns an array of possible values for the ResultType const type.
func PossibleResultTypeValues() []ResultType {
	return []ResultType{ResultTypeFacet, ResultTypeFacetError, ResultTypeFacetResult}
}

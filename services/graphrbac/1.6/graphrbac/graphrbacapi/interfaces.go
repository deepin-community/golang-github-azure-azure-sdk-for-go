package graphrbacapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
)

// SignedInUserClientAPI contains the set of methods on the SignedInUserClient type.
type SignedInUserClientAPI interface {
	Get(ctx context.Context) (result graphrbac.User, err error)
	ListOwnedObjects(ctx context.Context) (result graphrbac.DirectoryObjectListResultPage, err error)
	ListOwnedObjectsComplete(ctx context.Context) (result graphrbac.DirectoryObjectListResultIterator, err error)
	ListOwnedObjectsNext(ctx context.Context, nextLink string) (result graphrbac.DirectoryObjectListResult, err error)
}

var _ SignedInUserClientAPI = (*graphrbac.SignedInUserClient)(nil)

// ApplicationsClientAPI contains the set of methods on the ApplicationsClient type.
type ApplicationsClientAPI interface {
	AddOwner(ctx context.Context, applicationObjectID string, parameters graphrbac.AddOwnerParameters) (result autorest.Response, err error)
	Create(ctx context.Context, parameters graphrbac.ApplicationCreateParameters) (result graphrbac.Application, err error)
	Delete(ctx context.Context, applicationObjectID string) (result autorest.Response, err error)
	Get(ctx context.Context, applicationObjectID string) (result graphrbac.Application, err error)
	GetServicePrincipalsIDByAppID(ctx context.Context, applicationID string) (result graphrbac.ServicePrincipalObjectResult, err error)
	List(ctx context.Context, filter string) (result graphrbac.ApplicationListResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result graphrbac.ApplicationListResultIterator, err error)
	ListKeyCredentials(ctx context.Context, applicationObjectID string) (result graphrbac.KeyCredentialListResult, err error)
	ListNext(ctx context.Context, nextLink string) (result graphrbac.ApplicationListResult, err error)
	ListOwners(ctx context.Context, applicationObjectID string) (result graphrbac.DirectoryObjectListResultPage, err error)
	ListOwnersComplete(ctx context.Context, applicationObjectID string) (result graphrbac.DirectoryObjectListResultIterator, err error)
	ListPasswordCredentials(ctx context.Context, applicationObjectID string) (result graphrbac.PasswordCredentialListResult, err error)
	Patch(ctx context.Context, applicationObjectID string, parameters graphrbac.ApplicationUpdateParameters) (result autorest.Response, err error)
	RemoveOwner(ctx context.Context, applicationObjectID string, ownerObjectID string) (result autorest.Response, err error)
	UpdateKeyCredentials(ctx context.Context, applicationObjectID string, parameters graphrbac.KeyCredentialsUpdateParameters) (result autorest.Response, err error)
	UpdatePasswordCredentials(ctx context.Context, applicationObjectID string, parameters graphrbac.PasswordCredentialsUpdateParameters) (result autorest.Response, err error)
}

var _ ApplicationsClientAPI = (*graphrbac.ApplicationsClient)(nil)

// DeletedApplicationsClientAPI contains the set of methods on the DeletedApplicationsClient type.
type DeletedApplicationsClientAPI interface {
	HardDelete(ctx context.Context, applicationObjectID string) (result autorest.Response, err error)
	List(ctx context.Context, filter string) (result graphrbac.ApplicationListResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result graphrbac.ApplicationListResultIterator, err error)
	ListNext(ctx context.Context, nextLink string) (result graphrbac.ApplicationListResult, err error)
	Restore(ctx context.Context, objectID string) (result graphrbac.Application, err error)
}

var _ DeletedApplicationsClientAPI = (*graphrbac.DeletedApplicationsClient)(nil)

// GroupsClientAPI contains the set of methods on the GroupsClient type.
type GroupsClientAPI interface {
	AddMember(ctx context.Context, groupObjectID string, parameters graphrbac.GroupAddMemberParameters) (result autorest.Response, err error)
	AddOwner(ctx context.Context, objectID string, parameters graphrbac.AddOwnerParameters) (result autorest.Response, err error)
	Create(ctx context.Context, parameters graphrbac.GroupCreateParameters) (result graphrbac.ADGroup, err error)
	Delete(ctx context.Context, objectID string) (result autorest.Response, err error)
	Get(ctx context.Context, objectID string) (result graphrbac.ADGroup, err error)
	GetGroupMembers(ctx context.Context, objectID string) (result graphrbac.DirectoryObjectListResultPage, err error)
	GetGroupMembersComplete(ctx context.Context, objectID string) (result graphrbac.DirectoryObjectListResultIterator, err error)
	GetGroupMembersNext(ctx context.Context, nextLink string) (result graphrbac.DirectoryObjectListResult, err error)
	GetMemberGroups(ctx context.Context, objectID string, parameters graphrbac.GroupGetMemberGroupsParameters) (result graphrbac.GroupGetMemberGroupsResult, err error)
	IsMemberOf(ctx context.Context, parameters graphrbac.CheckGroupMembershipParameters) (result graphrbac.CheckGroupMembershipResult, err error)
	List(ctx context.Context, filter string) (result graphrbac.GroupListResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result graphrbac.GroupListResultIterator, err error)
	ListNext(ctx context.Context, nextLink string) (result graphrbac.GroupListResult, err error)
	ListOwners(ctx context.Context, objectID string) (result graphrbac.DirectoryObjectListResultPage, err error)
	ListOwnersComplete(ctx context.Context, objectID string) (result graphrbac.DirectoryObjectListResultIterator, err error)
	RemoveMember(ctx context.Context, groupObjectID string, memberObjectID string) (result autorest.Response, err error)
	RemoveOwner(ctx context.Context, objectID string, ownerObjectID string) (result autorest.Response, err error)
}

var _ GroupsClientAPI = (*graphrbac.GroupsClient)(nil)

// ServicePrincipalsClientAPI contains the set of methods on the ServicePrincipalsClient type.
type ServicePrincipalsClientAPI interface {
	Create(ctx context.Context, parameters graphrbac.ServicePrincipalCreateParameters) (result graphrbac.ServicePrincipal, err error)
	Delete(ctx context.Context, objectID string) (result autorest.Response, err error)
	Get(ctx context.Context, objectID string) (result graphrbac.ServicePrincipal, err error)
	List(ctx context.Context, filter string) (result graphrbac.ServicePrincipalListResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result graphrbac.ServicePrincipalListResultIterator, err error)
	ListKeyCredentials(ctx context.Context, objectID string) (result graphrbac.KeyCredentialListResult, err error)
	ListNext(ctx context.Context, nextLink string) (result graphrbac.ServicePrincipalListResult, err error)
	ListOwners(ctx context.Context, objectID string) (result graphrbac.DirectoryObjectListResultPage, err error)
	ListOwnersComplete(ctx context.Context, objectID string) (result graphrbac.DirectoryObjectListResultIterator, err error)
	ListPasswordCredentials(ctx context.Context, objectID string) (result graphrbac.PasswordCredentialListResult, err error)
	Update(ctx context.Context, objectID string, parameters graphrbac.ServicePrincipalUpdateParameters) (result autorest.Response, err error)
	UpdateKeyCredentials(ctx context.Context, objectID string, parameters graphrbac.KeyCredentialsUpdateParameters) (result autorest.Response, err error)
	UpdatePasswordCredentials(ctx context.Context, objectID string, parameters graphrbac.PasswordCredentialsUpdateParameters) (result autorest.Response, err error)
}

var _ ServicePrincipalsClientAPI = (*graphrbac.ServicePrincipalsClient)(nil)

// UsersClientAPI contains the set of methods on the UsersClient type.
type UsersClientAPI interface {
	Create(ctx context.Context, parameters graphrbac.UserCreateParameters) (result graphrbac.User, err error)
	Delete(ctx context.Context, upnOrObjectID string) (result autorest.Response, err error)
	Get(ctx context.Context, upnOrObjectID string) (result graphrbac.User, err error)
	GetMemberGroups(ctx context.Context, objectID string, parameters graphrbac.UserGetMemberGroupsParameters) (result graphrbac.UserGetMemberGroupsResult, err error)
	List(ctx context.Context, filter string, expand string) (result graphrbac.UserListResultPage, err error)
	ListComplete(ctx context.Context, filter string, expand string) (result graphrbac.UserListResultIterator, err error)
	ListNext(ctx context.Context, nextLink string) (result graphrbac.UserListResult, err error)
	Update(ctx context.Context, upnOrObjectID string, parameters graphrbac.UserUpdateParameters) (result autorest.Response, err error)
}

var _ UsersClientAPI = (*graphrbac.UsersClient)(nil)

// ObjectsClientAPI contains the set of methods on the ObjectsClient type.
type ObjectsClientAPI interface {
	GetObjectsByObjectIds(ctx context.Context, parameters graphrbac.GetObjectsParameters) (result graphrbac.DirectoryObjectListResultPage, err error)
	GetObjectsByObjectIdsComplete(ctx context.Context, parameters graphrbac.GetObjectsParameters) (result graphrbac.DirectoryObjectListResultIterator, err error)
	GetObjectsByObjectIdsNext(ctx context.Context, nextLink string) (result graphrbac.DirectoryObjectListResult, err error)
}

var _ ObjectsClientAPI = (*graphrbac.ObjectsClient)(nil)

// DomainsClientAPI contains the set of methods on the DomainsClient type.
type DomainsClientAPI interface {
	Get(ctx context.Context, domainName string) (result graphrbac.Domain, err error)
	List(ctx context.Context, filter string) (result graphrbac.DomainListResult, err error)
}

var _ DomainsClientAPI = (*graphrbac.DomainsClient)(nil)

// OAuth2PermissionGrantClientAPI contains the set of methods on the OAuth2PermissionGrantClient type.
type OAuth2PermissionGrantClientAPI interface {
	Create(ctx context.Context, body *graphrbac.OAuth2PermissionGrant) (result graphrbac.OAuth2PermissionGrant, err error)
	Delete(ctx context.Context, objectID string) (result autorest.Response, err error)
	List(ctx context.Context, filter string) (result graphrbac.OAuth2PermissionGrantListResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result graphrbac.OAuth2PermissionGrantListResultIterator, err error)
	ListNext(ctx context.Context, nextLink string) (result graphrbac.OAuth2PermissionGrantListResult, err error)
}

var _ OAuth2PermissionGrantClientAPI = (*graphrbac.OAuth2PermissionGrantClient)(nil)
/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-keycloak/apis/group/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1"
	v1alpha12 "github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DefaultGroups.
func (mg *DefaultGroups) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.GroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.GroupIdsRefs,
		Selector:      mg.Spec.ForProvider.GroupIdsSelector,
		To: reference.To{
			List:    &v1alpha1.GroupList{},
			Managed: &v1alpha1.Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupIds")
	}
	mg.Spec.ForProvider.GroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.GroupIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha11.RealmList{},
			Managed: &v1alpha11.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.GroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.GroupIdsRefs,
		Selector:      mg.Spec.InitProvider.GroupIdsSelector,
		To: reference.To{
			List:    &v1alpha1.GroupList{},
			Managed: &v1alpha1.Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GroupIds")
	}
	mg.Spec.InitProvider.GroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.GroupIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha11.RealmList{},
			Managed: &v1alpha11.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Roles.
func (mg *Roles) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.DefaultRoles),
		Extract:       resource.ExtractParamPath("name", false),
		References:    mg.Spec.ForProvider.DefaultRolesRefs,
		Selector:      mg.Spec.ForProvider.DefaultRolesSelector,
		To: reference.To{
			List:    &v1alpha12.RoleList{},
			Managed: &v1alpha12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DefaultRoles")
	}
	mg.Spec.ForProvider.DefaultRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.DefaultRolesRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha11.RealmList{},
			Managed: &v1alpha11.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.DefaultRoles),
		Extract:       resource.ExtractParamPath("name", false),
		References:    mg.Spec.InitProvider.DefaultRolesRefs,
		Selector:      mg.Spec.InitProvider.DefaultRolesSelector,
		To: reference.To{
			List:    &v1alpha12.RoleList{},
			Managed: &v1alpha12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DefaultRoles")
	}
	mg.Spec.InitProvider.DefaultRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.DefaultRolesRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha11.RealmList{},
			Managed: &v1alpha11.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

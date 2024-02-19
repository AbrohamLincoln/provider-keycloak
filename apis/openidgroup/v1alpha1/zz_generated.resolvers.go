/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this GroupMembershipProtocolMapper.
func (mg *GroupMembershipProtocolMapper) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClientIDRef,
		Selector:     mg.Spec.ForProvider.ClientIDSelector,
		To: reference.To{
			List:    &v1alpha1.ClientList{},
			Managed: &v1alpha1.Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientScopeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClientScopeIDRef,
		Selector:     mg.Spec.ForProvider.ClientScopeIDSelector,
		To: reference.To{
			List:    &v1alpha1.ClientScopeList{},
			Managed: &v1alpha1.ClientScope{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientScopeID")
	}
	mg.Spec.ForProvider.ClientScopeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientScopeIDRef = rsp.ResolvedReference

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

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClientIDRef,
		Selector:     mg.Spec.InitProvider.ClientIDSelector,
		To: reference.To{
			List:    &v1alpha1.ClientList{},
			Managed: &v1alpha1.Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientID")
	}
	mg.Spec.InitProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientScopeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClientScopeIDRef,
		Selector:     mg.Spec.InitProvider.ClientScopeIDSelector,
		To: reference.To{
			List:    &v1alpha1.ClientScopeList{},
			Managed: &v1alpha1.ClientScope{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientScopeID")
	}
	mg.Spec.InitProvider.ClientScopeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientScopeIDRef = rsp.ResolvedReference

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

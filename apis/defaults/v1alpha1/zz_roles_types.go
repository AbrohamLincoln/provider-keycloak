/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RolesInitParameters struct {

	// Realm level roles assigned to new users by default.
	// Realm level roles assigned to new users.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name", false)
	// +listType=set
	DefaultRoles []*string `json:"defaultRoles,omitempty" tf:"default_roles,omitempty"`

	// References to Role in role to populate defaultRoles.
	// +kubebuilder:validation:Optional
	DefaultRolesRefs []v1.Reference `json:"defaultRolesRefs,omitempty" tf:"-"`

	// Selector for a list of Role in role to populate defaultRoles.
	// +kubebuilder:validation:Optional
	DefaultRolesSelector *v1.Selector `json:"defaultRolesSelector,omitempty" tf:"-"`

	// The realm this role exists within.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type RolesObservation struct {

	// Realm level roles assigned to new users by default.
	// Realm level roles assigned to new users.
	// +listType=set
	DefaultRoles []*string `json:"defaultRoles,omitempty" tf:"default_roles,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The realm this role exists within.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type RolesParameters struct {

	// Realm level roles assigned to new users by default.
	// Realm level roles assigned to new users.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name", false)
	// +kubebuilder:validation:Optional
	// +listType=set
	DefaultRoles []*string `json:"defaultRoles,omitempty" tf:"default_roles,omitempty"`

	// References to Role in role to populate defaultRoles.
	// +kubebuilder:validation:Optional
	DefaultRolesRefs []v1.Reference `json:"defaultRolesRefs,omitempty" tf:"-"`

	// Selector for a list of Role in role to populate defaultRoles.
	// +kubebuilder:validation:Optional
	DefaultRolesSelector *v1.Selector `json:"defaultRolesSelector,omitempty" tf:"-"`

	// The realm this role exists within.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// RolesSpec defines the desired state of Roles
type RolesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RolesParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RolesInitParameters `json:"initProvider,omitempty"`
}

// RolesStatus defines the observed state of Roles.
type RolesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RolesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Roles is the Schema for the Roless API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type Roles struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RolesSpec   `json:"spec"`
	Status            RolesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RolesList contains a list of Roless
type RolesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Roles `json:"items"`
}

// Repository type metadata.
var (
	Roles_Kind             = "Roles"
	Roles_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Roles_Kind}.String()
	Roles_KindAPIVersion   = Roles_Kind + "." + CRDGroupVersion.String()
	Roles_GroupVersionKind = CRDGroupVersion.WithKind(Roles_Kind)
)

func init() {
	SchemeBuilder.Register(&Roles{}, &RolesList{})
}

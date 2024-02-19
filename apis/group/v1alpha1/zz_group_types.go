// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type GroupInitParameters struct {

	// A map representing attributes for the group. In order to add multivalue attributes, use ## to seperate the values. Max length for each value is 255 chars
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The name of the group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of this group's parent. If omitted, this group will be defined at the root level.
	// +crossplane:generate:reference:type=Group
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Reference to a Group to populate parentId.
	// +kubebuilder:validation:Optional
	ParentIDRef *v1.Reference `json:"parentIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate parentId.
	// +kubebuilder:validation:Optional
	ParentIDSelector *v1.Selector `json:"parentIdSelector,omitempty" tf:"-"`

	// The realm this group exists in.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type GroupObservation struct {

	// A map representing attributes for the group. In order to add multivalue attributes, use ## to seperate the values. Max length for each value is 255 chars
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of this group's parent. If omitted, this group will be defined at the root level.
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// (Computed) The complete path of the group. For example, the child group's path in the example configuration would be /parent-group/child-group.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The realm this group exists in.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type GroupParameters struct {

	// A map representing attributes for the group. In order to add multivalue attributes, use ## to seperate the values. Max length for each value is 255 chars
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The name of the group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of this group's parent. If omitted, this group will be defined at the root level.
	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Reference to a Group to populate parentId.
	// +kubebuilder:validation:Optional
	ParentIDRef *v1.Reference `json:"parentIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate parentId.
	// +kubebuilder:validation:Optional
	ParentIDSelector *v1.Selector `json:"parentIdSelector,omitempty" tf:"-"`

	// The realm this group exists in.
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

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
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
	InitProvider GroupInitParameters `json:"initProvider,omitempty"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Group is the Schema for the Groups API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   GroupSpec   `json:"spec"`
	Status GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}

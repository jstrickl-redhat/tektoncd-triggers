package v1alpha1

/*
Copyright 2019 The Tekton Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Check that TriggerBinding may be validated and defaulted.
//var _ apis.Validatable = (*TriggerBinding)(nil)

// +k8s:deepcopy-gen=true
type TriggerBindingSpec struct {
	Templates []TriggerTemplateRef `json:"templates,omitempty"`
}

// +k8s:deepcopy-gen=true
type TriggerTemplateRef struct {
	TemplateRef string  `json:"templateref,omitempty"`
	Event       string  `json:"event,omitempty"`
	Params      []Param `json:"params,omitempty"` // TODO: do we need default params?
}

type TriggerBindingStatus struct{}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TriggerBinding
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type TriggerBinding struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec holds the desired state of the EventListener from the client
	// +optional
	Spec TriggerBindingSpec `json:"spec"`
	// +optional
	Status TriggerBindingStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TriggerBindingList contains a list of TriggerBinding
// +k8s:deepcopy-gen=true
type TriggerBindingList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerBinding `json:"items"`
}
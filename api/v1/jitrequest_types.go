/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// JitRequestSpec defines the desired state of JitRequest.
type JitRequestSpec struct {
	// Requestor's username/email
	Reporter string `json:"reporter" validate:"required"`
	// Approver's user mail
	Approver string `json:"approver" validate:"required"`
	// Product Owner's user email
	ProductOwner string `json:"productOwner" validate:"required"`
	// Justification
	Justification string `json:"justification" validate:"required"`
	// Role to bind
	ClusterRole string `json:"clusterRole" validate:"required"`
	// Namespace to bind role and user
	Namespace string `json:"namespace" validate:"required"`
	// Start time for the JIT access, i.e. "2024-12-04T21:00:00Z"
	// ISO 8601 format
	StartTime metav1.Time `json:"startTime" validate:"required"`
	// End time for the JIT access, i.e. "2024-12-04T22:00:00Z"
	// ISO 8601 format
	EndTime metav1.Time `json:"endTime" validate:"required"`
}

// JitRequestStatus defines the observed state of JitRequest.
type JitRequestStatus struct {
	// Status of jit request
	// +kubebuilder:default:=Pending
	State string `json:"state,omitempty"`
	// Detailed message of jit request
	Message string `json:"message,omitempty"`
	// Jira ticket for jit request
	JiraTicket string `json:"jiraTicket,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName=jitreq
// +kubebuilder:printcolumn:name="User",type=string,JSONPath=`.spec.user`
// +kubebuilder:printcolumn:name="Cluster Role",type=string,JSONPath=`.spec.clusterRole`
// +kubebuilder:printcolumn:name="Namespace",type=string,JSONPath=`.spec.namespace`
// +kubebuilder:printcolumn:name="Start Time",type=string,JSONPath=`.spec.startTime`
// +kubebuilder:printcolumn:name="End Time",type=string,JSONPath=`.spec.endTime`

// JitRequest is the Schema for the jitrequests API.
type JitRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JitRequestSpec   `json:"spec,omitempty"`
	Status JitRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JitRequestList contains a list of JitRequest.
type JitRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JitRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&JitRequest{}, &JitRequestList{})
}

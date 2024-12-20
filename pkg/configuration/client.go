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

package configuration

import (
	"context"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	justintimev1 "jira-jit-rbac-operator/api/v1"
)

// jitRbacOperatorConfiguration
type jitRbacOperatorConfiguration struct {
	retrievalFn func() *justintimev1.JustInTimeConfig
}

func NewJitRbacOperatorConfiguration(ctx context.Context, client client.Client, name string) Configuration {
	return &jitRbacOperatorConfiguration{retrievalFn: func() *justintimev1.JustInTimeConfig {
		config := &justintimev1.JustInTimeConfig{}

		if err := client.Get(ctx, types.NamespacedName{Name: name}, config); err != nil {
			if apierrors.IsNotFound(err) {
				return &justintimev1.JustInTimeConfig{
					Spec: justintimev1.JustInTimeConfigSpec{
						AllowedClusterRoles:  []string{"edit"},
						RejectedTransitionID: "21",
						JiraProject:          "IAM",
						JiraIssueType:        "Access Request",
						ApprovedTransitionID: "41",
						CustomFields: &justintimev1.CustomFieldsSpec{
							Reporter:      justintimev1.CustomFieldSettings{Type: "text", JiraCustomField: "customfield_10113"},
							Approver:      justintimev1.CustomFieldSettings{Type: "user", JiraCustomField: "customfield_10114"},
							ProductOwner:  justintimev1.CustomFieldSettings{Type: "user", JiraCustomField: "customfield_10115"},
							Justification: justintimev1.CustomFieldSettings{Type: "text", JiraCustomField: "customfield_10116"},
							ClusterRole:   justintimev1.CustomFieldSettings{Type: "select", JiraCustomField: "customfield_10117"},
							StartTime:     justintimev1.CustomFieldSettings{Type: "date", JiraCustomField: "customfield_10118"},
							EndTime:       justintimev1.CustomFieldSettings{Type: "date", JiraCustomField: "customfield_10119"},
						},
					},
				}
			}
			panic(errors.Wrap(err, "Cannot retrieve configuration with name "+name))
		}

		return config
	}}
}

func (c *jitRbacOperatorConfiguration) AllowedClusterRoles() []string {
	return c.retrievalFn().Spec.AllowedClusterRoles
}

func (c *jitRbacOperatorConfiguration) RejectedTransitionID() string {
	return c.retrievalFn().Spec.RejectedTransitionID
}

func (c *jitRbacOperatorConfiguration) JiraProject() string {
	return c.retrievalFn().Spec.JiraProject
}

func (c *jitRbacOperatorConfiguration) JiraIssueType() string {
	return c.retrievalFn().Spec.JiraIssueType
}

func (c *jitRbacOperatorConfiguration) ApprovedTransitionID() string {
	return c.retrievalFn().Spec.ApprovedTransitionID
}

func (c *jitRbacOperatorConfiguration) CustomFields() *justintimev1.CustomFieldsSpec {
	return c.retrievalFn().Spec.CustomFields
}

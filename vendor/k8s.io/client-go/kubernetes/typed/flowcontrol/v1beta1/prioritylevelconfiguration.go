/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	flowcontrolv1beta1 "k8s.io/api/flowcontrol/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsflowcontrolv1beta1 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// PriorityLevelConfigurationsGetter has a method to return a PriorityLevelConfigurationInterface.
// A group's client should implement this interface.
type PriorityLevelConfigurationsGetter interface {
	PriorityLevelConfigurations() PriorityLevelConfigurationInterface
}

// PriorityLevelConfigurationInterface has methods to work with PriorityLevelConfiguration resources.
type PriorityLevelConfigurationInterface interface {
	Create(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta1.PriorityLevelConfiguration, opts v1.CreateOptions) (*flowcontrolv1beta1.PriorityLevelConfiguration, error)
	Update(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta1.PriorityLevelConfiguration, opts v1.UpdateOptions) (*flowcontrolv1beta1.PriorityLevelConfiguration, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta1.PriorityLevelConfiguration, opts v1.UpdateOptions) (*flowcontrolv1beta1.PriorityLevelConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*flowcontrolv1beta1.PriorityLevelConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*flowcontrolv1beta1.PriorityLevelConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *flowcontrolv1beta1.PriorityLevelConfiguration, err error)
	Apply(ctx context.Context, priorityLevelConfiguration *applyconfigurationsflowcontrolv1beta1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *flowcontrolv1beta1.PriorityLevelConfiguration, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, priorityLevelConfiguration *applyconfigurationsflowcontrolv1beta1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *flowcontrolv1beta1.PriorityLevelConfiguration, err error)
	PriorityLevelConfigurationExpansion
}

// priorityLevelConfigurations implements PriorityLevelConfigurationInterface
type priorityLevelConfigurations struct {
	*gentype.ClientWithListAndApply[*flowcontrolv1beta1.PriorityLevelConfiguration, *flowcontrolv1beta1.PriorityLevelConfigurationList, *applyconfigurationsflowcontrolv1beta1.PriorityLevelConfigurationApplyConfiguration]
}

// newPriorityLevelConfigurations returns a PriorityLevelConfigurations
func newPriorityLevelConfigurations(c *FlowcontrolV1beta1Client) *priorityLevelConfigurations {
	return &priorityLevelConfigurations{
		gentype.NewClientWithListAndApply[*flowcontrolv1beta1.PriorityLevelConfiguration, *flowcontrolv1beta1.PriorityLevelConfigurationList, *applyconfigurationsflowcontrolv1beta1.PriorityLevelConfigurationApplyConfiguration](
			"prioritylevelconfigurations",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *flowcontrolv1beta1.PriorityLevelConfiguration {
				return &flowcontrolv1beta1.PriorityLevelConfiguration{}
			},
			func() *flowcontrolv1beta1.PriorityLevelConfigurationList {
				return &flowcontrolv1beta1.PriorityLevelConfigurationList{}
			},
			gentype.PrefersProtobuf[*flowcontrolv1beta1.PriorityLevelConfiguration](),
		),
	}
}
/*
Copyright 2019 Rancher Labs.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	autoscaleriocattleiov1 "github.com/rancher/rio/pkg/apis/autoscale.rio.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceScaleRecommendations implements ServiceScaleRecommendationInterface
type FakeServiceScaleRecommendations struct {
	Fake *FakeAutoscaleV1
	ns   string
}

var servicescalerecommendationsResource = schema.GroupVersionResource{Group: "autoscale.rio.cattle.io", Version: "v1", Resource: "servicescalerecommendations"}

var servicescalerecommendationsKind = schema.GroupVersionKind{Group: "autoscale.rio.cattle.io", Version: "v1", Kind: "ServiceScaleRecommendation"}

// Get takes name of the serviceScaleRecommendation, and returns the corresponding serviceScaleRecommendation object, and an error if there is any.
func (c *FakeServiceScaleRecommendations) Get(name string, options v1.GetOptions) (result *autoscaleriocattleiov1.ServiceScaleRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicescalerecommendationsResource, c.ns, name), &autoscaleriocattleiov1.ServiceScaleRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaleriocattleiov1.ServiceScaleRecommendation), err
}

// List takes label and field selectors, and returns the list of ServiceScaleRecommendations that match those selectors.
func (c *FakeServiceScaleRecommendations) List(opts v1.ListOptions) (result *autoscaleriocattleiov1.ServiceScaleRecommendationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicescalerecommendationsResource, servicescalerecommendationsKind, c.ns, opts), &autoscaleriocattleiov1.ServiceScaleRecommendationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &autoscaleriocattleiov1.ServiceScaleRecommendationList{ListMeta: obj.(*autoscaleriocattleiov1.ServiceScaleRecommendationList).ListMeta}
	for _, item := range obj.(*autoscaleriocattleiov1.ServiceScaleRecommendationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceScaleRecommendations.
func (c *FakeServiceScaleRecommendations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicescalerecommendationsResource, c.ns, opts))

}

// Create takes the representation of a serviceScaleRecommendation and creates it.  Returns the server's representation of the serviceScaleRecommendation, and an error, if there is any.
func (c *FakeServiceScaleRecommendations) Create(serviceScaleRecommendation *autoscaleriocattleiov1.ServiceScaleRecommendation) (result *autoscaleriocattleiov1.ServiceScaleRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicescalerecommendationsResource, c.ns, serviceScaleRecommendation), &autoscaleriocattleiov1.ServiceScaleRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaleriocattleiov1.ServiceScaleRecommendation), err
}

// Update takes the representation of a serviceScaleRecommendation and updates it. Returns the server's representation of the serviceScaleRecommendation, and an error, if there is any.
func (c *FakeServiceScaleRecommendations) Update(serviceScaleRecommendation *autoscaleriocattleiov1.ServiceScaleRecommendation) (result *autoscaleriocattleiov1.ServiceScaleRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicescalerecommendationsResource, c.ns, serviceScaleRecommendation), &autoscaleriocattleiov1.ServiceScaleRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaleriocattleiov1.ServiceScaleRecommendation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceScaleRecommendations) UpdateStatus(serviceScaleRecommendation *autoscaleriocattleiov1.ServiceScaleRecommendation) (*autoscaleriocattleiov1.ServiceScaleRecommendation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicescalerecommendationsResource, "status", c.ns, serviceScaleRecommendation), &autoscaleriocattleiov1.ServiceScaleRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaleriocattleiov1.ServiceScaleRecommendation), err
}

// Delete takes name of the serviceScaleRecommendation and deletes it. Returns an error if one occurs.
func (c *FakeServiceScaleRecommendations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicescalerecommendationsResource, c.ns, name), &autoscaleriocattleiov1.ServiceScaleRecommendation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceScaleRecommendations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicescalerecommendationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &autoscaleriocattleiov1.ServiceScaleRecommendationList{})
	return err
}

// Patch applies the patch and returns the patched serviceScaleRecommendation.
func (c *FakeServiceScaleRecommendations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *autoscaleriocattleiov1.ServiceScaleRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicescalerecommendationsResource, c.ns, name, pt, data, subresources...), &autoscaleriocattleiov1.ServiceScaleRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaleriocattleiov1.ServiceScaleRecommendation), err
}

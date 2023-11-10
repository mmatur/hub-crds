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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/traefik/hub-crds/hub/v1alpha1"
	scheme "github.com/traefik/hub-crds/hub/v1alpha1/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// APIPortalsGetter has a method to return a APIPortalInterface.
// A group's client should implement this interface.
type APIPortalsGetter interface {
	APIPortals() APIPortalInterface
}

// APIPortalInterface has methods to work with APIPortal resources.
type APIPortalInterface interface {
	Create(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.CreateOptions) (*v1alpha1.APIPortal, error)
	Update(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.UpdateOptions) (*v1alpha1.APIPortal, error)
	UpdateStatus(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.UpdateOptions) (*v1alpha1.APIPortal, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIPortal, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIPortalList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIPortal, err error)
	APIPortalExpansion
}

// aPIPortals implements APIPortalInterface
type aPIPortals struct {
	client rest.Interface
}

// newAPIPortals returns a APIPortals
func newAPIPortals(c *HubV1alpha1Client) *aPIPortals {
	return &aPIPortals{
		client: c.RESTClient(),
	}
}

// Get takes name of the aPIPortal, and returns the corresponding aPIPortal object, and an error if there is any.
func (c *aPIPortals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIPortal, err error) {
	result = &v1alpha1.APIPortal{}
	err = c.client.Get().
		Resource("apiportals").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APIPortals that match those selectors.
func (c *aPIPortals) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIPortalList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.APIPortalList{}
	err = c.client.Get().
		Resource("apiportals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPIPortals.
func (c *aPIPortals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apiportals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPIPortal and creates it.  Returns the server's representation of the aPIPortal, and an error, if there is any.
func (c *aPIPortals) Create(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.CreateOptions) (result *v1alpha1.APIPortal, err error) {
	result = &v1alpha1.APIPortal{}
	err = c.client.Post().
		Resource("apiportals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIPortal).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPIPortal and updates it. Returns the server's representation of the aPIPortal, and an error, if there is any.
func (c *aPIPortals) Update(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.UpdateOptions) (result *v1alpha1.APIPortal, err error) {
	result = &v1alpha1.APIPortal{}
	err = c.client.Put().
		Resource("apiportals").
		Name(aPIPortal.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIPortal).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aPIPortals) UpdateStatus(ctx context.Context, aPIPortal *v1alpha1.APIPortal, opts v1.UpdateOptions) (result *v1alpha1.APIPortal, err error) {
	result = &v1alpha1.APIPortal{}
	err = c.client.Put().
		Resource("apiportals").
		Name(aPIPortal.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIPortal).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPIPortal and deletes it. Returns an error if one occurs.
func (c *aPIPortals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apiportals").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPIPortals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apiportals").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPIPortal.
func (c *aPIPortals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIPortal, err error) {
	result = &v1alpha1.APIPortal{}
	err = c.client.Patch(pt).
		Resource("apiportals").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
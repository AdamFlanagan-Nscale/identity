/*
Copyright 2022-2024 EscherCloud.

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

	scheme "github.com/eschercloudai/unikorn/generated/clientset/unikorn/scheme"
	v1alpha1 "github.com/eschercloudai/unikorn/pkg/apis/unikorn/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UserApplicationBundlesGetter has a method to return a UserApplicationBundleInterface.
// A group's client should implement this interface.
type UserApplicationBundlesGetter interface {
	UserApplicationBundles(namespace string) UserApplicationBundleInterface
}

// UserApplicationBundleInterface has methods to work with UserApplicationBundle resources.
type UserApplicationBundleInterface interface {
	Create(ctx context.Context, userApplicationBundle *v1alpha1.UserApplicationBundle, opts v1.CreateOptions) (*v1alpha1.UserApplicationBundle, error)
	Update(ctx context.Context, userApplicationBundle *v1alpha1.UserApplicationBundle, opts v1.UpdateOptions) (*v1alpha1.UserApplicationBundle, error)
	UpdateStatus(ctx context.Context, userApplicationBundle *v1alpha1.UserApplicationBundle, opts v1.UpdateOptions) (*v1alpha1.UserApplicationBundle, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.UserApplicationBundle, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.UserApplicationBundleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserApplicationBundle, err error)
	UserApplicationBundleExpansion
}

// userApplicationBundles implements UserApplicationBundleInterface
type userApplicationBundles struct {
	client rest.Interface
	ns     string
}

// newUserApplicationBundles returns a UserApplicationBundles
func newUserApplicationBundles(c *UnikornV1alpha1Client, namespace string) *userApplicationBundles {
	return &userApplicationBundles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the userApplicationBundle, and returns the corresponding userApplicationBundle object, and an error if there is any.
func (c *userApplicationBundles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UserApplicationBundle, err error) {
	result = &v1alpha1.UserApplicationBundle{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("userapplicationbundles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UserApplicationBundles that match those selectors.
func (c *userApplicationBundles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserApplicationBundleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UserApplicationBundleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("userapplicationbundles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested userApplicationBundles.
func (c *userApplicationBundles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("userapplicationbundles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a userApplicationBundle and creates it.  Returns the server's representation of the userApplicationBundle, and an error, if there is any.
func (c *userApplicationBundles) Create(ctx context.Context, userApplicationBundle *v1alpha1.UserApplicationBundle, opts v1.CreateOptions) (result *v1alpha1.UserApplicationBundle, err error) {
	result = &v1alpha1.UserApplicationBundle{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("userapplicationbundles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userApplicationBundle).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a userApplicationBundle and updates it. Returns the server's representation of the userApplicationBundle, and an error, if there is any.
func (c *userApplicationBundles) Update(ctx context.Context, userApplicationBundle *v1alpha1.UserApplicationBundle, opts v1.UpdateOptions) (result *v1alpha1.UserApplicationBundle, err error) {
	result = &v1alpha1.UserApplicationBundle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("userapplicationbundles").
		Name(userApplicationBundle.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userApplicationBundle).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *userApplicationBundles) UpdateStatus(ctx context.Context, userApplicationBundle *v1alpha1.UserApplicationBundle, opts v1.UpdateOptions) (result *v1alpha1.UserApplicationBundle, err error) {
	result = &v1alpha1.UserApplicationBundle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("userapplicationbundles").
		Name(userApplicationBundle.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userApplicationBundle).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the userApplicationBundle and deletes it. Returns an error if one occurs.
func (c *userApplicationBundles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("userapplicationbundles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *userApplicationBundles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("userapplicationbundles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched userApplicationBundle.
func (c *userApplicationBundles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserApplicationBundle, err error) {
	result = &v1alpha1.UserApplicationBundle{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("userapplicationbundles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

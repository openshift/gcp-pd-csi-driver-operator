// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OpenShiftAPIServersGetter has a method to return a OpenShiftAPIServerInterface.
// A group's client should implement this interface.
type OpenShiftAPIServersGetter interface {
	OpenShiftAPIServers() OpenShiftAPIServerInterface
}

// OpenShiftAPIServerInterface has methods to work with OpenShiftAPIServer resources.
type OpenShiftAPIServerInterface interface {
	Create(ctx context.Context, openShiftAPIServer *v1.OpenShiftAPIServer, opts metav1.CreateOptions) (*v1.OpenShiftAPIServer, error)
	Update(ctx context.Context, openShiftAPIServer *v1.OpenShiftAPIServer, opts metav1.UpdateOptions) (*v1.OpenShiftAPIServer, error)
	UpdateStatus(ctx context.Context, openShiftAPIServer *v1.OpenShiftAPIServer, opts metav1.UpdateOptions) (*v1.OpenShiftAPIServer, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.OpenShiftAPIServer, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.OpenShiftAPIServerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OpenShiftAPIServer, err error)
	Apply(ctx context.Context, openShiftAPIServer *operatorv1.OpenShiftAPIServerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OpenShiftAPIServer, err error)
	ApplyStatus(ctx context.Context, openShiftAPIServer *operatorv1.OpenShiftAPIServerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OpenShiftAPIServer, err error)
	OpenShiftAPIServerExpansion
}

// openShiftAPIServers implements OpenShiftAPIServerInterface
type openShiftAPIServers struct {
	client rest.Interface
}

// newOpenShiftAPIServers returns a OpenShiftAPIServers
func newOpenShiftAPIServers(c *OperatorV1Client) *openShiftAPIServers {
	return &openShiftAPIServers{
		client: c.RESTClient(),
	}
}

// Get takes name of the openShiftAPIServer, and returns the corresponding openShiftAPIServer object, and an error if there is any.
func (c *openShiftAPIServers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OpenShiftAPIServer, err error) {
	result = &v1.OpenShiftAPIServer{}
	err = c.client.Get().
		Resource("openshiftapiservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpenShiftAPIServers that match those selectors.
func (c *openShiftAPIServers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OpenShiftAPIServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.OpenShiftAPIServerList{}
	err = c.client.Get().
		Resource("openshiftapiservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested openShiftAPIServers.
func (c *openShiftAPIServers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("openshiftapiservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a openShiftAPIServer and creates it.  Returns the server's representation of the openShiftAPIServer, and an error, if there is any.
func (c *openShiftAPIServers) Create(ctx context.Context, openShiftAPIServer *v1.OpenShiftAPIServer, opts metav1.CreateOptions) (result *v1.OpenShiftAPIServer, err error) {
	result = &v1.OpenShiftAPIServer{}
	err = c.client.Post().
		Resource("openshiftapiservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(openShiftAPIServer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a openShiftAPIServer and updates it. Returns the server's representation of the openShiftAPIServer, and an error, if there is any.
func (c *openShiftAPIServers) Update(ctx context.Context, openShiftAPIServer *v1.OpenShiftAPIServer, opts metav1.UpdateOptions) (result *v1.OpenShiftAPIServer, err error) {
	result = &v1.OpenShiftAPIServer{}
	err = c.client.Put().
		Resource("openshiftapiservers").
		Name(openShiftAPIServer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(openShiftAPIServer).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *openShiftAPIServers) UpdateStatus(ctx context.Context, openShiftAPIServer *v1.OpenShiftAPIServer, opts metav1.UpdateOptions) (result *v1.OpenShiftAPIServer, err error) {
	result = &v1.OpenShiftAPIServer{}
	err = c.client.Put().
		Resource("openshiftapiservers").
		Name(openShiftAPIServer.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(openShiftAPIServer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the openShiftAPIServer and deletes it. Returns an error if one occurs.
func (c *openShiftAPIServers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("openshiftapiservers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *openShiftAPIServers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("openshiftapiservers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched openShiftAPIServer.
func (c *openShiftAPIServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OpenShiftAPIServer, err error) {
	result = &v1.OpenShiftAPIServer{}
	err = c.client.Patch(pt).
		Resource("openshiftapiservers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied openShiftAPIServer.
func (c *openShiftAPIServers) Apply(ctx context.Context, openShiftAPIServer *operatorv1.OpenShiftAPIServerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OpenShiftAPIServer, err error) {
	if openShiftAPIServer == nil {
		return nil, fmt.Errorf("openShiftAPIServer provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(openShiftAPIServer)
	if err != nil {
		return nil, err
	}
	name := openShiftAPIServer.Name
	if name == nil {
		return nil, fmt.Errorf("openShiftAPIServer.Name must be provided to Apply")
	}
	result = &v1.OpenShiftAPIServer{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("openshiftapiservers").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *openShiftAPIServers) ApplyStatus(ctx context.Context, openShiftAPIServer *operatorv1.OpenShiftAPIServerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OpenShiftAPIServer, err error) {
	if openShiftAPIServer == nil {
		return nil, fmt.Errorf("openShiftAPIServer provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(openShiftAPIServer)
	if err != nil {
		return nil, err
	}

	name := openShiftAPIServer.Name
	if name == nil {
		return nil, fmt.Errorf("openShiftAPIServer.Name must be provided to Apply")
	}

	result = &v1.OpenShiftAPIServer{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("openshiftapiservers").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

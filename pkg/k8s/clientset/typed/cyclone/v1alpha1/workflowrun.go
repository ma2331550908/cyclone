/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/cyclone/pkg/apis/cyclone/v1alpha1"
	scheme "github.com/caicloud/cyclone/pkg/k8s/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkflowRunsGetter has a method to return a WorkflowRunInterface.
// A group's client should implement this interface.
type WorkflowRunsGetter interface {
	WorkflowRuns(namespace string) WorkflowRunInterface
}

// WorkflowRunInterface has methods to work with WorkflowRun resources.
type WorkflowRunInterface interface {
	Create(*v1alpha1.WorkflowRun) (*v1alpha1.WorkflowRun, error)
	Update(*v1alpha1.WorkflowRun) (*v1alpha1.WorkflowRun, error)
	UpdateStatus(*v1alpha1.WorkflowRun) (*v1alpha1.WorkflowRun, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WorkflowRun, error)
	List(opts v1.ListOptions) (*v1alpha1.WorkflowRunList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WorkflowRun, err error)
	WorkflowRunExpansion
}

// workflowRuns implements WorkflowRunInterface
type workflowRuns struct {
	client rest.Interface
	ns     string
}

// newWorkflowRuns returns a WorkflowRuns
func newWorkflowRuns(c *CycloneV1alpha1Client, namespace string) *workflowRuns {
	return &workflowRuns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the workflowRun, and returns the corresponding workflowRun object, and an error if there is any.
func (c *workflowRuns) Get(name string, options v1.GetOptions) (result *v1alpha1.WorkflowRun, err error) {
	result = &v1alpha1.WorkflowRun{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workflowruns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkflowRuns that match those selectors.
func (c *workflowRuns) List(opts v1.ListOptions) (result *v1alpha1.WorkflowRunList, err error) {
	result = &v1alpha1.WorkflowRunList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workflowruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workflowRuns.
func (c *workflowRuns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("workflowruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a workflowRun and creates it.  Returns the server's representation of the workflowRun, and an error, if there is any.
func (c *workflowRuns) Create(workflowRun *v1alpha1.WorkflowRun) (result *v1alpha1.WorkflowRun, err error) {
	result = &v1alpha1.WorkflowRun{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("workflowruns").
		Body(workflowRun).
		Do().
		Into(result)
	return
}

// Update takes the representation of a workflowRun and updates it. Returns the server's representation of the workflowRun, and an error, if there is any.
func (c *workflowRuns) Update(workflowRun *v1alpha1.WorkflowRun) (result *v1alpha1.WorkflowRun, err error) {
	result = &v1alpha1.WorkflowRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workflowruns").
		Name(workflowRun.Name).
		Body(workflowRun).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *workflowRuns) UpdateStatus(workflowRun *v1alpha1.WorkflowRun) (result *v1alpha1.WorkflowRun, err error) {
	result = &v1alpha1.WorkflowRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workflowruns").
		Name(workflowRun.Name).
		SubResource("status").
		Body(workflowRun).
		Do().
		Into(result)
	return
}

// Delete takes name of the workflowRun and deletes it. Returns an error if one occurs.
func (c *workflowRuns) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workflowruns").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workflowRuns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workflowruns").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched workflowRun.
func (c *workflowRuns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WorkflowRun, err error) {
	result = &v1alpha1.WorkflowRun{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("workflowruns").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

/*
Copyright 2019 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/kubeflow/arena/pkg/volcano-operator/apis/batch/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeJobs implements JobInterface
type FakeJobs struct {
	Fake *FakeBatchV1alpha1
	ns   string
}

var jobsResource = schema.GroupVersionResource{Group: "batch", Version: "v1alpha1", Resource: "jobs"}

var jobsKind = schema.GroupVersionKind{Group: "batch", Version: "v1alpha1", Kind: "Job"}

// Get takes name of the job, and returns the corresponding job object, and an error if there is any.
func (c *FakeJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(jobsResource, c.ns, name), &v1alpha1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Job), err
}

// List takes label and field selectors, and returns the list of Jobs that match those selectors.
func (c *FakeJobs) List(opts v1.ListOptions) (result *v1alpha1.JobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(jobsResource, jobsKind, c.ns, opts), &v1alpha1.JobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.JobList{}
	for _, item := range obj.(*v1alpha1.JobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jobs.
func (c *FakeJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(jobsResource, c.ns, opts))

}

// Create takes the representation of a job and creates it.  Returns the server's representation of the job, and an error, if there is any.
func (c *FakeJobs) Create(job *v1alpha1.Job) (result *v1alpha1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(jobsResource, c.ns, job), &v1alpha1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Job), err
}

// Update takes the representation of a job and updates it. Returns the server's representation of the job, and an error, if there is any.
func (c *FakeJobs) Update(job *v1alpha1.Job) (result *v1alpha1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(jobsResource, c.ns, job), &v1alpha1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Job), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJobs) UpdateStatus(job *v1alpha1.Job) (*v1alpha1.Job, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(jobsResource, "status", c.ns, job), &v1alpha1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Job), err
}

// Delete takes name of the job and deletes it. Returns an error if one occurs.
func (c *FakeJobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(jobsResource, c.ns, name), &v1alpha1.Job{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(jobsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.JobList{})
	return err
}

// Patch applies the patch and returns the patched job.
func (c *FakeJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(jobsResource, c.ns, name, data, subresources...), &v1alpha1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Job), err
}

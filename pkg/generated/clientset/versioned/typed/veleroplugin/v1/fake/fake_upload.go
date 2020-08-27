/*
Copyright the Velero contributors.

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

package fake

import (
	"context"

	veleropluginv1 "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/apis/veleroplugin/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUploads implements UploadInterface
type FakeUploads struct {
	Fake *FakeVeleropluginV1
	ns   string
}

var uploadsResource = schema.GroupVersionResource{Group: "veleroplugin.io", Version: "v1", Resource: "uploads"}

var uploadsKind = schema.GroupVersionKind{Group: "veleroplugin.io", Version: "v1", Kind: "Upload"}

// Get takes name of the upload, and returns the corresponding upload object, and an error if there is any.
func (c *FakeUploads) Get(ctx context.Context, name string, options v1.GetOptions) (result *veleropluginv1.Upload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(uploadsResource, c.ns, name), &veleropluginv1.Upload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Upload), err
}

// List takes label and field selectors, and returns the list of Uploads that match those selectors.
func (c *FakeUploads) List(ctx context.Context, opts v1.ListOptions) (result *veleropluginv1.UploadList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(uploadsResource, uploadsKind, c.ns, opts), &veleropluginv1.UploadList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &veleropluginv1.UploadList{ListMeta: obj.(*veleropluginv1.UploadList).ListMeta}
	for _, item := range obj.(*veleropluginv1.UploadList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested uploads.
func (c *FakeUploads) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(uploadsResource, c.ns, opts))

}

// Create takes the representation of a upload and creates it.  Returns the server's representation of the upload, and an error, if there is any.
func (c *FakeUploads) Create(ctx context.Context, upload *veleropluginv1.Upload, opts v1.CreateOptions) (result *veleropluginv1.Upload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(uploadsResource, c.ns, upload), &veleropluginv1.Upload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Upload), err
}

// Update takes the representation of a upload and updates it. Returns the server's representation of the upload, and an error, if there is any.
func (c *FakeUploads) Update(ctx context.Context, upload *veleropluginv1.Upload, opts v1.UpdateOptions) (result *veleropluginv1.Upload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(uploadsResource, c.ns, upload), &veleropluginv1.Upload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Upload), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUploads) UpdateStatus(ctx context.Context, upload *veleropluginv1.Upload, opts v1.UpdateOptions) (*veleropluginv1.Upload, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(uploadsResource, "status", c.ns, upload), &veleropluginv1.Upload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Upload), err
}

// Delete takes name of the upload and deletes it. Returns an error if one occurs.
func (c *FakeUploads) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(uploadsResource, c.ns, name), &veleropluginv1.Upload{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUploads) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(uploadsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &veleropluginv1.UploadList{})
	return err
}

// Patch applies the patch and returns the patched upload.
func (c *FakeUploads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *veleropluginv1.Upload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(uploadsResource, c.ns, name, pt, data, subresources...), &veleropluginv1.Upload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*veleropluginv1.Upload), err
}

/*
The GNU AFFERO GENERAL PUBLIC LICENSE

Copyright (c) 2020-2024 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EdgeIngressLister helps list EdgeIngresses.
// All objects returned here must be treated as read-only.
type EdgeIngressLister interface {
	// List lists all EdgeIngresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EdgeIngress, err error)
	// EdgeIngresses returns an object that can list and get EdgeIngresses.
	EdgeIngresses(namespace string) EdgeIngressNamespaceLister
	EdgeIngressListerExpansion
}

// edgeIngressLister implements the EdgeIngressLister interface.
type edgeIngressLister struct {
	indexer cache.Indexer
}

// NewEdgeIngressLister returns a new EdgeIngressLister.
func NewEdgeIngressLister(indexer cache.Indexer) EdgeIngressLister {
	return &edgeIngressLister{indexer: indexer}
}

// List lists all EdgeIngresses in the indexer.
func (s *edgeIngressLister) List(selector labels.Selector) (ret []*v1alpha1.EdgeIngress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EdgeIngress))
	})
	return ret, err
}

// EdgeIngresses returns an object that can list and get EdgeIngresses.
func (s *edgeIngressLister) EdgeIngresses(namespace string) EdgeIngressNamespaceLister {
	return edgeIngressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EdgeIngressNamespaceLister helps list and get EdgeIngresses.
// All objects returned here must be treated as read-only.
type EdgeIngressNamespaceLister interface {
	// List lists all EdgeIngresses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EdgeIngress, err error)
	// Get retrieves the EdgeIngress from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EdgeIngress, error)
	EdgeIngressNamespaceListerExpansion
}

// edgeIngressNamespaceLister implements the EdgeIngressNamespaceLister
// interface.
type edgeIngressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EdgeIngresses in the indexer for a given namespace.
func (s edgeIngressNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EdgeIngress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EdgeIngress))
	})
	return ret, err
}

// Get retrieves the EdgeIngress from the indexer for a given namespace and name.
func (s edgeIngressNamespaceLister) Get(name string) (*v1alpha1.EdgeIngress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("edgeingress"), name)
	}
	return obj.(*v1alpha1.EdgeIngress), nil
}
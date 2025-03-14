// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	authorizationv1 "github.com/openshift/api/authorization/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterRoleBindingLister helps list ClusterRoleBindings.
// All objects returned here must be treated as read-only.
type ClusterRoleBindingLister interface {
	// List lists all ClusterRoleBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*authorizationv1.ClusterRoleBinding, err error)
	// Get retrieves the ClusterRoleBinding from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*authorizationv1.ClusterRoleBinding, error)
	ClusterRoleBindingListerExpansion
}

// clusterRoleBindingLister implements the ClusterRoleBindingLister interface.
type clusterRoleBindingLister struct {
	listers.ResourceIndexer[*authorizationv1.ClusterRoleBinding]
}

// NewClusterRoleBindingLister returns a new ClusterRoleBindingLister.
func NewClusterRoleBindingLister(indexer cache.Indexer) ClusterRoleBindingLister {
	return &clusterRoleBindingLister{listers.New[*authorizationv1.ClusterRoleBinding](indexer, authorizationv1.Resource("clusterrolebinding"))}
}

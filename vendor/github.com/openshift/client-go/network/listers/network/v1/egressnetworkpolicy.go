// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	networkv1 "github.com/openshift/api/network/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// EgressNetworkPolicyLister helps list EgressNetworkPolicies.
// All objects returned here must be treated as read-only.
type EgressNetworkPolicyLister interface {
	// List lists all EgressNetworkPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkv1.EgressNetworkPolicy, err error)
	// EgressNetworkPolicies returns an object that can list and get EgressNetworkPolicies.
	EgressNetworkPolicies(namespace string) EgressNetworkPolicyNamespaceLister
	EgressNetworkPolicyListerExpansion
}

// egressNetworkPolicyLister implements the EgressNetworkPolicyLister interface.
type egressNetworkPolicyLister struct {
	listers.ResourceIndexer[*networkv1.EgressNetworkPolicy]
}

// NewEgressNetworkPolicyLister returns a new EgressNetworkPolicyLister.
func NewEgressNetworkPolicyLister(indexer cache.Indexer) EgressNetworkPolicyLister {
	return &egressNetworkPolicyLister{listers.New[*networkv1.EgressNetworkPolicy](indexer, networkv1.Resource("egressnetworkpolicy"))}
}

// EgressNetworkPolicies returns an object that can list and get EgressNetworkPolicies.
func (s *egressNetworkPolicyLister) EgressNetworkPolicies(namespace string) EgressNetworkPolicyNamespaceLister {
	return egressNetworkPolicyNamespaceLister{listers.NewNamespaced[*networkv1.EgressNetworkPolicy](s.ResourceIndexer, namespace)}
}

// EgressNetworkPolicyNamespaceLister helps list and get EgressNetworkPolicies.
// All objects returned here must be treated as read-only.
type EgressNetworkPolicyNamespaceLister interface {
	// List lists all EgressNetworkPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkv1.EgressNetworkPolicy, err error)
	// Get retrieves the EgressNetworkPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*networkv1.EgressNetworkPolicy, error)
	EgressNetworkPolicyNamespaceListerExpansion
}

// egressNetworkPolicyNamespaceLister implements the EgressNetworkPolicyNamespaceLister
// interface.
type egressNetworkPolicyNamespaceLister struct {
	listers.ResourceIndexer[*networkv1.EgressNetworkPolicy]
}

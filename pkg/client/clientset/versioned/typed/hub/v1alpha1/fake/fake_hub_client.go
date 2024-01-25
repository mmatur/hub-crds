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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/traefik/hub-crds/pkg/client/clientset/versioned/typed/hub/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHubV1alpha1 struct {
	*testing.Fake
}

func (c *FakeHubV1alpha1) APIs(namespace string) v1alpha1.APIInterface {
	return &FakeAPIs{c, namespace}
}

func (c *FakeHubV1alpha1) APIAccesses() v1alpha1.APIAccessInterface {
	return &FakeAPIAccesses{c}
}

func (c *FakeHubV1alpha1) APICollections() v1alpha1.APICollectionInterface {
	return &FakeAPICollections{c}
}

func (c *FakeHubV1alpha1) APIGateways() v1alpha1.APIGatewayInterface {
	return &FakeAPIGateways{c}
}

func (c *FakeHubV1alpha1) APIPortals() v1alpha1.APIPortalInterface {
	return &FakeAPIPortals{c}
}

func (c *FakeHubV1alpha1) APIRateLimits() v1alpha1.APIRateLimitInterface {
	return &FakeAPIRateLimits{c}
}

func (c *FakeHubV1alpha1) APIVersions(namespace string) v1alpha1.APIVersionInterface {
	return &FakeAPIVersions{c, namespace}
}

func (c *FakeHubV1alpha1) AccessControlPolicies() v1alpha1.AccessControlPolicyInterface {
	return &FakeAccessControlPolicies{c}
}

func (c *FakeHubV1alpha1) EdgeIngresses(namespace string) v1alpha1.EdgeIngressInterface {
	return &FakeEdgeIngresses{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHubV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
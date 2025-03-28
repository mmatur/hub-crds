/*
Copyright (C) 2022-2025 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package v1alpha1_test

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestAPIPortal_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  title: title
  description: description
  trustedUrls: ["https://example.com"]
  ui:
    logoUrl: https://example.com/logo.png
  `),
		},
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
spec:
  trustedUrls: ["https://example.com"]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: .non-dns-compliant-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com"]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-portal", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: ""
  namespace: default
spec:
  trustedUrls: ["https://example.com"]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: portal-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default
spec:
  trustedUrls: ["https://example.com"]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "portal-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "missing trustedUrls",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec: {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.trustedUrls", BadValue: ""}},
		},
		{
			desc: "empty trustedUrls",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: []`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.trustedUrls", BadValue: int64(0), Detail: "spec.trustedUrls in body should have at least 1 items"}},
		},
		{
			desc: "too many trustedUrls",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APIPortal
metadata:
  name: my-portal
  namespace: default
spec:
  trustedUrls: ["https://example.com", https://another.example.com]`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeTooMany, Field: "spec.trustedUrls", BadValue: 2, Detail: "must have at most 1 items"}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}

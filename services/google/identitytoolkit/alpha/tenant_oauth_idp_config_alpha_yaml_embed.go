// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_tenant_oauth_idp_config blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/identitytoolkit/alpha/tenant_oauth_idp_config.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/identitytoolkit/alpha/tenant_oauth_idp_config.yaml
var YAML_tenant_oauth_idp_config = []byte("info:\n  title: IdentityToolkit/TenantOAuthIdpConfig\n  description: DCL Specification for the IdentityToolkit TenantOAuthIdpConfig resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a TenantOAuthIdpConfig\n    parameters:\n    - name: TenantOAuthIdpConfig\n      required: true\n      description: A full instance of a TenantOAuthIdpConfig\n  apply:\n    description: The function used to apply information about a TenantOAuthIdpConfig\n    parameters:\n    - name: TenantOAuthIdpConfig\n      required: true\n      description: A full instance of a TenantOAuthIdpConfig\n  delete:\n    description: The function used to delete a TenantOAuthIdpConfig\n    parameters:\n    - name: TenantOAuthIdpConfig\n      required: true\n      description: A full instance of a TenantOAuthIdpConfig\n  deleteAll:\n    description: The function used to delete all TenantOAuthIdpConfig\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: tenant\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many TenantOAuthIdpConfig\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: tenant\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    TenantOAuthIdpConfig:\n      title: TenantOAuthIdpConfig\n      x-dcl-id: projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - project\n      - tenant\n      properties:\n        clientId:\n          type: string\n          x-dcl-go-name: ClientId\n          description: The client id of an OAuth client.\n        clientSecret:\n          type: string\n          x-dcl-go-name: ClientSecret\n          description: The client secret of the OAuth client, to enable OIDC code\n            flow.\n          x-dcl-sensitive: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: The config's display name set by developers.\n        enabled:\n          type: boolean\n          x-dcl-go-name: Enabled\n          description: True if allows the user to sign in with the provider.\n        issuer:\n          type: string\n          x-dcl-go-name: Issuer\n          description: For OIDC Idps, the issuer identifier.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the Config resource\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        responseType:\n          type: object\n          x-dcl-go-name: ResponseType\n          x-dcl-go-type: TenantOAuthIdpConfigResponseType\n          description: 'The multiple response type to request for in the OAuth authorization\n            flow. This can possibly be a combination of set bits (e.g.: {id\\_token,\n            token}).'\n          properties:\n            code:\n              type: boolean\n              x-dcl-go-name: Code\n              description: If true, authorization code is returned from IdP's authorization\n                endpoint.\n            idToken:\n              type: boolean\n              x-dcl-go-name: IdToken\n              description: If true, ID token is returned from IdP's authorization\n                endpoint.\n            token:\n              type: boolean\n              x-dcl-go-name: Token\n              description: If true, access token is returned from IdP's authorization\n                endpoint.\n        tenant:\n          type: string\n          x-dcl-go-name: Tenant\n          description: The tenant for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Identitytoolkit/Tenant\n            field: name\n            parent: true\n")

// 4046 bytes
// MD5: db9768d21cbc41e8c65c56de6e22e528

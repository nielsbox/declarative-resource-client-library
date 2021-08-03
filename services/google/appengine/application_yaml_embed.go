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
// gen_go_data -package appengine -var YAML_application blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/appengine/application.yaml

package appengine

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/appengine/application.yaml
var YAML_application = []byte("info:\n  title: AppEngine/Application\n  description: DCL Specification for the AppEngine Application resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Application\n    parameters:\n    - name: Application\n      required: true\n      description: A full instance of a Application\n  apply:\n    description: The function used to apply information about a Application\n    parameters:\n    - name: Application\n      required: true\n      description: A full instance of a Application\ncomponents:\n  schemas:\n    Application:\n      title: Application\n      x-dcl-id: apps/{{name}}\n      x-dcl-uses-state-hint: true\n      type: object\n      properties:\n        authDomain:\n          type: string\n          x-dcl-go-name: AuthDomain\n          description: Google Apps authentication domain that controls which users\n            can access this application. Defaults to open access for any Google Account.\n        codeBucket:\n          type: string\n          x-dcl-go-name: CodeBucket\n          readOnly: true\n          description: Google Cloud Storage bucket that can be used for storing files\n            associated with this application. This bucket is associated with the application\n            and can be used by the gcloud deployment commands. @OutputOnly\n          x-kubernetes-immutable: true\n        databaseType:\n          type: string\n          x-dcl-go-name: DatabaseType\n          x-dcl-go-type: ApplicationDatabaseTypeEnum\n          description: 'The type of the Cloud Firestore or Cloud Datastore database\n            associated with this application. Possible values: DATABASE_TYPE_UNSPECIFIED,\n            CLOUD_DATASTORE, CLOUD_FIRESTORE, CLOUD_DATASTORE_COMPATIBILITY'\n          enum:\n          - DATABASE_TYPE_UNSPECIFIED\n          - CLOUD_DATASTORE\n          - CLOUD_FIRESTORE\n          - CLOUD_DATASTORE_COMPATIBILITY\n        defaultBucket:\n          type: string\n          x-dcl-go-name: DefaultBucket\n          readOnly: true\n          description: Google Cloud Storage bucket that can be used by this application\n            to store content. @OutputOnly\n          x-kubernetes-immutable: true\n        defaultCookieExpiration:\n          type: string\n          x-dcl-go-name: DefaultCookieExpiration\n          description: 'Cookie expiration policy for this application. A duration\n            in seconds with up to nine fractional digits, terminated by ''s''. Example:\n            \"3.5s\".'\n        defaultHostname:\n          type: string\n          x-dcl-go-name: DefaultHostname\n          readOnly: true\n          description: Hostname used to reach this application, as resolved by App\n            Engine. @OutputOnly\n          x-kubernetes-immutable: true\n        dispatchRules:\n          type: array\n          x-dcl-go-name: DispatchRules\n          description: HTTP path dispatch rules for requests to the application that\n            do not explicitly target a service or version. Rules are order-dependent.\n            Up to 20 dispatch rules can be supported.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: ApplicationDispatchRules\n            properties:\n              domain:\n                type: string\n                x-dcl-go-name: Domain\n                description: 'Domain name to match against. The wildcard \"`*`\" is\n                  supported if specified before a period: \"`*.`\". Defaults to matching\n                  all domains: \"`*`\".'\n              path:\n                type: string\n                x-dcl-go-name: Path\n                description: Pathname within the host. Must start with a \"`/`\". A\n                  single \"`*`\" can be included at the end of the path. The sum of\n                  the lengths of the domain and path may not exceed 100 characters.\n              service:\n                type: string\n                x-dcl-go-name: Service\n                description: 'Resource ID of a service in this application that should\n                  serve the matched request. The service must already exist. Example:\n                  `default`.'\n        featureSettings:\n          type: object\n          x-dcl-go-name: FeatureSettings\n          x-dcl-go-type: ApplicationFeatureSettings\n          description: The feature specific settings to be used in the application.\n          properties:\n            splitHealthChecks:\n              type: boolean\n              x-dcl-go-name: SplitHealthChecks\n              description: Boolean value indicating if split health checks should\n                be used instead of the legacy health checks. At an app.yaml level,\n                this means defaulting to 'readiness_check' and 'liveness_check' values\n                instead of 'health_check' ones. Once the legacy 'health_check' behavior\n                is deprecated, and this value is always true, this setting can be\n                removed.\n            useContainerOptimizedOS:\n              type: boolean\n              x-dcl-go-name: UseContainerOptimizedOS\n              description: If true, use [Container-Optimized OS](https://cloud.google.com/container-optimized-os/)\n                base image for VMs, rather than a base Debian image.\n        gcrDomain:\n          type: string\n          x-dcl-go-name: GcrDomain\n          description: The Google Container Registry domain used for storing managed\n            build docker images for this application.\n        iap:\n          type: object\n          x-dcl-go-name: Iap\n          x-dcl-go-type: ApplicationIap\n          description: Identity Aware Proxy.\n          x-kubernetes-immutable: true\n          properties:\n            enabled:\n              type: boolean\n              x-dcl-go-name: Enabled\n              description: Whether the serving infrastructure will authenticate and\n                authorize all incoming requests. If true, the `oauth2_client_id` and\n                `oauth2_client_secret` fields must be non-empty.\n              x-kubernetes-immutable: true\n            oauth2ClientId:\n              type: string\n              x-dcl-go-name: OAuth2ClientId\n              description: OAuth2 client ID to use for the authentication flow.\n              x-kubernetes-immutable: true\n            oauth2ClientSecret:\n              type: string\n              x-dcl-go-name: OAuth2ClientSecret\n              description: OAuth2 client secret to use for the authentication flow.\n                For security reasons, this value cannot be retrieved via the API.\n                Instead, the SHA-256 hash of the value is returned in the `oauth2_client_secret_sha256`\n                field. @InputOnly\n              x-kubernetes-immutable: true\n              x-dcl-sensitive: true\n              x-dcl-mutable-unreadable: true\n            oauth2ClientSecretSha256:\n              type: string\n              x-dcl-go-name: OAuth2ClientSecretSha256\n              readOnly: true\n              description: Hex-encoded SHA-256 hash of the client secret. @OutputOnly\n              x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: Location from which this application runs. Application instances\n            run out of the data centers in the specified location, which is also where\n            all of the application's end user content is stored. Defaults to `us-central`.\n            View the list of [supported locations](https://cloud.google.com/appengine/docs/locations).\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the application\n          x-kubernetes-immutable: true\n        servingStatus:\n          type: string\n          x-dcl-go-name: ServingStatus\n          x-dcl-go-type: ApplicationServingStatusEnum\n          description: 'Serving status of this application. Possible values: UNSPECIFIED,\n            SERVING, USER_DISABLED, SYSTEM_DISABLED'\n          enum:\n          - UNSPECIFIED\n          - SERVING\n          - USER_DISABLED\n          - SYSTEM_DISABLED\n")

// 8105 bytes
// MD5: 8466a08bef321c640ad0c0c6344b9fb2
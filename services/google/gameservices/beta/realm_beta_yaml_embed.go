// Copyright 2022 Google LLC. All Rights Reserved.
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
// gen_go_data -package beta -var YAML_realm blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gameservices/beta/realm.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gameservices/beta/realm.yaml
var YAML_realm = []byte("info:\n  title: GameServices/Realm\n  description: The GameServices Realm resource\n  x-dcl-struct-name: Realm\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Realm\n    parameters:\n    - name: Realm\n      required: true\n      description: A full instance of a Realm\n  apply:\n    description: The function used to apply information about a Realm\n    parameters:\n    - name: Realm\n      required: true\n      description: A full instance of a Realm\n  delete:\n    description: The function used to delete a Realm\n    parameters:\n    - name: Realm\n      required: true\n      description: A full instance of a Realm\n  deleteAll:\n    description: The function used to delete all Realm\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Realm\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Realm:\n      title: Realm\n      x-dcl-id: projects/{{project}}/locations/{{location}}/realms/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - timeZone\n      - location\n      - project\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The creation time.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Human readable description of the realm.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: The labels associated with this realm. Each label is a key-value\n            pair.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for this realm.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name of the realm.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for this realm.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        timeZone:\n          type: string\n          x-dcl-go-name: TimeZone\n          description: 'Required. Time zone where all policies targeting this realm\n            are evaluated. The value of this field must be from the IANA time zone\n            database: https://www.iana.org/time-zones.'\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The last-modified time.\n          x-kubernetes-immutable: true\n")

// 3220 bytes
// MD5: 111c18b1383a50cdaac567ce12b5a2c5

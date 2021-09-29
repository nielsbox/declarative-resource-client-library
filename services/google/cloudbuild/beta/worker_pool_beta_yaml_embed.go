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
// gen_go_data -package beta -var YAML_worker_pool blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudbuild/beta/worker_pool.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudbuild/beta/worker_pool.yaml
var YAML_worker_pool = []byte("info:\n  title: CloudBuild/WorkerPool\n  description: DCL Specification for the CloudBuild WorkerPool resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a WorkerPool\n    parameters:\n    - name: WorkerPool\n      required: true\n      description: A full instance of a WorkerPool\n  apply:\n    description: The function used to apply information about a WorkerPool\n    parameters:\n    - name: WorkerPool\n      required: true\n      description: A full instance of a WorkerPool\n  delete:\n    description: The function used to delete a WorkerPool\n    parameters:\n    - name: WorkerPool\n      required: true\n      description: A full instance of a WorkerPool\n  deleteAll:\n    description: The function used to delete all WorkerPool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many WorkerPool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    WorkerPool:\n      title: WorkerPool\n      x-dcl-id: projects/{{project}}/locations/{{location}}/workerPools/{{name}}\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Time at which the request to create the `WorkerPool`\n            was received.\n          x-kubernetes-immutable: true\n        deleteTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: DeleteTime\n          readOnly: true\n          description: Output only. Time at which the request to delete the `WorkerPool`\n            was received.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: User-defined name of the `WorkerPool`.\n          x-kubernetes-immutable: true\n        networkConfig:\n          type: object\n          x-dcl-go-name: NetworkConfig\n          x-dcl-go-type: WorkerPoolNetworkConfig\n          description: Network configuration for the `WorkerPool`.\n          x-kubernetes-immutable: true\n          required:\n          - peeredNetwork\n          properties:\n            peeredNetwork:\n              type: string\n              x-dcl-go-name: PeeredNetwork\n              description: Required. Immutable. The network definition that the workers\n                are peered to. If this section is left empty, the workers will be\n                peered to `WorkerPool.project_id` on the service producer network.\n                Must be in the format `projects/{project}/global/networks/{network}`,\n                where `{project}` is a project number, such as `12345`, and `{network}`\n                is the name of a VPC network in the project. See [Understanding network\n                configuration options](https://cloud.google.com/cloud-build/docs/custom-workers/set-up-custom-worker-pool-environment#understanding_the_network_configuration_options)\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Compute/Network\n                field: selfLink\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: WorkerPoolStateEnum\n          readOnly: true\n          description: 'Output only. `WorkerPool` state. Possible values: STATE_UNSPECIFIED,\n            PENDING, APPROVED, REJECTED, CANCELLED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - PENDING\n          - APPROVED\n          - REJECTED\n          - CANCELLED\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Time at which the request to update the `WorkerPool`\n            was received.\n          x-kubernetes-immutable: true\n        workerConfig:\n          type: object\n          x-dcl-go-name: WorkerConfig\n          x-dcl-go-type: WorkerPoolWorkerConfig\n          description: Configuration to be used for a creating workers in the `WorkerPool`.\n          x-dcl-server-default: true\n          properties:\n            diskSizeGb:\n              type: integer\n              format: int64\n              x-dcl-go-name: DiskSizeGb\n              description: Size of the disk attached to the worker, in GB. See [Worker\n                pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file).\n                Specify a value of up to 1000. If `0` is specified, Cloud Build will\n                use a standard disk size.\n            machineType:\n              type: string\n              x-dcl-go-name: MachineType\n              description: Machine type of a worker, such as `n1-standard-1`. See\n                [Worker pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file).\n                If left blank, Cloud Build will use `n1-standard-1`.\n            noExternalIP:\n              type: boolean\n              x-dcl-go-name: NoExternalIP\n              description: If true, workers are created without any public address,\n                which prevents network egress to public IPs.\n")

// 6014 bytes
// MD5: 5c202ccf0700a700653e8947e9fd5196

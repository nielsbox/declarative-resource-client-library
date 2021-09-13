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
// gen_go_data -package beta -var YAML_instance blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/file/beta/instance.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/file/beta/instance.yaml
var YAML_instance = []byte("info:\n  title: File/Instance\n  description: DCL Specification for the File Instance resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Instance\n    parameters:\n    - name: Instance\n      required: true\n      description: A full instance of a Instance\n  apply:\n    description: The function used to apply information about a Instance\n    parameters:\n    - name: Instance\n      required: true\n      description: A full instance of a Instance\n  delete:\n    description: The function used to delete a Instance\n    parameters:\n    - name: Instance\n      required: true\n      description: A full instance of a Instance\n  deleteAll:\n    description: The function used to delete all Instance\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Instance\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Instance:\n      title: Instance\n      x-dcl-id: projects/{{project}}/locations/{{location}}/instances/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time when the instance was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: The description of the instance (2048 characters or less).\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Server-specified ETag for the instance resource to prevent\n            simultaneous updates from overwriting each other.\n          x-kubernetes-immutable: true\n        fileShares:\n          type: array\n          x-dcl-go-name: FileShares\n          description: File system shares on the instance. For this version, only\n            a single file share is supported.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceFileShares\n            properties:\n              capacityGb:\n                type: integer\n                format: int64\n                x-dcl-go-name: CapacityGb\n                description: File share capacity in gigabytes (GB). Cloud Filestore\n                  defines 1 GB as 1024^3 bytes.\n              name:\n                type: string\n                x-dcl-go-name: Name\n                description: The name of the file share (must be 16 characters or\n                  less).\n              nfsExportOptions:\n                type: array\n                x-dcl-go-name: NfsExportOptions\n                description: Nfs Export Options. There is a limit of 10 export options\n                  per file share.\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: InstanceFileSharesNfsExportOptions\n                  properties:\n                    accessMode:\n                      type: string\n                      x-dcl-go-name: AccessMode\n                      x-dcl-go-type: InstanceFileSharesNfsExportOptionsAccessModeEnum\n                      description: 'Either READ_ONLY, for allowing only read requests\n                        on the exported directory, or READ_WRITE, for allowing both\n                        read and write requests. The default is READ_WRITE. Possible\n                        values: ACCESS_MODE_UNSPECIFIED, READ_ONLY, READ_WRITE'\n                      x-dcl-server-default: true\n                      enum:\n                      - ACCESS_MODE_UNSPECIFIED\n                      - READ_ONLY\n                      - READ_WRITE\n                    anonGid:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: AnonGid\n                      description: An integer representing the anonymous group id\n                        with a default value of 65534. Anon_gid may only be set with\n                        squash_mode of ROOT_SQUASH. An error will be returned if this\n                        field is specified for other squash_mode settings.\n                    anonUid:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: AnonUid\n                      description: An integer representing the anonymous user id with\n                        a default value of 65534. Anon_uid may only be set with squash_mode\n                        of ROOT_SQUASH. An error will be returned if this field is\n                        specified for other squash_mode settings.\n                    ipRanges:\n                      type: array\n                      x-dcl-go-name: IPRanges\n                      description: List of either an IPv4 addresses in the format\n                        `{octet1}.{octet2}.{octet3}.{octet4}` or CIDR ranges in the\n                        format `{octet1}.{octet2}.{octet3}.{octet4}/{mask size}` which\n                        may mount the file share. Overlapping IP ranges are not allowed,\n                        both within and across NfsExportOptions. An error will be\n                        returned. The limit is 64 IP ranges/addresses for each FileShareConfig\n                        among all NfsExportOptions.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n                    squashMode:\n                      type: string\n                      x-dcl-go-name: SquashMode\n                      x-dcl-go-type: InstanceFileSharesNfsExportOptionsSquashModeEnum\n                      description: 'Either NO_ROOT_SQUASH, for allowing root access\n                        on the exported directory, or ROOT_SQUASH, for not allowing\n                        root access. The default is NO_ROOT_SQUASH. Possible values:\n                        SQUASH_MODE_UNSPECIFIED, NO_ROOT_SQUASH, ROOT_SQUASH'\n                      x-dcl-server-default: true\n                      enum:\n                      - SQUASH_MODE_UNSPECIFIED\n                      - NO_ROOT_SQUASH\n                      - ROOT_SQUASH\n              sourceBackup:\n                type: string\n                x-dcl-go-name: SourceBackup\n                description: The resource name of the backup, in the format `projects/{project_number}/locations/{location_id}/backups/{backup_id}`,\n                  that this file share has been restored from.\n                x-dcl-references:\n                - resource: File/Backup\n                  field: name\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Resource labels to represent user provided metadata.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The resource name of the instance, in the format\n            `projects/{project}/locations/{location}/instances/{instance}`.\n          x-kubernetes-immutable: true\n        networks:\n          type: array\n          x-dcl-go-name: Networks\n          description: VPC networks to which the instance is connected. For this version,\n            only a single network is supported.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceNetworks\n            properties:\n              ipAddresses:\n                type: array\n                x-dcl-go-name: IPAddresses\n                readOnly: true\n                description: Output only. IPv4 addresses in the format `{octet1}.{octet2}.{octet3}.{octet4}`\n                  or IPv6 addresses in the format `{block1}:{block2}:{block3}:{block4}:{block5}:{block6}:{block7}:{block8}`.\n                x-kubernetes-immutable: true\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: string\n              modes:\n                type: array\n                x-dcl-go-name: Modes\n                description: Internet protocol versions for which the instance has\n                  IP addresses assigned. For this version, only MODE_IPV4 is supported.\n                x-kubernetes-immutable: true\n                x-dcl-server-default: true\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: InstanceNetworksModesEnum\n                  enum:\n                  - ADDRESS_MODE_UNSPECIFIED\n                  - MODE_IPV4\n              network:\n                type: string\n                x-dcl-go-name: Network\n                description: The name of the Google Compute Engine [VPC network](https://cloud.google.com/vpc/docs/vpc)\n                  to which the instance is connected.\n                x-kubernetes-immutable: true\n                x-dcl-references:\n                - resource: Compute/Network\n                  field: name\n              reservedIPRange:\n                type: string\n                x-dcl-go-name: ReservedIPRange\n                description: A /29 CIDR block in one of the [internal IP address ranges](https://www.arin.net/reference/research/statistics/address_filters/)\n                  that identifies the range of IP addresses reserved for this instance.\n                  For example, 10.0.0.0/29 or 192.168.0.0/29. The range you specify\n                  can't overlap with either existing subnets or assigned IP address\n                  ranges for other Cloud Filestore instances in the selected VPC network.\n                x-kubernetes-immutable: true\n                x-dcl-server-default: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: InstanceStateEnum\n          readOnly: true\n          description: 'Output only. The instance state. Possible values: STATE_UNSPECIFIED,\n            CREATING, READY, REPAIRING, DELETING, ERROR'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - CREATING\n          - READY\n          - REPAIRING\n          - DELETING\n          - ERROR\n        statusMessage:\n          type: string\n          x-dcl-go-name: StatusMessage\n          readOnly: true\n          description: Output only. Additional information about the instance state,\n            if available.\n          x-kubernetes-immutable: true\n        tier:\n          type: string\n          x-dcl-go-name: Tier\n          x-dcl-go-type: InstanceTierEnum\n          description: 'The service tier of the instance. Possible values: TIER_UNSPECIFIED,\n            STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD'\n          x-kubernetes-immutable: true\n          enum:\n          - TIER_UNSPECIFIED\n          - STANDARD\n          - PREMIUM\n          - BASIC_HDD\n          - BASIC_SSD\n          - HIGH_SCALE_SSD\n")

// 12035 bytes
// MD5: 3ada45aac740c24c91bace9eb1347ee1

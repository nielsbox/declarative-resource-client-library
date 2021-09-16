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
// gen_go_data -package alpha -var YAML_metric_descriptor blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/alpha/metric_descriptor.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/alpha/metric_descriptor.yaml
var YAML_metric_descriptor = []byte("info:\n  title: Monitoring/MetricDescriptor\n  description: DCL Specification for the Monitoring MetricDescriptor resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a MetricDescriptor\n    parameters:\n    - name: MetricDescriptor\n      required: true\n      description: A full instance of a MetricDescriptor\n  apply:\n    description: The function used to apply information about a MetricDescriptor\n    parameters:\n    - name: MetricDescriptor\n      required: true\n      description: A full instance of a MetricDescriptor\n  delete:\n    description: The function used to delete a MetricDescriptor\n    parameters:\n    - name: MetricDescriptor\n      required: true\n      description: A full instance of a MetricDescriptor\n  deleteAll:\n    description: The function used to delete all MetricDescriptor\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many MetricDescriptor\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    MetricDescriptor:\n      title: MetricDescriptor\n      x-dcl-id: projects/{{project}}/metricDescriptors/{{type}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - type\n      - metricKind\n      - valueType\n      - project\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: A detailed description of the metric, which can be used in\n            documentation.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: A concise name for the metric, which can be displayed in user\n            interfaces. Use sentence case without an ending period, for example \"Request\n            count\". This field is optional but it is recommended to be set for any\n            metrics associated with user-visible concepts, such as Quota.\n          x-kubernetes-immutable: true\n        labels:\n          type: array\n          x-dcl-go-name: Labels\n          description: The set of labels that can be used to describe a specific instance\n            of this metric type. For example, the `appengine.googleapis.com/http/server/response_latencies`\n            metric type has a label for the HTTP response code, `response_code`, so\n            you can look at latencies for successful responses or just for responses\n            that failed.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: object\n            x-dcl-go-type: MetricDescriptorLabels\n            properties:\n              description:\n                type: string\n                x-dcl-go-name: Description\n                description: A human-readable description for the label.\n                x-kubernetes-immutable: true\n              key:\n                type: string\n                x-dcl-go-name: Key\n                description: 'The key for this label. The key must meet the following\n                  criteria: * Does not exceed 100 characters. * Matches the following\n                  regular expression: `a-zA-Z*` * The first character must be an upper-\n                  or lower-case letter. * The remaining characters must be letters,\n                  digits, or underscores.'\n                x-kubernetes-immutable: true\n              valueType:\n                type: string\n                x-dcl-go-name: ValueType\n                x-dcl-go-type: MetricDescriptorLabelsValueTypeEnum\n                description: 'The type of data that can be assigned to the label.\n                  Possible values: STRING, BOOL, INT64'\n                x-kubernetes-immutable: true\n                enum:\n                - STRING\n                - BOOL\n                - INT64\n        launchStage:\n          type: string\n          x-dcl-go-name: LaunchStage\n          x-dcl-go-type: MetricDescriptorLaunchStageEnum\n          description: 'Optional. The launch stage of the metric definition. Possible\n            values: LAUNCH_STAGE_UNSPECIFIED, UNIMPLEMENTED, PRELAUNCH, EARLY_ACCESS,\n            ALPHA, BETA, GA, DEPRECATED'\n          x-kubernetes-immutable: true\n          enum:\n          - LAUNCH_STAGE_UNSPECIFIED\n          - UNIMPLEMENTED\n          - PRELAUNCH\n          - EARLY_ACCESS\n          - ALPHA\n          - BETA\n          - GA\n          - DEPRECATED\n          x-dcl-mutable-unreadable: true\n        metadata:\n          type: object\n          x-dcl-go-name: Metadata\n          x-dcl-go-type: MetricDescriptorMetadata\n          description: Optional. Metadata which can be used to guide usage of the\n            metric.\n          x-kubernetes-immutable: true\n          x-dcl-mutable-unreadable: true\n          properties:\n            ingestDelay:\n              type: string\n              x-dcl-go-name: IngestDelay\n              description: The delay of data points caused by ingestion. Data points\n                older than this age are guaranteed to be ingested and available to\n                be read, excluding data loss due to errors.\n              x-kubernetes-immutable: true\n            launchStage:\n              type: string\n              x-dcl-go-name: LaunchStage\n              x-dcl-go-type: MetricDescriptorMetadataLaunchStageEnum\n              description: 'Deprecated. Must use the MetricDescriptor.launch_stage\n                instead. Possible values: LAUNCH_STAGE_UNSPECIFIED, UNIMPLEMENTED,\n                PRELAUNCH, EARLY_ACCESS, ALPHA, BETA, GA, DEPRECATED'\n              x-kubernetes-immutable: true\n              enum:\n              - LAUNCH_STAGE_UNSPECIFIED\n              - UNIMPLEMENTED\n              - PRELAUNCH\n              - EARLY_ACCESS\n              - ALPHA\n              - BETA\n              - GA\n              - DEPRECATED\n            samplePeriod:\n              type: string\n              x-dcl-go-name: SamplePeriod\n              description: The sampling period of metric data points. For metrics\n                which are written periodically, consecutive data points are stored\n                at this time interval, excluding data loss due to errors. Metrics\n                with a higher granularity have a smaller sampling period.\n              x-kubernetes-immutable: true\n        metricKind:\n          type: string\n          x-dcl-go-name: MetricKind\n          x-dcl-go-type: MetricDescriptorMetricKindEnum\n          description: 'Whether the metric records instantaneous values, changes to\n            a value, etc. Some combinations of `metric_kind` and `value_type` might\n            not be supported. Possible values: METRIC_KIND_UNSPECIFIED, GAUGE, DELTA,\n            CUMULATIVE'\n          x-kubernetes-immutable: true\n          enum:\n          - METRIC_KIND_UNSPECIFIED\n          - GAUGE\n          - DELTA\n          - CUMULATIVE\n        monitoredResourceTypes:\n          type: array\n          x-dcl-go-name: MonitoredResourceTypes\n          readOnly: true\n          description: Read-only. If present, then a time series, which is identified\n            partially by a metric type and a MonitoredResourceDescriptor, that is\n            associated with this metric type can only be associated with one of the\n            monitored resource types listed here.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: The resource name of the metric descriptor.\n          x-kubernetes-immutable: true\n        type:\n          type: string\n          x-dcl-go-name: Type\n          description: 'The metric type, including its DNS name prefix. The type is\n            not URL-encoded. All user-defined metric types have the DNS name `custom.googleapis.com`\n            or `external.googleapis.com`. Metric types should use a natural hierarchical\n            grouping. For example: \"custom.googleapis.com/invoice/paid/amount\" \"external.googleapis.com/prometheus/up\"\n            \"appengine.googleapis.com/http/server/response_latencies\"'\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n        unit:\n          type: string\n          x-dcl-go-name: Unit\n          description: 'The units in which the metric value is reported. It is only\n            applicable if the `value_type` is `INT64`, `DOUBLE`, or `DISTRIBUTION`.\n            The `unit` defines the representation of the stored metric values. Different\n            systems might scale the values to be more easily displayed (so a value\n            of `0.02kBy` _might_ be displayed as `20By`, and a value of `3523kBy`\n            _might_ be displayed as `3.5MBy`). However, if the `unit` is `kBy`, then\n            the value of the metric is always in thousands of bytes, no matter how\n            it might be displayed. If you want a custom metric to record the exact\n            number of CPU-seconds used by a job, you can create an `INT64 CUMULATIVE`\n            metric whose `unit` is `s{CPU}` (or equivalently `1s{CPU}` or just `s`).\n            If the job uses 12,005 CPU-seconds, then the value is written as `12005`.\n            Alternatively, if you want a custom metric to record data in a more granular\n            way, you can create a `DOUBLE CUMULATIVE` metric whose `unit` is `ks{CPU}`,\n            and then write the value `12.005` (which is `12005/1000`), or use `Kis{CPU}`\n            and write `11.723` (which is `12005/1024`). The supported units are a\n            subset of [The Unified Code for Units of Measure](https://unitsofmeasure.org/ucum.html)\n            standard: **Basic units (UNIT)** * `bit` bit * `By` byte * `s` second\n            * `min` minute * `h` hour * `d` day * `1` dimensionless **Prefixes (PREFIX)**\n            * `k` kilo (10^3) * `M` mega (10^6) * `G` giga (10^9) * `T` tera (10^12)\n            * `P` peta (10^15) * `E` exa (10^18) * `Z` zetta (10^21) * `Y` yotta (10^24)\n            * `m` milli (10^-3) * `u` micro (10^-6) * `n` nano (10^-9) * `p` pico\n            (10^-12) * `f` femto (10^-15) * `a` atto (10^-18) * `z` zepto (10^-21)\n            * `y` yocto (10^-24) * `Ki` kibi (2^10) * `Mi` mebi (2^20) * `Gi` gibi\n            (2^30) * `Ti` tebi (2^40) * `Pi` pebi (2^50) **Grammar** The grammar also\n            includes these connectors: * `/` division or ratio (as an infix operator).\n            For examples, `kBy/{email}` or `MiBy/10ms` (although you should almost\n            never have `/s` in a metric `unit`; rates should always be computed at\n            query time from the underlying cumulative or delta value). * `.` multiplication\n            or composition (as an infix operator). For examples, `GBy.d` or `k{watt}.h`.\n            The grammar for a unit is as follows: Expression = Component: { \".\" Component\n            } { \"/\" Component } ; Component = ( [ PREFIX ] UNIT | \"%\" ) [ Annotation\n            ] | Annotation | \"1\" ; Annotation = \"{\" NAME \"}\" ; Notes: * `Annotation`\n            is just a comment if it follows a `UNIT`. If the annotation is used alone,\n            then the unit is equivalent to `1`. For examples, `{request}/s == 1/s`,\n            `By{transmitted}/s == By/s`. * `NAME` is a sequence of non-blank printable\n            ASCII characters not containing `{` or `}`. * `1` represents a unitary\n            [dimensionless unit](https://en.wikipedia.org/wiki/Dimensionless_quantity)\n            of 1, such as in `1/s`. It is typically used when none of the basic units\n            are appropriate. For example, \"new users per day\" can be represented as\n            `1/d` or `{new-users}/d` (and a metric value `5` would mean \"5 new users).\n            Alternatively, \"thousands of page views per day\" would be represented\n            as `1000/d` or `k1/d` or `k{page_views}/d` (and a metric value of `5.3`\n            would mean \"5300 page views per day\"). * `%` represents dimensionless\n            value of 1/100, and annotates values giving a percentage (so the metric\n            values are typically in the range of 0..100, and a metric value `3` means\n            \"3 percent\"). * `10^2.%` indicates a metric contains a ratio, typically\n            in the range 0..1, that will be multiplied by 100 and displayed as a percentage\n            (so a metric value `0.03` means \"3 percent\").'\n          x-kubernetes-immutable: true\n        valueType:\n          type: string\n          x-dcl-go-name: ValueType\n          x-dcl-go-type: MetricDescriptorValueTypeEnum\n          description: 'Whether the measurement is an integer, a floating-point number,\n            etc. Some combinations of `metric_kind` and `value_type` might not be\n            supported. Possible values: STRING, BOOL, INT64'\n          x-kubernetes-immutable: true\n          enum:\n          - STRING\n          - BOOL\n          - INT64\n")

// 13396 bytes
// MD5: e63f67cc50f2c7b6bab66cb9f8316c26

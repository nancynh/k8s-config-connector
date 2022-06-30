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
// gen_go_data -package beta -var YAML_uptime_check_config blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/beta/uptime_check_config.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/beta/uptime_check_config.yaml
var YAML_uptime_check_config = []byte("info:\n  title: Monitoring/UptimeCheckConfig\n  description: The Monitoring UptimeCheckConfig resource\n  x-dcl-struct-name: UptimeCheckConfig\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a UptimeCheckConfig\n    parameters:\n    - name: UptimeCheckConfig\n      required: true\n      description: A full instance of a UptimeCheckConfig\n    timeoutSecs: 0\n  apply:\n    description: The function used to apply information about a UptimeCheckConfig\n    parameters:\n    - name: UptimeCheckConfig\n      required: true\n      description: A full instance of a UptimeCheckConfig\n    timeoutSecs: 0\n  delete:\n    description: The function used to delete a UptimeCheckConfig\n    parameters:\n    - name: UptimeCheckConfig\n      required: true\n      description: A full instance of a UptimeCheckConfig\n    timeoutSecs: 0\n  deleteAll:\n    description: The function used to delete all UptimeCheckConfig\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\n  list:\n    description: The function used to list information about many UptimeCheckConfig\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\ncomponents:\n  schemas:\n    UptimeCheckConfig:\n      title: UptimeCheckConfig\n      x-dcl-id: projects/{{project}}/uptimeCheckConfigs/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - displayName\n      - timeout\n      properties:\n        contentMatchers:\n          type: array\n          x-dcl-go-name: ContentMatchers\n          description: The content that is expected to appear in the data returned\n            by the target server against which the check is run.  Currently, only\n            the first entry in the `content_matchers` list is supported, and additional\n            entries will be ignored. This field is optional and should only be specified\n            if a content match is required as part of the/ Uptime check.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: UptimeCheckConfigContentMatchers\n            required:\n            - content\n            properties:\n              content:\n                type: string\n                x-dcl-go-name: Content\n              matcher:\n                type: string\n                x-dcl-go-name: Matcher\n                x-dcl-go-type: UptimeCheckConfigContentMatchersMatcherEnum\n                description: ' Possible values: CONTENT_MATCHER_OPTION_UNSPECIFIED,\n                  CONTAINS_STRING, NOT_CONTAINS_STRING, MATCHES_REGEX, NOT_MATCHES_REGEX'\n                default: CONTAINS_STRING\n                enum:\n                - CONTENT_MATCHER_OPTION_UNSPECIFIED\n                - CONTAINS_STRING\n                - NOT_CONTAINS_STRING\n                - MATCHES_REGEX\n                - NOT_MATCHES_REGEX\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: A human-friendly name for the Uptime check configuration. The\n            display name should be unique within a Stackdriver Workspace in order\n            to make it easier to identify; however, uniqueness is not enforced. Required.\n        httpCheck:\n          type: object\n          x-dcl-go-name: HttpCheck\n          x-dcl-go-type: UptimeCheckConfigHttpCheck\n          description: Contains information needed to make an HTTP or HTTPS check.\n          x-dcl-conflicts:\n          - tcpCheck\n          properties:\n            authInfo:\n              type: object\n              x-dcl-go-name: AuthInfo\n              x-dcl-go-type: UptimeCheckConfigHttpCheckAuthInfo\n              description: The authentication information. Optional when creating\n                an HTTP check; defaults to empty.\n              required:\n              - username\n              - password\n              properties:\n                password:\n                  type: string\n                  x-dcl-go-name: Password\n                  x-dcl-sensitive: true\n                  x-dcl-mutable-unreadable: true\n                username:\n                  type: string\n                  x-dcl-go-name: Username\n            body:\n              type: string\n              x-dcl-go-name: Body\n              description: 'The request body associated with the HTTP POST request.\n                If `content_type` is `URL_ENCODED`, the body passed in must be URL-encoded.\n                Users can provide a `Content-Length` header via the `headers` field\n                or the API will do so. If the `request_method` is `GET` and `body`\n                is not empty, the API will return an error. The maximum byte size\n                is 1 megabyte. Note: As with all `bytes` fields JSON representations\n                are base64 encoded. e.g.: \"foo=bar\" in URL-encoded form is \"foo%3Dbar\"\n                and in base64 encoding is \"Zm9vJTI1M0RiYXI=\".'\n            contentType:\n              type: string\n              x-dcl-go-name: ContentType\n              x-dcl-go-type: UptimeCheckConfigHttpCheckContentTypeEnum\n              description: 'The content type to use for the check.  Possible values:\n                TYPE_UNSPECIFIED, URL_ENCODED'\n              x-kubernetes-immutable: true\n              enum:\n              - TYPE_UNSPECIFIED\n              - URL_ENCODED\n            headers:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Headers\n              description: The list of headers to send as part of the Uptime check\n                request. If two headers have the same key and different values, they\n                should be entered as a single header, with the value being a comma-separated\n                list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt\n                (page 31). Entering two separate headers with the same key in a Create\n                call will cause the first to be overwritten by the second. The maximum\n                number of headers allowed is 100.\n              x-dcl-server-default: true\n              x-dcl-mutable-unreadable: true\n            maskHeaders:\n              type: boolean\n              x-dcl-go-name: MaskHeaders\n              description: Boolean specifying whether to encrypt the header information.\n                Encryption should be specified for any headers related to authentication\n                that you do not wish to be seen when retrieving the configuration.\n                The server will be responsible for encrypting the headers. On Get/List\n                calls, if `mask_headers` is set to `true` then the headers will be\n                obscured with `******.`\n              x-kubernetes-immutable: true\n            path:\n              type: string\n              x-dcl-go-name: Path\n              description: Optional (defaults to \"/\"). The path to the page against\n                which to run the check. Will be combined with the `host` (specified\n                within the `monitored_resource`) and `port` to construct the full\n                URL. If the provided path does not begin with \"/\", a \"/\" will be prepended\n                automatically.\n              default: /\n            port:\n              type: integer\n              format: int64\n              x-dcl-go-name: Port\n              description: Optional (defaults to 80 when `use_ssl` is `false`, and\n                443 when `use_ssl` is `true`). The TCP port on the HTTP server against\n                which to run the check. Will be combined with host (specified within\n                the `monitored_resource`) and `path` to construct the full URL.\n              x-dcl-server-default: true\n            requestMethod:\n              type: string\n              x-dcl-go-name: RequestMethod\n              x-dcl-go-type: UptimeCheckConfigHttpCheckRequestMethodEnum\n              description: The HTTP request method to use for the check. If set to\n                `METHOD_UNSPECIFIED` then `request_method` defaults to `GET`.\n              x-kubernetes-immutable: true\n              default: GET\n              enum:\n              - METHOD_UNSPECIFIED\n              - GET\n              - POST\n            useSsl:\n              type: boolean\n              x-dcl-go-name: UseSsl\n              description: If `true`, use HTTPS instead of HTTP to run the check.\n            validateSsl:\n              type: boolean\n              x-dcl-go-name: ValidateSsl\n              description: Boolean specifying whether to include SSL certificate validation\n                as a part of the Uptime check. Only applies to checks where `monitored_resource`\n                is set to `uptime_url`. If `use_ssl` is `false`, setting `validate_ssl`\n                to `true` has no effect.\n        monitoredResource:\n          type: object\n          x-dcl-go-name: MonitoredResource\n          x-dcl-go-type: UptimeCheckConfigMonitoredResource\n          description: 'The [monitored resource](https://cloud.google.com/monitoring/api/resources)\n            associated with the configuration. The following monitored resource types\n            are supported for Uptime checks:   `uptime_url`,   `gce_instance`,   `gae_app`,   `aws_ec2_instance`,   `aws_elb_load_balancer`'\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - resourceGroup\n          required:\n          - type\n          - filterLabels\n          properties:\n            filterLabels:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: FilterLabels\n              x-kubernetes-immutable: true\n            type:\n              type: string\n              x-dcl-go-name: Type\n              x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'A unique resource name for this Uptime check configuration.\n            The format is: projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID]\n            This field should be omitted when creating the Uptime check configuration;\n            on create, the resource name is assigned by the server and included in\n            the response.'\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        period:\n          type: string\n          x-dcl-go-name: Period\n          description: How often, in seconds, the Uptime check is performed. Currently,\n            the only supported values are `60s` (1 minute), `300s` (5 minutes), `600s`\n            (10 minutes), and `900s` (15 minutes). Optional, defaults to `60s`.\n          default: 60s\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for this uptime check config.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        resourceGroup:\n          type: object\n          x-dcl-go-name: ResourceGroup\n          x-dcl-go-type: UptimeCheckConfigResourceGroup\n          description: The group resource associated with the configuration.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - monitoredResource\n          properties:\n            groupId:\n              type: string\n              x-dcl-go-name: GroupId\n              description: The group of resources being monitored. Should be only\n                the `[GROUP_ID]`, and not the full-path `projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID]`.\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Monitoring/Group\n                field: name\n            resourceType:\n              type: string\n              x-dcl-go-name: ResourceType\n              x-dcl-go-type: UptimeCheckConfigResourceGroupResourceTypeEnum\n              description: 'The resource type of the group members. Possible values:\n                RESOURCE_TYPE_UNSPECIFIED, INSTANCE, AWS_ELB_LOAD_BALANCER'\n              x-kubernetes-immutable: true\n              enum:\n              - RESOURCE_TYPE_UNSPECIFIED\n              - INSTANCE\n              - AWS_ELB_LOAD_BALANCER\n        selectedRegions:\n          type: array\n          x-dcl-go-name: SelectedRegions\n          description: The list of regions from which the check will be run. Some\n            regions contain one location, and others contain more than one. If this\n            field is specified, enough regions must be provided to include a minimum\n            of 3 locations.  Not specifying this field will result in Uptime checks\n            running from all available regions.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        tcpCheck:\n          type: object\n          x-dcl-go-name: TcpCheck\n          x-dcl-go-type: UptimeCheckConfigTcpCheck\n          description: Contains information needed to make a TCP check.\n          x-dcl-conflicts:\n          - httpCheck\n          required:\n          - port\n          properties:\n            port:\n              type: integer\n              format: int64\n              x-dcl-go-name: Port\n              description: The TCP port on the server against which to run the check.\n                Will be combined with host (specified within the `monitored_resource`)\n                to construct the full URL. Required.\n        timeout:\n          type: string\n          x-dcl-go-name: Timeout\n          description: The maximum amount of time to wait for the request to complete\n            (must be between 1 and 60 seconds). Required.\n")

// 13852 bytes
// MD5: 8b6229e892ca346bced438e9540c3516

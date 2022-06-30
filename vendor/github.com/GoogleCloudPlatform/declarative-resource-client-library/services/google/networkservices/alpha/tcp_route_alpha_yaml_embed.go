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
// gen_go_data -package alpha -var YAML_tcp_route blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/tcp_route.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/tcp_route.yaml
var YAML_tcp_route = []byte("info:\n  title: NetworkServices/TcpRoute\n  description: The NetworkServices TcpRoute resource\n  x-dcl-struct-name: TcpRoute\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a TcpRoute\n    parameters:\n    - name: TcpRoute\n      required: true\n      description: A full instance of a TcpRoute\n    timeoutSecs: 0\n  apply:\n    description: The function used to apply information about a TcpRoute\n    parameters:\n    - name: TcpRoute\n      required: true\n      description: A full instance of a TcpRoute\n    timeoutSecs: 0\n  delete:\n    description: The function used to delete a TcpRoute\n    parameters:\n    - name: TcpRoute\n      required: true\n      description: A full instance of a TcpRoute\n    timeoutSecs: 0\n  deleteAll:\n    description: The function used to delete all TcpRoute\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\n  list:\n    description: The function used to list information about many TcpRoute\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\ncomponents:\n  schemas:\n    TcpRoute:\n      title: TcpRoute\n      x-dcl-id: projects/{{project}}/locations/{{location}}/tcpRoutes/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - rules\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A free-text description of the resource. Max length\n            1024 characters.\n        gateways:\n          type: array\n          x-dcl-go-name: Gateways\n          description: 'Optional. Gateways defines a list of gateways this TcpRoute\n            is attached to, as one of the routing rules to route the requests served\n            by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Networkservices/Gateway\n              field: selfLink\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the TcpRoute resource.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        meshes:\n          type: array\n          x-dcl-go-name: Meshes\n          description: 'Optional. Meshes defines a list of meshes this TcpRoute is\n            attached to, as one of the routing rules to route the requests served\n            by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`\n            The attached Mesh should be of a type SIDECAR'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Networkservices/Mesh\n              field: selfLink\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the TcpRoute resource. It matches pattern\n            `projects/*/locations/global/tcpRoutes/tcp_route_name>`.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        routers:\n          type: array\n          x-dcl-go-name: Routers\n          description: 'Optional. Routers define a list of routers this TcpRoute should\n            be served by. Each router reference should match the pattern: `projects/*/locations/global/routers/`\n            The attached Router should be of a type PROXY'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        rules:\n          type: array\n          x-dcl-go-name: Rules\n          description: Required. Rules that define how traffic is routed and handled.\n            At least one RouteRule must be supplied. If there are multiple rules then\n            the action taken will be the first rule to match.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: TcpRouteRules\n            required:\n            - action\n            properties:\n              action:\n                type: object\n                x-dcl-go-name: Action\n                x-dcl-go-type: TcpRouteRulesAction\n                description: Required. The detailed rule defining how to route matched\n                  traffic.\n                properties:\n                  destinations:\n                    type: array\n                    x-dcl-go-name: Destinations\n                    description: Optional. The destination services to which traffic\n                      should be forwarded. At least one destination service is required.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: object\n                      x-dcl-go-type: TcpRouteRulesActionDestinations\n                      required:\n                      - serviceName\n                      properties:\n                        serviceName:\n                          type: string\n                          x-dcl-go-name: ServiceName\n                          description: Required. The URL of a BackendService to route\n                            traffic to.\n                          x-dcl-references:\n                          - resource: Compute/BackendService\n                            field: name\n                            format: projects/{{project}}/locations/global/backendServices/{{name}}\n                        weight:\n                          type: integer\n                          format: int64\n                          x-dcl-go-name: Weight\n                          description: 'Optional. Specifies the proportion of requests\n                            forwarded to the backend referenced by the serviceName\n                            field. This is computed as: weight/Sum(weights in this\n                            destination list). For non-zero values, there may be some\n                            epsilon from the exact proportion defined here depending\n                            on the precision an implementation supports. If only one\n                            serviceName is specified and it has a weight greater than\n                            0, 100% of the traffic is forwarded to that backend. If\n                            weights are specified for any one service name, they need\n                            to be specified for all of them. If weights are unspecified\n                            for all services, then, traffic is distributed in equal\n                            proportions to all of them.'\n                  originalDestination:\n                    type: boolean\n                    x-dcl-go-name: OriginalDestination\n                    description: Optional. If true, Router will use the destination\n                      IP and port of the original connection as the destination of\n                      the request. Default is false.\n              matches:\n                type: array\n                x-dcl-go-name: Matches\n                description: Optional. RouteMatch defines the predicate used to match\n                  requests to a given action. Multiple match types are “OR”ed for\n                  evaluation. If no routeMatch field is specified, this rule will\n                  unconditionally match traffic.\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: TcpRouteRulesMatches\n                  required:\n                  - address\n                  - port\n                  properties:\n                    address:\n                      type: string\n                      x-dcl-go-name: Address\n                      description: 'Required. Must be specified in the CIDR range\n                        format. A CIDR range consists of an IP Address and a prefix\n                        length to construct the subnet mask. By default, the prefix\n                        length is 32 (i.e. matches a single IP address). Only IPV4\n                        addresses are supported. Examples: “10.0.0.1” - matches against\n                        this exact IP address. “10.0.0.0/8\" - matches against any\n                        IP address within the 10.0.0.0 subnet and 255.255.255.0 mask.\n                        \"0.0.0.0/0\" - matches against any IP address''.'\n                    port:\n                      type: string\n                      x-dcl-go-name: Port\n                      description: Required. Specifies the destination port to match\n                        against.\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Output only. Server-defined URL of this resource\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 10261 bytes
// MD5: fb5635be1e0e1bab01812d3ed2630a6b

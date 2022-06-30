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
// gen_go_data -package beta -var YAML_environment_group_attachment blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/beta/environment_group_attachment.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/beta/environment_group_attachment.yaml
var YAML_environment_group_attachment = []byte("info:\n  title: Apigee/EnvironmentGroupAttachment\n  description: The Apigee EnvironmentGroupAttachment resource\n  x-dcl-struct-name: EnvironmentGroupAttachment\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a EnvironmentGroupAttachment\n    parameters:\n    - name: EnvironmentGroupAttachment\n      required: true\n      description: A full instance of a EnvironmentGroupAttachment\n    timeoutSecs: 0\n  apply:\n    description: The function used to apply information about a EnvironmentGroupAttachment\n    parameters:\n    - name: EnvironmentGroupAttachment\n      required: true\n      description: A full instance of a EnvironmentGroupAttachment\n    timeoutSecs: 0\n  delete:\n    description: The function used to delete a EnvironmentGroupAttachment\n    parameters:\n    - name: EnvironmentGroupAttachment\n      required: true\n      description: A full instance of a EnvironmentGroupAttachment\n    timeoutSecs: 0\n  deleteAll:\n    description: The function used to delete all EnvironmentGroupAttachment\n    parameters:\n    - name: envgroup\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\n  list:\n    description: The function used to list information about many EnvironmentGroupAttachment\n    parameters:\n    - name: envgroup\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\ncomponents:\n  schemas:\n    EnvironmentGroupAttachment:\n      title: EnvironmentGroupAttachment\n      x-dcl-id: '{{envgroup}}/attachments/{{name}}'\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - environment\n      - envgroup\n      properties:\n        createdAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: CreatedAt\n          readOnly: true\n          x-kubernetes-immutable: true\n        envgroup:\n          type: string\n          x-dcl-go-name: Envgroup\n          description: The environment group for the resource\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n          x-dcl-references:\n          - resource: Apigee/EnvironmentGroup\n            field: name\n            parent: true\n        environment:\n          type: string\n          x-dcl-go-name: Environment\n          description: Required. ID of the attached environment.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Apigee/Environment\n            field: name\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: ID of the environment group attachment.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n")

// 2667 bytes
// MD5: 96d625568855b5802db9dbbf9414b0fe

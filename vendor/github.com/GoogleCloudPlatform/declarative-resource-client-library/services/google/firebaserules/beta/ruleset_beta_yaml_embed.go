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
// gen_go_data -package beta -var YAML_ruleset blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebaserules/beta/ruleset.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebaserules/beta/ruleset.yaml
var YAML_ruleset = []byte("info:\n  title: Firebaserules/Ruleset\n  description: \"\"\n  x-dcl-struct-name: Ruleset\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: Firebase Ruleset API Documentation\n    url: https://firebase.google.com/docs/reference/rules/rest#rest-resource:-v1.projects.rulesets\n  x-dcl-guides:\n  - text: Get started with Firebase Security Rules\n    url: https://firebase.google.com/docs/rules/get-started\npaths:\n  get:\n    description: The function used to get information about a Ruleset\n    parameters:\n    - name: Ruleset\n      required: true\n      description: A full instance of a Ruleset\n    timeoutSecs: 0\n  apply:\n    description: The function used to apply information about a Ruleset\n    parameters:\n    - name: Ruleset\n      required: true\n      description: A full instance of a Ruleset\n    timeoutSecs: 0\n  delete:\n    description: The function used to delete a Ruleset\n    parameters:\n    - name: Ruleset\n      required: true\n      description: A full instance of a Ruleset\n    timeoutSecs: 0\n  deleteAll:\n    description: The function used to delete all Ruleset\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\n  list:\n    description: The function used to list information about many Ruleset\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    timeoutSecs: 0\ncomponents:\n  schemas:\n    Ruleset:\n      title: Ruleset\n      x-dcl-id: projects/{{project}}/rulesets/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - source\n      - project\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Time the `Ruleset` was created.\n          x-kubernetes-immutable: true\n        metadata:\n          type: object\n          x-dcl-go-name: Metadata\n          x-dcl-go-type: RulesetMetadata\n          readOnly: true\n          description: Output only. The metadata for this ruleset.\n          x-kubernetes-immutable: true\n          properties:\n            services:\n              type: array\n              x-dcl-go-name: Services\n              description: Services that this ruleset has declarations for (e.g.,\n                \"cloud.firestore\"). There may be 0+ of these.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output only. Name of the `Ruleset`. The ruleset_id is auto\n            generated by the service. Format: `projects/{project_id}/rulesets/{ruleset_id}`'\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        source:\n          type: object\n          x-dcl-go-name: Source\n          x-dcl-go-type: RulesetSource\n          description: '`Source` for the `Ruleset`.'\n          x-kubernetes-immutable: true\n          required:\n          - files\n          properties:\n            files:\n              type: array\n              x-dcl-go-name: Files\n              description: '`File` set constituting the `Source` bundle.'\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: RulesetSourceFiles\n                required:\n                - content\n                - name\n                properties:\n                  content:\n                    type: string\n                    x-dcl-go-name: Content\n                    description: Textual Content.\n                    x-kubernetes-immutable: true\n                  fingerprint:\n                    type: string\n                    x-dcl-go-name: Fingerprint\n                    description: Fingerprint (e.g. github sha) associated with the\n                      `File`.\n                    x-kubernetes-immutable: true\n                  name:\n                    type: string\n                    x-dcl-go-name: Name\n                    description: File name.\n                    x-kubernetes-immutable: true\n            language:\n              type: string\n              x-dcl-go-name: Language\n              x-dcl-go-type: RulesetSourceLanguageEnum\n              description: '`Language` of the `Source` bundle. If unspecified, the\n                language will default to `FIREBASE_RULES`. Possible values: LANGUAGE_UNSPECIFIED,\n                FIREBASE_RULES, EVENT_FLOW_TRIGGERS'\n              x-kubernetes-immutable: true\n              enum:\n              - LANGUAGE_UNSPECIFIED\n              - FIREBASE_RULES\n              - EVENT_FLOW_TRIGGERS\n")

// 5165 bytes
// MD5: fbeadbf74431b8ba6005ea5a9448e25e

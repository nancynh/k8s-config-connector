// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/fileutil"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

type templateData struct {
	Resources []schema.GroupVersionKind
}

func main() {
	if err := clearGeneratedDocDir(); err != nil {
		log.Fatalf("error clearing generated doc dir: %v", err)
	}

	smLoader, err := servicemappingloader.New()
	if err != nil {
		log.Fatalf("error creating service mapping loader: %v", err)
	}
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		log.Fatalf("error creating a DCL schema loader: %v", err)
	}

	if err := generateListOfResourcesWithServiceGeneratedResourceID(smLoader, serviceMetadataLoader, dclSchemaLoader); err != nil {
		log.Fatal(err)
	}
	if err := generateListOfUnacquirableResources(smLoader); err != nil {
		log.Fatal(err)
	}
}

func generateListOfResourcesWithServiceGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) error {
	// Use "service-generated" instead of "server-generated" in the output file
	// name since the former is the terminology used by the public docs.
	const outputFileName = "_resources-with-service-generated-resource-id.html"
	resources, err := resourcesWithServerGeneratedResourceID(smLoader, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return fmt.Errorf("eror getting resources with a server-generated resource ID: %v", err)
	}
	return generateListOfResources(resources, outputFileName)
}

func generateListOfUnacquirableResources(smLoader *servicemappingloader.ServiceMappingLoader) error {
	const outputFileName = "_unacquirable-resources.html"
	resources := unacquirableResources(smLoader)
	return generateListOfResources(resources, outputFileName)
}

func generateListOfResources(resources []schema.GroupVersionKind, outputFileName string) error {
	outputFilePath := path.Join(repo.GetG3ResourceListsGeneratedPath(), outputFileName)
	outputFile, err := fileutil.NewEmptyFile(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating empty file for output: %v", err)
	}

	templateFilePath := repo.GetG3ResourceListsTemplatePath()
	template, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("error parsing template file: %v", err)
	}
	templateData := &templateData{
		Resources: resources,
	}
	if err := template.Execute(outputFile, templateData); err != nil {
		return fmt.Errorf("error while executing template: %v", err)
	}
	return nil
}

func resourcesWithServerGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	tfResources := tfBasedResourcesWithServerGeneratedResourceID(smLoader)
	dclResources, err := dclBasedResourcesWithServerGeneratedResourceID(serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting DCL-based resources with a server-generated resource ID: %v", err)
	}
	return k8s.SortGVKsByKind(append(tfResources, dclResources...)), nil
}

func tfBasedResourcesWithServerGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	gvkList := make([]schema.GroupVersionKind, 0)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			if krmtotf.SupportsResourceIDField(&rc) && krmtotf.IsResourceIDFieldServerGenerated(&rc) {
				gvk := krmtotf.GVKForResource(&sm, &rc)
				gvkList = append(gvkList, gvk)
			}
		}
	}
	return gvkList
}

func dclBasedResourcesWithServerGeneratedResourceID(serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	gvkList := make([]schema.GroupVersionKind, 0)
	for _, s := range serviceMetadataLoader.GetAllServiceMetadata() {
		for _, r := range s.Resources {
			if !r.Releasable {
				continue
			}
			gvk := dclmetadata.GVKForResource(s, r)
			schema, err := dclschemaloader.GetDCLSchemaForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
			if err != nil {
				return nil, fmt.Errorf("error getting DCL schema for GroupVersionKind %v: %v", gvk, err)
			}
			nameFieldSchema, ok := dclextension.GetNameFieldSchema(schema)
			if !ok {
				// Resource doesn't support the resourceID field
				continue
			}
			isServerGenerated, err := dclextension.IsResourceIDFieldServerGenerated(nameFieldSchema)
			if err != nil {
				return nil, fmt.Errorf("error determining if resourceID is server-generated for GroupVersionKind %v: %v", gvk, err)
			}
			if !isServerGenerated {
				continue
			}
			gvkList = append(gvkList, gvk)
		}
	}
	return gvkList, nil
}

func unacquirableResources(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	// Only TF resources can be unacquirable because only TF resources can have
	// a server-generated ID that cannot be set via the resourceID field.
	gvkList := make([]schema.GroupVersionKind, 0)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			if !krmtotf.SupportsResourceIDField(&rc) && krmtotf.SupportsServerGeneratedIDField(&rc) {
				gvk := krmtotf.GVKForResource(&sm, &rc)
				gvkList = append(gvkList, gvk)
			}
		}
	}
	return k8s.SortGVKsByKind(gvkList)
}

func clearGeneratedDocDir() error {
	docDir := repo.GetG3ResourceListsGeneratedPath()
	if err := os.RemoveAll(docDir); err != nil {
		return fmt.Errorf("error deleting generated doc dir at %v: %v", docDir, err)
	}
	if err := os.Mkdir(docDir, 0700); err != nil {
		return fmt.Errorf("error recreating generated doc dir at %v: %v", docDir, err)
	}
	return nil
}

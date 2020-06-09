// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinition) DeepCopyInto(out *ComponentDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.CustomTypes != nil {
		in, out := &in.CustomTypes, &out.CustomTypes
		*out = make([]CustomType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Imports != nil {
		in, out := &in.Imports, &out.Imports
		*out = make([]DefinitionImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Exports != nil {
		in, out := &in.Exports, &out.Exports
		*out = make([]DefinitionExport, len(*in))
		copy(*out, *in)
	}
	if in.DefinitionReferences != nil {
		in, out := &in.DefinitionReferences, &out.DefinitionReferences
		*out = make([]DefinitionReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinition.
func (in *ComponentDefinition) DeepCopy() *ComponentDefinition {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentInstallation) DeepCopyInto(out *ComponentInstallation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentInstallation.
func (in *ComponentInstallation) DeepCopy() *ComponentInstallation {
	if in == nil {
		return nil
	}
	out := new(ComponentInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentInstallation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentInstallationList) DeepCopyInto(out *ComponentInstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComponentInstallation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentInstallationList.
func (in *ComponentInstallationList) DeepCopy() *ComponentInstallationList {
	if in == nil {
		return nil
	}
	out := new(ComponentInstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentInstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentInstallationSpec) DeepCopyInto(out *ComponentInstallationSpec) {
	*out = *in
	if in.Imports != nil {
		in, out := &in.Imports, &out.Imports
		*out = make([]DefinitionImportMapping, len(*in))
		copy(*out, *in)
	}
	if in.Exports != nil {
		in, out := &in.Exports, &out.Exports
		*out = make([]DefinitionExportMapping, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentInstallationSpec.
func (in *ComponentInstallationSpec) DeepCopy() *ComponentInstallationSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentInstallationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentInstallationStatus) DeepCopyInto(out *ComponentInstallationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ExportReference = in.ExportReference
	if in.Imports != nil {
		in, out := &in.Imports, &out.Imports
		*out = make([]ImportState, len(*in))
		copy(*out, *in)
	}
	if in.InstallationReferences != nil {
		in, out := &in.InstallationReferences, &out.InstallationReferences
		*out = make([]NamedObjectReference, len(*in))
		copy(*out, *in)
	}
	out.ExecutionReference = in.ExecutionReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentInstallationStatus.
func (in *ComponentInstallationStatus) DeepCopy() *ComponentInstallationStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentInstallationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	if in.Codes != nil {
		in, out := &in.Codes, &out.Codes
		*out = make([]ErrorCode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomType) DeepCopyInto(out *CustomType) {
	*out = *in
	in.OpenAPIV3Schema.DeepCopyInto(&out.OpenAPIV3Schema)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomType.
func (in *CustomType) DeepCopy() *CustomType {
	if in == nil {
		return nil
	}
	out := new(CustomType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataType) DeepCopyInto(out *DataType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Scheme.DeepCopyInto(&out.Scheme)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataType.
func (in *DataType) DeepCopy() *DataType {
	if in == nil {
		return nil
	}
	out := new(DataType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataTypeList) DeepCopyInto(out *DataTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataTypeList.
func (in *DataTypeList) DeepCopy() *DataTypeList {
	if in == nil {
		return nil
	}
	out := new(DataTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataTypeScheme) DeepCopyInto(out *DataTypeScheme) {
	*out = *in
	in.OpenAPIV3Schema.DeepCopyInto(&out.OpenAPIV3Schema)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataTypeScheme.
func (in *DataTypeScheme) DeepCopy() *DataTypeScheme {
	if in == nil {
		return nil
	}
	out := new(DataTypeScheme)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Default) DeepCopyInto(out *Default) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Default.
func (in *Default) DeepCopy() *Default {
	if in == nil {
		return nil
	}
	out := new(Default)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionExport) DeepCopyInto(out *DefinitionExport) {
	*out = *in
	out.DefinitionFieldValue = in.DefinitionFieldValue
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionExport.
func (in *DefinitionExport) DeepCopy() *DefinitionExport {
	if in == nil {
		return nil
	}
	out := new(DefinitionExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionExportMapping) DeepCopyInto(out *DefinitionExportMapping) {
	*out = *in
	out.DefinitionFieldMapping = in.DefinitionFieldMapping
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionExportMapping.
func (in *DefinitionExportMapping) DeepCopy() *DefinitionExportMapping {
	if in == nil {
		return nil
	}
	out := new(DefinitionExportMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionFieldMapping) DeepCopyInto(out *DefinitionFieldMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionFieldMapping.
func (in *DefinitionFieldMapping) DeepCopy() *DefinitionFieldMapping {
	if in == nil {
		return nil
	}
	out := new(DefinitionFieldMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionFieldValue) DeepCopyInto(out *DefinitionFieldValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionFieldValue.
func (in *DefinitionFieldValue) DeepCopy() *DefinitionFieldValue {
	if in == nil {
		return nil
	}
	out := new(DefinitionFieldValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionImport) DeepCopyInto(out *DefinitionImport) {
	*out = *in
	out.DefinitionFieldValue = in.DefinitionFieldValue
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
	in.Default.DeepCopyInto(&out.Default)
	if in.ConditionalImports != nil {
		in, out := &in.ConditionalImports, &out.ConditionalImports
		*out = make([]DefinitionImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionImport.
func (in *DefinitionImport) DeepCopy() *DefinitionImport {
	if in == nil {
		return nil
	}
	out := new(DefinitionImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionImportMapping) DeepCopyInto(out *DefinitionImportMapping) {
	*out = *in
	out.DefinitionFieldMapping = in.DefinitionFieldMapping
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionImportMapping.
func (in *DefinitionImportMapping) DeepCopy() *DefinitionImportMapping {
	if in == nil {
		return nil
	}
	out := new(DefinitionImportMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionReference) DeepCopyInto(out *DefinitionReference) {
	*out = *in
	if in.Imports != nil {
		in, out := &in.Imports, &out.Imports
		*out = make([]DefinitionImportMapping, len(*in))
		copy(*out, *in)
	}
	if in.Exports != nil {
		in, out := &in.Exports, &out.Exports
		*out = make([]DefinitionExportMapping, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionReference.
func (in *DefinitionReference) DeepCopy() *DefinitionReference {
	if in == nil {
		return nil
	}
	out := new(DefinitionReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployItem) DeepCopyInto(out *DeployItem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployItem.
func (in *DeployItem) DeepCopy() *DeployItem {
	if in == nil {
		return nil
	}
	out := new(DeployItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeployItem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployItemList) DeepCopyInto(out *DeployItemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeployItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployItemList.
func (in *DeployItemList) DeepCopy() *DeployItemList {
	if in == nil {
		return nil
	}
	out := new(DeployItemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeployItemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployItemSpec) DeepCopyInto(out *DeployItemSpec) {
	*out = *in
	out.ImportReferences = in.ImportReferences
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployItemSpec.
func (in *DeployItemSpec) DeepCopy() *DeployItemSpec {
	if in == nil {
		return nil
	}
	out := new(DeployItemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployItemStatus) DeepCopyInto(out *DeployItemStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExportReference != nil {
		in, out := &in.ExportReference, &out.ExportReference
		*out = new(ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployItemStatus.
func (in *DeployItemStatus) DeepCopy() *DeployItemStatus {
	if in == nil {
		return nil
	}
	out := new(DeployItemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Execution) DeepCopyInto(out *Execution) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Execution.
func (in *Execution) DeepCopy() *Execution {
	if in == nil {
		return nil
	}
	out := new(Execution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Execution) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionItem) DeepCopyInto(out *ExecutionItem) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionItem.
func (in *ExecutionItem) DeepCopy() *ExecutionItem {
	if in == nil {
		return nil
	}
	out := new(ExecutionItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionList) DeepCopyInto(out *ExecutionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Execution, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionList.
func (in *ExecutionList) DeepCopy() *ExecutionList {
	if in == nil {
		return nil
	}
	out := new(ExecutionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExecutionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionSpec) DeepCopyInto(out *ExecutionSpec) {
	*out = *in
	if in.Executions != nil {
		in, out := &in.Executions, &out.Executions
		*out = make([]ExecutionItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionSpec.
func (in *ExecutionSpec) DeepCopy() *ExecutionSpec {
	if in == nil {
		return nil
	}
	out := new(ExecutionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionStatus) DeepCopyInto(out *ExecutionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ExportReference = in.ExportReference
	if in.DeployItemReferences != nil {
		in, out := &in.DeployItemReferences, &out.DeployItemReferences
		*out = make([]NamedObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionStatus.
func (in *ExecutionStatus) DeepCopy() *ExecutionStatus {
	if in == nil {
		return nil
	}
	out := new(ExecutionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportState) DeepCopyInto(out *ImportState) {
	*out = *in
	out.InstallationRef = in.InstallationRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportState.
func (in *ImportState) DeepCopy() *ImportState {
	if in == nil {
		return nil
	}
	out := new(ImportState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscapeConfiguration) DeepCopyInto(out *LandscapeConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscapeConfiguration.
func (in *LandscapeConfiguration) DeepCopy() *LandscapeConfiguration {
	if in == nil {
		return nil
	}
	out := new(LandscapeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LandscapeConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscapeConfigurationSpec) DeepCopyInto(out *LandscapeConfigurationSpec) {
	*out = *in
	if in.SecretReferences != nil {
		in, out := &in.SecretReferences, &out.SecretReferences
		*out = make([]ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscapeConfigurationSpec.
func (in *LandscapeConfigurationSpec) DeepCopy() *LandscapeConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(LandscapeConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscapeConfigurationStatus) DeepCopyInto(out *LandscapeConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigReference != nil {
		in, out := &in.ConfigReference, &out.ConfigReference
		*out = new(ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscapeConfigurationStatus.
func (in *LandscapeConfigurationStatus) DeepCopy() *LandscapeConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(LandscapeConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedObjectReference) DeepCopyInto(out *NamedObjectReference) {
	*out = *in
	out.Reference = in.Reference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedObjectReference.
func (in *NamedObjectReference) DeepCopy() *NamedObjectReference {
	if in == nil {
		return nil
	}
	out := new(NamedObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenAPIV3Schema) DeepCopyInto(out *OpenAPIV3Schema) {
	*out = *in
	in.JSONSchemaProps.DeepCopyInto(&out.JSONSchemaProps)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenAPIV3Schema.
func (in *OpenAPIV3Schema) DeepCopy() *OpenAPIV3Schema {
	if in == nil {
		return nil
	}
	out := new(OpenAPIV3Schema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}
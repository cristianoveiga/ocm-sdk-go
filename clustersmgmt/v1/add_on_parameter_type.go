/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// AddOnParameterKind is the name of the type used to represent objects
// of type 'add_on_parameter'.
const AddOnParameterKind = "AddOnParameter"

// AddOnParameterLinkKind is the name of the type used to represent links
// to objects of type 'add_on_parameter'.
const AddOnParameterLinkKind = "AddOnParameterLink"

// AddOnParameterNilKind is the name of the type used to nil references
// to objects of type 'add_on_parameter'.
const AddOnParameterNilKind = "AddOnParameterNil"

// AddOnParameter represents the values of the 'add_on_parameter' type.
//
// Representation of an add-on parameter.
type AddOnParameter struct {
	id           *string
	href         *string
	link         bool
	addon        *AddOn
	defaultValue *string
	description  *string
	editable     *bool
	enabled      *bool
	name         *string
	required     *bool
	validation   *string
	valueType    *string
}

// Kind returns the name of the type of the object.
func (o *AddOnParameter) Kind() string {
	if o == nil {
		return AddOnParameterNilKind
	}
	if o.link {
		return AddOnParameterLinkKind
	}
	return AddOnParameterKind
}

// ID returns the identifier of the object.
func (o *AddOnParameter) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *AddOnParameter) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *AddOnParameter) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *AddOnParameter) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *AddOnParameter) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *AddOnParameter) Empty() bool {
	return o == nil || (o.id == nil &&
		o.defaultValue == nil &&
		o.description == nil &&
		o.editable == nil &&
		o.enabled == nil &&
		o.name == nil &&
		o.required == nil &&
		o.validation == nil &&
		o.valueType == nil &&
		true)
}

// Addon returns the value of the 'addon' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Link to add-on.
func (o *AddOnParameter) Addon() *AddOn {
	if o == nil {
		return nil
	}
	return o.addon
}

// GetAddon returns the value of the 'addon' attribute and
// a flag indicating if the attribute has a value.
//
// Link to add-on.
func (o *AddOnParameter) GetAddon() (value *AddOn, ok bool) {
	ok = o != nil && o.addon != nil
	if ok {
		value = o.addon
	}
	return
}

// DefaultValue returns the value of the 'default_value' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates the value default for the add-on parameter
func (o *AddOnParameter) DefaultValue() string {
	if o != nil && o.defaultValue != nil {
		return *o.defaultValue
	}
	return ""
}

// GetDefaultValue returns the value of the 'default_value' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates the value default for the add-on parameter
func (o *AddOnParameter) GetDefaultValue() (value string, ok bool) {
	ok = o != nil && o.defaultValue != nil
	if ok {
		value = *o.defaultValue
	}
	return
}

// Description returns the value of the 'description' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Description of the add-on parameter.
func (o *AddOnParameter) Description() string {
	if o != nil && o.description != nil {
		return *o.description
	}
	return ""
}

// GetDescription returns the value of the 'description' attribute and
// a flag indicating if the attribute has a value.
//
// Description of the add-on parameter.
func (o *AddOnParameter) GetDescription() (value string, ok bool) {
	ok = o != nil && o.description != nil
	if ok {
		value = *o.description
	}
	return
}

// Editable returns the value of the 'editable' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates if this parameter can be edited after creation.
func (o *AddOnParameter) Editable() bool {
	if o != nil && o.editable != nil {
		return *o.editable
	}
	return false
}

// GetEditable returns the value of the 'editable' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates if this parameter can be edited after creation.
func (o *AddOnParameter) GetEditable() (value bool, ok bool) {
	ok = o != nil && o.editable != nil
	if ok {
		value = *o.editable
	}
	return
}

// Enabled returns the value of the 'enabled' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates if this parameter is enabled for the add-on.
func (o *AddOnParameter) Enabled() bool {
	if o != nil && o.enabled != nil {
		return *o.enabled
	}
	return false
}

// GetEnabled returns the value of the 'enabled' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates if this parameter is enabled for the add-on.
func (o *AddOnParameter) GetEnabled() (value bool, ok bool) {
	ok = o != nil && o.enabled != nil
	if ok {
		value = *o.enabled
	}
	return
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Name of the add-on parameter.
func (o *AddOnParameter) Name() string {
	if o != nil && o.name != nil {
		return *o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
// Name of the add-on parameter.
func (o *AddOnParameter) GetName() (value string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = *o.name
	}
	return
}

// Required returns the value of the 'required' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Indicates if this parameter is required by the add-on.
func (o *AddOnParameter) Required() bool {
	if o != nil && o.required != nil {
		return *o.required
	}
	return false
}

// GetRequired returns the value of the 'required' attribute and
// a flag indicating if the attribute has a value.
//
// Indicates if this parameter is required by the add-on.
func (o *AddOnParameter) GetRequired() (value bool, ok bool) {
	ok = o != nil && o.required != nil
	if ok {
		value = *o.required
	}
	return
}

// Validation returns the value of the 'validation' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Validation rule for the add-on parameter.
func (o *AddOnParameter) Validation() string {
	if o != nil && o.validation != nil {
		return *o.validation
	}
	return ""
}

// GetValidation returns the value of the 'validation' attribute and
// a flag indicating if the attribute has a value.
//
// Validation rule for the add-on parameter.
func (o *AddOnParameter) GetValidation() (value string, ok bool) {
	ok = o != nil && o.validation != nil
	if ok {
		value = *o.validation
	}
	return
}

// ValueType returns the value of the 'value_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Type of value of the add-on parameter.
func (o *AddOnParameter) ValueType() string {
	if o != nil && o.valueType != nil {
		return *o.valueType
	}
	return ""
}

// GetValueType returns the value of the 'value_type' attribute and
// a flag indicating if the attribute has a value.
//
// Type of value of the add-on parameter.
func (o *AddOnParameter) GetValueType() (value string, ok bool) {
	ok = o != nil && o.valueType != nil
	if ok {
		value = *o.valueType
	}
	return
}

// AddOnParameterListKind is the name of the type used to represent list of objects of
// type 'add_on_parameter'.
const AddOnParameterListKind = "AddOnParameterList"

// AddOnParameterListLinkKind is the name of the type used to represent links to list
// of objects of type 'add_on_parameter'.
const AddOnParameterListLinkKind = "AddOnParameterListLink"

// AddOnParameterNilKind is the name of the type used to nil lists of objects of
// type 'add_on_parameter'.
const AddOnParameterListNilKind = "AddOnParameterListNil"

// AddOnParameterList is a list of values of the 'add_on_parameter' type.
type AddOnParameterList struct {
	href  *string
	link  bool
	items []*AddOnParameter
}

// Kind returns the name of the type of the object.
func (l *AddOnParameterList) Kind() string {
	if l == nil {
		return AddOnParameterListNilKind
	}
	if l.link {
		return AddOnParameterListLinkKind
	}
	return AddOnParameterListKind
}

// Link returns true iif this is a link.
func (l *AddOnParameterList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *AddOnParameterList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *AddOnParameterList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *AddOnParameterList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *AddOnParameterList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *AddOnParameterList) Get(i int) *AddOnParameter {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *AddOnParameterList) Slice() []*AddOnParameter {
	var slice []*AddOnParameter
	if l == nil {
		slice = make([]*AddOnParameter, 0)
	} else {
		slice = make([]*AddOnParameter, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *AddOnParameterList) Each(f func(item *AddOnParameter) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *AddOnParameterList) Range(f func(index int, item *AddOnParameter) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}

/*
Todos API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Todo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Todo{}

// Todo struct for Todo
type Todo struct {
	Id *string `json:"id,omitempty"`
	Title string `json:"title"`
	Description *string `json:"description,omitempty"`
}

// NewTodo instantiates a new Todo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTodo(title string) *Todo {
	this := Todo{}
	this.Title = title
	return &this
}

// NewTodoWithDefaults instantiates a new Todo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTodoWithDefaults() *Todo {
	this := Todo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Todo) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Todo) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Todo) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Todo) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value
func (o *Todo) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Todo) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Todo) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Todo) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Todo) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Todo) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Todo) SetDescription(v string) {
	o.Description = &v
}

func (o Todo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Todo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["title"] = o.Title
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableTodo struct {
	value *Todo
	isSet bool
}

func (v NullableTodo) Get() *Todo {
	return v.value
}

func (v *NullableTodo) Set(val *Todo) {
	v.value = val
	v.isSet = true
}

func (v NullableTodo) IsSet() bool {
	return v.isSet
}

func (v *NullableTodo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTodo(val *Todo) *NullableTodo {
	return &NullableTodo{value: val, isSet: true}
}

func (v NullableTodo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTodo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



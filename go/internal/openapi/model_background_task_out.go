/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BackgroundTaskOut struct for BackgroundTaskOut
type BackgroundTaskOut struct {
	Data map[string]interface{} `json:"data"`
	Id string `json:"id"`
	Status BackgroundTaskStatus `json:"status"`
	Task BackgroundTaskType `json:"task"`
}

// NewBackgroundTaskOut instantiates a new BackgroundTaskOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundTaskOut(data map[string]interface{}, id string, status BackgroundTaskStatus, task BackgroundTaskType) *BackgroundTaskOut {
	this := BackgroundTaskOut{}
	this.Data = data
	this.Id = id
	this.Status = status
	this.Task = task
	return &this
}

// NewBackgroundTaskOutWithDefaults instantiates a new BackgroundTaskOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundTaskOutWithDefaults() *BackgroundTaskOut {
	this := BackgroundTaskOut{}
	return &this
}

// GetData returns the Data field value
func (o *BackgroundTaskOut) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BackgroundTaskOut) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BackgroundTaskOut) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetId returns the Id field value
func (o *BackgroundTaskOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BackgroundTaskOut) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BackgroundTaskOut) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *BackgroundTaskOut) GetStatus() BackgroundTaskStatus {
	if o == nil {
		var ret BackgroundTaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BackgroundTaskOut) GetStatusOk() (*BackgroundTaskStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BackgroundTaskOut) SetStatus(v BackgroundTaskStatus) {
	o.Status = v
}

// GetTask returns the Task field value
func (o *BackgroundTaskOut) GetTask() BackgroundTaskType {
	if o == nil {
		var ret BackgroundTaskType
		return ret
	}

	return o.Task
}

// GetTaskOk returns a tuple with the Task field value
// and a boolean to check if the value has been set.
func (o *BackgroundTaskOut) GetTaskOk() (*BackgroundTaskType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Task, true
}

// SetTask sets field value
func (o *BackgroundTaskOut) SetTask(v BackgroundTaskType) {
	o.Task = v
}

func (o BackgroundTaskOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["task"] = o.Task
	}
	return json.Marshal(toSerialize)
}

type NullableBackgroundTaskOut struct {
	value *BackgroundTaskOut
	isSet bool
}

func (v NullableBackgroundTaskOut) Get() *BackgroundTaskOut {
	return v.value
}

func (v *NullableBackgroundTaskOut) Set(val *BackgroundTaskOut) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundTaskOut) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundTaskOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundTaskOut(val *BackgroundTaskOut) *NullableBackgroundTaskOut {
	return &NullableBackgroundTaskOut{value: val, isSet: true}
}

func (v NullableBackgroundTaskOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundTaskOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



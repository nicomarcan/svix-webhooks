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

// MessageAttemptRecoveredEvent Sent on a successful dispatch after an earlier failure op webhook has already been sent.
type MessageAttemptRecoveredEvent struct {
	Data MessageAttemptRecoveredEventData `json:"data"`
	Type string `json:"type"`
}

// NewMessageAttemptRecoveredEvent instantiates a new MessageAttemptRecoveredEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttemptRecoveredEvent(data MessageAttemptRecoveredEventData, type_ string) *MessageAttemptRecoveredEvent {
	this := MessageAttemptRecoveredEvent{}
	this.Data = data
	this.Type = type_
	return &this
}

// NewMessageAttemptRecoveredEventWithDefaults instantiates a new MessageAttemptRecoveredEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttemptRecoveredEventWithDefaults() *MessageAttemptRecoveredEvent {
	this := MessageAttemptRecoveredEvent{}
	var type_ string = "message.attempt.recovered"
	this.Type = type_
	return &this
}

// GetData returns the Data field value
func (o *MessageAttemptRecoveredEvent) GetData() MessageAttemptRecoveredEventData {
	if o == nil {
		var ret MessageAttemptRecoveredEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptRecoveredEvent) GetDataOk() (*MessageAttemptRecoveredEventData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MessageAttemptRecoveredEvent) SetData(v MessageAttemptRecoveredEventData) {
	o.Data = v
}

// GetType returns the Type field value
func (o *MessageAttemptRecoveredEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptRecoveredEvent) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageAttemptRecoveredEvent) SetType(v string) {
	o.Type = v
}

func (o MessageAttemptRecoveredEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMessageAttemptRecoveredEvent struct {
	value *MessageAttemptRecoveredEvent
	isSet bool
}

func (v NullableMessageAttemptRecoveredEvent) Get() *MessageAttemptRecoveredEvent {
	return v.value
}

func (v *NullableMessageAttemptRecoveredEvent) Set(val *MessageAttemptRecoveredEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttemptRecoveredEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttemptRecoveredEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttemptRecoveredEvent(val *MessageAttemptRecoveredEvent) *NullableMessageAttemptRecoveredEvent {
	return &NullableMessageAttemptRecoveredEvent{value: val, isSet: true}
}

func (v NullableMessageAttemptRecoveredEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttemptRecoveredEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// DashboardAccessOut struct for DashboardAccessOut
type DashboardAccessOut struct {
	Token string `json:"token"`
	Url string `json:"url"`
}

// NewDashboardAccessOut instantiates a new DashboardAccessOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardAccessOut(token string, url string) *DashboardAccessOut {
	this := DashboardAccessOut{}
	this.Token = token
	this.Url = url
	return &this
}

// NewDashboardAccessOutWithDefaults instantiates a new DashboardAccessOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardAccessOutWithDefaults() *DashboardAccessOut {
	this := DashboardAccessOut{}
	return &this
}

// GetToken returns the Token field value
func (o *DashboardAccessOut) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DashboardAccessOut) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DashboardAccessOut) SetToken(v string) {
	o.Token = v
}

// GetUrl returns the Url field value
func (o *DashboardAccessOut) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DashboardAccessOut) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DashboardAccessOut) SetUrl(v string) {
	o.Url = v
}

func (o DashboardAccessOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardAccessOut struct {
	value *DashboardAccessOut
	isSet bool
}

func (v NullableDashboardAccessOut) Get() *DashboardAccessOut {
	return v.value
}

func (v *NullableDashboardAccessOut) Set(val *DashboardAccessOut) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardAccessOut) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardAccessOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardAccessOut(val *DashboardAccessOut) *NullableDashboardAccessOut {
	return &NullableDashboardAccessOut{value: val, isSet: true}
}

func (v NullableDashboardAccessOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardAccessOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



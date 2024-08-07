/*
SnapTrade

Connect brokerage accounts to your app for live positions and trading

API version: 1.0.0
Contact: api@snaptrade.com
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package snaptrade

import (
	"encoding/json"
)

// Model403FailedRequestResponse Example for failed request response
type Model403FailedRequestResponse struct {
	DefaultDetail interface{} `json:"default_detail,omitempty"`
	DefaultCode interface{} `json:"default_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Model403FailedRequestResponse Model403FailedRequestResponse

// NewModel403FailedRequestResponse instantiates a new Model403FailedRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel403FailedRequestResponse() *Model403FailedRequestResponse {
	this := Model403FailedRequestResponse{}
	return &this
}

// NewModel403FailedRequestResponseWithDefaults instantiates a new Model403FailedRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel403FailedRequestResponseWithDefaults() *Model403FailedRequestResponse {
	this := Model403FailedRequestResponse{}
	return &this
}

// GetDefaultDetail returns the DefaultDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Model403FailedRequestResponse) GetDefaultDetail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultDetail
}

// GetDefaultDetailOk returns a tuple with the DefaultDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model403FailedRequestResponse) GetDefaultDetailOk() (*interface{}, bool) {
	if o == nil || isNil(o.DefaultDetail) {
    return nil, false
	}
	return &o.DefaultDetail, true
}

// HasDefaultDetail returns a boolean if a field has been set.
func (o *Model403FailedRequestResponse) HasDefaultDetail() bool {
	if o != nil && isNil(o.DefaultDetail) {
		return true
	}

	return false
}

// SetDefaultDetail gets a reference to the given interface{} and assigns it to the DefaultDetail field.
func (o *Model403FailedRequestResponse) SetDefaultDetail(v interface{}) {
	o.DefaultDetail = v
}

// GetDefaultCode returns the DefaultCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Model403FailedRequestResponse) GetDefaultCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultCode
}

// GetDefaultCodeOk returns a tuple with the DefaultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model403FailedRequestResponse) GetDefaultCodeOk() (*interface{}, bool) {
	if o == nil || isNil(o.DefaultCode) {
    return nil, false
	}
	return &o.DefaultCode, true
}

// HasDefaultCode returns a boolean if a field has been set.
func (o *Model403FailedRequestResponse) HasDefaultCode() bool {
	if o != nil && isNil(o.DefaultCode) {
		return true
	}

	return false
}

// SetDefaultCode gets a reference to the given interface{} and assigns it to the DefaultCode field.
func (o *Model403FailedRequestResponse) SetDefaultCode(v interface{}) {
	o.DefaultCode = v
}

func (o Model403FailedRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultDetail != nil {
		toSerialize["default_detail"] = o.DefaultDetail
	}
	if o.DefaultCode != nil {
		toSerialize["default_code"] = o.DefaultCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Model403FailedRequestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varModel403FailedRequestResponse := _Model403FailedRequestResponse{}

	if err = json.Unmarshal(bytes, &varModel403FailedRequestResponse); err == nil {
		*o = Model403FailedRequestResponse(varModel403FailedRequestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default_detail")
		delete(additionalProperties, "default_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModel403FailedRequestResponse struct {
	value *Model403FailedRequestResponse
	isSet bool
}

func (v NullableModel403FailedRequestResponse) Get() *Model403FailedRequestResponse {
	return v.value
}

func (v *NullableModel403FailedRequestResponse) Set(val *Model403FailedRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModel403FailedRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModel403FailedRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel403FailedRequestResponse(val *Model403FailedRequestResponse) *NullableModel403FailedRequestResponse {
	return &NullableModel403FailedRequestResponse{value: val, isSet: true}
}

func (v NullableModel403FailedRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel403FailedRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



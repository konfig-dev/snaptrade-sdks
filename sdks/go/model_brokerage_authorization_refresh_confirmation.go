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

// BrokerageAuthorizationRefreshConfirmation struct for BrokerageAuthorizationRefreshConfirmation
type BrokerageAuthorizationRefreshConfirmation struct {
	// Refresh confirmation details
	Detail *string `json:"detail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BrokerageAuthorizationRefreshConfirmation BrokerageAuthorizationRefreshConfirmation

// NewBrokerageAuthorizationRefreshConfirmation instantiates a new BrokerageAuthorizationRefreshConfirmation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerageAuthorizationRefreshConfirmation() *BrokerageAuthorizationRefreshConfirmation {
	this := BrokerageAuthorizationRefreshConfirmation{}
	return &this
}

// NewBrokerageAuthorizationRefreshConfirmationWithDefaults instantiates a new BrokerageAuthorizationRefreshConfirmation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerageAuthorizationRefreshConfirmationWithDefaults() *BrokerageAuthorizationRefreshConfirmation {
	this := BrokerageAuthorizationRefreshConfirmation{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *BrokerageAuthorizationRefreshConfirmation) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageAuthorizationRefreshConfirmation) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
    return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *BrokerageAuthorizationRefreshConfirmation) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *BrokerageAuthorizationRefreshConfirmation) SetDetail(v string) {
	o.Detail = &v
}

func (o BrokerageAuthorizationRefreshConfirmation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BrokerageAuthorizationRefreshConfirmation) UnmarshalJSON(bytes []byte) (err error) {
	varBrokerageAuthorizationRefreshConfirmation := _BrokerageAuthorizationRefreshConfirmation{}

	if err = json.Unmarshal(bytes, &varBrokerageAuthorizationRefreshConfirmation); err == nil {
		*o = BrokerageAuthorizationRefreshConfirmation(varBrokerageAuthorizationRefreshConfirmation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrokerageAuthorizationRefreshConfirmation struct {
	value *BrokerageAuthorizationRefreshConfirmation
	isSet bool
}

func (v NullableBrokerageAuthorizationRefreshConfirmation) Get() *BrokerageAuthorizationRefreshConfirmation {
	return v.value
}

func (v *NullableBrokerageAuthorizationRefreshConfirmation) Set(val *BrokerageAuthorizationRefreshConfirmation) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerageAuthorizationRefreshConfirmation) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerageAuthorizationRefreshConfirmation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerageAuthorizationRefreshConfirmation(val *BrokerageAuthorizationRefreshConfirmation) *NullableBrokerageAuthorizationRefreshConfirmation {
	return &NullableBrokerageAuthorizationRefreshConfirmation{value: val, isSet: true}
}

func (v NullableBrokerageAuthorizationRefreshConfirmation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerageAuthorizationRefreshConfirmation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// BrokerageType Type of brokerage
type BrokerageType struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BrokerageType BrokerageType

// NewBrokerageType instantiates a new BrokerageType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerageType() *BrokerageType {
	this := BrokerageType{}
	return &this
}

// NewBrokerageTypeWithDefaults instantiates a new BrokerageType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerageTypeWithDefaults() *BrokerageType {
	this := BrokerageType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrokerageType) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageType) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrokerageType) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BrokerageType) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BrokerageType) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageType) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BrokerageType) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BrokerageType) SetName(v string) {
	o.Name = &v
}

func (o BrokerageType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BrokerageType) UnmarshalJSON(bytes []byte) (err error) {
	varBrokerageType := _BrokerageType{}

	if err = json.Unmarshal(bytes, &varBrokerageType); err == nil {
		*o = BrokerageType(varBrokerageType)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrokerageType struct {
	value *BrokerageType
	isSet bool
}

func (v NullableBrokerageType) Get() *BrokerageType {
	return v.value
}

func (v *NullableBrokerageType) Set(val *BrokerageType) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerageType) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerageType(val *BrokerageType) *NullableBrokerageType {
	return &NullableBrokerageType{value: val, isSet: true}
}

func (v NullableBrokerageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



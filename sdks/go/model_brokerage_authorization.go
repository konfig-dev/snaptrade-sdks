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

// BrokerageAuthorization struct for BrokerageAuthorization
type BrokerageAuthorization struct {
	Id *string `json:"id,omitempty"`
	// Time
	CreatedDate *string `json:"created_date,omitempty"`
	// Time
	UpdatedDate *string `json:"updated_date,omitempty"`
	Brokerage *Brokerage `json:"brokerage,omitempty"`
	// Connection Name
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	// Disabled date
	DisabledDate NullableString `json:"disabled_date,omitempty"`
	// Additional data about brokerage authorization
	Meta map[string]interface{} `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BrokerageAuthorization BrokerageAuthorization

// NewBrokerageAuthorization instantiates a new BrokerageAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerageAuthorization() *BrokerageAuthorization {
	this := BrokerageAuthorization{}
	return &this
}

// NewBrokerageAuthorizationWithDefaults instantiates a new BrokerageAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerageAuthorizationWithDefaults() *BrokerageAuthorization {
	this := BrokerageAuthorization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrokerageAuthorization) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageAuthorization) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrokerageAuthorization) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BrokerageAuthorization) SetId(v string) {
	o.Id = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *BrokerageAuthorization) GetCreatedDate() string {
	if o == nil || isNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageAuthorization) GetCreatedDateOk() (*string, bool) {
	if o == nil || isNil(o.CreatedDate) {
    return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *BrokerageAuthorization) HasCreatedDate() bool {
	if o != nil && !isNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *BrokerageAuthorization) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *BrokerageAuthorization) GetUpdatedDate() string {
	if o == nil || isNil(o.UpdatedDate) {
		var ret string
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageAuthorization) GetUpdatedDateOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedDate) {
    return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *BrokerageAuthorization) HasUpdatedDate() bool {
	if o != nil && !isNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given string and assigns it to the UpdatedDate field.
func (o *BrokerageAuthorization) SetUpdatedDate(v string) {
	o.UpdatedDate = &v
}

// GetBrokerage returns the Brokerage field value if set, zero value otherwise.
func (o *BrokerageAuthorization) GetBrokerage() Brokerage {
	if o == nil || isNil(o.Brokerage) {
		var ret Brokerage
		return ret
	}
	return *o.Brokerage
}

// GetBrokerageOk returns a tuple with the Brokerage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageAuthorization) GetBrokerageOk() (*Brokerage, bool) {
	if o == nil || isNil(o.Brokerage) {
    return nil, false
	}
	return o.Brokerage, true
}

// HasBrokerage returns a boolean if a field has been set.
func (o *BrokerageAuthorization) HasBrokerage() bool {
	if o != nil && !isNil(o.Brokerage) {
		return true
	}

	return false
}

// SetBrokerage gets a reference to the given Brokerage and assigns it to the Brokerage field.
func (o *BrokerageAuthorization) SetBrokerage(v Brokerage) {
	o.Brokerage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BrokerageAuthorization) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageAuthorization) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BrokerageAuthorization) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BrokerageAuthorization) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BrokerageAuthorization) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageAuthorization) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BrokerageAuthorization) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BrokerageAuthorization) SetType(v string) {
	o.Type = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *BrokerageAuthorization) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageAuthorization) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
    return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *BrokerageAuthorization) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *BrokerageAuthorization) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDisabledDate returns the DisabledDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrokerageAuthorization) GetDisabledDate() string {
	if o == nil || isNil(o.DisabledDate.Get()) {
		var ret string
		return ret
	}
	return *o.DisabledDate.Get()
}

// GetDisabledDateOk returns a tuple with the DisabledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrokerageAuthorization) GetDisabledDateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DisabledDate.Get(), o.DisabledDate.IsSet()
}

// HasDisabledDate returns a boolean if a field has been set.
func (o *BrokerageAuthorization) HasDisabledDate() bool {
	if o != nil && o.DisabledDate.IsSet() {
		return true
	}

	return false
}

// SetDisabledDate gets a reference to the given NullableString and assigns it to the DisabledDate field.
func (o *BrokerageAuthorization) SetDisabledDate(v string) {
	o.DisabledDate.Set(&v)
}
// SetDisabledDateNil sets the value for DisabledDate to be an explicit nil
func (o *BrokerageAuthorization) SetDisabledDateNil() {
	o.DisabledDate.Set(nil)
}

// UnsetDisabledDate ensures that no value is present for DisabledDate, not even an explicit nil
func (o *BrokerageAuthorization) UnsetDisabledDate() {
	o.DisabledDate.Unset()
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BrokerageAuthorization) GetMeta() map[string]interface{} {
	if o == nil || isNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerageAuthorization) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Meta) {
    return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BrokerageAuthorization) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *BrokerageAuthorization) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o BrokerageAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreatedDate) {
		toSerialize["created_date"] = o.CreatedDate
	}
	if !isNil(o.UpdatedDate) {
		toSerialize["updated_date"] = o.UpdatedDate
	}
	if !isNil(o.Brokerage) {
		toSerialize["brokerage"] = o.Brokerage
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if o.DisabledDate.IsSet() {
		toSerialize["disabled_date"] = o.DisabledDate.Get()
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BrokerageAuthorization) UnmarshalJSON(bytes []byte) (err error) {
	varBrokerageAuthorization := _BrokerageAuthorization{}

	if err = json.Unmarshal(bytes, &varBrokerageAuthorization); err == nil {
		*o = BrokerageAuthorization(varBrokerageAuthorization)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_date")
		delete(additionalProperties, "updated_date")
		delete(additionalProperties, "brokerage")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "disabled_date")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrokerageAuthorization struct {
	value *BrokerageAuthorization
	isSet bool
}

func (v NullableBrokerageAuthorization) Get() *BrokerageAuthorization {
	return v.value
}

func (v *NullableBrokerageAuthorization) Set(val *BrokerageAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerageAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerageAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerageAuthorization(val *BrokerageAuthorization) *NullableBrokerageAuthorization {
	return &NullableBrokerageAuthorization{value: val, isSet: true}
}

func (v NullableBrokerageAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerageAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



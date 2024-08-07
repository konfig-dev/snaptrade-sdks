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
	"fmt"
)

// ActionStrict Trade Action
type ActionStrict string

// List of ActionStrict
const (
	ACTIONSTRICT_BUY ActionStrict = "BUY"
	ACTIONSTRICT_SELL ActionStrict = "SELL"
)

// All allowed values of ActionStrict enum
var AllowedActionStrictEnumValues = []ActionStrict{
	"BUY",
	"SELL",
}

func (v *ActionStrict) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActionStrict(value)
	for _, existing := range AllowedActionStrictEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActionStrict", value)
}

// NewActionStrictFromValue returns a pointer to a valid ActionStrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionStrictFromValue(v string) (*ActionStrict, error) {
	ev := ActionStrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActionStrict: valid values are %v", v, AllowedActionStrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionStrict) IsValid() bool {
	for _, existing := range AllowedActionStrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionStrict value
func (v ActionStrict) Ptr() *ActionStrict {
	return &v
}

type NullableActionStrict struct {
	value *ActionStrict
	isSet bool
}

func (v NullableActionStrict) Get() *ActionStrict {
	return v.value
}

func (v *NullableActionStrict) Set(val *ActionStrict) {
	v.value = val
	v.isSet = true
}

func (v NullableActionStrict) IsSet() bool {
	return v.isSet
}

func (v *NullableActionStrict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionStrict(val *ActionStrict) *NullableActionStrict {
	return &NullableActionStrict{value: val, isSet: true}
}

func (v NullableActionStrict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionStrict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


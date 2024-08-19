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

// DeleteUserResponse struct for DeleteUserResponse
type DeleteUserResponse struct {
	// This is always `deleted` when a user is queued for deletion.
	Status *string `json:"status,omitempty"`
	// Human friendly message about the deletion status.
	Detail *string `json:"detail,omitempty"`
	// SnapTrade User ID. This is chosen by the API partner and can be any string that is a) unique to the user, and b) immutable for the user. It is recommended to NOT use email addresses for this property because they are usually not immutable.
	UserId *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteUserResponse DeleteUserResponse

// NewDeleteUserResponse instantiates a new DeleteUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUserResponse() *DeleteUserResponse {
	this := DeleteUserResponse{}
	return &this
}

// NewDeleteUserResponseWithDefaults instantiates a new DeleteUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUserResponseWithDefaults() *DeleteUserResponse {
	this := DeleteUserResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeleteUserResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeleteUserResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeleteUserResponse) SetStatus(v string) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *DeleteUserResponse) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserResponse) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
    return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *DeleteUserResponse) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *DeleteUserResponse) SetDetail(v string) {
	o.Detail = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DeleteUserResponse) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserResponse) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
    return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DeleteUserResponse) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *DeleteUserResponse) SetUserId(v string) {
	o.UserId = &v
}

func (o DeleteUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !isNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeleteUserResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteUserResponse := _DeleteUserResponse{}

	if err = json.Unmarshal(bytes, &varDeleteUserResponse); err == nil {
		*o = DeleteUserResponse(varDeleteUserResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteUserResponse struct {
	value *DeleteUserResponse
	isSet bool
}

func (v NullableDeleteUserResponse) Get() *DeleteUserResponse {
	return v.value
}

func (v *NullableDeleteUserResponse) Set(val *DeleteUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserResponse(val *DeleteUserResponse) *NullableDeleteUserResponse {
	return &NullableDeleteUserResponse{value: val, isSet: true}
}

func (v NullableDeleteUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



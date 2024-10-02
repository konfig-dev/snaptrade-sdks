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
	"time"
)

// HoldingsStatus Status of account holdings sync. SnapTrade syncs holdings from the brokerage under the following conditions: 1. Initial connection - SnapTrade syncs all holdings (positions, balances, recent orders, and transactions) immediately after the connection is established. 2. Daily sync - Once a day SnapTrade refreshes all holdings from the brokerage. 3. Manual sync - You can trigger a refresh of holdings with the [manual refresh](/reference/Connections/Connections_refreshBrokerageAuthorization) endpoint. 
type HoldingsStatus struct {
	// Indicates if the initial sync of holdings has been completed. For accounts with a large number of positions/orders/transactions, the initial sync may take a while to complete.
	InitialSyncCompleted *bool `json:"initial_sync_completed,omitempty"`
	// The last time holdings were successfully synced by SnapTrade.
	LastSuccessfulSync NullableTime `json:"last_successful_sync,omitempty"`
}

// NewHoldingsStatus instantiates a new HoldingsStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldingsStatus() *HoldingsStatus {
	this := HoldingsStatus{}
	return &this
}

// NewHoldingsStatusWithDefaults instantiates a new HoldingsStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldingsStatusWithDefaults() *HoldingsStatus {
	this := HoldingsStatus{}
	return &this
}

// GetInitialSyncCompleted returns the InitialSyncCompleted field value if set, zero value otherwise.
func (o *HoldingsStatus) GetInitialSyncCompleted() bool {
	if o == nil || isNil(o.InitialSyncCompleted) {
		var ret bool
		return ret
	}
	return *o.InitialSyncCompleted
}

// GetInitialSyncCompletedOk returns a tuple with the InitialSyncCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HoldingsStatus) GetInitialSyncCompletedOk() (*bool, bool) {
	if o == nil || isNil(o.InitialSyncCompleted) {
    return nil, false
	}
	return o.InitialSyncCompleted, true
}

// HasInitialSyncCompleted returns a boolean if a field has been set.
func (o *HoldingsStatus) HasInitialSyncCompleted() bool {
	if o != nil && !isNil(o.InitialSyncCompleted) {
		return true
	}

	return false
}

// SetInitialSyncCompleted gets a reference to the given bool and assigns it to the InitialSyncCompleted field.
func (o *HoldingsStatus) SetInitialSyncCompleted(v bool) {
	o.InitialSyncCompleted = &v
}

// GetLastSuccessfulSync returns the LastSuccessfulSync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HoldingsStatus) GetLastSuccessfulSync() time.Time {
	if o == nil || isNil(o.LastSuccessfulSync.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSuccessfulSync.Get()
}

// GetLastSuccessfulSyncOk returns a tuple with the LastSuccessfulSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldingsStatus) GetLastSuccessfulSyncOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.LastSuccessfulSync.Get(), o.LastSuccessfulSync.IsSet()
}

// HasLastSuccessfulSync returns a boolean if a field has been set.
func (o *HoldingsStatus) HasLastSuccessfulSync() bool {
	if o != nil && o.LastSuccessfulSync.IsSet() {
		return true
	}

	return false
}

// SetLastSuccessfulSync gets a reference to the given NullableTime.Time and assigns it to the LastSuccessfulSync field.
func (o *HoldingsStatus) SetLastSuccessfulSync(v time.Time) {
	o.LastSuccessfulSync.Set(&v)
}
// SetLastSuccessfulSyncNil sets the value for LastSuccessfulSync to be an explicit nil
func (o *HoldingsStatus) SetLastSuccessfulSyncNil() {
	o.LastSuccessfulSync.Set(nil)
}

// UnsetLastSuccessfulSync ensures that no value is present for LastSuccessfulSync, not even an explicit nil
func (o *HoldingsStatus) UnsetLastSuccessfulSync() {
	o.LastSuccessfulSync.Unset()
}

func (o HoldingsStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InitialSyncCompleted) {
		toSerialize["initial_sync_completed"] = o.InitialSyncCompleted
	}
	if o.LastSuccessfulSync.IsSet() {
		toSerialize["last_successful_sync"] = o.LastSuccessfulSync.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHoldingsStatus struct {
	value *HoldingsStatus
	isSet bool
}

func (v NullableHoldingsStatus) Get() *HoldingsStatus {
	return v.value
}

func (v *NullableHoldingsStatus) Set(val *HoldingsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldingsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldingsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldingsStatus(val *HoldingsStatus) *NullableHoldingsStatus {
	return &NullableHoldingsStatus{value: val, isSet: true}
}

func (v NullableHoldingsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldingsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



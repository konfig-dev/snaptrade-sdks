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

// SnapTradeHoldingsAccountAccountId A single brokerage account at a financial institution.
type SnapTradeHoldingsAccountAccountId struct {
	// Unique identifier for the connected brokerage account. This is the UUID used to reference the account in SnapTrade.
	Id *string `json:"id,omitempty"`
	// Unique identifier for the connection (brokerage authorization). This is the UUID used to reference the connection in SnapTrade.
	BrokerageAuthorization *string `json:"brokerage_authorization,omitempty"`
	// Portfolio Group ID. Portfolio Groups have been deprecated. Please contact support if you have a usecase for it.
	// Deprecated
	PortfolioGroup *string `json:"portfolio_group,omitempty"`
	// A display name for the account. Either assigned by the user or by the financial institution itself. For certain institutions, SnapTrade appends the institution name to the account name for clarity.
	Name NullableString `json:"name,omitempty"`
	// The account number assigned by the financial institution.
	Number *string `json:"number,omitempty"`
	// The name of the financial institution that holds the account.
	InstitutionName *string `json:"institution_name,omitempty"`
	Balance NullableSnapTradeHoldingsAccountAccountIdBalance `json:"balance,omitempty"`
	// Additional information about the account, such as account type, status, etc. This information is specific to the financial institution and there's no standard format for this data. Please use at your own risk.
	// Deprecated
	Meta map[string]interface{} `json:"meta,omitempty"`
	// This field is deprecated.
	// Deprecated
	CashRestrictions []CashRestriction `json:"cash_restrictions,omitempty"`
	// Timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format indicating when the account was created in SnapTrade. This is _not_ the account opening date at the financial institution.
	CreatedDate *time.Time `json:"created_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnapTradeHoldingsAccountAccountId SnapTradeHoldingsAccountAccountId

// NewSnapTradeHoldingsAccountAccountId instantiates a new SnapTradeHoldingsAccountAccountId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapTradeHoldingsAccountAccountId() *SnapTradeHoldingsAccountAccountId {
	this := SnapTradeHoldingsAccountAccountId{}
	return &this
}

// NewSnapTradeHoldingsAccountAccountIdWithDefaults instantiates a new SnapTradeHoldingsAccountAccountId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapTradeHoldingsAccountAccountIdWithDefaults() *SnapTradeHoldingsAccountAccountId {
	this := SnapTradeHoldingsAccountAccountId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SnapTradeHoldingsAccountAccountId) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapTradeHoldingsAccountAccountId) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SnapTradeHoldingsAccountAccountId) SetId(v string) {
	o.Id = &v
}

// GetBrokerageAuthorization returns the BrokerageAuthorization field value if set, zero value otherwise.
func (o *SnapTradeHoldingsAccountAccountId) GetBrokerageAuthorization() string {
	if o == nil || isNil(o.BrokerageAuthorization) {
		var ret string
		return ret
	}
	return *o.BrokerageAuthorization
}

// GetBrokerageAuthorizationOk returns a tuple with the BrokerageAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapTradeHoldingsAccountAccountId) GetBrokerageAuthorizationOk() (*string, bool) {
	if o == nil || isNil(o.BrokerageAuthorization) {
    return nil, false
	}
	return o.BrokerageAuthorization, true
}

// HasBrokerageAuthorization returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasBrokerageAuthorization() bool {
	if o != nil && !isNil(o.BrokerageAuthorization) {
		return true
	}

	return false
}

// SetBrokerageAuthorization gets a reference to the given string and assigns it to the BrokerageAuthorization field.
func (o *SnapTradeHoldingsAccountAccountId) SetBrokerageAuthorization(v string) {
	o.BrokerageAuthorization = &v
}

// GetPortfolioGroup returns the PortfolioGroup field value if set, zero value otherwise.
// Deprecated
func (o *SnapTradeHoldingsAccountAccountId) GetPortfolioGroup() string {
	if o == nil || isNil(o.PortfolioGroup) {
		var ret string
		return ret
	}
	return *o.PortfolioGroup
}

// GetPortfolioGroupOk returns a tuple with the PortfolioGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SnapTradeHoldingsAccountAccountId) GetPortfolioGroupOk() (*string, bool) {
	if o == nil || isNil(o.PortfolioGroup) {
    return nil, false
	}
	return o.PortfolioGroup, true
}

// HasPortfolioGroup returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasPortfolioGroup() bool {
	if o != nil && !isNil(o.PortfolioGroup) {
		return true
	}

	return false
}

// SetPortfolioGroup gets a reference to the given string and assigns it to the PortfolioGroup field.
// Deprecated
func (o *SnapTradeHoldingsAccountAccountId) SetPortfolioGroup(v string) {
	o.PortfolioGroup = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapTradeHoldingsAccountAccountId) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapTradeHoldingsAccountAccountId) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SnapTradeHoldingsAccountAccountId) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SnapTradeHoldingsAccountAccountId) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SnapTradeHoldingsAccountAccountId) UnsetName() {
	o.Name.Unset()
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SnapTradeHoldingsAccountAccountId) GetNumber() string {
	if o == nil || isNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapTradeHoldingsAccountAccountId) GetNumberOk() (*string, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *SnapTradeHoldingsAccountAccountId) SetNumber(v string) {
	o.Number = &v
}

// GetInstitutionName returns the InstitutionName field value if set, zero value otherwise.
func (o *SnapTradeHoldingsAccountAccountId) GetInstitutionName() string {
	if o == nil || isNil(o.InstitutionName) {
		var ret string
		return ret
	}
	return *o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapTradeHoldingsAccountAccountId) GetInstitutionNameOk() (*string, bool) {
	if o == nil || isNil(o.InstitutionName) {
    return nil, false
	}
	return o.InstitutionName, true
}

// HasInstitutionName returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasInstitutionName() bool {
	if o != nil && !isNil(o.InstitutionName) {
		return true
	}

	return false
}

// SetInstitutionName gets a reference to the given string and assigns it to the InstitutionName field.
func (o *SnapTradeHoldingsAccountAccountId) SetInstitutionName(v string) {
	o.InstitutionName = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapTradeHoldingsAccountAccountId) GetBalance() SnapTradeHoldingsAccountAccountIdBalance {
	if o == nil || isNil(o.Balance.Get()) {
		var ret SnapTradeHoldingsAccountAccountIdBalance
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapTradeHoldingsAccountAccountId) GetBalanceOk() (*SnapTradeHoldingsAccountAccountIdBalance, bool) {
	if o == nil {
    return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableSnapTradeHoldingsAccountAccountIdBalance and assigns it to the Balance field.
func (o *SnapTradeHoldingsAccountAccountId) SetBalance(v SnapTradeHoldingsAccountAccountIdBalance) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *SnapTradeHoldingsAccountAccountId) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *SnapTradeHoldingsAccountAccountId) UnsetBalance() {
	o.Balance.Unset()
}

// GetMeta returns the Meta field value if set, zero value otherwise.
// Deprecated
func (o *SnapTradeHoldingsAccountAccountId) GetMeta() map[string]interface{} {
	if o == nil || isNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SnapTradeHoldingsAccountAccountId) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Meta) {
    return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
// Deprecated
func (o *SnapTradeHoldingsAccountAccountId) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetCashRestrictions returns the CashRestrictions field value if set, zero value otherwise.
// Deprecated
func (o *SnapTradeHoldingsAccountAccountId) GetCashRestrictions() []CashRestriction {
	if o == nil || isNil(o.CashRestrictions) {
		var ret []CashRestriction
		return ret
	}
	return o.CashRestrictions
}

// GetCashRestrictionsOk returns a tuple with the CashRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SnapTradeHoldingsAccountAccountId) GetCashRestrictionsOk() ([]CashRestriction, bool) {
	if o == nil || isNil(o.CashRestrictions) {
    return nil, false
	}
	return o.CashRestrictions, true
}

// HasCashRestrictions returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasCashRestrictions() bool {
	if o != nil && !isNil(o.CashRestrictions) {
		return true
	}

	return false
}

// SetCashRestrictions gets a reference to the given []CashRestriction and assigns it to the CashRestrictions field.
// Deprecated
func (o *SnapTradeHoldingsAccountAccountId) SetCashRestrictions(v []CashRestriction) {
	o.CashRestrictions = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *SnapTradeHoldingsAccountAccountId) GetCreatedDate() time.Time {
	if o == nil || isNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapTradeHoldingsAccountAccountId) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedDate) {
    return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *SnapTradeHoldingsAccountAccountId) HasCreatedDate() bool {
	if o != nil && !isNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *SnapTradeHoldingsAccountAccountId) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

func (o SnapTradeHoldingsAccountAccountId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.BrokerageAuthorization) {
		toSerialize["brokerage_authorization"] = o.BrokerageAuthorization
	}
	if !isNil(o.PortfolioGroup) {
		toSerialize["portfolio_group"] = o.PortfolioGroup
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !isNil(o.InstitutionName) {
		toSerialize["institution_name"] = o.InstitutionName
	}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.CashRestrictions) {
		toSerialize["cash_restrictions"] = o.CashRestrictions
	}
	if !isNil(o.CreatedDate) {
		toSerialize["created_date"] = o.CreatedDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SnapTradeHoldingsAccountAccountId) UnmarshalJSON(bytes []byte) (err error) {
	varSnapTradeHoldingsAccountAccountId := _SnapTradeHoldingsAccountAccountId{}

	if err = json.Unmarshal(bytes, &varSnapTradeHoldingsAccountAccountId); err == nil {
		*o = SnapTradeHoldingsAccountAccountId(varSnapTradeHoldingsAccountAccountId)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "brokerage_authorization")
		delete(additionalProperties, "portfolio_group")
		delete(additionalProperties, "name")
		delete(additionalProperties, "number")
		delete(additionalProperties, "institution_name")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "cash_restrictions")
		delete(additionalProperties, "created_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnapTradeHoldingsAccountAccountId struct {
	value *SnapTradeHoldingsAccountAccountId
	isSet bool
}

func (v NullableSnapTradeHoldingsAccountAccountId) Get() *SnapTradeHoldingsAccountAccountId {
	return v.value
}

func (v *NullableSnapTradeHoldingsAccountAccountId) Set(val *SnapTradeHoldingsAccountAccountId) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapTradeHoldingsAccountAccountId) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapTradeHoldingsAccountAccountId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapTradeHoldingsAccountAccountId(val *SnapTradeHoldingsAccountAccountId) *NullableSnapTradeHoldingsAccountAccountId {
	return &NullableSnapTradeHoldingsAccountAccountId{value: val, isSet: true}
}

func (v NullableSnapTradeHoldingsAccountAccountId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapTradeHoldingsAccountAccountId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



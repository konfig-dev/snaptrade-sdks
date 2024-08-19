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

// UniversalActivityOptionSymbol The option security for the transaction. The field is `null` if the transaction is not related to an option security (like a deposit, withdrawal, fee, etc). SnapTrade does a best effort to map the brokerage's option symbol. In cases where the brokerage option symbol is not recognized, the field will be set to `null`.
type UniversalActivityOptionSymbol struct {
	// Unique identifier for the option symbol within SnapTrade. This is the ID used to reference the symbol in SnapTrade API calls.
	Id string `json:"id"`
	// The [OCC symbol](https://en.wikipedia.org/wiki/Option_symbol) for the option.
	Ticker string `json:"ticker"`
	// The type of option. Either \"CALL\" or \"PUT\".
	OptionType string `json:"option_type"`
	// The option strike price.
	StrikePrice float32 `json:"strike_price"`
	// The option expiration date.
	ExpirationDate string `json:"expiration_date"`
	// Whether the option is a mini option. Mini options have 10 underlying shares per contract instead of the standard 100.
	IsMiniOption *bool `json:"is_mini_option,omitempty"`
	UnderlyingSymbol UnderlyingSymbol `json:"underlying_symbol"`
}

// NewUniversalActivityOptionSymbol instantiates a new UniversalActivityOptionSymbol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniversalActivityOptionSymbol(id string, ticker string, optionType string, strikePrice float32, expirationDate string, underlyingSymbol UnderlyingSymbol) *UniversalActivityOptionSymbol {
	this := UniversalActivityOptionSymbol{}
	this.Id = id
	this.Ticker = ticker
	this.OptionType = optionType
	this.StrikePrice = strikePrice
	this.ExpirationDate = expirationDate
	this.UnderlyingSymbol = underlyingSymbol
	return &this
}

// NewUniversalActivityOptionSymbolWithDefaults instantiates a new UniversalActivityOptionSymbol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniversalActivityOptionSymbolWithDefaults() *UniversalActivityOptionSymbol {
	this := UniversalActivityOptionSymbol{}
	return &this
}

// GetId returns the Id field value
func (o *UniversalActivityOptionSymbol) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UniversalActivityOptionSymbol) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UniversalActivityOptionSymbol) SetId(v string) {
	o.Id = v
}

// GetTicker returns the Ticker field value
func (o *UniversalActivityOptionSymbol) GetTicker() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticker
}

// GetTickerOk returns a tuple with the Ticker field value
// and a boolean to check if the value has been set.
func (o *UniversalActivityOptionSymbol) GetTickerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ticker, true
}

// SetTicker sets field value
func (o *UniversalActivityOptionSymbol) SetTicker(v string) {
	o.Ticker = v
}

// GetOptionType returns the OptionType field value
func (o *UniversalActivityOptionSymbol) GetOptionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptionType
}

// GetOptionTypeOk returns a tuple with the OptionType field value
// and a boolean to check if the value has been set.
func (o *UniversalActivityOptionSymbol) GetOptionTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OptionType, true
}

// SetOptionType sets field value
func (o *UniversalActivityOptionSymbol) SetOptionType(v string) {
	o.OptionType = v
}

// GetStrikePrice returns the StrikePrice field value
func (o *UniversalActivityOptionSymbol) GetStrikePrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value
// and a boolean to check if the value has been set.
func (o *UniversalActivityOptionSymbol) GetStrikePriceOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StrikePrice, true
}

// SetStrikePrice sets field value
func (o *UniversalActivityOptionSymbol) SetStrikePrice(v float32) {
	o.StrikePrice = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *UniversalActivityOptionSymbol) GetExpirationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *UniversalActivityOptionSymbol) GetExpirationDateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *UniversalActivityOptionSymbol) SetExpirationDate(v string) {
	o.ExpirationDate = v
}

// GetIsMiniOption returns the IsMiniOption field value if set, zero value otherwise.
func (o *UniversalActivityOptionSymbol) GetIsMiniOption() bool {
	if o == nil || isNil(o.IsMiniOption) {
		var ret bool
		return ret
	}
	return *o.IsMiniOption
}

// GetIsMiniOptionOk returns a tuple with the IsMiniOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalActivityOptionSymbol) GetIsMiniOptionOk() (*bool, bool) {
	if o == nil || isNil(o.IsMiniOption) {
    return nil, false
	}
	return o.IsMiniOption, true
}

// HasIsMiniOption returns a boolean if a field has been set.
func (o *UniversalActivityOptionSymbol) HasIsMiniOption() bool {
	if o != nil && !isNil(o.IsMiniOption) {
		return true
	}

	return false
}

// SetIsMiniOption gets a reference to the given bool and assigns it to the IsMiniOption field.
func (o *UniversalActivityOptionSymbol) SetIsMiniOption(v bool) {
	o.IsMiniOption = &v
}

// GetUnderlyingSymbol returns the UnderlyingSymbol field value
func (o *UniversalActivityOptionSymbol) GetUnderlyingSymbol() UnderlyingSymbol {
	if o == nil {
		var ret UnderlyingSymbol
		return ret
	}

	return o.UnderlyingSymbol
}

// GetUnderlyingSymbolOk returns a tuple with the UnderlyingSymbol field value
// and a boolean to check if the value has been set.
func (o *UniversalActivityOptionSymbol) GetUnderlyingSymbolOk() (*UnderlyingSymbol, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UnderlyingSymbol, true
}

// SetUnderlyingSymbol sets field value
func (o *UniversalActivityOptionSymbol) SetUnderlyingSymbol(v UnderlyingSymbol) {
	o.UnderlyingSymbol = v
}

func (o UniversalActivityOptionSymbol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ticker"] = o.Ticker
	}
	if true {
		toSerialize["option_type"] = o.OptionType
	}
	if true {
		toSerialize["strike_price"] = o.StrikePrice
	}
	if true {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if !isNil(o.IsMiniOption) {
		toSerialize["is_mini_option"] = o.IsMiniOption
	}
	if true {
		toSerialize["underlying_symbol"] = o.UnderlyingSymbol
	}
	return json.Marshal(toSerialize)
}

type NullableUniversalActivityOptionSymbol struct {
	value *UniversalActivityOptionSymbol
	isSet bool
}

func (v NullableUniversalActivityOptionSymbol) Get() *UniversalActivityOptionSymbol {
	return v.value
}

func (v *NullableUniversalActivityOptionSymbol) Set(val *UniversalActivityOptionSymbol) {
	v.value = val
	v.isSet = true
}

func (v NullableUniversalActivityOptionSymbol) IsSet() bool {
	return v.isSet
}

func (v *NullableUniversalActivityOptionSymbol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniversalActivityOptionSymbol(val *UniversalActivityOptionSymbol) *NullableUniversalActivityOptionSymbol {
	return &NullableUniversalActivityOptionSymbol{value: val, isSet: true}
}

func (v NullableUniversalActivityOptionSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniversalActivityOptionSymbol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



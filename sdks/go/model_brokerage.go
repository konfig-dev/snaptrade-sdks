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

// Brokerage struct for Brokerage
type Brokerage struct {
	Id *string `json:"id,omitempty"`
	// Full name of the brokerage.
	Name *string `json:"name,omitempty"`
	// A display-friendly name of the brokerage.
	DisplayName *string `json:"display_name,omitempty"`
	Description *string `json:"description,omitempty"`
	AwsS3LogoUrl *string `json:"aws_s3_logo_url,omitempty"`
	AwsS3SquareLogoUrl NullableString `json:"aws_s3_square_logo_url,omitempty"`
	OpenUrl NullableString `json:"open_url,omitempty"`
	// A unique identifier for that brokerage. It is usually the name of the brokerage in capital letters and will never change.
	Slug *string `json:"slug,omitempty"`
	Url *string `json:"url,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	MaintenanceMode *bool `json:"maintenance_mode,omitempty"`
	AllowsFractionalUnits NullableBool `json:"allows_fractional_units,omitempty"`
	AllowsTrading NullableBool `json:"allows_trading,omitempty"`
	HasReporting NullableBool `json:"has_reporting,omitempty"`
	IsRealTimeConnection *bool `json:"is_real_time_connection,omitempty"`
	AllowsTradingThroughSnaptradeApi NullableBool `json:"allows_trading_through_snaptrade_api,omitempty"`
	IsScrapingIntegration *bool `json:"is_scraping_integration,omitempty"`
	DefaultCurrency *string `json:"default_currency,omitempty"`
	BrokerageType *BrokerageType `json:"brokerage_type,omitempty"`
	// List of exchange ID supported by brokerage
	Exchanges []interface{} `json:"exchanges,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Brokerage Brokerage

// NewBrokerage instantiates a new Brokerage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerage() *Brokerage {
	this := Brokerage{}
	return &this
}

// NewBrokerageWithDefaults instantiates a new Brokerage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerageWithDefaults() *Brokerage {
	this := Brokerage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Brokerage) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Brokerage) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Brokerage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Brokerage) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Brokerage) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Brokerage) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Brokerage) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Brokerage) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Brokerage) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Brokerage) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Brokerage) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Brokerage) SetDescription(v string) {
	o.Description = &v
}

// GetAwsS3LogoUrl returns the AwsS3LogoUrl field value if set, zero value otherwise.
func (o *Brokerage) GetAwsS3LogoUrl() string {
	if o == nil || isNil(o.AwsS3LogoUrl) {
		var ret string
		return ret
	}
	return *o.AwsS3LogoUrl
}

// GetAwsS3LogoUrlOk returns a tuple with the AwsS3LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetAwsS3LogoUrlOk() (*string, bool) {
	if o == nil || isNil(o.AwsS3LogoUrl) {
    return nil, false
	}
	return o.AwsS3LogoUrl, true
}

// HasAwsS3LogoUrl returns a boolean if a field has been set.
func (o *Brokerage) HasAwsS3LogoUrl() bool {
	if o != nil && !isNil(o.AwsS3LogoUrl) {
		return true
	}

	return false
}

// SetAwsS3LogoUrl gets a reference to the given string and assigns it to the AwsS3LogoUrl field.
func (o *Brokerage) SetAwsS3LogoUrl(v string) {
	o.AwsS3LogoUrl = &v
}

// GetAwsS3SquareLogoUrl returns the AwsS3SquareLogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brokerage) GetAwsS3SquareLogoUrl() string {
	if o == nil || isNil(o.AwsS3SquareLogoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AwsS3SquareLogoUrl.Get()
}

// GetAwsS3SquareLogoUrlOk returns a tuple with the AwsS3SquareLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brokerage) GetAwsS3SquareLogoUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AwsS3SquareLogoUrl.Get(), o.AwsS3SquareLogoUrl.IsSet()
}

// HasAwsS3SquareLogoUrl returns a boolean if a field has been set.
func (o *Brokerage) HasAwsS3SquareLogoUrl() bool {
	if o != nil && o.AwsS3SquareLogoUrl.IsSet() {
		return true
	}

	return false
}

// SetAwsS3SquareLogoUrl gets a reference to the given NullableString and assigns it to the AwsS3SquareLogoUrl field.
func (o *Brokerage) SetAwsS3SquareLogoUrl(v string) {
	o.AwsS3SquareLogoUrl.Set(&v)
}
// SetAwsS3SquareLogoUrlNil sets the value for AwsS3SquareLogoUrl to be an explicit nil
func (o *Brokerage) SetAwsS3SquareLogoUrlNil() {
	o.AwsS3SquareLogoUrl.Set(nil)
}

// UnsetAwsS3SquareLogoUrl ensures that no value is present for AwsS3SquareLogoUrl, not even an explicit nil
func (o *Brokerage) UnsetAwsS3SquareLogoUrl() {
	o.AwsS3SquareLogoUrl.Unset()
}

// GetOpenUrl returns the OpenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brokerage) GetOpenUrl() string {
	if o == nil || isNil(o.OpenUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OpenUrl.Get()
}

// GetOpenUrlOk returns a tuple with the OpenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brokerage) GetOpenUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.OpenUrl.Get(), o.OpenUrl.IsSet()
}

// HasOpenUrl returns a boolean if a field has been set.
func (o *Brokerage) HasOpenUrl() bool {
	if o != nil && o.OpenUrl.IsSet() {
		return true
	}

	return false
}

// SetOpenUrl gets a reference to the given NullableString and assigns it to the OpenUrl field.
func (o *Brokerage) SetOpenUrl(v string) {
	o.OpenUrl.Set(&v)
}
// SetOpenUrlNil sets the value for OpenUrl to be an explicit nil
func (o *Brokerage) SetOpenUrlNil() {
	o.OpenUrl.Set(nil)
}

// UnsetOpenUrl ensures that no value is present for OpenUrl, not even an explicit nil
func (o *Brokerage) UnsetOpenUrl() {
	o.OpenUrl.Unset()
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Brokerage) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
    return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Brokerage) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Brokerage) SetSlug(v string) {
	o.Slug = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Brokerage) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Brokerage) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Brokerage) SetUrl(v string) {
	o.Url = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Brokerage) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Brokerage) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Brokerage) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaintenanceMode returns the MaintenanceMode field value if set, zero value otherwise.
func (o *Brokerage) GetMaintenanceMode() bool {
	if o == nil || isNil(o.MaintenanceMode) {
		var ret bool
		return ret
	}
	return *o.MaintenanceMode
}

// GetMaintenanceModeOk returns a tuple with the MaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetMaintenanceModeOk() (*bool, bool) {
	if o == nil || isNil(o.MaintenanceMode) {
    return nil, false
	}
	return o.MaintenanceMode, true
}

// HasMaintenanceMode returns a boolean if a field has been set.
func (o *Brokerage) HasMaintenanceMode() bool {
	if o != nil && !isNil(o.MaintenanceMode) {
		return true
	}

	return false
}

// SetMaintenanceMode gets a reference to the given bool and assigns it to the MaintenanceMode field.
func (o *Brokerage) SetMaintenanceMode(v bool) {
	o.MaintenanceMode = &v
}

// GetAllowsFractionalUnits returns the AllowsFractionalUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brokerage) GetAllowsFractionalUnits() bool {
	if o == nil || isNil(o.AllowsFractionalUnits.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowsFractionalUnits.Get()
}

// GetAllowsFractionalUnitsOk returns a tuple with the AllowsFractionalUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brokerage) GetAllowsFractionalUnitsOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowsFractionalUnits.Get(), o.AllowsFractionalUnits.IsSet()
}

// HasAllowsFractionalUnits returns a boolean if a field has been set.
func (o *Brokerage) HasAllowsFractionalUnits() bool {
	if o != nil && o.AllowsFractionalUnits.IsSet() {
		return true
	}

	return false
}

// SetAllowsFractionalUnits gets a reference to the given NullableBool and assigns it to the AllowsFractionalUnits field.
func (o *Brokerage) SetAllowsFractionalUnits(v bool) {
	o.AllowsFractionalUnits.Set(&v)
}
// SetAllowsFractionalUnitsNil sets the value for AllowsFractionalUnits to be an explicit nil
func (o *Brokerage) SetAllowsFractionalUnitsNil() {
	o.AllowsFractionalUnits.Set(nil)
}

// UnsetAllowsFractionalUnits ensures that no value is present for AllowsFractionalUnits, not even an explicit nil
func (o *Brokerage) UnsetAllowsFractionalUnits() {
	o.AllowsFractionalUnits.Unset()
}

// GetAllowsTrading returns the AllowsTrading field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brokerage) GetAllowsTrading() bool {
	if o == nil || isNil(o.AllowsTrading.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowsTrading.Get()
}

// GetAllowsTradingOk returns a tuple with the AllowsTrading field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brokerage) GetAllowsTradingOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowsTrading.Get(), o.AllowsTrading.IsSet()
}

// HasAllowsTrading returns a boolean if a field has been set.
func (o *Brokerage) HasAllowsTrading() bool {
	if o != nil && o.AllowsTrading.IsSet() {
		return true
	}

	return false
}

// SetAllowsTrading gets a reference to the given NullableBool and assigns it to the AllowsTrading field.
func (o *Brokerage) SetAllowsTrading(v bool) {
	o.AllowsTrading.Set(&v)
}
// SetAllowsTradingNil sets the value for AllowsTrading to be an explicit nil
func (o *Brokerage) SetAllowsTradingNil() {
	o.AllowsTrading.Set(nil)
}

// UnsetAllowsTrading ensures that no value is present for AllowsTrading, not even an explicit nil
func (o *Brokerage) UnsetAllowsTrading() {
	o.AllowsTrading.Unset()
}

// GetHasReporting returns the HasReporting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brokerage) GetHasReporting() bool {
	if o == nil || isNil(o.HasReporting.Get()) {
		var ret bool
		return ret
	}
	return *o.HasReporting.Get()
}

// GetHasReportingOk returns a tuple with the HasReporting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brokerage) GetHasReportingOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.HasReporting.Get(), o.HasReporting.IsSet()
}

// HasHasReporting returns a boolean if a field has been set.
func (o *Brokerage) HasHasReporting() bool {
	if o != nil && o.HasReporting.IsSet() {
		return true
	}

	return false
}

// SetHasReporting gets a reference to the given NullableBool and assigns it to the HasReporting field.
func (o *Brokerage) SetHasReporting(v bool) {
	o.HasReporting.Set(&v)
}
// SetHasReportingNil sets the value for HasReporting to be an explicit nil
func (o *Brokerage) SetHasReportingNil() {
	o.HasReporting.Set(nil)
}

// UnsetHasReporting ensures that no value is present for HasReporting, not even an explicit nil
func (o *Brokerage) UnsetHasReporting() {
	o.HasReporting.Unset()
}

// GetIsRealTimeConnection returns the IsRealTimeConnection field value if set, zero value otherwise.
func (o *Brokerage) GetIsRealTimeConnection() bool {
	if o == nil || isNil(o.IsRealTimeConnection) {
		var ret bool
		return ret
	}
	return *o.IsRealTimeConnection
}

// GetIsRealTimeConnectionOk returns a tuple with the IsRealTimeConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetIsRealTimeConnectionOk() (*bool, bool) {
	if o == nil || isNil(o.IsRealTimeConnection) {
    return nil, false
	}
	return o.IsRealTimeConnection, true
}

// HasIsRealTimeConnection returns a boolean if a field has been set.
func (o *Brokerage) HasIsRealTimeConnection() bool {
	if o != nil && !isNil(o.IsRealTimeConnection) {
		return true
	}

	return false
}

// SetIsRealTimeConnection gets a reference to the given bool and assigns it to the IsRealTimeConnection field.
func (o *Brokerage) SetIsRealTimeConnection(v bool) {
	o.IsRealTimeConnection = &v
}

// GetAllowsTradingThroughSnaptradeApi returns the AllowsTradingThroughSnaptradeApi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brokerage) GetAllowsTradingThroughSnaptradeApi() bool {
	if o == nil || isNil(o.AllowsTradingThroughSnaptradeApi.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowsTradingThroughSnaptradeApi.Get()
}

// GetAllowsTradingThroughSnaptradeApiOk returns a tuple with the AllowsTradingThroughSnaptradeApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brokerage) GetAllowsTradingThroughSnaptradeApiOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowsTradingThroughSnaptradeApi.Get(), o.AllowsTradingThroughSnaptradeApi.IsSet()
}

// HasAllowsTradingThroughSnaptradeApi returns a boolean if a field has been set.
func (o *Brokerage) HasAllowsTradingThroughSnaptradeApi() bool {
	if o != nil && o.AllowsTradingThroughSnaptradeApi.IsSet() {
		return true
	}

	return false
}

// SetAllowsTradingThroughSnaptradeApi gets a reference to the given NullableBool and assigns it to the AllowsTradingThroughSnaptradeApi field.
func (o *Brokerage) SetAllowsTradingThroughSnaptradeApi(v bool) {
	o.AllowsTradingThroughSnaptradeApi.Set(&v)
}
// SetAllowsTradingThroughSnaptradeApiNil sets the value for AllowsTradingThroughSnaptradeApi to be an explicit nil
func (o *Brokerage) SetAllowsTradingThroughSnaptradeApiNil() {
	o.AllowsTradingThroughSnaptradeApi.Set(nil)
}

// UnsetAllowsTradingThroughSnaptradeApi ensures that no value is present for AllowsTradingThroughSnaptradeApi, not even an explicit nil
func (o *Brokerage) UnsetAllowsTradingThroughSnaptradeApi() {
	o.AllowsTradingThroughSnaptradeApi.Unset()
}

// GetIsScrapingIntegration returns the IsScrapingIntegration field value if set, zero value otherwise.
func (o *Brokerage) GetIsScrapingIntegration() bool {
	if o == nil || isNil(o.IsScrapingIntegration) {
		var ret bool
		return ret
	}
	return *o.IsScrapingIntegration
}

// GetIsScrapingIntegrationOk returns a tuple with the IsScrapingIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetIsScrapingIntegrationOk() (*bool, bool) {
	if o == nil || isNil(o.IsScrapingIntegration) {
    return nil, false
	}
	return o.IsScrapingIntegration, true
}

// HasIsScrapingIntegration returns a boolean if a field has been set.
func (o *Brokerage) HasIsScrapingIntegration() bool {
	if o != nil && !isNil(o.IsScrapingIntegration) {
		return true
	}

	return false
}

// SetIsScrapingIntegration gets a reference to the given bool and assigns it to the IsScrapingIntegration field.
func (o *Brokerage) SetIsScrapingIntegration(v bool) {
	o.IsScrapingIntegration = &v
}

// GetDefaultCurrency returns the DefaultCurrency field value if set, zero value otherwise.
func (o *Brokerage) GetDefaultCurrency() string {
	if o == nil || isNil(o.DefaultCurrency) {
		var ret string
		return ret
	}
	return *o.DefaultCurrency
}

// GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetDefaultCurrencyOk() (*string, bool) {
	if o == nil || isNil(o.DefaultCurrency) {
    return nil, false
	}
	return o.DefaultCurrency, true
}

// HasDefaultCurrency returns a boolean if a field has been set.
func (o *Brokerage) HasDefaultCurrency() bool {
	if o != nil && !isNil(o.DefaultCurrency) {
		return true
	}

	return false
}

// SetDefaultCurrency gets a reference to the given string and assigns it to the DefaultCurrency field.
func (o *Brokerage) SetDefaultCurrency(v string) {
	o.DefaultCurrency = &v
}

// GetBrokerageType returns the BrokerageType field value if set, zero value otherwise.
func (o *Brokerage) GetBrokerageType() BrokerageType {
	if o == nil || isNil(o.BrokerageType) {
		var ret BrokerageType
		return ret
	}
	return *o.BrokerageType
}

// GetBrokerageTypeOk returns a tuple with the BrokerageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetBrokerageTypeOk() (*BrokerageType, bool) {
	if o == nil || isNil(o.BrokerageType) {
    return nil, false
	}
	return o.BrokerageType, true
}

// HasBrokerageType returns a boolean if a field has been set.
func (o *Brokerage) HasBrokerageType() bool {
	if o != nil && !isNil(o.BrokerageType) {
		return true
	}

	return false
}

// SetBrokerageType gets a reference to the given BrokerageType and assigns it to the BrokerageType field.
func (o *Brokerage) SetBrokerageType(v BrokerageType) {
	o.BrokerageType = &v
}

// GetExchanges returns the Exchanges field value if set, zero value otherwise.
func (o *Brokerage) GetExchanges() []interface{} {
	if o == nil || isNil(o.Exchanges) {
		var ret []interface{}
		return ret
	}
	return o.Exchanges
}

// GetExchangesOk returns a tuple with the Exchanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brokerage) GetExchangesOk() ([]interface{}, bool) {
	if o == nil || isNil(o.Exchanges) {
    return nil, false
	}
	return o.Exchanges, true
}

// HasExchanges returns a boolean if a field has been set.
func (o *Brokerage) HasExchanges() bool {
	if o != nil && !isNil(o.Exchanges) {
		return true
	}

	return false
}

// SetExchanges gets a reference to the given []interface{} and assigns it to the Exchanges field.
func (o *Brokerage) SetExchanges(v []interface{}) {
	o.Exchanges = v
}

func (o Brokerage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.AwsS3LogoUrl) {
		toSerialize["aws_s3_logo_url"] = o.AwsS3LogoUrl
	}
	if o.AwsS3SquareLogoUrl.IsSet() {
		toSerialize["aws_s3_square_logo_url"] = o.AwsS3SquareLogoUrl.Get()
	}
	if o.OpenUrl.IsSet() {
		toSerialize["open_url"] = o.OpenUrl.Get()
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.MaintenanceMode) {
		toSerialize["maintenance_mode"] = o.MaintenanceMode
	}
	if o.AllowsFractionalUnits.IsSet() {
		toSerialize["allows_fractional_units"] = o.AllowsFractionalUnits.Get()
	}
	if o.AllowsTrading.IsSet() {
		toSerialize["allows_trading"] = o.AllowsTrading.Get()
	}
	if o.HasReporting.IsSet() {
		toSerialize["has_reporting"] = o.HasReporting.Get()
	}
	if !isNil(o.IsRealTimeConnection) {
		toSerialize["is_real_time_connection"] = o.IsRealTimeConnection
	}
	if o.AllowsTradingThroughSnaptradeApi.IsSet() {
		toSerialize["allows_trading_through_snaptrade_api"] = o.AllowsTradingThroughSnaptradeApi.Get()
	}
	if !isNil(o.IsScrapingIntegration) {
		toSerialize["is_scraping_integration"] = o.IsScrapingIntegration
	}
	if !isNil(o.DefaultCurrency) {
		toSerialize["default_currency"] = o.DefaultCurrency
	}
	if !isNil(o.BrokerageType) {
		toSerialize["brokerage_type"] = o.BrokerageType
	}
	if !isNil(o.Exchanges) {
		toSerialize["exchanges"] = o.Exchanges
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Brokerage) UnmarshalJSON(bytes []byte) (err error) {
	varBrokerage := _Brokerage{}

	if err = json.Unmarshal(bytes, &varBrokerage); err == nil {
		*o = Brokerage(varBrokerage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "aws_s3_logo_url")
		delete(additionalProperties, "aws_s3_square_logo_url")
		delete(additionalProperties, "open_url")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "url")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "maintenance_mode")
		delete(additionalProperties, "allows_fractional_units")
		delete(additionalProperties, "allows_trading")
		delete(additionalProperties, "has_reporting")
		delete(additionalProperties, "is_real_time_connection")
		delete(additionalProperties, "allows_trading_through_snaptrade_api")
		delete(additionalProperties, "is_scraping_integration")
		delete(additionalProperties, "default_currency")
		delete(additionalProperties, "brokerage_type")
		delete(additionalProperties, "exchanges")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrokerage struct {
	value *Brokerage
	isSet bool
}

func (v NullableBrokerage) Get() *Brokerage {
	return v.value
}

func (v *NullableBrokerage) Set(val *Brokerage) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerage) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerage(val *Brokerage) *NullableBrokerage {
	return &NullableBrokerage{value: val, isSet: true}
}

func (v NullableBrokerage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



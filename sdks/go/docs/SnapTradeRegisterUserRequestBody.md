# SnapTradeRegisterUserRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | SnapTrade User ID. This is chosen by the API partner and can be any string that is a) unique to the user, and b) immutable for the user. It is recommended to NOT use email addresses for this property because they are usually not immutable. | [optional] 

## Methods

### NewSnapTradeRegisterUserRequestBody

`func NewSnapTradeRegisterUserRequestBody() *SnapTradeRegisterUserRequestBody`

NewSnapTradeRegisterUserRequestBody instantiates a new SnapTradeRegisterUserRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapTradeRegisterUserRequestBodyWithDefaults

`func NewSnapTradeRegisterUserRequestBodyWithDefaults() *SnapTradeRegisterUserRequestBody`

NewSnapTradeRegisterUserRequestBodyWithDefaults instantiates a new SnapTradeRegisterUserRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SnapTradeRegisterUserRequestBody) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SnapTradeRegisterUserRequestBody) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SnapTradeRegisterUserRequestBody) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SnapTradeRegisterUserRequestBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



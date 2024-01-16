# SnapTrade.Net.Model.ManualTradeForm
Manual Trade Form

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | [optional] 
**_Action** | **ModelAction** |  | [optional] 
**OrderType** | **OrderType** |  | [optional] 
**Price** | **double?** | Trade Price if limit or stop limit order | [optional] 
**Stop** | **double?** | Stop Price. If stop loss or stop limit order, the price to trigger the stop | [optional] 
**TimeInForce** | **TimeInForce** |  | [optional] 
**Units** | **double?** | Trade Units. Cannot work with notional value. | [optional] 
**UniversalSymbolId** | **string** |  | [optional] 
**NotionalValue** | **double?** | Dollar amount to trade. Cannot work with units. Can only work for market order types and day for time in force. **Only available for Alpaca and Alpaca Paper. Please contact support to get access to place notional trades** | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


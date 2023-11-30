# SnapTrade.Net - C#

[![NuGet](https://img.shields.io/badge/NuGet-4.0.6-blue.svg)](https://www.nuget.org/packages/SnapTrade.Net/4.0.6)
[![More Info](https://img.shields.io/badge/More%20Info-Click%20Here-orange)](https://snaptrade.com/)

Connect brokerage accounts to your app for live positions and trading

## Frameworks supported
- .NET Core >=1.0
- .NET Framework >=4.6
- Mono/Xamarin >=vNext

## Installation

Using the [.NET Core command-line interface (CLI) tools][dotnet-core-cli-tools]:

```sh
dotnet add package SnapTrade.Net
```

Using the [NuGet Command Line Interface (CLI)][nuget-cli]:

```sh
nuget install SnapTrade.Net
```

Using the [Package Manager Console][package-manager-console]:

```powershell
Install-Package SnapTrade.Net
```

From within Visual Studio:

1. Open the Solution Explorer.
2. Right-click on a project within your solution.
3. Click on *Manage NuGet Packages...*
4. Click on the *Browse* tab and search for "SnapTrade.Net".
5. Click on the "SnapTrade.Net" package, select the appropriate version in the
   right-tab and click *Install*.

## Getting Started

```csharp
using System;
using System.Collections.Generic;
using System.Diagnostics;
using SnapTrade.Net.Client;
using SnapTrade.Net.Model;

namespace Example
{
    public class GetAllUserHoldingsExample
    {
        public static void Main()
        {
            Snaptrade client = new Snaptrade();
            // Configure custom BasePath if desired
            // client.SetBasePath("https://api.snaptrade.com/api/v1");
            client.SetClientId(System.Environment.GetEnvironmentVariable("SNAPTRADE_CLIENT_ID"));
            client.SetConsumerKey(System.Environment.GetEnvironmentVariable("SNAPTRADE_CONSUMER_KEY"));

            var userId = "userId_example";
            var userSecret = "userSecret_example";
            var brokerageAuthorizations = "917c8734-8470-4a3e-a18f-57c3f2ee6631"; // Optional. Comma seperated list of authorization IDs (only use if filtering is needed on one or more authorizations). (optional) 
            
            try
            {
                // List all accounts for the user, plus balances, positions, and orders for each account.
                List<AccountHoldings> result = client.AccountInformation.GetAllUserHoldings(userId, userSecret, brokerageAuthorizations);
                Console.WriteLine(result);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling AccountInformationApi.GetAllUserHoldings: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}
```

## Documentation for API Endpoints

All URIs are relative to *https://api.snaptrade.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountInformationApi* | [**GetAllUserHoldings**](docs/AccountInformationApi.md#getalluserholdings) | **GET** /holdings | List all accounts for the user, plus balances, positions, and orders for each account.
*AccountInformationApi* | [**GetUserAccountBalance**](docs/AccountInformationApi.md#getuseraccountbalance) | **GET** /accounts/{accountId}/balances | List account balances
*AccountInformationApi* | [**GetUserAccountDetails**](docs/AccountInformationApi.md#getuseraccountdetails) | **GET** /accounts/{accountId} | Return details of a specific investment account
*AccountInformationApi* | [**GetUserAccountOrders**](docs/AccountInformationApi.md#getuseraccountorders) | **GET** /accounts/{accountId}/orders | List account orders
*AccountInformationApi* | [**GetUserAccountPositions**](docs/AccountInformationApi.md#getuseraccountpositions) | **GET** /accounts/{accountId}/positions | List account positions
*AccountInformationApi* | [**GetUserHoldings**](docs/AccountInformationApi.md#getuserholdings) | **GET** /accounts/{accountId}/holdings | List balances, positions and orders for the specified account
*AccountInformationApi* | [**ListUserAccounts**](docs/AccountInformationApi.md#listuseraccounts) | **GET** /accounts | List accounts
*AccountInformationApi* | [**UpdateUserAccount**](docs/AccountInformationApi.md#updateuseraccount) | **PUT** /accounts/{accountId} | Update details of an investment account
*APIStatusApi* | [**Check**](docs/APIStatusApi.md#check) | **GET** / | Get API Status
*AuthenticationApi* | [**DeleteSnapTradeUser**](docs/AuthenticationApi.md#deletesnaptradeuser) | **DELETE** /snapTrade/deleteUser | Delete SnapTrade user
*AuthenticationApi* | [**GetUserJWT**](docs/AuthenticationApi.md#getuserjwt) | **GET** /snapTrade/encryptedJWT | Generate encrypted JWT token
*AuthenticationApi* | [**ListSnapTradeUsers**](docs/AuthenticationApi.md#listsnaptradeusers) | **GET** /snapTrade/listUsers | List SnapTrade users
*AuthenticationApi* | [**LoginSnapTradeUser**](docs/AuthenticationApi.md#loginsnaptradeuser) | **POST** /snapTrade/login | Login user & generate connection link
*AuthenticationApi* | [**RegisterSnapTradeUser**](docs/AuthenticationApi.md#registersnaptradeuser) | **POST** /snapTrade/registerUser | Create SnapTrade user
*ConnectionsApi* | [**DetailBrokerageAuthorization**](docs/ConnectionsApi.md#detailbrokerageauthorization) | **GET** /authorizations/{authorizationId} | Get brokerage authorization details
*ConnectionsApi* | [**ListBrokerageAuthorizations**](docs/ConnectionsApi.md#listbrokerageauthorizations) | **GET** /authorizations | List all brokerage authorizations for the user
*ConnectionsApi* | [**RemoveBrokerageAuthorization**](docs/ConnectionsApi.md#removebrokerageauthorization) | **DELETE** /authorizations/{authorizationId} | Delete brokerage authorization
*ConnectionsApi* | [**SessionEvents**](docs/ConnectionsApi.md#sessionevents) | **GET** /sessionEvents | List all session events for the partner
*ErrorLogsApi* | [**ListUserErrors**](docs/ErrorLogsApi.md#listusererrors) | **GET** /snapTrade/listUserErrors | Retrieve error logs on behalf of your SnapTrade users
*OptionsApi* | [**GetOptionStrategy**](docs/OptionsApi.md#getoptionstrategy) | **POST** /accounts/{accountId}/optionStrategy | Creates an option strategy object that will be used to place an option strategy order
*OptionsApi* | [**GetOptionsChain**](docs/OptionsApi.md#getoptionschain) | **GET** /accounts/{accountId}/optionsChain | Get the options chain
*OptionsApi* | [**GetOptionsStrategyQuote**](docs/OptionsApi.md#getoptionsstrategyquote) | **GET** /accounts/{accountId}/optionStrategy/{optionStrategyId} | Get latest market data of option strategy
*OptionsApi* | [**ListOptionHoldings**](docs/OptionsApi.md#listoptionholdings) | **GET** /accounts/{accountId}/options | Get the options holdings in the account
*OptionsApi* | [**PlaceOptionStrategy**](docs/OptionsApi.md#placeoptionstrategy) | **POST** /accounts/{accountId}/optionStrategy/{optionStrategyId}/execute | Place an option strategy order on the brokerage
*ReferenceDataApi* | [**GetCurrencyExchangeRatePair**](docs/ReferenceDataApi.md#getcurrencyexchangeratepair) | **GET** /currencies/rates/{currencyPair} | Return the exchange rate of a currency pair
*ReferenceDataApi* | [**GetPartnerInfo**](docs/ReferenceDataApi.md#getpartnerinfo) | **GET** /snapTrade/partners | Get metadata related to Snaptrade partner
*ReferenceDataApi* | [**GetSecurityTypes**](docs/ReferenceDataApi.md#getsecuritytypes) | **GET** /securityTypes | List of all security types
*ReferenceDataApi* | [**GetStockExchanges**](docs/ReferenceDataApi.md#getstockexchanges) | **GET** /exchanges | List exchanges
*ReferenceDataApi* | [**GetSymbols**](docs/ReferenceDataApi.md#getsymbols) | **POST** /symbols | Search for symbols
*ReferenceDataApi* | [**GetSymbolsByTicker**](docs/ReferenceDataApi.md#getsymbolsbyticker) | **GET** /symbols/{ticker} | Get details of a symbol by the ticker
*ReferenceDataApi* | [**ListAllBrokerageAuthorizationType**](docs/ReferenceDataApi.md#listallbrokerageauthorizationtype) | **GET** /brokerageAuthorizationTypes | List of all brokerage authorization types
*ReferenceDataApi* | [**ListAllBrokerages**](docs/ReferenceDataApi.md#listallbrokerages) | **GET** /brokerages | List brokerages
*ReferenceDataApi* | [**ListAllCurrencies**](docs/ReferenceDataApi.md#listallcurrencies) | **GET** /currencies | List currencies
*ReferenceDataApi* | [**ListAllCurrenciesRates**](docs/ReferenceDataApi.md#listallcurrenciesrates) | **GET** /currencies/rates | List currency exchange rates
*ReferenceDataApi* | [**SymbolSearchUserAccount**](docs/ReferenceDataApi.md#symbolsearchuseraccount) | **POST** /accounts/{accountId}/symbols | Search for symbols available in an account
*TradingApi* | [**CancelUserAccountOrder**](docs/TradingApi.md#canceluseraccountorder) | **POST** /accounts/{accountId}/orders/cancel | Cancel open order in account
*TradingApi* | [**GetOrderImpact**](docs/TradingApi.md#getorderimpact) | **POST** /trade/impact | Check impact of trades on account.
*TradingApi* | [**GetUserAccountQuotes**](docs/TradingApi.md#getuseraccountquotes) | **GET** /accounts/{accountId}/quotes | Get symbol quotes
*TradingApi* | [**PlaceForceOrder**](docs/TradingApi.md#placeforceorder) | **POST** /trade/place | Place a trade with NO validation.
*TradingApi* | [**PlaceOCOOrder**](docs/TradingApi.md#placeocoorder) | **POST** /trade/oco | Place a OCO (One Cancels Other) order
*TradingApi* | [**PlaceOrder**](docs/TradingApi.md#placeorder) | **POST** /trade/{tradeId} | Place order
*TransactionsAndReportingApi* | [**GetActivities**](docs/TransactionsAndReportingApi.md#getactivities) | **GET** /activities | Get transaction history for a user
*TransactionsAndReportingApi* | [**GetReportingCustomRange**](docs/TransactionsAndReportingApi.md#getreportingcustomrange) | **GET** /performance/custom | Get performance information for a specific timeframe


## Documentation for Models

 - [Account](docs/Account.md)
 - [AccountBalance](docs/AccountBalance.md)
 - [AccountBalanceTotal](docs/AccountBalanceTotal.md)
 - [AccountHoldings](docs/AccountHoldings.md)
 - [AccountHoldingsAccount](docs/AccountHoldingsAccount.md)
 - [AccountOrderRecord](docs/AccountOrderRecord.md)
 - [AccountOrderRecordStatus](docs/AccountOrderRecordStatus.md)
 - [AccountSimple](docs/AccountSimple.md)
 - [AccountSyncStatus](docs/AccountSyncStatus.md)
 - [Amount](docs/Amount.md)
 - [AuthenticationLoginSnapTradeUser200Response](docs/AuthenticationLoginSnapTradeUser200Response.md)
 - [Balance](docs/Balance.md)
 - [Brokerage](docs/Brokerage.md)
 - [BrokerageAuthorization](docs/BrokerageAuthorization.md)
 - [BrokerageAuthorizationType](docs/BrokerageAuthorizationType.md)
 - [BrokerageAuthorizationTypeReadOnly](docs/BrokerageAuthorizationTypeReadOnly.md)
 - [BrokerageAuthorizationTypeReadOnlyBrokerage](docs/BrokerageAuthorizationTypeReadOnlyBrokerage.md)
 - [BrokerageSymbol](docs/BrokerageSymbol.md)
 - [BrokerageSymbolOptionSymbol](docs/BrokerageSymbolOptionSymbol.md)
 - [BrokerageSymbolSymbol](docs/BrokerageSymbolSymbol.md)
 - [BrokerageType](docs/BrokerageType.md)
 - [CalculatedTrade](docs/CalculatedTrade.md)
 - [CashRestriction](docs/CashRestriction.md)
 - [ConnectionsSessionEvents200ResponseInner](docs/ConnectionsSessionEvents200ResponseInner.md)
 - [Currency](docs/Currency.md)
 - [DeleteUserResponse](docs/DeleteUserResponse.md)
 - [DividendAtDate](docs/DividendAtDate.md)
 - [EncryptedResponse](docs/EncryptedResponse.md)
 - [EncryptedResponseEncryptedMessageData](docs/EncryptedResponseEncryptedMessageData.md)
 - [Exchange](docs/Exchange.md)
 - [ExchangeRatePairs](docs/ExchangeRatePairs.md)
 - [ExcludedAsset](docs/ExcludedAsset.md)
 - [JWT](docs/JWT.md)
 - [LoginRedirectURI](docs/LoginRedirectURI.md)
 - [ManualTrade](docs/ManualTrade.md)
 - [ManualTradeAndImpact](docs/ManualTradeAndImpact.md)
 - [ManualTradeBalance](docs/ManualTradeBalance.md)
 - [ManualTradeForm](docs/ManualTradeForm.md)
 - [ManualTradeSymbol](docs/ManualTradeSymbol.md)
 - [Model400FailedRequestResponse](docs/Model400FailedRequestResponse.md)
 - [Model401FailedRequestResponse](docs/Model401FailedRequestResponse.md)
 - [Model403FailedRequestResponse](docs/Model403FailedRequestResponse.md)
 - [Model404FailedRequestResponse](docs/Model404FailedRequestResponse.md)
 - [Model500UnexpectedExceptionResponse](docs/Model500UnexpectedExceptionResponse.md)
 - [ModelAction](docs/ModelAction.md)
 - [ModelAssetClass](docs/ModelAssetClass.md)
 - [ModelAssetClassDetails](docs/ModelAssetClassDetails.md)
 - [ModelAssetClassTarget](docs/ModelAssetClassTarget.md)
 - [ModelPortfolio](docs/ModelPortfolio.md)
 - [ModelPortfolioAssetClass](docs/ModelPortfolioAssetClass.md)
 - [ModelPortfolioDetails](docs/ModelPortfolioDetails.md)
 - [ModelPortfolioSecurity](docs/ModelPortfolioSecurity.md)
 - [MonthlyDividends](docs/MonthlyDividends.md)
 - [NetContributions](docs/NetContributions.md)
 - [NetDividend](docs/NetDividend.md)
 - [OptionChainInner](docs/OptionChainInner.md)
 - [OptionChainInnerChainPerRootInner](docs/OptionChainInnerChainPerRootInner.md)
 - [OptionChainInnerChainPerRootInnerChainPerStrikePriceInner](docs/OptionChainInnerChainPerRootInnerChainPerStrikePriceInner.md)
 - [OptionLeg](docs/OptionLeg.md)
 - [OptionStrategy](docs/OptionStrategy.md)
 - [OptionStrategyLegsInner](docs/OptionStrategyLegsInner.md)
 - [OptionsGetOptionStrategyRequest](docs/OptionsGetOptionStrategyRequest.md)
 - [OptionsHoldings](docs/OptionsHoldings.md)
 - [OptionsPlaceOptionStrategyRequest](docs/OptionsPlaceOptionStrategyRequest.md)
 - [OptionsPosition](docs/OptionsPosition.md)
 - [OptionsPositionCurrency](docs/OptionsPositionCurrency.md)
 - [OptionsSymbol](docs/OptionsSymbol.md)
 - [OrderType](docs/OrderType.md)
 - [PartnerData](docs/PartnerData.md)
 - [PastValue](docs/PastValue.md)
 - [PerformanceCustom](docs/PerformanceCustom.md)
 - [PortfolioGroup](docs/PortfolioGroup.md)
 - [PortfolioGroupInfo](docs/PortfolioGroupInfo.md)
 - [PortfolioGroupPosition](docs/PortfolioGroupPosition.md)
 - [PortfolioGroupSettings](docs/PortfolioGroupSettings.md)
 - [Position](docs/Position.md)
 - [PositionSymbol](docs/PositionSymbol.md)
 - [RedirectTokenandPin](docs/RedirectTokenandPin.md)
 - [SecurityType](docs/SecurityType.md)
 - [SessionEvent](docs/SessionEvent.md)
 - [SnapTradeAPIDisclaimerAcceptStatus](docs/SnapTradeAPIDisclaimerAcceptStatus.md)
 - [SnapTradeHoldingsAccount](docs/SnapTradeHoldingsAccount.md)
 - [SnapTradeHoldingsAccountAccountId](docs/SnapTradeHoldingsAccountAccountId.md)
 - [SnapTradeHoldingsTotalValue](docs/SnapTradeHoldingsTotalValue.md)
 - [SnapTradeLoginUserRequestBody](docs/SnapTradeLoginUserRequestBody.md)
 - [SnapTradeRegisterUserRequestBody](docs/SnapTradeRegisterUserRequestBody.md)
 - [Status](docs/Status.md)
 - [StrategyImpact](docs/StrategyImpact.md)
 - [StrategyImpactLegsInner](docs/StrategyImpactLegsInner.md)
 - [StrategyOrderPlace](docs/StrategyOrderPlace.md)
 - [StrategyOrderPlaceOrdersInner](docs/StrategyOrderPlaceOrdersInner.md)
 - [StrategyOrderPlaceOrdersInnerLegsInner](docs/StrategyOrderPlaceOrdersInnerLegsInner.md)
 - [StrategyOrderRecord](docs/StrategyOrderRecord.md)
 - [StrategyQuotes](docs/StrategyQuotes.md)
 - [StrategyQuotesGreek](docs/StrategyQuotesGreek.md)
 - [SubPeriodReturnRate](docs/SubPeriodReturnRate.md)
 - [Symbol](docs/Symbol.md)
 - [SymbolQuery](docs/SymbolQuery.md)
 - [SymbolsQuotesInner](docs/SymbolsQuotesInner.md)
 - [TargetAsset](docs/TargetAsset.md)
 - [TimeInForce](docs/TimeInForce.md)
 - [Trade](docs/Trade.md)
 - [TradeExecutionStatus](docs/TradeExecutionStatus.md)
 - [TradeImpact](docs/TradeImpact.md)
 - [TradingCancelUserAccountOrderRequest](docs/TradingCancelUserAccountOrderRequest.md)
 - [TradingPlaceOCOOrderRequest](docs/TradingPlaceOCOOrderRequest.md)
 - [TransactionsStatus](docs/TransactionsStatus.md)
 - [USExchange](docs/USExchange.md)
 - [UnderlyingSymbol](docs/UnderlyingSymbol.md)
 - [UniversalActivity](docs/UniversalActivity.md)
 - [UniversalSymbol](docs/UniversalSymbol.md)
 - [UniversalSymbolTicker](docs/UniversalSymbolTicker.md)
 - [UserErrorLog](docs/UserErrorLog.md)
 - [UserIDandSecret](docs/UserIDandSecret.md)
 - [UserSettings](docs/UserSettings.md)


## Dependencies

- [RestSharp](https://www.nuget.org/packages/RestSharp) - 106.13.0 or later
- [Json.NET](https://www.nuget.org/packages/Newtonsoft.Json/) - 13.0.1 or later
- [JsonSubTypes](https://www.nuget.org/packages/JsonSubTypes/) - 1.8.0 or later
- [System.ComponentModel.Annotations](https://www.nuget.org/packages/System.ComponentModel.Annotations) - 5.0.0 or later

## Author
This C# package is automatically generated by [Konfig](https://konfigthis.com)

[dotnet-core-cli-tools]: https://docs.microsoft.com/en-us/dotnet/core/tools/
[nuget-cli]: https://docs.microsoft.com/en-us/nuget/tools/nuget-exe-cli-reference
[package-manager-console]: https://docs.microsoft.com/en-us/nuget/tools/package-manager-console

<div align="center">

[![Visit SnapTrade](https://raw.githubusercontent.com/passiv/snaptrade-sdks/HEAD/sdks/java/header.png)](https://snaptrade.com)

# [SnapTrade](https://snaptrade.com)

Connect brokerage accounts to your app for live positions and trading

[![Maven Central](https://img.shields.io/badge/Maven%20Central-v5.0.31-blue)](https://central.sonatype.com/artifact/com.konfigthis/snaptrade-java-sdk/5.0.31)
[![More Info](https://img.shields.io/badge/More%20Info-Click%20Here-orange)](https://snaptrade.com/)

</div>

## Requirements

Building the API client library requires:

1. Java 1.8+
2. Maven (3.8.3+)/Gradle (7.2+)

If you are adding this library to an Android Application or Library:

3. Android 8.0+ (API Level 26+)

## Installation

To install the API client library to your local Maven repository, simply execute:

```shell
mvn clean install
```

To deploy it to a remote Maven repository instead, configure the settings of the repository and execute:

```shell
mvn clean deploy
```

Refer to the [OSSRH Guide](http://central.sonatype.org/pages/ossrh-guide.html) for more information.

### Maven users

Add this dependency to your project's POM:

```xml
<dependency>
  <groupId>com.konfigthis</groupId>
  <artifactId>snaptrade-java-sdk</artifactId>
  <version>5.0.31</version>
  <scope>compile</scope>
</dependency>
```

### Gradle users

Add this dependency to your `build.gradle`:

```groovy
// build.gradle
repositories {
  mavenCentral()
}

dependencies {
   implementation "com.konfigthis:snaptrade-java-sdk:5.0.31"
}
```

### Android users

Make sure your `build.gradle` file as a `minSdk` version of at least 26:
```groovy
// build.gradle
android {
    defaultConfig {
        minSdk 26
    }
}
```

Also make sure your library or application has internet permissions in your `AndroidManifest.xml`:

```xml
<!--AndroidManifest.xml-->
<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:tools="http://schemas.android.com/tools">
    <uses-permission android:name="android.permission.INTERNET"/>
</manifest>
```

### Others

At first generate the JAR by executing:

```shell
mvn clean package
```

Then manually install the following JARs:

* `target/snaptrade-java-sdk-5.0.31.jar`
* `target/lib/*.jar`

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Java code:

```java
import com.konfigthis.client.ApiException;
import com.konfigthis.client.Configuration;
import com.konfigthis.client.Snaptrade;
import com.konfigthis.client.model.*;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {

        public static void main(String[] args) {
                // 1) Initialize default client with clientID and consumerKey
                Configuration configuration = new Configuration();
                configuration.clientId = System.getenv("SNAPTRADE_CLIENT_ID");
                configuration.consumerKey = System.getenv("SNAPTRADE_CONSUMER_KEY");
                Snaptrade snaptrade = new Snaptrade(configuration);

                // 2) Check that the client is able to make a request to the API server
                Status status = snaptrade.apiStatus.check().execute();
                System.out.printf("SnapTrade is online: %s\n", status.getOnline());

                // 3) Create a new user on SnapTrade
                // The userId should be provided by you and refer to permanent value such as a
                // database row ID
                UUID userId = UUID.randomUUID();
                UserIDandSecret userIDandSecret = snaptrade.authentication.registerSnapTradeUser()
                                .userId(userId.toString()).execute();

                new SnapTradeRegisterUserRequestBody().userId(userId.toString());
                // Note: A user secret is only generated once. It's required to access resources
                // for certain endpoints
                System.out.printf("userID: %s, userSecret: %s\n", userIDandSecret.getUserId(),
                                userIDandSecret.getUserSecret());

                // 4) Get a redirect URI. Users will need this to connect their brokerage to the
                // SnapTrade server
                Map response = (Map) snaptrade.authentication
                                .loginSnapTradeUser(userIDandSecret.getUserId(),
                                                userIDandSecret.getUserSecret())
                                .execute();
                System.out.println(response.get("redirectURI"));

                // 5) Make a portfolio group and query
                List<PortfolioGroup> portfolioGroupsFromPost = snaptrade.portfolioManagement.create(
                                userIDandSecret.getUserId(), userIDandSecret.getUserSecret()).id(UUID.randomUUID())
                                .name("MyPortfolio").execute();
                System.out.println(portfolioGroupsFromPost);
                List<PortfolioGroup> portfolioGroups = snaptrade.portfolioManagement.list(userIDandSecret.getUserId(),
                                userIDandSecret.getUserSecret()).execute();
                System.out.println(portfolioGroups);

                // 7) Query holdings and available brokerages
                List<AccountHoldings> holdings = snaptrade.accountInformation
                                .getAllUserHoldings(userIDandSecret.getUserId(),
                                                userIDandSecret.getUserSecret())
                                .execute();
                System.out.println(holdings);
                List<Account> accounts = snaptrade.accountInformation.listUserAccounts(userIDandSecret.getUserId(),
                                userIDandSecret.getUserSecret()).execute();
                System.out.println(accounts);
                List<Brokerage> brokerages = snaptrade.referenceData.listAllBrokerages().execute();
                System.out.println(brokerages);

                // 8) Deleting a user
                DeleteUserResponse deleteUserResponse = snaptrade.authentication
                                .deleteSnapTradeUser(userIDandSecret.getUserId()).execute();
                System.out.println(deleteUserResponse);
        }
}

```

## Documentation for API Endpoints

All URIs are relative to *https://api.snaptrade.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountInformationApi* | [**getAllUserHoldings**](docs/AccountInformationApi.md#getAllUserHoldings) | **GET** /holdings | List all accounts for the user, plus balances, positions, and orders for each account.
*AccountInformationApi* | [**getUserAccountBalance**](docs/AccountInformationApi.md#getUserAccountBalance) | **GET** /accounts/{accountId}/balances | List account balances
*AccountInformationApi* | [**getUserAccountDetails**](docs/AccountInformationApi.md#getUserAccountDetails) | **GET** /accounts/{accountId} | Return details of a specific investment account
*AccountInformationApi* | [**getUserAccountOrders**](docs/AccountInformationApi.md#getUserAccountOrders) | **GET** /accounts/{accountId}/orders | List account orders
*AccountInformationApi* | [**getUserAccountPositions**](docs/AccountInformationApi.md#getUserAccountPositions) | **GET** /accounts/{accountId}/positions | List account positions
*AccountInformationApi* | [**getUserHoldings**](docs/AccountInformationApi.md#getUserHoldings) | **GET** /accounts/{accountId}/holdings | List account holdings
*AccountInformationApi* | [**listUserAccounts**](docs/AccountInformationApi.md#listUserAccounts) | **GET** /accounts | List accounts
*AccountInformationApi* | [**updateUserAccount**](docs/AccountInformationApi.md#updateUserAccount) | **PUT** /accounts/{accountId} | Update details of an investment account
*ApiStatusApi* | [**check**](docs/ApiStatusApi.md#check) | **GET** / | Get API Status
*AuthenticationApi* | [**deleteSnapTradeUser**](docs/AuthenticationApi.md#deleteSnapTradeUser) | **DELETE** /snapTrade/deleteUser | Delete SnapTrade user
*AuthenticationApi* | [**listSnapTradeUsers**](docs/AuthenticationApi.md#listSnapTradeUsers) | **GET** /snapTrade/listUsers | List SnapTrade users
*AuthenticationApi* | [**loginSnapTradeUser**](docs/AuthenticationApi.md#loginSnapTradeUser) | **POST** /snapTrade/login | Login user &amp; generate connection link
*AuthenticationApi* | [**registerSnapTradeUser**](docs/AuthenticationApi.md#registerSnapTradeUser) | **POST** /snapTrade/registerUser | Create SnapTrade user
*AuthenticationApi* | [**resetSnapTradeUserSecret**](docs/AuthenticationApi.md#resetSnapTradeUserSecret) | **POST** /snapTrade/resetUserSecret | Obtain a new user secret for a user
*ConnectionsApi* | [**detailBrokerageAuthorization**](docs/ConnectionsApi.md#detailBrokerageAuthorization) | **GET** /authorizations/{authorizationId} | Get brokerage authorization details
*ConnectionsApi* | [**disableBrokerageAuthorization**](docs/ConnectionsApi.md#disableBrokerageAuthorization) | **POST** /authorizations/{authorizationId}/disable | Manually disable a connection for testing
*ConnectionsApi* | [**listBrokerageAuthorizations**](docs/ConnectionsApi.md#listBrokerageAuthorizations) | **GET** /authorizations | List all brokerage authorizations for the User
*ConnectionsApi* | [**refreshBrokerageAuthorization**](docs/ConnectionsApi.md#refreshBrokerageAuthorization) | **POST** /authorizations/{authorizationId}/refresh | Refresh holdings for a connection
*ConnectionsApi* | [**removeBrokerageAuthorization**](docs/ConnectionsApi.md#removeBrokerageAuthorization) | **DELETE** /authorizations/{authorizationId} | Delete brokerage authorization
*ConnectionsApi* | [**sessionEvents**](docs/ConnectionsApi.md#sessionEvents) | **GET** /sessionEvents | Get all session events for a user
*OptionsApi* | [**getOptionStrategy**](docs/OptionsApi.md#getOptionStrategy) | **POST** /accounts/{accountId}/optionStrategy | Create options strategy
*OptionsApi* | [**getOptionsChain**](docs/OptionsApi.md#getOptionsChain) | **GET** /accounts/{accountId}/optionsChain | Get the options chain for a symbol
*OptionsApi* | [**getOptionsStrategyQuote**](docs/OptionsApi.md#getOptionsStrategyQuote) | **GET** /accounts/{accountId}/optionStrategy/{optionStrategyId} | Get options strategy quotes
*OptionsApi* | [**listOptionHoldings**](docs/OptionsApi.md#listOptionHoldings) | **GET** /accounts/{accountId}/options | Get account option holdings
*OptionsApi* | [**placeOptionStrategy**](docs/OptionsApi.md#placeOptionStrategy) | **POST** /accounts/{accountId}/optionStrategy/{optionStrategyId}/execute | Place an option strategy order
*ReferenceDataApi* | [**getCurrencyExchangeRatePair**](docs/ReferenceDataApi.md#getCurrencyExchangeRatePair) | **GET** /currencies/rates/{currencyPair} | Get exchange rate of a currency pair
*ReferenceDataApi* | [**getPartnerInfo**](docs/ReferenceDataApi.md#getPartnerInfo) | **GET** /snapTrade/partners | Get metadata related to Snaptrade partner
*ReferenceDataApi* | [**getSecurityTypes**](docs/ReferenceDataApi.md#getSecurityTypes) | **GET** /securityTypes | List of all security types
*ReferenceDataApi* | [**getStockExchanges**](docs/ReferenceDataApi.md#getStockExchanges) | **GET** /exchanges | Get exchanges
*ReferenceDataApi* | [**getSymbols**](docs/ReferenceDataApi.md#getSymbols) | **POST** /symbols | Search for symbols
*ReferenceDataApi* | [**getSymbolsByTicker**](docs/ReferenceDataApi.md#getSymbolsByTicker) | **GET** /symbols/{query} | Get details of a symbol
*ReferenceDataApi* | [**listAllBrokerageAuthorizationType**](docs/ReferenceDataApi.md#listAllBrokerageAuthorizationType) | **GET** /brokerageAuthorizationTypes | Get all brokerage authorization types
*ReferenceDataApi* | [**listAllBrokerages**](docs/ReferenceDataApi.md#listAllBrokerages) | **GET** /brokerages | Get brokerages
*ReferenceDataApi* | [**listAllCurrencies**](docs/ReferenceDataApi.md#listAllCurrencies) | **GET** /currencies | Get currencies
*ReferenceDataApi* | [**listAllCurrenciesRates**](docs/ReferenceDataApi.md#listAllCurrenciesRates) | **GET** /currencies/rates | Get currency exchange rates
*ReferenceDataApi* | [**symbolSearchUserAccount**](docs/ReferenceDataApi.md#symbolSearchUserAccount) | **POST** /accounts/{accountId}/symbols | Search for symbols available in an account
*TradingApi* | [**cancelUserAccountOrder**](docs/TradingApi.md#cancelUserAccountOrder) | **POST** /accounts/{accountId}/orders/cancel | Cancel open order in account
*TradingApi* | [**getOrderImpact**](docs/TradingApi.md#getOrderImpact) | **POST** /trade/impact | Check the impact of a trade on an account
*TradingApi* | [**getUserAccountQuotes**](docs/TradingApi.md#getUserAccountQuotes) | **GET** /accounts/{accountId}/quotes | Get symbol quotes
*TradingApi* | [**placeForceOrder**](docs/TradingApi.md#placeForceOrder) | **POST** /trade/place | Place a trade with NO validation.
*TradingApi* | [**placeOrder**](docs/TradingApi.md#placeOrder) | **POST** /trade/{tradeId} | Place order
*TransactionsAndReportingApi* | [**getActivities**](docs/TransactionsAndReportingApi.md#getActivities) | **GET** /activities | Get transaction history for a user
*TransactionsAndReportingApi* | [**getReportingCustomRange**](docs/TransactionsAndReportingApi.md#getReportingCustomRange) | **GET** /performance/custom | Get performance information for a specific timeframe


## Documentation for Models

 - [Account](docs/Account.md)
 - [AccountBalance](docs/AccountBalance.md)
 - [AccountBalanceNullable](docs/AccountBalanceNullable.md)
 - [AccountHoldings](docs/AccountHoldings.md)
 - [AccountHoldingsAccount](docs/AccountHoldingsAccount.md)
 - [AccountOrderRecord](docs/AccountOrderRecord.md)
 - [AccountOrderRecordStatus](docs/AccountOrderRecordStatus.md)
 - [AccountSimple](docs/AccountSimple.md)
 - [AccountSyncStatus](docs/AccountSyncStatus.md)
 - [ActionStrict](docs/ActionStrict.md)
 - [Amount](docs/Amount.md)
 - [AmountNullable](docs/AmountNullable.md)
 - [Balance](docs/Balance.md)
 - [Brokerage](docs/Brokerage.md)
 - [BrokerageAuthorization](docs/BrokerageAuthorization.md)
 - [BrokerageAuthorizationDisabledConfirmation](docs/BrokerageAuthorizationDisabledConfirmation.md)
 - [BrokerageAuthorizationRefreshConfirmation](docs/BrokerageAuthorizationRefreshConfirmation.md)
 - [BrokerageAuthorizationTypeReadOnly](docs/BrokerageAuthorizationTypeReadOnly.md)
 - [BrokerageAuthorizationTypeReadOnlyBrokerage](docs/BrokerageAuthorizationTypeReadOnlyBrokerage.md)
 - [BrokerageType](docs/BrokerageType.md)
 - [CashRestriction](docs/CashRestriction.md)
 - [Currency](docs/Currency.md)
 - [CurrencyNullable](docs/CurrencyNullable.md)
 - [DeleteUserResponse](docs/DeleteUserResponse.md)
 - [DividendAtDate](docs/DividendAtDate.md)
 - [EncryptedResponse](docs/EncryptedResponse.md)
 - [EncryptedResponseEncryptedMessageData](docs/EncryptedResponseEncryptedMessageData.md)
 - [Exchange](docs/Exchange.md)
 - [ExchangeRatePairs](docs/ExchangeRatePairs.md)
 - [FigiInstrument](docs/FigiInstrument.md)
 - [FigiInstrumentNullable](docs/FigiInstrumentNullable.md)
 - [HoldingsStatus](docs/HoldingsStatus.md)
 - [LoginRedirectURI](docs/LoginRedirectURI.md)
 - [ManualTrade](docs/ManualTrade.md)
 - [ManualTradeAndImpact](docs/ManualTradeAndImpact.md)
 - [ManualTradeBalance](docs/ManualTradeBalance.md)
 - [ManualTradeForm](docs/ManualTradeForm.md)
 - [ManualTradeSymbol](docs/ManualTradeSymbol.md)
 - [MonthlyDividends](docs/MonthlyDividends.md)
 - [NetContributions](docs/NetContributions.md)
 - [NetDividend](docs/NetDividend.md)
 - [OptionBrokerageSymbol](docs/OptionBrokerageSymbol.md)
 - [OptionChainInner](docs/OptionChainInner.md)
 - [OptionChainInnerChainPerRootInner](docs/OptionChainInnerChainPerRootInner.md)
 - [OptionChainInnerChainPerRootInnerChainPerStrikePriceInner](docs/OptionChainInnerChainPerRootInnerChainPerStrikePriceInner.md)
 - [OptionLeg](docs/OptionLeg.md)
 - [OptionStrategy](docs/OptionStrategy.md)
 - [OptionStrategyLegsInner](docs/OptionStrategyLegsInner.md)
 - [OptionsGetOptionStrategyRequest](docs/OptionsGetOptionStrategyRequest.md)
 - [OptionsPlaceOptionStrategyRequest](docs/OptionsPlaceOptionStrategyRequest.md)
 - [OptionsPosition](docs/OptionsPosition.md)
 - [OptionsSymbol](docs/OptionsSymbol.md)
 - [OrderTypeStrict](docs/OrderTypeStrict.md)
 - [PartnerData](docs/PartnerData.md)
 - [PastValue](docs/PastValue.md)
 - [PerformanceCustom](docs/PerformanceCustom.md)
 - [Position](docs/Position.md)
 - [PositionSymbol](docs/PositionSymbol.md)
 - [SecurityType](docs/SecurityType.md)
 - [SessionEvent](docs/SessionEvent.md)
 - [SnapTradeHoldingsAccount](docs/SnapTradeHoldingsAccount.md)
 - [SnapTradeHoldingsAccountAccountId](docs/SnapTradeHoldingsAccountAccountId.md)
 - [SnapTradeHoldingsTotalValue](docs/SnapTradeHoldingsTotalValue.md)
 - [SnapTradeLoginUserRequestBody](docs/SnapTradeLoginUserRequestBody.md)
 - [SnapTradeRegisterUserRequestBody](docs/SnapTradeRegisterUserRequestBody.md)
 - [Status](docs/Status.md)
 - [StrategyOrderRecord](docs/StrategyOrderRecord.md)
 - [StrategyQuotes](docs/StrategyQuotes.md)
 - [StrategyQuotesGreek](docs/StrategyQuotesGreek.md)
 - [SubPeriodReturnRate](docs/SubPeriodReturnRate.md)
 - [Symbol](docs/Symbol.md)
 - [SymbolQuery](docs/SymbolQuery.md)
 - [SymbolsQuotesInner](docs/SymbolsQuotesInner.md)
 - [TimeInForceStrict](docs/TimeInForceStrict.md)
 - [TradingCancelUserAccountOrderRequest](docs/TradingCancelUserAccountOrderRequest.md)
 - [TransactionsStatus](docs/TransactionsStatus.md)
 - [USExchange](docs/USExchange.md)
 - [UnderlyingSymbol](docs/UnderlyingSymbol.md)
 - [UnderlyingSymbolExchange](docs/UnderlyingSymbolExchange.md)
 - [UnderlyingSymbolType](docs/UnderlyingSymbolType.md)
 - [UniversalActivity](docs/UniversalActivity.md)
 - [UniversalSymbol](docs/UniversalSymbol.md)
 - [UniversalSymbolCurrency](docs/UniversalSymbolCurrency.md)
 - [UniversalSymbolExchange](docs/UniversalSymbolExchange.md)
 - [UserIDandSecret](docs/UserIDandSecret.md)
 - [ValidatedTradeBody](docs/ValidatedTradeBody.md)


## Author
This Java package is automatically generated by [Konfig](https://konfigthis.com)

<div align="center">

[![Visit SnapTrade](./header.png)](https://snaptrade.com)

# [SnapTrade](https://snaptrade.com)<a id="snaptrade"></a>

Connect brokerage accounts to your app for live positions and trading

[![npm](https://img.shields.io/badge/npm-v8.29.5-blue)](https://www.npmjs.com/package/snaptrade-typescript-sdk/v/8.29.5)
[![More Info](https://img.shields.io/badge/More%20Info-Click%20Here-orange)](https://snaptrade.com/)

</div>

## Table of Contents<a id="table-of-contents"></a>

<!-- toc -->

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Reference](#reference)
  * [`snaptrade.accountInformation.getAllUserHoldings`](#snaptradeaccountinformationgetalluserholdings)
  * [`snaptrade.accountInformation.getUserAccountBalance`](#snaptradeaccountinformationgetuseraccountbalance)
  * [`snaptrade.accountInformation.getUserAccountDetails`](#snaptradeaccountinformationgetuseraccountdetails)
  * [`snaptrade.accountInformation.getUserAccountOrders`](#snaptradeaccountinformationgetuseraccountorders)
  * [`snaptrade.accountInformation.getUserAccountPositions`](#snaptradeaccountinformationgetuseraccountpositions)
  * [`snaptrade.accountInformation.getUserHoldings`](#snaptradeaccountinformationgetuserholdings)
  * [`snaptrade.accountInformation.listUserAccounts`](#snaptradeaccountinformationlistuseraccounts)
  * [`snaptrade.accountInformation.updateUserAccount`](#snaptradeaccountinformationupdateuseraccount)
  * [`snaptrade.apiStatus.check`](#snaptradeapistatuscheck)
  * [`snaptrade.authentication.deleteSnapTradeUser`](#snaptradeauthenticationdeletesnaptradeuser)
  * [`snaptrade.authentication.getUserJWT`](#snaptradeauthenticationgetuserjwt)
  * [`snaptrade.authentication.listSnapTradeUsers`](#snaptradeauthenticationlistsnaptradeusers)
  * [`snaptrade.authentication.loginSnapTradeUser`](#snaptradeauthenticationloginsnaptradeuser)
  * [`snaptrade.authentication.registerSnapTradeUser`](#snaptradeauthenticationregistersnaptradeuser)
  * [`snaptrade.connections.detailBrokerageAuthorization`](#snaptradeconnectionsdetailbrokerageauthorization)
  * [`snaptrade.connections.listBrokerageAuthorizations`](#snaptradeconnectionslistbrokerageauthorizations)
  * [`snaptrade.connections.removeBrokerageAuthorization`](#snaptradeconnectionsremovebrokerageauthorization)
  * [`snaptrade.connections.sessionEvents`](#snaptradeconnectionssessionevents)
  * [`snaptrade.errorLogs.listUserErrors`](#snaptradeerrorlogslistusererrors)
  * [`snaptrade.options.getOptionStrategy`](#snaptradeoptionsgetoptionstrategy)
  * [`snaptrade.options.getOptionsChain`](#snaptradeoptionsgetoptionschain)
  * [`snaptrade.options.getOptionsStrategyQuote`](#snaptradeoptionsgetoptionsstrategyquote)
  * [`snaptrade.options.listOptionHoldings`](#snaptradeoptionslistoptionholdings)
  * [`snaptrade.options.placeOptionStrategy`](#snaptradeoptionsplaceoptionstrategy)
  * [`snaptrade.referenceData.getCurrencyExchangeRatePair`](#snaptradereferencedatagetcurrencyexchangeratepair)
  * [`snaptrade.referenceData.getPartnerInfo`](#snaptradereferencedatagetpartnerinfo)
  * [`snaptrade.referenceData.getSecurityTypes`](#snaptradereferencedatagetsecuritytypes)
  * [`snaptrade.referenceData.getStockExchanges`](#snaptradereferencedatagetstockexchanges)
  * [`snaptrade.referenceData.getSymbols`](#snaptradereferencedatagetsymbols)
  * [`snaptrade.referenceData.getSymbolsByTicker`](#snaptradereferencedatagetsymbolsbyticker)
  * [`snaptrade.referenceData.listAllBrokerageAuthorizationType`](#snaptradereferencedatalistallbrokerageauthorizationtype)
  * [`snaptrade.referenceData.listAllBrokerages`](#snaptradereferencedatalistallbrokerages)
  * [`snaptrade.referenceData.listAllCurrencies`](#snaptradereferencedatalistallcurrencies)
  * [`snaptrade.referenceData.listAllCurrenciesRates`](#snaptradereferencedatalistallcurrenciesrates)
  * [`snaptrade.referenceData.symbolSearchUserAccount`](#snaptradereferencedatasymbolsearchuseraccount)
  * [`snaptrade.trading.cancelUserAccountOrder`](#snaptradetradingcanceluseraccountorder)
  * [`snaptrade.trading.getOrderImpact`](#snaptradetradinggetorderimpact)
  * [`snaptrade.trading.getUserAccountQuotes`](#snaptradetradinggetuseraccountquotes)
  * [`snaptrade.trading.placeForceOrder`](#snaptradetradingplaceforceorder)
  * [`snaptrade.trading.placeOCOOrder`](#snaptradetradingplaceocoorder)
  * [`snaptrade.trading.placeOrder`](#snaptradetradingplaceorder)
  * [`snaptrade.transactionsAndReporting.getActivities`](#snaptradetransactionsandreportinggetactivities)
  * [`snaptrade.transactionsAndReporting.getReportingCustomRange`](#snaptradetransactionsandreportinggetreportingcustomrange)

<!-- tocstop -->

## Installation<a id="installation"></a>

<table>
<tr>
<th width="292px"><code>npm</code></th>
<th width="293px"><code>pnpm</code></th>
<th width="292px"><code>yarn</code></th>
</tr>
<tr>
<td>

```bash
npm i snaptrade-typescript-sdk
```

</td>
<td>

```bash
pnpm i snaptrade-typescript-sdk
```

</td>
<td>

```bash
yarn add snaptrade-typescript-sdk
```

</td>
</tr>
</table>

## Getting Started<a id="getting-started"></a>

```typescript
const { Snaptrade } = require("snaptrade-typescript-sdk");

async function main() {
  // 1) Initialize a client with your clientID and consumerKey.
  const snaptrade = new Snaptrade({
    consumerKey: process.env.SNAPTRADE_CONSUMER_KEY,
    clientId: process.env.SNAPTRADE_CLIENT_ID,
  });

  // 2) Check that the client is able to make a request to the API server.
  const status = await snaptrade.apiStatus.check();
  console.log("status:", status.data);

  // 3) Create a new user on SnapTrade
  const userId = uuid();
  const { userSecret } = (
    await snaptrade.authentication.registerSnapTradeUser({
      userId,
    })
  ).data;

  // Note: A user secret is only generated once. It's required to access
  // resources for certain endpoints.
  console.log("userSecret:", userSecret);

  // 4) Get a redirect URI. Users will need this to connect
  const data = (
    await snaptrade.authentication.loginSnapTradeUser({ userId, userSecret })
  ).data;
  if (!("redirectURI" in data)) throw Error("Should have gotten redirect URI");
  console.log("redirectURI:", data.redirectURI);

  // 5) Obtaining account holdings data
  const holdings = (
    await snaptrade.accountInformation.getAllUserHoldings({
      userId,
      userSecret,
    })
  ).data;
  console.log("holdings:", holdings);

  // 6) Deleting a user
  const deleteResponse = (
    await snaptrade.authentication.deleteSnapTradeUser({ userId })
  ).data;
  console.log("deleteResponse:", deleteResponse);
}

// Should be replaced with function to get user ID
function getUserId() {
  var d = new Date().getTime(); //Timestamp
  var d2 =
    (typeof performance !== "undefined" &&
      performance.now &&
      performance.now() * 1000) ||
    0; //Time in microseconds since page-load or 0 if unsupported
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
    var r = Math.random() * 16; //random number between 0 and 16
    if (d > 0) {
      //Use timestamp until depleted
      r = (d + r) % 16 | 0;
      d = Math.floor(d / 16);
    } else {
      //Use microseconds since page-load if supported
      r = (d2 + r) % 16 | 0;
      d2 = Math.floor(d2 / 16);
    }
    return (c === "x" ? r : (r & 0x3) | 0x8).toString(16);
  });
}

main();
```

## Reference<a id="reference"></a>


### `snaptrade.accountInformation.getAllUserHoldings`<a id="snaptradeaccountinformationgetalluserholdings"></a>

List all accounts for the user, plus balances, positions, and orders for each account.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getAllUserHoldingsResponse =
  await snaptrade.accountInformation.getAllUserHoldings({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    brokerageAuthorizations: "917c8734-8470-4a3e-a18f-57c3f2ee6631",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### brokerageAuthorizations: `string`<a id="brokerageauthorizations-string"></a>

Optional. Comma seperated list of authorization IDs (only use if filtering is needed on one or more authorizations).

#### 🔄 Return<a id="🔄-return"></a>

[AccountHoldings](./models/account-holdings.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/holdings` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.accountInformation.getUserAccountBalance`<a id="snaptradeaccountinformationgetuseraccountbalance"></a>

A list of account balances for the specified account (one per currency that the account holds).

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getUserAccountBalanceResponse =
  await snaptrade.accountInformation.getUserAccountBalance({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    accountId: "accountId_example",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to get balances.

#### 🔄 Return<a id="🔄-return"></a>

[Balance](./models/balance.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/balances` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.accountInformation.getUserAccountDetails`<a id="snaptradeaccountinformationgetuseraccountdetails"></a>

Return details of a specific investment account

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getUserAccountDetailsResponse =
  await snaptrade.accountInformation.getUserAccountDetails({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    accountId: "accountId_example",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to get detail of.

#### 🔄 Return<a id="🔄-return"></a>

[Account](./models/account.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.accountInformation.getUserAccountOrders`<a id="snaptradeaccountinformationgetuseraccountorders"></a>

Fetch all recent orders from a user\'s account.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getUserAccountOrdersResponse =
  await snaptrade.accountInformation.getUserAccountOrders({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    state: "all",
    days: 30,
    accountId: "accountId_example",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to get orders.

##### state: `'all' | 'open' | 'executed'`<a id="state-all--open--executed"></a>

defaults value is set to \"all\"

##### days: `number`<a id="days-number"></a>

Number of days in the past to fetch the most recent orders. Defaults to the last 90 days if no value is passed in.

#### 🔄 Return<a id="🔄-return"></a>

[AccountOrderRecord](./models/account-order-record.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/orders` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.accountInformation.getUserAccountPositions`<a id="snaptradeaccountinformationgetuseraccountpositions"></a>

List account positions

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getUserAccountPositionsResponse =
  await snaptrade.accountInformation.getUserAccountPositions({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    accountId: "accountId_example",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to get positions.

#### 🔄 Return<a id="🔄-return"></a>

[Position](./models/position.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/positions` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.accountInformation.getUserHoldings`<a id="snaptradeaccountinformationgetuserholdings"></a>

List balances, positions and orders for the specified account

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getUserHoldingsResponse =
  await snaptrade.accountInformation.getUserHoldings({
    accountId: "917c8734-8470-4a3e-a18f-57c3f2ee6631",
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to fetch holdings for.

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[AccountHoldingsAccount](./models/account-holdings-account.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/holdings` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.accountInformation.listUserAccounts`<a id="snaptradeaccountinformationlistuseraccounts"></a>

List accounts

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const listUserAccountsResponse =
  await snaptrade.accountInformation.listUserAccounts({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[Account](./models/account.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.accountInformation.updateUserAccount`<a id="snaptradeaccountinformationupdateuseraccount"></a>

Update details of an investment account

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const updateUserAccountResponse =
  await snaptrade.accountInformation.updateUserAccount({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    accountId: "accountId_example",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to update.

#### 🔄 Return<a id="🔄-return"></a>

[Account](./models/account.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}` `PUT`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.apiStatus.check`<a id="snaptradeapistatuscheck"></a>

Check whether the API is operational and verify timestamps.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const checkResponse = await snaptrade.apiStatus.check();
```

#### 🔄 Return<a id="🔄-return"></a>

[Status](./models/status.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.authentication.deleteSnapTradeUser`<a id="snaptradeauthenticationdeletesnaptradeuser"></a>

Deletes a user you\'ve registered over the SnapTrade API, and any data associated with them or their investment accounts.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const deleteSnapTradeUserResponse =
  await snaptrade.authentication.deleteSnapTradeUser({
    userId: "John.doe@snaptrade.com",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[DeleteUserResponse](./models/delete-user-response.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/snapTrade/deleteUser` `DELETE`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.authentication.getUserJWT`<a id="snaptradeauthenticationgetuserjwt"></a>

Generate encrypted JWT token

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getUserJWTResponse = await snaptrade.authentication.getUserJWT({
  userId: "John.doe@snaptrade.com",
  userSecret: "USERSECRET123",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[EncryptedResponse](./models/encrypted-response.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/snapTrade/encryptedJWT` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.authentication.listSnapTradeUsers`<a id="snaptradeauthenticationlistsnaptradeusers"></a>

Returns a list of users you\'ve registered over the SnapTrade API.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const listSnapTradeUsersResponse =
  await snaptrade.authentication.listSnapTradeUsers();
```

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/snapTrade/listUsers` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.authentication.loginSnapTradeUser`<a id="snaptradeauthenticationloginsnaptradeuser"></a>

Logs in a SnapTrade user and returns an authenticated connection portal URL for them to use to connect a brokerage account.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const loginSnapTradeUserResponse =
  await snaptrade.authentication.loginSnapTradeUser({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    broker: "ALPACA",
    immediateRedirect: true,
    customRedirect: "https://snaptrade.com",
    reconnect: "8b5f262d-4bb9-365d-888a-202bd3b15fa1",
    connectionType: "read",
    connectionPortalVersion: "v2",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### broker: `string`<a id="broker-string"></a>

Slug of the brokerage to connect the user to

##### immediateRedirect: `boolean`<a id="immediateredirect-boolean"></a>

When set to True, user will be redirected back to the partner\\\'s site instead of the connection portal

##### customRedirect: `string`<a id="customredirect-string"></a>

URL to redirect the user to after the user connects their brokerage account

##### reconnect: `string`<a id="reconnect-string"></a>

The UUID of the brokerage connection to be reconnected. This parameter should be left empty unless you are reconnecting a disabled connection. See ‘Reconnecting Accounts’ for more information.

##### connectionType: `string`<a id="connectiontype-string"></a>

Sets whether the connection should be read or trade

##### connectionPortalVersion: `string`<a id="connectionportalversion-string"></a>

Sets the version of the connection portal to render, with a default to \\\'v2\\\'

#### 🔄 Return<a id="🔄-return"></a>

[AuthenticationLoginSnapTradeUser200Response](./models/authentication-login-snap-trade-user200-response.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/snapTrade/login` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.authentication.registerSnapTradeUser`<a id="snaptradeauthenticationregistersnaptradeuser"></a>

Create SnapTrade user

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const registerSnapTradeUserResponse =
  await snaptrade.authentication.registerSnapTradeUser({
    userId: "snaptrade-user-123",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

SnapTrade User ID. Provided by SnapTrade Partner. Can be any string, as long as it\\\'s unique to a user

#### 🔄 Return<a id="🔄-return"></a>

[UserIDandSecret](./models/user-idand-secret.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/snapTrade/registerUser` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.connections.detailBrokerageAuthorization`<a id="snaptradeconnectionsdetailbrokerageauthorization"></a>

Get brokerage authorization details

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const detailBrokerageAuthorizationResponse =
  await snaptrade.connections.detailBrokerageAuthorization({
    authorizationId: "2bcd7cc3-e922-4976-bce1-9858296801c3",
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### authorizationId: `string`<a id="authorizationid-string"></a>

The ID of a brokerage authorization object.

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[BrokerageAuthorization](./models/brokerage-authorization.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/authorizations/{authorizationId}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.connections.listBrokerageAuthorizations`<a id="snaptradeconnectionslistbrokerageauthorizations"></a>

List all brokerage authorizations for the user

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const listBrokerageAuthorizationsResponse =
  await snaptrade.connections.listBrokerageAuthorizations({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[BrokerageAuthorization](./models/brokerage-authorization.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/authorizations` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.connections.removeBrokerageAuthorization`<a id="snaptradeconnectionsremovebrokerageauthorization"></a>

Delete brokerage authorization

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const removeBrokerageAuthorizationResponse =
  await snaptrade.connections.removeBrokerageAuthorization({
    authorizationId: "2bcd7cc3-e922-4976-bce1-9858296801c3",
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### authorizationId: `string`<a id="authorizationid-string"></a>

The ID of the Authorization to delete.

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/authorizations/{authorizationId}` `DELETE`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.connections.sessionEvents`<a id="snaptradeconnectionssessionevents"></a>

List all session events for the partner

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const sessionEventsResponse = await snaptrade.connections.sessionEvents({
  partnerClientId: "SNAPTRADETEST",
  userId:
    "917c8734-8470-4a3e-a18f-57c3f2ee6631,65e839a3-9103-4cfb-9b72-2071ef80c5f2",
  sessionId:
    "917c8734-8470-4a3e-a18f-57c3f2ee6631,65e839a3-9103-4cfb-9b72-2071ef80c5f2",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### partnerClientId: `string`<a id="partnerclientid-string"></a>

##### userId: `string`<a id="userid-string"></a>

Optional comma seperated list of user IDs used to filter the request on specific users

##### sessionId: `string`<a id="sessionid-string"></a>

Optional comma seperated list of session IDs used to filter the request on specific users

#### 🔄 Return<a id="🔄-return"></a>

[ConnectionsSessionEvents200ResponseInner](./models/connections-session-events200-response-inner.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/sessionEvents` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.errorLogs.listUserErrors`<a id="snaptradeerrorlogslistusererrors"></a>

Retrieve error logs on behalf of your SnapTrade users

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const listUserErrorsResponse = await snaptrade.errorLogs.listUserErrors({
  userId: "John.doe@snaptrade.com",
  userSecret: "USERSECRET123",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[UserErrorLog](./models/user-error-log.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/snapTrade/listUserErrors` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.options.getOptionStrategy`<a id="snaptradeoptionsgetoptionstrategy"></a>

Creates an option strategy object that will be used to place an option strategy order

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getOptionStrategyResponse = await snaptrade.options.getOptionStrategy({
  userId: "John.doe@snaptrade.com",
  userSecret: "USERSECRET123",
  accountId: "accountId_example",
  underlying_symbol_id: "2bcd7cc3-e922-4976-bce1-9858296801c3",
  legs: [
    {
      action: "BUY_TO_OPEN",
      option_symbol_id: "SPY220819P00200000",
      quantity: 1,
    },
  ],
  strategy_type: "CUSTOM",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### underlying_symbol_id: `string`<a id="underlying_symbol_id-string"></a>

##### legs: [`OptionLeg`](./models/option-leg.ts)[]<a id="legs-optionlegmodelsoption-legts"></a>

##### strategy_type: `string`<a id="strategy_type-string"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to create the option strategy object in.

#### 🔄 Return<a id="🔄-return"></a>

[StrategyQuotes](./models/strategy-quotes.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/optionStrategy` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.options.getOptionsChain`<a id="snaptradeoptionsgetoptionschain"></a>

Get the options chain

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getOptionsChainResponse = await snaptrade.options.getOptionsChain({
  userId: "John.doe@snaptrade.com",
  userSecret: "USERSECRET123",
  accountId: "accountId_example",
  symbol: "symbol_example",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to get the options chain from.

##### symbol: `string`<a id="symbol-string"></a>

Universal symbol ID if symbol

#### 🔄 Return<a id="🔄-return"></a>

[OptionChainInner](./models/option-chain-inner.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/optionsChain` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.options.getOptionsStrategyQuote`<a id="snaptradeoptionsgetoptionsstrategyquote"></a>

Get latest market data of option strategy

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getOptionsStrategyQuoteResponse =
  await snaptrade.options.getOptionsStrategyQuote({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    accountId: "accountId_example",
    optionStrategyId: "2bcd7cc3-e922-4976-bce1-9858296801c3",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account the strategy will be placed in.

##### optionStrategyId: `string`<a id="optionstrategyid-string"></a>

Option strategy id obtained from response when creating option strategy object

#### 🔄 Return<a id="🔄-return"></a>

[StrategyQuotes](./models/strategy-quotes.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/optionStrategy/{optionStrategyId}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.options.listOptionHoldings`<a id="snaptradeoptionslistoptionholdings"></a>

Get the options holdings in the account

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const listOptionHoldingsResponse = await snaptrade.options.listOptionHoldings({
  userId: "John.doe@snaptrade.com",
  userSecret: "USERSECRET123",
  accountId: "accountId_example",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to fetch options holdings for.

#### 🔄 Return<a id="🔄-return"></a>

[OptionsPosition](./models/options-position.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/options` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.options.placeOptionStrategy`<a id="snaptradeoptionsplaceoptionstrategy"></a>

Place an option strategy order on the brokerage

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const placeOptionStrategyResponse = await snaptrade.options.placeOptionStrategy(
  {
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    accountId: "2bcd7cc3-e922-4976-bce1-9858296801c3",
    optionStrategyId: "2bcd7cc3-e922-4976-bce1-9858296801c3",
    order_type: "Limit",
    time_in_force: "Day",
    price: 31.33,
  }
);
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### order_type:<a id="order_type"></a>

##### time_in_force:<a id="time_in_force"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to execute the strategy in.

##### optionStrategyId: `string`<a id="optionstrategyid-string"></a>

Option strategy id obtained from response when creating option strategy object

##### price: `number`<a id="price-number"></a>

Trade Price if limit or stop limit order

#### 🔄 Return<a id="🔄-return"></a>

[StrategyOrderRecord](./models/strategy-order-record.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/optionStrategy/{optionStrategyId}/execute` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.getCurrencyExchangeRatePair`<a id="snaptradereferencedatagetcurrencyexchangeratepair"></a>

Return the exchange rate of a currency pair

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getCurrencyExchangeRatePairResponse =
  await snaptrade.referenceData.getCurrencyExchangeRatePair({
    currencyPair: "currencyPair_example",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### currencyPair: `string`<a id="currencypair-string"></a>

A currency pair based on currency code for example, {CAD-USD}

#### 🔄 Return<a id="🔄-return"></a>

[ExchangeRatePairs](./models/exchange-rate-pairs.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/currencies/rates/{currencyPair}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.getPartnerInfo`<a id="snaptradereferencedatagetpartnerinfo"></a>

Get metadata related to Snaptrade partner

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getPartnerInfoResponse = await snaptrade.referenceData.getPartnerInfo();
```

#### 🔄 Return<a id="🔄-return"></a>

[PartnerData](./models/partner-data.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/snapTrade/partners` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.getSecurityTypes`<a id="snaptradereferencedatagetsecuritytypes"></a>

List security types available on SnapTrade.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getSecurityTypesResponse =
  await snaptrade.referenceData.getSecurityTypes();
```

#### 🔄 Return<a id="🔄-return"></a>

[SecurityType](./models/security-type.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/securityTypes` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.getStockExchanges`<a id="snaptradereferencedatagetstockexchanges"></a>

List exchanges

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getStockExchangesResponse =
  await snaptrade.referenceData.getStockExchanges();
```

#### 🔄 Return<a id="🔄-return"></a>

[Exchange](./models/exchange.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/exchanges` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.getSymbols`<a id="snaptradereferencedatagetsymbols"></a>

Search for symbols

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getSymbolsResponse = await snaptrade.referenceData.getSymbols({
  substring: "apple",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### substring: `string`<a id="substring-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[UniversalSymbol](./models/universal-symbol.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/symbols` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.getSymbolsByTicker`<a id="snaptradereferencedatagetsymbolsbyticker"></a>

Get details of a symbol by the ticker

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getSymbolsByTickerResponse =
  await snaptrade.referenceData.getSymbolsByTicker({
    ticker: "ticker_example",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### ticker: `string`<a id="ticker-string"></a>

The ticker of the UniversalSymbol to get.

##### symbolId: `string`<a id="symbolid-string"></a>

OPTIONAL IN PATH Can be used instead of the ticker ; The ID of the UniversalSymbol to get.

#### 🔄 Return<a id="🔄-return"></a>

[UniversalSymbol](./models/universal-symbol.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/symbols/{ticker}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.listAllBrokerageAuthorizationType`<a id="snaptradereferencedatalistallbrokerageauthorizationtype"></a>

List of all brokerage authorization types

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const listAllBrokerageAuthorizationTypeResponse =
  await snaptrade.referenceData.listAllBrokerageAuthorizationType({
    brokerage: "QUESTRADE,ALPACA",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### brokerage: `string`<a id="brokerage-string"></a>

Comma separated value of brokerage slugs

#### 🔄 Return<a id="🔄-return"></a>

[BrokerageAuthorizationTypeReadOnly](./models/brokerage-authorization-type-read-only.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/brokerageAuthorizationTypes` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.listAllBrokerages`<a id="snaptradereferencedatalistallbrokerages"></a>

List brokerages

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const listAllBrokeragesResponse =
  await snaptrade.referenceData.listAllBrokerages();
```

#### 🔄 Return<a id="🔄-return"></a>

[Brokerage](./models/brokerage.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/brokerages` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.listAllCurrencies`<a id="snaptradereferencedatalistallcurrencies"></a>

List currencies

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const listAllCurrenciesResponse =
  await snaptrade.referenceData.listAllCurrencies();
```

#### 🔄 Return<a id="🔄-return"></a>

[Currency](./models/currency.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/currencies` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.listAllCurrenciesRates`<a id="snaptradereferencedatalistallcurrenciesrates"></a>

List currency exchange rates

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const listAllCurrenciesRatesResponse =
  await snaptrade.referenceData.listAllCurrenciesRates();
```

#### 🔄 Return<a id="🔄-return"></a>

[ExchangeRatePairs](./models/exchange-rate-pairs.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/currencies/rates` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.referenceData.symbolSearchUserAccount`<a id="snaptradereferencedatasymbolsearchuseraccount"></a>

Search for symbols available in an account

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const symbolSearchUserAccountResponse =
  await snaptrade.referenceData.symbolSearchUserAccount({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    accountId: "accountId_example",
    substring: "apple",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to search for symbols within.

##### substring: `string`<a id="substring-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[UniversalSymbol](./models/universal-symbol.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/symbols` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.trading.cancelUserAccountOrder`<a id="snaptradetradingcanceluseraccountorder"></a>

Cancel open order in account

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const cancelUserAccountOrderResponse =
  await snaptrade.trading.cancelUserAccountOrder({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    accountId: "accountId_example",
    brokerage_order_id: "2bcd7cc3-e922-4976-bce1-9858296801c3",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to cancel the order in.

##### brokerage_order_id: `string`<a id="brokerage_order_id-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[AccountOrderRecord](./models/account-order-record.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/orders/cancel` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.trading.getOrderImpact`<a id="snaptradetradinggetorderimpact"></a>

Check impact of trades on account.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getOrderImpactResponse = await snaptrade.trading.getOrderImpact({
  userId: "John.doe@snaptrade.com",
  userSecret: "USERSECRET123",
  account_id: "2bcd7cc3-e922-4976-bce1-9858296801c3",
  action: "BUY",
  order_type: "Limit",
  price: 31.33,
  stop: 31.33,
  time_in_force: "Day",
  universal_symbol_id: "2bcd7cc3-e922-4976-bce1-9858296801c3",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### account_id: `string`<a id="account_id-string"></a>

##### action:<a id="action"></a>

##### order_type:<a id="order_type"></a>

##### price: `number`<a id="price-number"></a>

Trade Price if limit or stop limit order

##### stop: `number`<a id="stop-number"></a>

Stop Price. If stop loss or stop limit order, the price to trigger the stop

##### time_in_force:<a id="time_in_force"></a>

##### units: `number`<a id="units-number"></a>

Trade Units

##### universal_symbol_id: `string`<a id="universal_symbol_id-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[ManualTradeAndImpact](./models/manual-trade-and-impact.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/trade/impact` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.trading.getUserAccountQuotes`<a id="snaptradetradinggetuseraccountquotes"></a>

Get symbol quotes

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getUserAccountQuotesResponse =
  await snaptrade.trading.getUserAccountQuotes({
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
    symbols: "symbols_example",
    accountId: "accountId_example",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### symbols: `string`<a id="symbols-string"></a>

List of universal_symbol_id or tickers to get quotes for.

##### accountId: `string`<a id="accountid-string"></a>

The ID of the account to get quotes.

##### useTicker: `boolean`<a id="useticker-boolean"></a>

Should be set to True if providing tickers.

#### 🔄 Return<a id="🔄-return"></a>

[SymbolsQuotesInner](./models/symbols-quotes-inner.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/accounts/{accountId}/quotes` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.trading.placeForceOrder`<a id="snaptradetradingplaceforceorder"></a>

Place a trade with NO validation.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const placeForceOrderResponse = await snaptrade.trading.placeForceOrder({
  userId: "John.doe@snaptrade.com",
  userSecret: "USERSECRET123",
  account_id: "2bcd7cc3-e922-4976-bce1-9858296801c3",
  action: "BUY",
  order_type: "Limit",
  price: 31.33,
  stop: 31.33,
  time_in_force: "Day",
  universal_symbol_id: "2bcd7cc3-e922-4976-bce1-9858296801c3",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### account_id: `string`<a id="account_id-string"></a>

##### action:<a id="action"></a>

##### order_type:<a id="order_type"></a>

##### price: `number`<a id="price-number"></a>

Trade Price if limit or stop limit order

##### stop: `number`<a id="stop-number"></a>

Stop Price. If stop loss or stop limit order, the price to trigger the stop

##### time_in_force:<a id="time_in_force"></a>

##### units: `number`<a id="units-number"></a>

Trade Units

##### universal_symbol_id: `string`<a id="universal_symbol_id-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[AccountOrderRecord](./models/account-order-record.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/trade/place` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.trading.placeOCOOrder`<a id="snaptradetradingplaceocoorder"></a>

Place a OCO (One Cancels Other) order

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const placeOCOOrderResponse = await snaptrade.trading.placeOCOOrder({
  userId: "John.doe@snaptrade.com",
  userSecret: "USERSECRET123",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### first_trade_id:<a id="first_trade_id"></a>

The ID of first trade object obtained from trade/impact endpoint

##### second_trade_id:<a id="second_trade_id"></a>

The ID of second trade object obtained from trade/impact endpoint

#### 🔄 Return<a id="🔄-return"></a>

[AccountOrderRecord](./models/account-order-record.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/trade/oco` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.trading.placeOrder`<a id="snaptradetradingplaceorder"></a>

Place order

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const placeOrderResponse = await snaptrade.trading.placeOrder({
  tradeId: "tradeId_example",
  userId: "John.doe@snaptrade.com",
  userSecret: "USERSECRET123",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### tradeId: `string`<a id="tradeid-string"></a>

The ID of trade object obtained from trade/impact endpoint

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[AccountOrderRecord](./models/account-order-record.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/trade/{tradeId}` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.transactionsAndReporting.getActivities`<a id="snaptradetransactionsandreportinggetactivities"></a>

Returns activities (transactions) for a user. Specifying start and end date is highly recommended for better performance

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getActivitiesResponse =
  await snaptrade.transactionsAndReporting.getActivities({
    startDate: "2022-01-24",
    endDate: "2022-01-24",
    accounts:
      "917c8734-8470-4a3e-a18f-57c3f2ee6631,65e839a3-9103-4cfb-9b72-2071ef80c5f2",
    brokerageAuthorizations:
      "917c8734-8470-4a3e-a18f-57c3f2ee6631,65e839a3-9103-4cfb-9b72-2071ef80c5f2",
    type: "DIVIDEND",
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### startDate: `string | Date`<a id="startdate-string--date"></a>

##### endDate: `string | Date`<a id="enddate-string--date"></a>

##### accounts: `string`<a id="accounts-string"></a>

Optional comma seperated list of account IDs used to filter the request on specific accounts

##### brokerageAuthorizations: `string`<a id="brokerageauthorizations-string"></a>

Optional comma seperated list of brokerage authorization IDs used to filter the request on only accounts that belong to those authorizations

##### type: `string`<a id="type-string"></a>

Optional comma seperated list of types to filter activities by. This is not an exhaustive list, if we fail to match to these types, we will return the raw description from the brokerage. Potential values include - DIVIDEND - BUY - SELL - CONTRIBUTION - WITHDRAWAL - EXTERNAL_ASSET_TRANSFER_IN - EXTERNAL_ASSET_TRANSFER_OUT - INTERNAL_CASH_TRANSFER_IN - INTERNAL_CASH_TRANSFER_OUT - INTERNAL_ASSET_TRANSFER_IN - INTERNAL_ASSET_TRANSFER_OUT - INTEREST - REBATE - GOV_GRANT - TAX - FEE - REI - FXT

#### 🔄 Return<a id="🔄-return"></a>

[UniversalActivity](./models/universal-activity.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/activities` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `snaptrade.transactionsAndReporting.getReportingCustomRange`<a id="snaptradetransactionsandreportinggetreportingcustomrange"></a>

Returns performance information (contributions, dividends, rate of return, etc) for a specific timeframe. Please note that Total Equity Timeframe and Rate of Returns are experimental features. Please contact support@snaptrade.com if you notice any inconsistencies.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getReportingCustomRangeResponse =
  await snaptrade.transactionsAndReporting.getReportingCustomRange({
    startDate: "2022-01-24",
    endDate: "2022-01-24",
    accounts:
      "917c8734-8470-4a3e-a18f-57c3f2ee6631,65e839a3-9103-4cfb-9b72-2071ef80c5f2",
    detailed: true,
    frequency: "monthly",
    userId: "John.doe@snaptrade.com",
    userSecret: "USERSECRET123",
  });
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### startDate: `string | Date`<a id="startdate-string--date"></a>

##### endDate: `string | Date`<a id="enddate-string--date"></a>

##### userId: `string`<a id="userid-string"></a>

##### userSecret: `string`<a id="usersecret-string"></a>

##### accounts: `string`<a id="accounts-string"></a>

Optional comma seperated list of account IDs used to filter the request on specific accounts

##### detailed: `boolean`<a id="detailed-boolean"></a>

Optional, increases frequency of data points for the total value and contribution charts if set to true

##### frequency: `string`<a id="frequency-string"></a>

Optional frequency for the rate of return chart (defaults to monthly). Possible values are daily, weekly, monthly, quarterly, yearly.

#### 🔄 Return<a id="🔄-return"></a>

[PerformanceCustom](./models/performance-custom.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/performance/custom` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


## Author<a id="author"></a>
This TypeScript package is automatically generated by [Konfig](https://konfigthis.com)

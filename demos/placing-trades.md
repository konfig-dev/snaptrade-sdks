# Placing Trades

:::info
This demo is for after a user is created and connected. See [Getting Started](/snaptrade/snaptrade-demos/getting-started) or
[Registering Users](/snaptrade/snaptrade-demos/registering-users) for creating and connecting a user to SnapTrade.
:::

### 1) Create user and Get redirect URL

You can get your `clientId` and `consumerKey` by contacting [api@snaptrade.com](mailto:api@snaptrade.com)

:::form

::input{name=SNAPTRADE_CLIENT_ID label="Client ID" placeholder="YOUR_CLIENT_ID" type="password"}

::input{name=SNAPTRADE_CONSUMER_KEY label="Consumer Key" placeholder="YOUR_CONSUMER_KEY" type="password"}

```python
from snaptrade_client import SnapTrade
import json
import uuid
import os

snaptrade = SnapTrade(
  consumer_key=SNAPTRADE_CONSUMER_KEY,
  client_id=SNAPTRADE_CLIENT_ID,
)

user_id = f"konfigdemo:{uuid.uuid4()}"
register_response = snaptrade.authentication.register_snap_trade_user(
  user_id=user_id
)
user_secret = register_response.body["userSecret"]

redirect_uri = snaptrade.authentication.login_snap_trade_user(
  user_id=user_id, user_secret=user_secret, connection_type="trade"
)
print(json.dumps(redirect_uri.body, indent=2))
```

::button[Create User and Get Redirect URL]

:::

### 2) List accounts

After connecting an account you can see the account IDs by calling `list_user_accounts`.

:::form

```python
accounts = snaptrade.account_information.list_user_accounts(
    user_id=user_id,
    user_secret=user_secret
)
print(json.dumps(accounts.body, indent=2))

for account in accounts.body:
    print("::SAVE[ACCOUNTS]/{}".format(account["id"]))
```

::button[List Accounts]

:::

### 3) Get Order Impact

Check the account has enough cash to place trades.

:::form{skippable}

::enum{name=ACCOUNT_ID label="Account ID" placeholder="ACCOUNT_ID" savedData=ACCOUNTS description="The ID of the account to pull holdings for."}
::input{name=UNIVERSAL_SYMBOL label="Universal Symbol" defaultValue="c15a817e-7171-4940-9ae7-f7b4a95408ee"}
::enum{name=ACTION label="Action" data="BUY,SELL" defaultValue=BUY}
::enum{name=ORDER_TYPE label="Order Type" data="Limit,Market,StopLimit,StopLoss" defaultValue="Limit"}
::enum{name=TIME_IN_FORCE label="Time in Force" data="Day,FOK,GTC" defaultValue="Day"}
::number{name=PRICE label="Price" defaultValue=10, step=0.01 precision=2}
::number{name=UNITS label="Units" defaultValue=1}

```python
result = snaptrade.trading.get_order_impact(
  user_id=user_id,
  user_secret=user_secret,
  action=ACTION,
  account_id=ACCOUNT_ID,
  universal_symbol_id=UNIVERSAL_SYMBOL,
  order_type=ORDER_TYPE,
  time_in_force=TIME_IN_FORCE,
  price=PRICE,
  units=UNITS,
)
print(json.dumps(result.body, indent=2))
```

::button[Get Order Impact]

:::

### 4) Place Force Order

Place the order without checking impact

:::form

::enum{name=ACCOUNT_ID label="Account ID" placeholder="ACCOUNT_ID" savedData=ACCOUNTS description="The ID of the account to pull holdings for."}
::input{name=UNIVERSAL_SYMBOL label="Universal Symbol" defaultValue="c15a817e-7171-4940-9ae7-f7b4a95408ee"}
::enum{name=ACTION label="Action" data="BUY,SELL" defaultValue=BUY}
::enum{name=ORDER_TYPE label="Order Type" data="Limit,Market,StopLimit,StopLoss" defaultValue="Limit"}
::enum{name=TIME_IN_FORCE label="Time in Force" data="Day,FOK,GTC" defaultValue="Day"}
::number{name=PRICE label="Price" defaultValue=10 step=0.01 precision=2}
::number{name=UNITS label="Units" defaultValue=1}

```python
result = snaptrade.trading.place_force_order(
  user_id=user_id,
  user_secret=user_secret,
  action=ACTION,
  account_id=ACCOUNT_ID,
  universal_symbol_id=UNIVERSAL_SYMBOL,
  order_type=ORDER_TYPE,
  time_in_force=TIME_IN_FORCE,
  price=PRICE,
  units=UNITS,
)
print(json.dumps(result.body, indent=2))
```

::button[Place Force Order]

:::

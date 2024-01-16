# coding: utf-8

"""
    SnapTrade

    Connect brokerage accounts to your app for live positions and trading

    The version of the OpenAPI document: 1.0.0
    Contact: api@snaptrade.com
    Created by: https://snaptrade.com/
"""

from datetime import datetime, date
import typing
from enum import Enum
from typing_extensions import TypedDict, Literal, TYPE_CHECKING

from snaptrade_client.type.action import Action
from snaptrade_client.type.order_type import OrderType
from snaptrade_client.type.price import Price
from snaptrade_client.type.stop_price import StopPrice
from snaptrade_client.type.time_in_force import TimeInForce

class RequiredManualTradeForm(TypedDict):
    pass

class OptionalManualTradeForm(TypedDict, total=False):
    account_id: str

    action: Action

    order_type: OrderType

    price: Price

    stop: StopPrice

    time_in_force: TimeInForce

    units: float

    universal_symbol_id: str

    notional_value: float

class ManualTradeForm(RequiredManualTradeForm, OptionalManualTradeForm):
    pass

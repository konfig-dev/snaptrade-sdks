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
from typing_extensions import TypedDict, Literal

from snaptrade_client.type.cash_restriction import CashRestriction

class RequiredSnapTradeHoldingsAccountAccountId(TypedDict):
    pass

class OptionalSnapTradeHoldingsAccountAccountId(TypedDict, total=False):
    id: str

    brokerage_authorization: str

    portfolio_group: str

    name: typing.Optional[str]

    number: str

    institution_name: str

    meta: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]

    cash_restrictions: typing.List[CashRestriction]

    # Time
    created_date: str

class SnapTradeHoldingsAccountAccountId(RequiredSnapTradeHoldingsAccountAccountId, OptionalSnapTradeHoldingsAccountAccountId):
    pass

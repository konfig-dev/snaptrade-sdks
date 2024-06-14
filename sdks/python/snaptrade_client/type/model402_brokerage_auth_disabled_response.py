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


class RequiredModel402BrokerageAuthDisabledResponse(TypedDict):
    pass

class OptionalModel402BrokerageAuthDisabledResponse(TypedDict, total=False):
    detail: typing.Union[bool, date, datetime, dict, float, int, list, str, None]

    code: typing.Union[bool, date, datetime, dict, float, int, list, str, None]

class Model402BrokerageAuthDisabledResponse(RequiredModel402BrokerageAuthDisabledResponse, OptionalModel402BrokerageAuthDisabledResponse):
    pass

# coding: utf-8

"""
    SnapTrade

    Connect brokerage accounts to your app for live positions and trading

    The version of the OpenAPI document: 1.0.0
    Contact: api@snaptrade.com
    Created by: https://snaptrade.com/
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from snaptrade_client import schemas  # noqa: F401


class SyncStatusDate(
    schemas.DateBase,
    schemas.StrBase,
    schemas.NoneBase,
    schemas.Schema,
    schemas.NoneStrMixin
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)

    Date in YYYY-MM-DD format or null
    """


    class MetaOapg:
        format = 'date'


    def __new__(
        cls,
        *args: typing.Union[None, str, date, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
    ) -> 'SyncStatusDate':
        return super().__new__(
            cls,
            *args,
            _configuration=_configuration,
        )

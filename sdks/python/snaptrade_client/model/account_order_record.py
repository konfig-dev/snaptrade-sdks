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


class AccountOrderRecord(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)

    Record of order in brokerageaccount
    """


    class MetaOapg:
        
        class properties:
            brokerage_order_id = schemas.StrSchema
        
            @staticmethod
            def status() -> typing.Type['AccountOrderRecordStatus']:
                return AccountOrderRecordStatus
            symbol = schemas.UUIDSchema
        
            @staticmethod
            def universal_symbol() -> typing.Type['UniversalSymbol']:
                return UniversalSymbol
        
            @staticmethod
            def option_symbol() -> typing.Type['OptionsSymbol']:
                return OptionsSymbol
        
            @staticmethod
            def action() -> typing.Type['Action']:
                return Action
            total_quantity = schemas.NumberSchema
            open_quantity = schemas.NumberSchema
            canceled_quantity = schemas.NumberSchema
            filled_quantity = schemas.NumberSchema
        
            @staticmethod
            def execution_price() -> typing.Type['Price']:
                return Price
        
            @staticmethod
            def limit_price() -> typing.Type['Price']:
                return Price
        
            @staticmethod
            def stop_price() -> typing.Type['Price']:
                return Price
        
            @staticmethod
            def order_type() -> typing.Type['OrderType']:
                return OrderType
        
            @staticmethod
            def time_in_force() -> typing.Type['TimeInForce']:
                return TimeInForce
            time_placed = schemas.StrSchema
            time_updated = schemas.StrSchema
            expiry_date = schemas.StrSchema
            __annotations__ = {
                "brokerage_order_id": brokerage_order_id,
                "status": status,
                "symbol": symbol,
                "universal_symbol": universal_symbol,
                "option_symbol": option_symbol,
                "action": action,
                "total_quantity": total_quantity,
                "open_quantity": open_quantity,
                "canceled_quantity": canceled_quantity,
                "filled_quantity": filled_quantity,
                "execution_price": execution_price,
                "limit_price": limit_price,
                "stop_price": stop_price,
                "order_type": order_type,
                "time_in_force": time_in_force,
                "time_placed": time_placed,
                "time_updated": time_updated,
                "expiry_date": expiry_date,
            }
        additional_properties = schemas.AnyTypeSchema
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["brokerage_order_id"]) -> MetaOapg.properties.brokerage_order_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["status"]) -> 'AccountOrderRecordStatus': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["symbol"]) -> MetaOapg.properties.symbol: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["universal_symbol"]) -> 'UniversalSymbol': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["option_symbol"]) -> 'OptionsSymbol': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["action"]) -> 'Action': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["total_quantity"]) -> MetaOapg.properties.total_quantity: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["open_quantity"]) -> MetaOapg.properties.open_quantity: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["canceled_quantity"]) -> MetaOapg.properties.canceled_quantity: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["filled_quantity"]) -> MetaOapg.properties.filled_quantity: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["execution_price"]) -> 'Price': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["limit_price"]) -> 'Price': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["stop_price"]) -> 'Price': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["order_type"]) -> 'OrderType': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["time_in_force"]) -> 'TimeInForce': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["time_placed"]) -> MetaOapg.properties.time_placed: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["time_updated"]) -> MetaOapg.properties.time_updated: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["expiry_date"]) -> MetaOapg.properties.expiry_date: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> MetaOapg.additional_properties: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["brokerage_order_id"], typing_extensions.Literal["status"], typing_extensions.Literal["symbol"], typing_extensions.Literal["universal_symbol"], typing_extensions.Literal["option_symbol"], typing_extensions.Literal["action"], typing_extensions.Literal["total_quantity"], typing_extensions.Literal["open_quantity"], typing_extensions.Literal["canceled_quantity"], typing_extensions.Literal["filled_quantity"], typing_extensions.Literal["execution_price"], typing_extensions.Literal["limit_price"], typing_extensions.Literal["stop_price"], typing_extensions.Literal["order_type"], typing_extensions.Literal["time_in_force"], typing_extensions.Literal["time_placed"], typing_extensions.Literal["time_updated"], typing_extensions.Literal["expiry_date"], str, ]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["brokerage_order_id"]) -> typing.Union[MetaOapg.properties.brokerage_order_id, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["status"]) -> typing.Union['AccountOrderRecordStatus', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["symbol"]) -> typing.Union[MetaOapg.properties.symbol, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["universal_symbol"]) -> typing.Union['UniversalSymbol', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["option_symbol"]) -> typing.Union['OptionsSymbol', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["action"]) -> typing.Union['Action', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["total_quantity"]) -> typing.Union[MetaOapg.properties.total_quantity, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["open_quantity"]) -> typing.Union[MetaOapg.properties.open_quantity, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["canceled_quantity"]) -> typing.Union[MetaOapg.properties.canceled_quantity, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["filled_quantity"]) -> typing.Union[MetaOapg.properties.filled_quantity, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["execution_price"]) -> typing.Union['Price', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["limit_price"]) -> typing.Union['Price', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["stop_price"]) -> typing.Union['Price', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["order_type"]) -> typing.Union['OrderType', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["time_in_force"]) -> typing.Union['TimeInForce', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["time_placed"]) -> typing.Union[MetaOapg.properties.time_placed, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["time_updated"]) -> typing.Union[MetaOapg.properties.time_updated, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["expiry_date"]) -> typing.Union[MetaOapg.properties.expiry_date, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[MetaOapg.additional_properties, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["brokerage_order_id"], typing_extensions.Literal["status"], typing_extensions.Literal["symbol"], typing_extensions.Literal["universal_symbol"], typing_extensions.Literal["option_symbol"], typing_extensions.Literal["action"], typing_extensions.Literal["total_quantity"], typing_extensions.Literal["open_quantity"], typing_extensions.Literal["canceled_quantity"], typing_extensions.Literal["filled_quantity"], typing_extensions.Literal["execution_price"], typing_extensions.Literal["limit_price"], typing_extensions.Literal["stop_price"], typing_extensions.Literal["order_type"], typing_extensions.Literal["time_in_force"], typing_extensions.Literal["time_placed"], typing_extensions.Literal["time_updated"], typing_extensions.Literal["expiry_date"], str, ]):
        return super().get_item_oapg(name)

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        brokerage_order_id: typing.Union[MetaOapg.properties.brokerage_order_id, str, schemas.Unset] = schemas.unset,
        status: typing.Union['AccountOrderRecordStatus', schemas.Unset] = schemas.unset,
        symbol: typing.Union[MetaOapg.properties.symbol, str, uuid.UUID, schemas.Unset] = schemas.unset,
        universal_symbol: typing.Union['UniversalSymbol', schemas.Unset] = schemas.unset,
        option_symbol: typing.Union['OptionsSymbol', schemas.Unset] = schemas.unset,
        action: typing.Union['Action', schemas.Unset] = schemas.unset,
        total_quantity: typing.Union[MetaOapg.properties.total_quantity, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        open_quantity: typing.Union[MetaOapg.properties.open_quantity, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        canceled_quantity: typing.Union[MetaOapg.properties.canceled_quantity, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        filled_quantity: typing.Union[MetaOapg.properties.filled_quantity, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        execution_price: typing.Union['Price', schemas.Unset] = schemas.unset,
        limit_price: typing.Union['Price', schemas.Unset] = schemas.unset,
        stop_price: typing.Union['Price', schemas.Unset] = schemas.unset,
        order_type: typing.Union['OrderType', schemas.Unset] = schemas.unset,
        time_in_force: typing.Union['TimeInForce', schemas.Unset] = schemas.unset,
        time_placed: typing.Union[MetaOapg.properties.time_placed, str, schemas.Unset] = schemas.unset,
        time_updated: typing.Union[MetaOapg.properties.time_updated, str, schemas.Unset] = schemas.unset,
        expiry_date: typing.Union[MetaOapg.properties.expiry_date, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[MetaOapg.additional_properties, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, bool, None, list, tuple, bytes, io.FileIO, io.BufferedReader, ],
    ) -> 'AccountOrderRecord':
        return super().__new__(
            cls,
            *args,
            brokerage_order_id=brokerage_order_id,
            status=status,
            symbol=symbol,
            universal_symbol=universal_symbol,
            option_symbol=option_symbol,
            action=action,
            total_quantity=total_quantity,
            open_quantity=open_quantity,
            canceled_quantity=canceled_quantity,
            filled_quantity=filled_quantity,
            execution_price=execution_price,
            limit_price=limit_price,
            stop_price=stop_price,
            order_type=order_type,
            time_in_force=time_in_force,
            time_placed=time_placed,
            time_updated=time_updated,
            expiry_date=expiry_date,
            _configuration=_configuration,
            **kwargs,
        )

from snaptrade_client.model.account_order_record_status import AccountOrderRecordStatus
from snaptrade_client.model.action import Action
from snaptrade_client.model.options_symbol import OptionsSymbol
from snaptrade_client.model.order_type import OrderType
from snaptrade_client.model.price import Price
from snaptrade_client.model.time_in_force import TimeInForce
from snaptrade_client.model.universal_symbol import UniversalSymbol

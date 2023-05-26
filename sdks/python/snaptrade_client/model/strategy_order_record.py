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


class StrategyOrderRecord(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)

    Strategy order record
    """


    class MetaOapg:
        
        class properties:
        
            @staticmethod
            def strategy() -> typing.Type['OptionStrategy']:
                return OptionStrategy
            
            
            class status(
                schemas.EnumBase,
                schemas.StrSchema
            ):
            
            
                class MetaOapg:
                    enum_value_to_name = {
                        "PENDING": "PENDING",
                        "ACCEPTED": "ACCEPTED",
                        "FAILED": "FAILED",
                        "REJECTED": "REJECTED",
                        "CANCELED": "CANCELED",
                        "PARTIAL_CANCELED": "PARTIAL_CANCELED",
                        "CANCEL_PENDING": "CANCEL_PENDING",
                        "EXECUTED": "EXECUTED",
                        "PARTIAL": "PARTIAL",
                        "REPLACE_PENDING": "REPLACE_PENDING",
                        "REPLACED": "REPLACED",
                        "STOPPED": "STOPPED",
                        "SUSPENDED": "SUSPENDED",
                        "EXPIRED": "EXPIRED",
                        "QUEUED": "QUEUED",
                        "TRIGGERED": "TRIGGERED",
                        "ACTIVATED": "ACTIVATED",
                        "PENDING_RISK_REVIEW": "PENDING_RISK_REVIEW",
                        "CONTINGENT_ORDER": "CONTINGENT_ORDER",
                    }
                
                @schemas.classproperty
                def PENDING(cls):
                    return cls("PENDING")
                
                @schemas.classproperty
                def ACCEPTED(cls):
                    return cls("ACCEPTED")
                
                @schemas.classproperty
                def FAILED(cls):
                    return cls("FAILED")
                
                @schemas.classproperty
                def REJECTED(cls):
                    return cls("REJECTED")
                
                @schemas.classproperty
                def CANCELED(cls):
                    return cls("CANCELED")
                
                @schemas.classproperty
                def PARTIAL_CANCELED(cls):
                    return cls("PARTIAL_CANCELED")
                
                @schemas.classproperty
                def CANCEL_PENDING(cls):
                    return cls("CANCEL_PENDING")
                
                @schemas.classproperty
                def EXECUTED(cls):
                    return cls("EXECUTED")
                
                @schemas.classproperty
                def PARTIAL(cls):
                    return cls("PARTIAL")
                
                @schemas.classproperty
                def REPLACE_PENDING(cls):
                    return cls("REPLACE_PENDING")
                
                @schemas.classproperty
                def REPLACED(cls):
                    return cls("REPLACED")
                
                @schemas.classproperty
                def STOPPED(cls):
                    return cls("STOPPED")
                
                @schemas.classproperty
                def SUSPENDED(cls):
                    return cls("SUSPENDED")
                
                @schemas.classproperty
                def EXPIRED(cls):
                    return cls("EXPIRED")
                
                @schemas.classproperty
                def QUEUED(cls):
                    return cls("QUEUED")
                
                @schemas.classproperty
                def TRIGGERED(cls):
                    return cls("TRIGGERED")
                
                @schemas.classproperty
                def ACTIVATED(cls):
                    return cls("ACTIVATED")
                
                @schemas.classproperty
                def PENDING_RISK_REVIEW(cls):
                    return cls("PENDING_RISK_REVIEW")
                
                @schemas.classproperty
                def CONTINGENT_ORDER(cls):
                    return cls("CONTINGENT_ORDER")
            filled_quantity = schemas.NumberSchema
            open_quantity = schemas.NumberSchema
            closed_quantity = schemas.NumberSchema
            
            
            class order_type(
                schemas.EnumBase,
                schemas.StrSchema
            ):
            
            
                class MetaOapg:
                    enum_value_to_name = {
                        "Limit": "LIMIT",
                        "Market": "MARKET",
                        "NetDebit": "NET_DEBIT",
                        "NetCredit": "NET_CREDIT",
                    }
                
                @schemas.classproperty
                def LIMIT(cls):
                    return cls("Limit")
                
                @schemas.classproperty
                def MARKET(cls):
                    return cls("Market")
                
                @schemas.classproperty
                def NET_DEBIT(cls):
                    return cls("NetDebit")
                
                @schemas.classproperty
                def NET_CREDIT(cls):
                    return cls("NetCredit")
            
            
            class time_in_force(
                schemas.EnumBase,
                schemas.StrSchema
            ):
            
            
                class MetaOapg:
                    enum_value_to_name = {
                        "DAY": "DAY",
                        "GTC": "GTC",
                    }
                
                @schemas.classproperty
                def DAY(cls):
                    return cls("DAY")
                
                @schemas.classproperty
                def GTC(cls):
                    return cls("GTC")
            limit_price = schemas.NumberSchema
            execution_price = schemas.NumberSchema
            time_placed = schemas.StrSchema
            time_updated = schemas.StrSchema
            __annotations__ = {
                "strategy": strategy,
                "status": status,
                "filled_quantity": filled_quantity,
                "open_quantity": open_quantity,
                "closed_quantity": closed_quantity,
                "order_type": order_type,
                "time_in_force": time_in_force,
                "limit_price": limit_price,
                "execution_price": execution_price,
                "time_placed": time_placed,
                "time_updated": time_updated,
            }
        additional_properties = schemas.AnyTypeSchema
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["strategy"]) -> 'OptionStrategy': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["status"]) -> MetaOapg.properties.status: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["filled_quantity"]) -> MetaOapg.properties.filled_quantity: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["open_quantity"]) -> MetaOapg.properties.open_quantity: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["closed_quantity"]) -> MetaOapg.properties.closed_quantity: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["order_type"]) -> MetaOapg.properties.order_type: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["time_in_force"]) -> MetaOapg.properties.time_in_force: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["limit_price"]) -> MetaOapg.properties.limit_price: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["execution_price"]) -> MetaOapg.properties.execution_price: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["time_placed"]) -> MetaOapg.properties.time_placed: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["time_updated"]) -> MetaOapg.properties.time_updated: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> MetaOapg.additional_properties: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["strategy"], typing_extensions.Literal["status"], typing_extensions.Literal["filled_quantity"], typing_extensions.Literal["open_quantity"], typing_extensions.Literal["closed_quantity"], typing_extensions.Literal["order_type"], typing_extensions.Literal["time_in_force"], typing_extensions.Literal["limit_price"], typing_extensions.Literal["execution_price"], typing_extensions.Literal["time_placed"], typing_extensions.Literal["time_updated"], str, ]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["strategy"]) -> typing.Union['OptionStrategy', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["status"]) -> typing.Union[MetaOapg.properties.status, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["filled_quantity"]) -> typing.Union[MetaOapg.properties.filled_quantity, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["open_quantity"]) -> typing.Union[MetaOapg.properties.open_quantity, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["closed_quantity"]) -> typing.Union[MetaOapg.properties.closed_quantity, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["order_type"]) -> typing.Union[MetaOapg.properties.order_type, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["time_in_force"]) -> typing.Union[MetaOapg.properties.time_in_force, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["limit_price"]) -> typing.Union[MetaOapg.properties.limit_price, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["execution_price"]) -> typing.Union[MetaOapg.properties.execution_price, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["time_placed"]) -> typing.Union[MetaOapg.properties.time_placed, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["time_updated"]) -> typing.Union[MetaOapg.properties.time_updated, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[MetaOapg.additional_properties, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["strategy"], typing_extensions.Literal["status"], typing_extensions.Literal["filled_quantity"], typing_extensions.Literal["open_quantity"], typing_extensions.Literal["closed_quantity"], typing_extensions.Literal["order_type"], typing_extensions.Literal["time_in_force"], typing_extensions.Literal["limit_price"], typing_extensions.Literal["execution_price"], typing_extensions.Literal["time_placed"], typing_extensions.Literal["time_updated"], str, ]):
        return super().get_item_oapg(name)

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        strategy: typing.Union['OptionStrategy', schemas.Unset] = schemas.unset,
        status: typing.Union[MetaOapg.properties.status, str, schemas.Unset] = schemas.unset,
        filled_quantity: typing.Union[MetaOapg.properties.filled_quantity, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        open_quantity: typing.Union[MetaOapg.properties.open_quantity, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        closed_quantity: typing.Union[MetaOapg.properties.closed_quantity, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        order_type: typing.Union[MetaOapg.properties.order_type, str, schemas.Unset] = schemas.unset,
        time_in_force: typing.Union[MetaOapg.properties.time_in_force, str, schemas.Unset] = schemas.unset,
        limit_price: typing.Union[MetaOapg.properties.limit_price, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        execution_price: typing.Union[MetaOapg.properties.execution_price, decimal.Decimal, int, float, schemas.Unset] = schemas.unset,
        time_placed: typing.Union[MetaOapg.properties.time_placed, str, schemas.Unset] = schemas.unset,
        time_updated: typing.Union[MetaOapg.properties.time_updated, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[MetaOapg.additional_properties, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, bool, None, list, tuple, bytes, io.FileIO, io.BufferedReader, ],
    ) -> 'StrategyOrderRecord':
        return super().__new__(
            cls,
            *args,
            strategy=strategy,
            status=status,
            filled_quantity=filled_quantity,
            open_quantity=open_quantity,
            closed_quantity=closed_quantity,
            order_type=order_type,
            time_in_force=time_in_force,
            limit_price=limit_price,
            execution_price=execution_price,
            time_placed=time_placed,
            time_updated=time_updated,
            _configuration=_configuration,
            **kwargs,
        )

from snaptrade_client.model.option_strategy import OptionStrategy

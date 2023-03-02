# coding: utf-8

"""
    SnapTrade

    Connect brokerage accounts to your app for live positions and trading  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: api@snaptrade.com
    Generated by: https://konfigthis.com
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


class USExchange(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by Konfig.
    Ref: https://konfigthis.com

    Do not edit the class manually.

    US Stock Exchange
    """


    class MetaOapg:
        
        class properties:
            id = schemas.UUIDSchema
            code = schemas.StrSchema
            mic_code = schemas.StrSchema
            name = schemas.StrSchema
            timezone = schemas.StrSchema
            start_time = schemas.StrSchema
            close_time = schemas.StrSchema
            
            
            class suffix(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'suffix':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            allows_cryptocurrency_symbols = schemas.BoolSchema
            __annotations__ = {
                "id": id,
                "code": code,
                "mic_code": mic_code,
                "name": name,
                "timezone": timezone,
                "start_time": start_time,
                "close_time": close_time,
                "suffix": suffix,
                "allows_cryptocurrency_symbols": allows_cryptocurrency_symbols,
            }
        additional_properties = schemas.AnyTypeSchema
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["code"]) -> MetaOapg.properties.code: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["mic_code"]) -> MetaOapg.properties.mic_code: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["name"]) -> MetaOapg.properties.name: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["timezone"]) -> MetaOapg.properties.timezone: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["start_time"]) -> MetaOapg.properties.start_time: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["close_time"]) -> MetaOapg.properties.close_time: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["suffix"]) -> MetaOapg.properties.suffix: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["allows_cryptocurrency_symbols"]) -> MetaOapg.properties.allows_cryptocurrency_symbols: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> MetaOapg.additional_properties: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["id"], typing_extensions.Literal["code"], typing_extensions.Literal["mic_code"], typing_extensions.Literal["name"], typing_extensions.Literal["timezone"], typing_extensions.Literal["start_time"], typing_extensions.Literal["close_time"], typing_extensions.Literal["suffix"], typing_extensions.Literal["allows_cryptocurrency_symbols"], str, ]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["id"]) -> typing.Union[MetaOapg.properties.id, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["code"]) -> typing.Union[MetaOapg.properties.code, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["mic_code"]) -> typing.Union[MetaOapg.properties.mic_code, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["name"]) -> typing.Union[MetaOapg.properties.name, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["timezone"]) -> typing.Union[MetaOapg.properties.timezone, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["start_time"]) -> typing.Union[MetaOapg.properties.start_time, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["close_time"]) -> typing.Union[MetaOapg.properties.close_time, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["suffix"]) -> typing.Union[MetaOapg.properties.suffix, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["allows_cryptocurrency_symbols"]) -> typing.Union[MetaOapg.properties.allows_cryptocurrency_symbols, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[MetaOapg.additional_properties, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["id"], typing_extensions.Literal["code"], typing_extensions.Literal["mic_code"], typing_extensions.Literal["name"], typing_extensions.Literal["timezone"], typing_extensions.Literal["start_time"], typing_extensions.Literal["close_time"], typing_extensions.Literal["suffix"], typing_extensions.Literal["allows_cryptocurrency_symbols"], str, ]):
        return super().get_item_oapg(name)

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        id: typing.Union[MetaOapg.properties.id, str, uuid.UUID, schemas.Unset] = schemas.unset,
        code: typing.Union[MetaOapg.properties.code, str, schemas.Unset] = schemas.unset,
        mic_code: typing.Union[MetaOapg.properties.mic_code, str, schemas.Unset] = schemas.unset,
        name: typing.Union[MetaOapg.properties.name, str, schemas.Unset] = schemas.unset,
        timezone: typing.Union[MetaOapg.properties.timezone, str, schemas.Unset] = schemas.unset,
        start_time: typing.Union[MetaOapg.properties.start_time, str, schemas.Unset] = schemas.unset,
        close_time: typing.Union[MetaOapg.properties.close_time, str, schemas.Unset] = schemas.unset,
        suffix: typing.Union[MetaOapg.properties.suffix, None, str, schemas.Unset] = schemas.unset,
        allows_cryptocurrency_symbols: typing.Union[MetaOapg.properties.allows_cryptocurrency_symbols, bool, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[MetaOapg.additional_properties, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, bool, None, list, tuple, bytes, io.FileIO, io.BufferedReader, ],
    ) -> 'USExchange':
        return super().__new__(
            cls,
            *args,
            id=id,
            code=code,
            mic_code=mic_code,
            name=name,
            timezone=timezone,
            start_time=start_time,
            close_time=close_time,
            suffix=suffix,
            allows_cryptocurrency_symbols=allows_cryptocurrency_symbols,
            _configuration=_configuration,
            **kwargs,
        )

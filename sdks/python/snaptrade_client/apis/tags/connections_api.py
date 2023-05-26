# coding: utf-8

"""
    SnapTrade

    Connect brokerage accounts to your app for live positions and trading

    The version of the OpenAPI document: 1.0.0
    Contact: api@snaptrade.com
    Created by: https://snaptrade.com/
"""

from snaptrade_client.paths.authorizations_authorization_id.get import DetailBrokerageAuthorization
from snaptrade_client.paths.authorizations.get import ListBrokerageAuthorizations
from snaptrade_client.paths.authorizations_authorization_id.delete import RemoveBrokerageAuthorization
from snaptrade_client.paths.session_events.get import SessionEvents


class ConnectionsApi(
    DetailBrokerageAuthorization,
    ListBrokerageAuthorizations,
    RemoveBrokerageAuthorization,
    SessionEvents,
):
    """NOTE:
    This class is auto generated by Konfig (https://konfigthis.com)
    """
    pass

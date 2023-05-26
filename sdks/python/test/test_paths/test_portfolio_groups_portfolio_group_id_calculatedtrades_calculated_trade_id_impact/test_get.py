# coding: utf-8

"""
    SnapTrade

    Connect brokerage accounts to your app for live positions and trading

    The version of the OpenAPI document: 1.0.0
    Contact: api@snaptrade.com
    Created by: https://snaptrade.com/
"""

import unittest
from unittest.mock import patch

import urllib3

import snaptrade_client
from snaptrade_client.paths.portfolio_groups_portfolio_group_id_calculatedtrades_calculated_trade_id_impact import get
from snaptrade_client import configuration, schemas, api_client

from .. import ApiTestMixin


class TestPortfolioGroupsPortfolioGroupIdCalculatedtradesCalculatedTradeIdImpact(ApiTestMixin, unittest.TestCase):
    """
    PortfolioGroupsPortfolioGroupIdCalculatedtradesCalculatedTradeIdImpact unit test stubs
        Return the impact of placing a series of trades on the portfolio
    """

    def setUp(self):
        pass

    def tearDown(self):
        pass

    response_status = 200




if __name__ == '__main__':
    unittest.main()

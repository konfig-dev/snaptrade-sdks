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
from snaptrade_client.paths.model_portfolio_model_portfolio_id import get
from snaptrade_client import configuration, schemas, api_client

from .. import ApiTestMixin


class TestModelPortfolioModelPortfolioId(ApiTestMixin, unittest.TestCase):
    """
    ModelPortfolioModelPortfolioId unit test stubs
        Get details of a model portfolio
    """

    def setUp(self):
        pass

    def tearDown(self):
        pass

    response_status = 200




if __name__ == '__main__':
    unittest.main()

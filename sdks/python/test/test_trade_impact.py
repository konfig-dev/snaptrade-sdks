"""
    SnapTrade

    Connect brokerage accounts to your app for live positions and trading  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: api@snaptrade.com
    Generated by: https://konfigthis.com
"""


import sys
import unittest

import snaptrade_client
from snaptrade_client.model.account import Account
from snaptrade_client.model.currency import Currency
globals()['Account'] = Account
globals()['Currency'] = Currency
from snaptrade_client.model.trade_impact import TradeImpact


class TestTradeImpact(unittest.TestCase):
    """TradeImpact unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testTradeImpact(self):
        """Test TradeImpact"""
        # FIXME: construct object with mandatory attributes with example values
        # model = TradeImpact()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()

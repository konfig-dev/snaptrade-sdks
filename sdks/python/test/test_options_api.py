"""
    SnapTrade

    Connect brokerage accounts to your app for live positions and trading  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: api@snaptrade.com
    Generated by: https://konfigthis.com
"""


import unittest

import snaptrade_client
from snaptrade_client.api.options_api import OptionsApi  # noqa: E501


class TestOptionsApi(unittest.TestCase):
    """OptionsApi unit test stubs"""

    def setUp(self):
        self.api = OptionsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_accounts_account_id_options_chain_get(self):
        """Test case for accounts_account_id_options_chain_get

        Get the options chain  # noqa: E501
        """
        pass

    def test_accounts_account_id_options_get(self):
        """Test case for accounts_account_id_options_get

        Get the options holdings in the account  # noqa: E501
        """
        pass

    def test_accounts_account_id_options_search_get(self):
        """Test case for accounts_account_id_options_search_get

        Search for more specific option quotes from option chain. Date is required but can filter by min, max or min-max.  # noqa: E501
        """
        pass

    def test_accounts_account_id_strategy_impact_post(self):
        """Test case for accounts_account_id_strategy_impact_post

        Get a strategies impact on the account  # noqa: E501
        """
        pass

    def test_accounts_account_id_strategy_place_post(self):
        """Test case for accounts_account_id_strategy_place_post

        Place the strategy order; impact not required but the StrategyOrderQuotes object is  # noqa: E501
        """
        pass

    def test_accounts_account_id_strategy_quotes_post(self):
        """Test case for accounts_account_id_strategy_quotes_post

        Get a price quote for a strategy  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()

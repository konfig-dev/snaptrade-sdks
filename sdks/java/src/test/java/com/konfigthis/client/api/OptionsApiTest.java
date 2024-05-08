/*
 * SnapTrade
 * Connect brokerage accounts to your app for live positions and trading
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: api@snaptrade.com
 *
 * NOTE: This class is auto generated by Konfig (https://konfigthis.com).
 * Do not edit the class manually.
 */


package com.konfigthis.client.api;

import com.konfigthis.client.ApiException;
import com.konfigthis.client.ApiClient;
import com.konfigthis.client.ApiException;
import com.konfigthis.client.Configuration;
import com.konfigthis.client.model.OptionChainInner;
import com.konfigthis.client.model.OptionLeg;
import com.konfigthis.client.model.OptionsGetOptionStrategyRequest;
import com.konfigthis.client.model.OptionsPlaceOptionStrategyRequest;
import com.konfigthis.client.model.OptionsPosition;
import com.konfigthis.client.model.OrderType;
import com.konfigthis.client.model.StrategyOrderRecord;
import com.konfigthis.client.model.StrategyQuotes;
import com.konfigthis.client.model.TimeInForceStrict;
import java.util.UUID;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.BeforeAll;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for OptionsApi
 */
@Disabled
public class OptionsApiTest {

    private static OptionsApi api;

    
    @BeforeAll
    public static void beforeClass() {
        ApiClient apiClient = Configuration.getDefaultApiClient();
        api = new OptionsApi(apiClient);
    }

    /**
     * Create options strategy
     *
     * Creates an option strategy object that will be used to place an option strategy order. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getOptionStrategyTest() throws ApiException {
        UUID underlyingSymbolId = null;
        List<OptionLeg> legs = null;
        String strategyType = null;
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        StrategyQuotes response = api.getOptionStrategy(underlyingSymbolId, legs, strategyType, userId, userSecret, accountId)
                .execute();
        // TODO: test validations
    }

    /**
     * Get the options chain for a symbol
     *
     * Returns the option chain for the specified symbol in the specified account.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getOptionsChainTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        UUID symbol = null;
        List<OptionChainInner> response = api.getOptionsChain(userId, userSecret, accountId, symbol)
                .execute();
        // TODO: test validations
    }

    /**
     * Get options strategy quotes
     *
     * Returns a Strategy Quotes object which has latest market data of the specified option strategy. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getOptionsStrategyQuoteTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        UUID optionStrategyId = null;
        StrategyQuotes response = api.getOptionsStrategyQuote(userId, userSecret, accountId, optionStrategyId)
                .execute();
        // TODO: test validations
    }

    /**
     * Get account option holdings
     *
     * Returns a list of Options Positions. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listOptionHoldingsTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        List<OptionsPosition> response = api.listOptionHoldings(userId, userSecret, accountId)
                .execute();
        // TODO: test validations
    }

    /**
     * Place an option strategy order
     *
     * Places the option strategy order and returns the order record received from the brokerage.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void placeOptionStrategyTest() throws ApiException {
        OrderType orderType = null;
        TimeInForceStrict timeInForce = null;
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        UUID optionStrategyId = null;
        Double price = null;
        StrategyOrderRecord response = api.placeOptionStrategy(orderType, timeInForce, userId, userSecret, accountId, optionStrategyId)
                .price(price)
                .execute();
        // TODO: test validations
    }

}

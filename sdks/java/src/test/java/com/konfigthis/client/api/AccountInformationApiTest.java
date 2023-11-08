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
import com.konfigthis.client.model.Account;
import com.konfigthis.client.model.AccountHoldings;
import com.konfigthis.client.model.AccountHoldingsAccount;
import com.konfigthis.client.model.AccountOrderRecord;
import com.konfigthis.client.model.AccountOrderRecordStatus;
import com.konfigthis.client.model.AccountSyncStatus;
import com.konfigthis.client.model.Action;
import com.konfigthis.client.model.Balance;
import com.konfigthis.client.model.Brokerage;
import com.konfigthis.client.model.BrokerageAuthorization;
import com.konfigthis.client.model.BrokerageType;
import com.konfigthis.client.model.CashRestriction;
import com.konfigthis.client.model.Currency;
import com.konfigthis.client.model.Exchange;
import com.konfigthis.client.model.OptionsSymbol;
import com.konfigthis.client.model.OrderType;
import com.konfigthis.client.model.Position;
import com.konfigthis.client.model.PositionSymbol;
import com.konfigthis.client.model.SecurityType;
import com.konfigthis.client.model.SnapTradeHoldingsAccount;
import com.konfigthis.client.model.SnapTradeHoldingsAccountAccountId;
import com.konfigthis.client.model.SnapTradeHoldingsTotalValue;
import com.konfigthis.client.model.TimeInForce;
import com.konfigthis.client.model.USExchange;
import java.util.UUID;
import com.konfigthis.client.model.UnderlyingSymbol;
import com.konfigthis.client.model.UniversalSymbol;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.BeforeAll;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for AccountInformationApi
 */
@Disabled
public class AccountInformationApiTest {

    private static AccountInformationApi api;

    
    @BeforeAll
    public static void beforeClass() {
        ApiClient apiClient = Configuration.getDefaultApiClient();
        api = new AccountInformationApi(apiClient);
    }

    /**
     * List all accounts for the user, plus balances, positions, and orders for each account.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getAllUserHoldingsTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        UUID brokerageAuthorizations = null;
        List<AccountHoldings> response = api.getAllUserHoldings(userId, userSecret)
                .brokerageAuthorizations(brokerageAuthorizations)
                .execute();
        // TODO: test validations
    }

    /**
     * List account balances
     *
     * A list of account balances for the specified account (one per currency that the account holds).
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getUserAccountBalanceTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        List<Balance> response = api.getUserAccountBalance(userId, userSecret, accountId)
                .execute();
        // TODO: test validations
    }

    /**
     * Return details of a specific investment account
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getUserAccountDetailsTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        Account response = api.getUserAccountDetails(userId, userSecret, accountId)
                .execute();
        // TODO: test validations
    }

    /**
     * List account orders
     *
     * Fetch all recent orders from a user&#39;s account.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getUserAccountOrdersTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        String state = null;
        Integer days = null;
        List<AccountOrderRecord> response = api.getUserAccountOrders(userId, userSecret, accountId)
                .state(state)
                .days(days)
                .execute();
        // TODO: test validations
    }

    /**
     * List account positions
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getUserAccountPositionsTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        List<Position> response = api.getUserAccountPositions(userId, userSecret, accountId)
                .execute();
        // TODO: test validations
    }

    /**
     * List balances, positions and orders for the specified account
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getUserHoldingsTest() throws ApiException {
        UUID accountId = null;
        String userId = null;
        String userSecret = null;
        AccountHoldingsAccount response = api.getUserHoldings(accountId, userId, userSecret)
                .execute();
        // TODO: test validations
    }

    /**
     * List accounts
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listUserAccountsTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        List<Account> response = api.listUserAccounts(userId, userSecret)
                .execute();
        // TODO: test validations
    }

    /**
     * Update details of an investment account
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateUserAccountTest() throws ApiException {
        String userId = null;
        String userSecret = null;
        UUID accountId = null;
        List<Account> response = api.updateUserAccount(userId, userSecret, accountId)
                .execute();
        // TODO: test validations
    }

}

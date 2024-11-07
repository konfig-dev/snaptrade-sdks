/*
SnapTrade

Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { Account } from './account';
import { AccountOrderRecord } from './account-order-record';
import { Balance } from './balance';
import { OptionsPosition } from './options-position';
import { Position } from './position';
import { SnapTradeHoldingsTotalValue } from './snap-trade-holdings-total-value';

/**
 * A wrapper object containing holdings information for a single account.
 * @export
 * @interface AccountHoldingsAccount
 */
export interface AccountHoldingsAccount {
    [key: string]: any;

    /**
     * 
     * @type {Account}
     * @memberof AccountHoldingsAccount
     */
    'account'?: Account;
    /**
     * List of balances for the account. Each element of the list has a distinct currency. Some brokerages like Questrade [allows holding multiple currencies in the same account](https://www.questrade.com/learning/questrade-basics/balances-and-reports/understanding-your-account-balances).
     * @type {Array<Balance>}
     * @memberof AccountHoldingsAccount
     */
    'balances'?: Array<Balance> | null;
    /**
     * List of stock/ETF/crypto/mutual fund positions in the account.
     * @type {Array<Position>}
     * @memberof AccountHoldingsAccount
     */
    'positions'?: Array<Position> | null;
    /**
     * List of option positions in the account.
     * @type {Array<OptionsPosition>}
     * @memberof AccountHoldingsAccount
     */
    'option_positions'?: Array<OptionsPosition> | null;
    /**
     * List of recent orders in the account, including both pending and executed orders. Note that option orders are included in this list. Option orders will have a null `universal_symbol` field and a non-null `option_symbol` field.
     * @type {Array<AccountOrderRecord>}
     * @memberof AccountHoldingsAccount
     */
    'orders'?: Array<AccountOrderRecord> | null;
    /**
     * 
     * @type {SnapTradeHoldingsTotalValue}
     * @memberof AccountHoldingsAccount
     * @deprecated
     */
    'total_value'?: SnapTradeHoldingsTotalValue;
}


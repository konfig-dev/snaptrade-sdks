/*
SnapTrade

Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"


/**
 * Calculated based on the sum of the values of account positions and cash balances
 * @export
 * @interface SnapTradeHoldingsTotalValue
 */
export interface SnapTradeHoldingsTotalValue {
    [key: string]: any;

    /**
     * 
     * @type {number}
     * @memberof SnapTradeHoldingsTotalValue
     */
    'value'?: number | null;
    /**
     * 
     * @type {string}
     * @memberof SnapTradeHoldingsTotalValue
     */
    'currency'?: string | null;
}


/*
SnapTrade

Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { ActionStrictWithOptions } from './action-strict-with-options';
import { ManualTradeFormNotionalValue } from './manual-trade-form-notional-value';
import { ManualTradeFormWithOptionsStopLoss } from './manual-trade-form-with-options-stop-loss';
import { ManualTradeFormWithOptionsTakeProfit } from './manual-trade-form-with-options-take-profit';
import { OrderTypeStrict } from './order-type-strict';
import { TimeInForceStrict } from './time-in-force-strict';

/**
 * Inputs for placing an order with the brokerage.
 * @export
 * @interface ManualTradeFormWithOptions
 */
export interface ManualTradeFormWithOptions {
    /**
     * Unique identifier for the connected brokerage account. This is the UUID used to reference the account in SnapTrade.
     * @type {string}
     * @memberof ManualTradeFormWithOptions
     */
    'account_id': string;
    /**
     * The action describes the intent or side of a trade. This is either `BUY` or `SELL` for Equity symbols or `BUY_TO_OPEN`, `BUY_TO_CLOSE`, `SELL_TO_OPEN` or `SELL_TO_CLOSE` for Options.
     * @type {ActionStrictWithOptions}
     * @memberof ManualTradeFormWithOptions
     */
    'action': ActionStrictWithOptions;
    /**
     * The universal symbol ID of the security to trade. Must be \'null\' if `symbol` is provided, otherwise must be provided.
     * @type {string}
     * @memberof ManualTradeFormWithOptions
     */
    'universal_symbol_id'?: string | null;
    /**
     * The security\'s trading ticker symbol. This currently only support Options symbols in the 21 character OCC format. For example \"AAPL  131124C00240000\" represents a call option on AAPL expiring on 2024-11-13 with a strike price of $240. For more information on the OCC format, see [here](https://en.wikipedia.org/wiki/Option_symbol#OCC_format). If \'symbol\' is provided, then \'universal_symbol_id\' must be \'null\'.
     * @type {string}
     * @memberof ManualTradeFormWithOptions
     */
    'symbol'?: string | null;
    /**
     * The type of order to place.  - For `Limit` and `StopLimit` orders, the `price` field is required. - For `Stop` and `StopLimit` orders, the `stop` field is required. 
     * @type {OrderTypeStrict}
     * @memberof ManualTradeFormWithOptions
     */
    'order_type': OrderTypeStrict;
    /**
     * The Time in Force type for the order. This field indicates how long the order will remain active before it is executed or expires. Here are the supported values:   - `Day` - Day. The order is valid only for the trading day on which it is placed.   - `GTC` - Good Til Canceled. The order is valid until it is executed or canceled.   - `FOK` - Fill Or Kill. The order must be executed in its entirety immediately or be canceled completely. 
     * @type {TimeInForceStrict}
     * @memberof ManualTradeFormWithOptions
     */
    'time_in_force': TimeInForceStrict;
    /**
     * The limit price for `Limit` and `StopLimit` orders.
     * @type {number}
     * @memberof ManualTradeFormWithOptions
     */
    'price'?: number | null;
    /**
     * The price at which a stop order is triggered for `Stop` and `StopLimit` orders.
     * @type {number}
     * @memberof ManualTradeFormWithOptions
     */
    'stop'?: number | null;
    /**
     * For Equity orders, this represents the number of shares for the order. This can be a decimal for fractional orders. Must be `null` if `notional_value` is provided. If placing an Option order, this field represents the number of contracts to buy or sell. (e.g., 1 contract = 100 shares).
     * @type {number}
     * @memberof ManualTradeFormWithOptions
     */
    'units'?: number | null;
    /**
     * 
     * @type {ManualTradeFormNotionalValue}
     * @memberof ManualTradeFormWithOptions
     */
    'notional_value'?: ManualTradeFormNotionalValue | null;
    /**
     * The class of order intended to be placed. Defaults to SIMPLE for regular, one legged trades. Set to BRACKET if looking to place a bracket (One-triggers-a-one-cancels-the-other) order, then specify take profit and stop loss conditions. Bracket orders currently only supported on Alpaca, Tradier, and Tradestation, contact us for more details
     * @type {string}
     * @memberof ManualTradeFormWithOptions
     */
    'order_class'?: ManualTradeFormWithOptionsOrderClassEnum;
    /**
     * 
     * @type {ManualTradeFormWithOptionsStopLoss}
     * @memberof ManualTradeFormWithOptions
     */
    'stop_loss'?: ManualTradeFormWithOptionsStopLoss | null;
    /**
     * 
     * @type {ManualTradeFormWithOptionsTakeProfit}
     * @memberof ManualTradeFormWithOptions
     */
    'take_profit'?: ManualTradeFormWithOptionsTakeProfit | null;
}

type ManualTradeFormWithOptionsOrderClassEnum = 'SIMPLE' | 'BRACKET'



/*
SnapTrade

Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { ActionStrict } from './action-strict';
import { ManualTradeFormNotionalValue } from './manual-trade-form-notional-value';
import { OrderTypeStrict } from './order-type-strict';
import { TimeInForceStrict } from './time-in-force-strict';

/**
 * Inputs for placing an order with the brokerage.
 * @export
 * @interface ManualTradeForm
 */
export interface ManualTradeForm {
    /**
     * Unique identifier for the connected brokerage account. This is the UUID used to reference the account in SnapTrade.
     * @type {string}
     * @memberof ManualTradeForm
     */
    'account_id': string;
    /**
     * The action describes the intent or side of a trade. This is either `BUY` or `SELL`.
     * @type {ActionStrict}
     * @memberof ManualTradeForm
     */
    'action': ActionStrict;
    /**
     * Unique identifier for the symbol within SnapTrade. This is the ID used to reference the symbol in SnapTrade API calls.
     * @type {string}
     * @memberof ManualTradeForm
     */
    'universal_symbol_id': string;
    /**
     * The type of order to place.  - For `Limit` and `StopLimit` orders, the `price` field is required. - For `Stop` and `StopLimit` orders, the `stop` field is required. 
     * @type {OrderTypeStrict}
     * @memberof ManualTradeForm
     */
    'order_type': OrderTypeStrict;
    /**
     * The Time in Force type for the order. This field indicates how long the order will remain active before it is executed or expires. Here are the supported values:   - `Day` - Day. The order is valid only for the trading day on which it is placed.   - `GTC` - Good Til Canceled. The order is valid until it is executed or canceled.   - `FOK` - Fill Or Kill. The order must be executed in its entirety immediately or be canceled completely. 
     * @type {TimeInForceStrict}
     * @memberof ManualTradeForm
     */
    'time_in_force': TimeInForceStrict;
    /**
     * The limit price for `Limit` and `StopLimit` orders.
     * @type {number}
     * @memberof ManualTradeForm
     */
    'price'?: number | null;
    /**
     * The price at which a stop order is triggered for `Stop` and `StopLimit` orders.
     * @type {number}
     * @memberof ManualTradeForm
     */
    'stop'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ManualTradeForm
     */
    'units'?: number | null;
    /**
     * 
     * @type {ManualTradeFormNotionalValue}
     * @memberof ManualTradeForm
     */
    'notional_value'?: ManualTradeFormNotionalValue | null;
}


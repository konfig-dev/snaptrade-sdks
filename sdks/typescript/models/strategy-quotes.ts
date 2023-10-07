/*
SnapTrade

Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/

import { OptionStrategy } from './option-strategy';
import { StrategyQuotesGreek } from './strategy-quotes-greek';

/**
 * 
 * @export
 * @interface StrategyQuotes
 */
export interface StrategyQuotes {
    [key: string]: any;

    /**
     * 
     * @type {OptionStrategy}
     * @memberof StrategyQuotes
     */
    'strategy'?: OptionStrategy;
    /**
     * Trade Price if limit or stop limit order
     * @type {number}
     * @memberof StrategyQuotes
     */
    'open_price'?: number | null;
    /**
     * Trade Price if limit or stop limit order
     * @type {number}
     * @memberof StrategyQuotes
     */
    'bid_price'?: number | null;
    /**
     * Trade Price if limit or stop limit order
     * @type {number}
     * @memberof StrategyQuotes
     */
    'ask_price'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof StrategyQuotes
     */
    'volatility'?: number;
    /**
     * 
     * @type {StrategyQuotesGreek}
     * @memberof StrategyQuotes
     */
    'greek'?: StrategyQuotesGreek;
}


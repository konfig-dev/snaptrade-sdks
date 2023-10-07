/*
SnapTrade

Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/

import { OptionsSymbol } from './options-symbol';
import { UnderlyingSymbol } from './underlying-symbol';

/**
 * 
 * @export
 * @interface BrokerageSymbolOptionSymbol
 */
export interface BrokerageSymbolOptionSymbol {
    /**
     * 
     * @type {string}
     * @memberof BrokerageSymbolOptionSymbol
     */
    'id': string;
    /**
     * 
     * @type {string}
     * @memberof BrokerageSymbolOptionSymbol
     */
    'ticker': string;
    /**
     * 
     * @type {string}
     * @memberof BrokerageSymbolOptionSymbol
     */
    'option_type': BrokerageSymbolOptionSymbolOptionTypeEnum;
    /**
     * 
     * @type {number}
     * @memberof BrokerageSymbolOptionSymbol
     */
    'strike_price': number;
    /**
     * 
     * @type {string}
     * @memberof BrokerageSymbolOptionSymbol
     */
    'expiration_date': string;
    /**
     * 
     * @type {boolean}
     * @memberof BrokerageSymbolOptionSymbol
     */
    'is_mini_option'?: boolean;
    /**
     * 
     * @type {UnderlyingSymbol}
     * @memberof BrokerageSymbolOptionSymbol
     */
    'underlying_symbol': UnderlyingSymbol;
    /**
     * 
     * @type {string}
     * @memberof BrokerageSymbolOptionSymbol
     */
    'local_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof BrokerageSymbolOptionSymbol
     */
    'exchange_id'?: string;
}

type BrokerageSymbolOptionSymbolOptionTypeEnum = 'CALL' | 'PUT'



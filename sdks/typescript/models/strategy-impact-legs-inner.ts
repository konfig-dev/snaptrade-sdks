/* tslint:disable */
/* eslint-disable */
/**
 * SnapTrade
 * Connect brokerage accounts to your app for live positions and trading
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: api@snaptrade.com
 *
 * NOTE: This file is auto generated by Konfig (https://konfigthis.com).
 * Do not edit the class manually.
 */


/**
 * 
 * @export
 * @interface StrategyImpactLegsInner
 */
export interface StrategyImpactLegsInner {
    [key: string]: any;

    /**
     * 
     * @type {number}
     * @memberof StrategyImpactLegsInner
     */
    'legId'?: number;
    /**
     * 
     * @type {string}
     * @memberof StrategyImpactLegsInner
     */
    'symbol'?: string;
    /**
     * 
     * @type {number}
     * @memberof StrategyImpactLegsInner
     */
    'symbolId'?: number;
    /**
     * 
     * @type {number}
     * @memberof StrategyImpactLegsInner
     */
    'legRatioQuantity'?: number;
    /**
     * 
     * @type {string}
     * @memberof StrategyImpactLegsInner
     */
    'side'?: string;
    /**
     * 
     * @type {string}
     * @memberof StrategyImpactLegsInner
     */
    'avgExecPrice'?: string;
    /**
     * 
     * @type {string}
     * @memberof StrategyImpactLegsInner
     */
    'lastExecPrice'?: string;
}


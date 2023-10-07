/*
SnapTrade

Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/

import { MonthlyDividends } from './monthly-dividends';
import { NetContributions } from './net-contributions';
import { NetDividend } from './net-dividend';
import { PastValue } from './past-value';
import { SubPeriodReturnRate } from './sub-period-return-rate';

/**
 * Performance Custom Response Object
 * @export
 * @interface PerformanceCustom
 */
export interface PerformanceCustom {
    [key: string]: any;

    /**
     * 
     * @type {Array<PastValue>}
     * @memberof PerformanceCustom
     */
    'totalEquityTimeframe'?: Array<PastValue>;
    /**
     * 
     * @type {NetContributions}
     * @memberof PerformanceCustom
     */
    'contributions'?: NetContributions;
    /**
     * 
     * @type {Array<PastValue>}
     * @memberof PerformanceCustom
     */
    'contributionTimeframe'?: Array<PastValue>;
    /**
     * 
     * @type {Array<PastValue>}
     * @memberof PerformanceCustom
     */
    'contributionTimeframeCumulative'?: Array<PastValue>;
    /**
     * 
     * @type {Array<PastValue>}
     * @memberof PerformanceCustom
     */
    'withdrawalTimeframe'?: Array<PastValue>;
    /**
     * Current streak of cosecutive months where contributions were made
     * @type {number}
     * @memberof PerformanceCustom
     */
    'contributionStreak'?: number | null;
    /**
     * Number of months in the timeframe with contributions
     * @type {number}
     * @memberof PerformanceCustom
     */
    'contributionMonthsContributed'?: number | null;
    /**
     * Total months in timeframe
     * @type {number}
     * @memberof PerformanceCustom
     */
    'contributionTotalMonths'?: number | null;
    /**
     * 
     * @type {Array<NetDividend>}
     * @memberof PerformanceCustom
     */
    'dividends'?: Array<NetDividend>;
    /**
     * Total dividends received over the timeframe
     * @type {number}
     * @memberof PerformanceCustom
     */
    'dividendIncome'?: number | null;
    /**
     * Average dividends received per month over the timeframe
     * @type {number}
     * @memberof PerformanceCustom
     */
    'monthlyDividends'?: number | null;
    /**
     * list of tickers which may not be supported or may not have accurate price data
     * @type {Array<string>}
     * @memberof PerformanceCustom
     */
    'badTickers'?: Array<string>;
    /**
     * 
     * @type {Array<MonthlyDividends>}
     * @memberof PerformanceCustom
     */
    'dividendTimeline'?: Array<MonthlyDividends>;
    /**
     * commissions incurred during the timeframe
     * @type {number}
     * @memberof PerformanceCustom
     */
    'commissions'?: number | null;
    /**
     * forex fees incurred during the timeframe
     * @type {number}
     * @memberof PerformanceCustom
     */
    'forexFees'?: number | null;
    /**
     * other fees incurred during the timeframe
     * @type {number}
     * @memberof PerformanceCustom
     */
    'fees'?: number | null;
    /**
     * The return rate over the timeframe. Annualized if timeframe is longer than 1 year
     * @type {number}
     * @memberof PerformanceCustom
     */
    'rateOfReturn'?: number | null;
    /**
     * 
     * @type {Array<SubPeriodReturnRate>}
     * @memberof PerformanceCustom
     */
    'returnRateTimeframe'?: Array<SubPeriodReturnRate>;
    /**
     * Whether the user has detailed mode enabled (more frequent data points for totalEquity and contribution timeframes)
     * @type {boolean}
     * @memberof PerformanceCustom
     */
    'detailedMode'?: boolean;
}


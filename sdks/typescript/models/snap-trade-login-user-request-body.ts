/*
SnapTrade

Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"


/**
 * Data to login a user via SnapTrade Partner
 * @export
 * @interface SnapTradeLoginUserRequestBody
 */
export interface SnapTradeLoginUserRequestBody {
    /**
     * Slug of the brokerage to connect the user to. See [the integrations page](https://snaptrade.notion.site/66793431ad0b416489eaabaf248d0afb?v=3cfea70ef4254afc89704e47275a7a9a&pvs=4) for a list of supported brokerages and their slugs.
     * @type {string}
     * @memberof SnapTradeLoginUserRequestBody
     */
    'broker'?: string;
    /**
     * When set to `true`, user will be redirected back to the partner\'s site instead of the connection portal. This parameter is ignored if the connection portal is loaded inside an iframe. See the [guide on ways to integrate the connection portal](/docs/implement-connection-portal) for more information.
     * @type {boolean}
     * @memberof SnapTradeLoginUserRequestBody
     */
    'immediateRedirect'?: boolean;
    /**
     * URL to redirect the user to after the user connects their brokerage account. This parameter is ignored if the connection portal is loaded inside an iframe. See the [guide on ways to integrate the connection portal](/docs/implement-connection-portal) for more information.
     * @type {string}
     * @memberof SnapTradeLoginUserRequestBody
     */
    'customRedirect'?: string;
    /**
     * The UUID of the brokerage connection to be reconnected. This parameter should be left empty unless you are reconnecting a disabled connection. See the [guide on fixing broken connections](/docs/fix-broken-connections) for more information.
     * @type {string}
     * @memberof SnapTradeLoginUserRequestBody
     */
    'reconnect'?: string;
    /**
     * Sets whether the connection should be read-only or trade-enabled. Defaults to read-only if not specified.
     * @type {string}
     * @memberof SnapTradeLoginUserRequestBody
     */
    'connectionType'?: SnapTradeLoginUserRequestBodyConnectionTypeEnum;
    /**
     * Sets the version of the connection portal to render.
     * @type {string}
     * @memberof SnapTradeLoginUserRequestBody
     */
    'connectionPortalVersion'?: SnapTradeLoginUserRequestBodyConnectionPortalVersionEnum;
}

type SnapTradeLoginUserRequestBodyConnectionTypeEnum = 'read' | 'trade'
type SnapTradeLoginUserRequestBodyConnectionPortalVersionEnum = 'v4' | 'v3' | 'v2'



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

import globalAxios, { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction, isBrowser } from '../common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { Status } from '../models';
import { paginate } from "../pagination/paginate";
import { requestBeforeHook } from '../requestBeforeHook';
/**
 * ApiStatusApi - axios parameter creator
 * @export
 */
export const ApiStatusApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Check whether the API is operational and verify timestamps.
         * @summary Get API Status
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        check: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions: AxiosRequestConfig = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = configuration && !isBrowser() ? { "User-Agent": configuration.userAgent } : {} as any;
            const localVarQueryParameter = {} as any;


    
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            requestBeforeHook({
                queryParameters: localVarQueryParameter,
                requestConfig: localVarRequestOptions,
                path: localVarPath,
                configuration
            });

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * ApiStatusApi - functional programming interface
 * @export
 */
export const ApiStatusApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = ApiStatusApiAxiosParamCreator(configuration)
    return {
        /**
         * Check whether the API is operational and verify timestamps.
         * @summary Get API Status
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async check(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Status>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.check(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * ApiStatusApi - factory interface
 * @export
 */
export const ApiStatusApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = ApiStatusApiFp(configuration)
    return {
        /**
         * Check whether the API is operational and verify timestamps.
         * @summary Get API Status
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        check(options?: AxiosRequestConfig): AxiosPromise<Status> {
            return localVarFp.check(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * ApiStatusApiGenerated - object-oriented interface
 * @export
 * @class ApiStatusApiGenerated
 * @extends {BaseAPI}
 */
export class ApiStatusApiGenerated extends BaseAPI {
    /**
     * Check whether the API is operational and verify timestamps.
     * @summary Get API Status
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ApiStatusApiGenerated
     */
    public check(options?: AxiosRequestConfig) {
        return ApiStatusApiFp(this.configuration).check(options).then((request) => request(this.axios, this.basePath));
    }
}

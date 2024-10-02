/*
 * SnapTrade
 *
 * Connect brokerage accounts to your app for live positions and trading
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: api@snaptrade.com
 * Generated by: https://konfigthis.com
 */

using System;
using System.IO;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Reflection;
using RestSharp;
using Xunit;

using SnapTrade.Net.Client;
using SnapTrade.Net.Api;
using SnapTrade.Net.Model;

namespace SnapTrade.Net.Test.Api
{
    /// <summary>
    ///  Class for testing TransactionsAndReportingApi
    /// </summary>
    public class TransactionsAndReportingApiTests : IDisposable
    {
        private Snaptrade client;

        public TransactionsAndReportingApiTests()
        {
            client = new Snaptrade();
            client.SetBasePath("http://127.0.0.1:4010");
            // Configure custom BasePath if desired
            // client.SetBasePath("https://api.snaptrade.com/api/v1");
            client.SetClientId(System.Environment.GetEnvironmentVariable("SNAPTRADE_CLIENT_ID"));
            client.SetConsumerKey(System.Environment.GetEnvironmentVariable("SNAPTRADE_CONSUMER_KEY"));

        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test GetActivities
        /// </summary>
        [Fact]
        public void GetActivitiesTest()
        {
            var userId = "userId_example";
            var userSecret = "userSecret_example";
            var startDate = DateTime.Parse("2013-10-20"); // The start date (inclusive) of the transaction history to retrieve. If not provided, the default is the first transaction known to SnapTrade based on `trade_date`. (optional) 
            var endDate = DateTime.Parse("2013-10-20"); // The end date (inclusive) of the transaction history to retrieve. If not provided, the default is the last transaction known to SnapTrade based on `trade_date`. (optional) 
            var accounts = "accounts_example"; // Optional comma separated list of SnapTrade Account IDs used to filter the request to specific accounts. If not provided, the default is all known brokerage accounts for the user. The `brokerageAuthorizations` parameter takes precedence over this parameter. (optional) 
            var brokerageAuthorizations = "brokerageAuthorizations_example"; // Optional comma separated list of SnapTrade Connection (Brokerage Authorization) IDs used to filter the request to only accounts that belong to those connections. If not provided, the default is all connections for the user. This parameter takes precedence over the `accounts` parameter. (optional) 
            var type = "BUY,SELL,DIVIDEND"; // Optional comma separated list of transaction types to filter by. SnapTrade does a best effort to categorize brokerage transaction types into a common set of values. Here are some of the most popular values:   - `BUY` - Asset bought.   - `SELL` - Asset sold.   - `DIVIDEND` - Dividend payout.   - `CONTRIBUTION` - Cash contribution.   - `WITHDRAWAL` - Cash withdrawal.   - `REI` - Dividend reinvestment.   - `INTEREST` - Interest deposited into the account.   - `FEE` - Fee withdrawn from the account.  (optional) 
            
            try
            {
                // Get transaction history for a user
                List<UniversalActivity> result = client.TransactionsAndReporting.GetActivities(userId, userSecret, startDate, endDate, accounts, brokerageAuthorizations, type);
                Console.WriteLine(result);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling TransactionsAndReportingApi.GetActivities: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }

        /// <summary>
        /// Test GetReportingCustomRange
        /// </summary>
        [Fact]
        public void GetReportingCustomRangeTest()
        {
            var startDate = DateTime.Parse("2013-10-20");
            var endDate = DateTime.Parse("2013-10-20");
            var userId = "userId_example";
            var userSecret = "userSecret_example";
            var accounts = "accounts_example"; // Optional comma separated list of account IDs used to filter the request on specific accounts (optional) 
            var detailed = true; // Optional, increases frequency of data points for the total value and contribution charts if set to true (optional) 
            var frequency = "frequency_example"; // Optional frequency for the rate of return chart (defaults to monthly). Possible values are daily, weekly, monthly, quarterly, yearly. (optional) 
            
            try
            {
                // Get performance information for a specific timeframe
                PerformanceCustom result = client.TransactionsAndReporting.GetReportingCustomRange(startDate, endDate, userId, userSecret, accounts, detailed, frequency);
                Console.WriteLine(result);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling TransactionsAndReportingApi.GetReportingCustomRange: " + e.Message);
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
            catch (ClientException e)
            {
                Console.WriteLine(e.Response.StatusCode);
                Console.WriteLine(e.Response.RawContent);
                Console.WriteLine(e.InnerException);
            }
        }
    }
}

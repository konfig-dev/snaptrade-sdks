/*
SnapTrade

Testing TransactionsAndReportingApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package snaptrade

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // snaptrade "github.com/passiv/snaptrade-sdks/sdks/go"
)

func Test_snaptrade_TransactionsAndReportingApiService(t *testing.T) {

    // configuration := snaptrade.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetPartnerClientId(os.Getenv("SNAPTRADE_CLIENT_ID"))
    configuration.SetConsumerKey(os.Getenv("SNAPTRADE_CONSUMER_KEY"))
    client := snaptrade.NewAPIClient(configuration)
    */

    t.Run("Test TransactionsAndReportingApiService GetActivities", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.TransactionsAndReportingApi.GetActivities(
            "userId_example",
            "userSecret_example",
        )
        request.StartDate(2013-10-20)
        request.EndDate(2013-10-20)
        request.Accounts("accounts_example")
        request.BrokerageAuthorizations("brokerageAuthorizations_example")
        request.Type(""BUY,SELL,DIVIDEND"")
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test TransactionsAndReportingApiService GetReportingCustomRange", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.TransactionsAndReportingApi.GetReportingCustomRange(
            2013-10-20,
            2013-10-20,
            "userId_example",
            "userSecret_example",
        )
        request.Accounts("accounts_example")
        request.Detailed(true)
        request.Frequency("frequency_example")
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}
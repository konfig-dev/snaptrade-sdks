/*
SnapTrade

Testing TradingApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package snaptrade

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // snaptrade "github.com/passiv/snaptrade-sdks/sdks/go"
)

func Test_snaptrade_TradingApiService(t *testing.T) {

    // configuration := snaptrade.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetPartnerClientId(os.Getenv("SNAPTRADE_CLIENT_ID"))
    configuration.SetConsumerKey(os.Getenv("SNAPTRADE_CONSUMER_KEY"))
    client := snaptrade.NewAPIClient(configuration)
    */

    t.Run("Test TradingApiService CancelUserAccountOrder", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        tradingCancelUserAccountOrderRequest := *snaptrade.NewTradingCancelUserAccountOrderRequest()
        tradingCancelUserAccountOrderRequest.SetBrokerageOrderId("66a033fa-da74-4fcf-b527-feefdec9257e")
        
        request := client.TradingApi.CancelUserAccountOrder(
            "userId_example",
            "userSecret_example",
            ""38400000-8cf0-11bd-b23e-10b96e4ef00d"",
            tradingCancelUserAccountOrderRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test TradingApiService GetOrderImpact", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        units := *snaptrade.Newfloat32()
        notionalValue := *snaptrade.NewManualTradeFormNotionalValue()
        
        manualTradeForm := *snaptrade.NewManualTradeForm(
            "917c8734-8470-4a3e-a18f-57c3f2ee6631",
            null,
            "2bcd7cc3-e922-4976-bce1-9858296801c3",
            null,
            null,
        )
        manualTradeForm.SetPrice(31.33)
        manualTradeForm.SetStop(31.33)
        manualTradeForm.SetUnits(units)
        manualTradeForm.SetNotionalValue(notionalValue)
        
        request := client.TradingApi.GetOrderImpact(
            "userId_example",
            "userSecret_example",
            manualTradeForm,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test TradingApiService GetUserAccountQuotes", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.TradingApi.GetUserAccountQuotes(
            "userId_example",
            "userSecret_example",
            "symbols_example",
            ""38400000-8cf0-11bd-b23e-10b96e4ef00d"",
        )
        request.UseTicker(true)
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test TradingApiService PlaceForceOrder", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        universalSymbolId := *snaptrade.Newstring()
        units := *snaptrade.Newfloat32()
        notionalValue := *snaptrade.NewManualTradeFormNotionalValue()
        
        manualTradeFormWithOptions := *snaptrade.NewManualTradeFormWithOptions(
            "917c8734-8470-4a3e-a18f-57c3f2ee6631",
            null,
            null,
            null,
        )
        manualTradeFormWithOptions.SetUniversalSymbolId(universalSymbolId)
        manualTradeFormWithOptions.SetSymbol("AAPL  131124C00240000")
        manualTradeFormWithOptions.SetPrice(31.33)
        manualTradeFormWithOptions.SetStop(31.33)
        manualTradeFormWithOptions.SetUnits(units)
        manualTradeFormWithOptions.SetNotionalValue(notionalValue)
        
        request := client.TradingApi.PlaceForceOrder(
            "userId_example",
            "userSecret_example",
            manualTradeFormWithOptions,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test TradingApiService PlaceOrder", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        
        validatedTradeBody := *snaptrade.NewValidatedTradeBody()
        validatedTradeBody.SetWaitToConfirm(true)
        
        request := client.TradingApi.PlaceOrder(
            ""38400000-8cf0-11bd-b23e-10b96e4ef00d"",
            "userId_example",
            "userSecret_example",
        )
        request.ValidatedTradeBody(validatedTradeBody)
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}
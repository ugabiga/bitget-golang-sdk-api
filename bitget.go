package bitget

import (
	"net/http"

	"github.com/aexlab51/bitget-golang-sdk-api/internal/common"
	"github.com/aexlab51/bitget-golang-sdk-api/pkg/client/broker"
	"github.com/aexlab51/bitget-golang-sdk-api/pkg/client/mix"
	"github.com/aexlab51/bitget-golang-sdk-api/pkg/client/spot"
	"github.com/aexlab51/bitget-golang-sdk-api/pkg/client/ws"
)

// client
type Client struct {
	client 				*common.BitgetRestClient

	brokerService 		*broker.BrokerAccountClient

	mixAccountService	*mix.MixAccountClient
	mixMarketService	*mix.MixMarketClient
	mixOrderService		*mix.MixOrderClient
	mixPlanService		*mix.MixPlanClient
	mixPositionService	*mix.MixPositionClient
	mixTraceService		*mix.MixTraceClient

	spotAccountService	*spot.SpotAccountClient
	spotMarketService	*spot.SpotMarketClient
	spotOrderService	*spot.SpotOrderClient
	spotPublicService	*spot.SpotPublicClient
	//spotWalletService	*spot.SpotWalletClient			// @todo
}

func NewClient() *Client {
	bc := new(common.BitgetRestClient).Init()

	return &Client{
		client: bc,

		brokerService:      &broker.BrokerAccountClient{bc},
		mixAccountService:  &mix.MixAccountClient{bc},
		mixMarketService:   &mix.MixMarketClient{bc},
		mixOrderService:    &mix.MixOrderClient{bc},
		mixPlanService:     &mix.MixPlanClient{bc},
		mixPositionService: &mix.MixPositionClient{bc},
		mixTraceService:    &mix.MixTraceClient{bc},
		spotAccountService: &spot.SpotAccountClient{bc},
		spotMarketService:  &spot.SpotMarketClient{bc},
		spotOrderService:   &spot.SpotOrderClient{bc},
		spotPublicService:  &spot.SpotPublicClient{bc},
	}
}

func (c *Client) SetHttpClient(client *http.Client) *Client {
	c.client.HttpClient = client
	return c
}


// broker
func (c *Client) GetBrokerService() *broker.BrokerAccountClient {
	return c.brokerService
}

// mix
func (c *Client) GetMixAccountService() *mix.MixAccountClient {
	return c.mixAccountService
}
func (c *Client) GetMixMarketService() *mix.MixMarketClient {
	return c.mixMarketService
}
func (c *Client) GetMixOrderService() *mix.MixOrderClient {
	return c.mixOrderService
}
func (c *Client) GetMixPlanService() *mix.MixPlanClient {
	return c.mixPlanService
}
func (c *Client) GetMixPositionService() *mix.MixPositionClient {
	return c.mixPositionService
}
func (c *Client) GetMixTraceService() *mix.MixTraceClient {
	return c.mixTraceService
}

// spot
func (c *Client) GetSpotAccountService() *spot.SpotAccountClient {
	return c.spotAccountService
}
func (c *Client) GetSpotMarketService() *spot.SpotMarketClient {
	return c.spotMarketService
}
func (c *Client) GetSpotOrderService() *spot.SpotOrderClient {
	return c.spotOrderService
}
func (c *Client) GetSpotPublicService() *spot.SpotPublicClient {
	return c.spotPublicService
}
/* @todo
func (c *Client) NewSpotWalletService() *spot.SpotWalletClient {
	return c.spotWalletService
}
//*/


// ws - @todo
func NewWsClient() *ws.BitgetWsClient {
	return new(ws.BitgetWsClient)
}

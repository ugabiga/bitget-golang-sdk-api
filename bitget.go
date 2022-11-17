package bitget

import (
	"github.com/aexlab51/bitget-golang-sdk-api/pkg/client/broker"
	"github.com/aexlab51/bitget-golang-sdk-api/pkg/client/mix"
	"github.com/aexlab51/bitget-golang-sdk-api/pkg/client/spot"
	"github.com/aexlab51/bitget-golang-sdk-api/pkg/client/ws"
)

// broker
func NewBrokerAccountClient() *broker.BrokerAccountClient {
	return new(broker.BrokerAccountClient).Init()
}

// mix
func NewMixAccountClient() *mix.MixAccountClient {
	return new(mix.MixAccountClient).Init()
}
func NewMixMarketClient() *mix.MixMarketClient {
	return new(mix.MixMarketClient).Init()
}
func NewMixOrderClient() *mix.MixOrderClient {
	return new(mix.MixOrderClient).Init()
}
func NewMixPlanClient() *mix.MixPlanClient {
	return new(mix.MixPlanClient).Init()
}
func NewMixPositionClient() *mix.MixPositionClient {
	return new(mix.MixPositionClient).Init()
}
func NewMixTraceClient() *mix.MixTraceClient {
	return new(mix.MixTraceClient).Init()
}

// spot
func NewSpotAccountClient() *spot.SpotAccountClient {
	return new(spot.SpotAccountClient).Init()
}
func NewSpotMarketClient() *spot.SpotMarketClient {
	return new(spot.SpotMarketClient).Init()
}
func NewSpotOrderClient() *spot.SpotOrderClient {
	return new(spot.SpotOrderClient).Init()
}
func NewSpotPublicClient() *spot.SpotPublicClient {
	return new(spot.SpotPublicClient).Init()
}
/* @todo
func NewSpotWalletClient() *spot.SpotWalletClient {
	return new(spot.SpotWalletClient).Init()
}
//*/

// ws
func NewWsClient() *ws.BitgetWsClient {
	return new(ws.BitgetWsClient)
}

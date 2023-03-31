package plan

type PlaceTrailingStopReq struct {
	/**
	 * Currency pair
	 */
	Symbol string `json:"symbol"`
	/**
	 * Deposit currency
	 */
	MarginCoin string `json:"marginCoin"`
	/**
	 * Trigger price
	 */
	TriggerPrice string `json:"triggerPrice"`
	TriggerType  string `json:"triggerType"`
	/**
	 * Is this position long or short
	 */
	Size string `json:"size"`
	/**
	 * Is this position open_long or open_short
	 */
	Side string `json:"side"`
	/**
	 * "1" means 1.0% price correction, max "10"
	 */
	RangeRate string `json:"rangeRate"`
	/**
	 * Client ID
	 */
	ClientOid string `json:"clientOid"`
}

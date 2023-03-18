package account

type SetPositionModeReq struct {
	HoldMode    string `json:"holdMode"`
	ProductType string `json:"productType"`
}

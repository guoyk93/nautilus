package svc_mp_token

type GetResp struct {
	AccessToken string `json:"access_token"`
}

type RefreshCommand struct {
	CurrentToken string `json:"current_token"`
}

type RefreshResp struct {
	AccessToken string `json:"access_token"`
}

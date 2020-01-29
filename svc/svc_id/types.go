package svc_id

type NewIDQuery struct {
	Size int `json:"size" query:"size" default:"1"`
}

type NewIDResp struct {
	IDs []string `json:"ids"`
}

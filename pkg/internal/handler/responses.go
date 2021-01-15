package handler

type Meta struct {
	Version   string `json:"version"`
	APIStatus string `json:"api_status"`
}

type LoginResponse struct {
	Users []struct {
		Token        string `json:"token"`
		ExpiresAfter string `json:"expires_after"`
	} `json:"users"`
	Meta Meta `json:"meta"`
}

package flight

type Flight struct {
	GUID     string `json:"guid"`
	Origin   string `json:"origin"`
	Destiny  string `json:"destiny"`
	Aircraft string `json:"aircraft"`
}

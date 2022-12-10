package daoarangodb

// 帳號資料: 廠商帳號, 管理員帳號
type Account struct {
	Key         string `json:"_key,omitempty"`
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
	LastLogin   string `json:"lastLogin"`
	CreatTime   string `json:"creatTime"`
	Token       string `json:"token"`
	Permission  string `json:"permission"`
}

// 供應商簡介: 遊戲供應商, 網站供應商索引
type Provider struct {
	Key          string `json:"_key,omitempty"`
	Name         string `json:"name"`
	ProviderType string `json:"providerType"`
}

// 遊戲供應商詳資料
type GameProvider struct {
	Key        string   `json:"_key,omitempty"`
	ProviderId string   `json:"providerId"`
	Setting    *Setting `json:"setting"`
}

// 遊戲供應商詳資料
type WebSiteProvider struct {
	Key        string   `json:"_key,omitempty"`
	ProviderId string   `json:"providerId"`
	Setting    *Setting `json:"setting"`
	ProvidGame string   `json:"providGame"`
}

type Setting struct {
	Url string `json:"url"`
}

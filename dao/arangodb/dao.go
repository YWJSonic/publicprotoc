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

type Token struct {
	Key        string `json:"_key,omitempty"`
	AccountId  string `json:"accountId"` // Account Key
	Token      string `json:"token"`
	CreateTime int64  `json:"createTime"`
	ExpTime    int64  `json:"expTime"`
}

// 供應商簡介: 遊戲供應商, 網站供應商索引
type Provider struct {
	Key          string `json:"_key,omitempty"`
	Name         string `json:"name"`
	ProviderType string `json:"providerType"`
}

type Game struct {
	Key        string      `json:"_key,omitempty"`
	Name       string      `json:"name"`
	ProviderId string      `json:"providerId"`
	Setting    interface{} `json:"setting"`
}

type ProviderSetting struct {
	Key        string      `json:"_key,omitempty"`
	ProviderId string      `json:"providerId"`
	Setting    interface{} `json:"setting"`
}

type DefaultSetting struct {
	Key         string      `json:"_key,omitempty"`
	SettingType string      `json:"settingType"` // 此設定類型 provider setting, game setting
	OwnerKey    string      `json:"ownerKey"`    // 預設使用者
	Setting     interface{} `json:"setting"`
}

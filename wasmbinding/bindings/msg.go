package bindings

import "github.com/CosmWasm/wasmvm/types"


type OmniflixMsg struct {
	TestBinding *TestBinding `json:"test_binding"`
	CreateDenom *CreateDenom `json:"create_denom"`
	MsgCreateDenom *MsgCreateDenom `json:"msg_create_denom"`
	
}

type TestBinding struct {
	Test string `json:"test"`
}

type CreateDenom struct {
	Subdenom string `json:"subdenom"`
}

type MsgCreateDenom struct {

	Id          string     `json:"id"`
	Symbol      string     `json:"symbol"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	PreviewURI  string     `json:"preview_uri" yaml:"preview_uri"`
	Schema      string     `json:"schema"`
	Sender      string     `json:"sender"`
	CreationFee types.Coin `json:"creation_fee" yaml:"creation_fee"`
}

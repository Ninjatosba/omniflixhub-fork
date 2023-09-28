package bindings


type OmniflixMsg struct {
	TestBinding *TestBinding `json:"test_binding"`
	CreateDenom *CreateDenom `json:"create_denom"`
}

type TestBinding struct {
	Test string `json:"test"`
}

type CreateDenom struct {
	Subdenom string `json:"subdenom"`
}
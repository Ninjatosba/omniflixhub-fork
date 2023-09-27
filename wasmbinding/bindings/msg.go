package bindings


type OmniflixMsg struct {
	TestBinding *TestBinding `json:"test_binding"`
}

type TestBinding struct {
	test string `json:"test"`
}
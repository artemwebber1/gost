package gost

type translationApiResponse struct {
	SourceLang      string `json:"source-language"`
	SourceText      string `json:"source-text"`
	DestinationLang string `json:"destination-language"`
	DestinationText string `json:"destination-text"`
}

package analyzed_response

type AnalyzedResponse struct {
	Sentiment          string `json:"sentiment"`
	RephrasedStatement string `json:"rephrased_statement"`
	Benefits           string `json:"benefits"`
}

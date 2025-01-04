package precheck

type Request struct {
	Account string `json:"account"`
}

type Response struct {
	CorrelationID string `json:"correlation_id"`
	Answer        Answer
}

type Answer struct {
	FIO string `json:"FIO"`
}

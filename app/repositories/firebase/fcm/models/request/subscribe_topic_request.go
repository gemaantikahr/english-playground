package request

type SubScribeTopicRequest struct {
	Tokens []string `json:"tokens"`
	Name   string   `json:"name"`
}

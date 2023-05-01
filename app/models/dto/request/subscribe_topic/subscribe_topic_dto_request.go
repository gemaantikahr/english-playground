package subscribe_topic

type SubscribeTopicDTORequest struct {
	Name   string   `json:"name"`
	Tokens []string `json:"tokens"`
}

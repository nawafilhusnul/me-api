package domain

const (
	GETSuccess    = "success"
	PUTSuccess    = "item successfully updated"
	DELETESuccess = "item successfully deleted"
	POSTSuccess   = "item successfully created"
)

type MetaResponse struct {
	Message string `json:"message"`
}

package handler

type Handler struct {
	GetPingHandler
	GetHealthHandler
	GetHealthAuthHandler
}

func NewHandlers() *Handler {
	return &Handler{
		GetPingHandler:       GetPingHandler{},
		GetHealthHandler:     GetHealthHandler{},
		GetHealthAuthHandler: GetHealthAuthHandler{},
	}
}

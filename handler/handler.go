package handler

type Handler struct {
	GetPingHandler
	GetHealthHandler
}

func NewHandlers() *Handler {
	return &Handler{
		GetPingHandler:   GetPingHandler{},
		GetHealthHandler: GetHealthHandler{},
	}
}

package handler

type Handler struct {
	GetHealthHandler
}

func NewHandlers() *Handler {
	return &Handler{
		GetHealthHandler: GetHealthHandler{},
	}
}

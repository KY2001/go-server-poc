package handler

type Handler struct {
	GetPingHandler
	GetHealthHandler
	GetHealthAuthHandler
	PostAuthSignupHandler
	GetUserHandler
}

func NewHandlers() *Handler {
	return &Handler{
		GetPingHandler:        GetPingHandler{},
		GetHealthHandler:      GetHealthHandler{},
		GetHealthAuthHandler:  GetHealthAuthHandler{},
		PostAuthSignupHandler: PostAuthSignupHandler{},
		GetUserHandler:        GetUserHandler{},
	}
}

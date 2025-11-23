package middlewarerepository

type (
	MiddlewareRepositoryHandler interface{}

	middlewareRepository struct{}
)

func NewMiddlewareRepository() MiddlewareRepositoryHandler {
	return &middlewareRepository{}
}

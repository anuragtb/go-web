package server

const (
	InvalidUserID       = "invalidUserID"
	InternalError       = "internalError"
	UserNotFound        = "userNotFound"
	InvalidBindingModel = "invalidBindingModel"
	EntityCreationError = "entityCreationError"
)

type Booms struct {
	Errors []Boom
}

// boom represent the basic structure of an json error
type Boom struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

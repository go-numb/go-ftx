package rest

type Requester interface {
	Path() string
	Method() string
	Query() string
	Payload() []byte
}

package v1

type Controllers struct {
	Index *ControllerIndex
}

func NewControllers() *Controllers {
	return &Controllers{
		Index:  NewControllerIndex(),
	}
}

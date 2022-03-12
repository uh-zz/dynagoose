package dynagoose

type Schema interface {
	Select() string
	Update() string
	Create() string
	Delete() string
}

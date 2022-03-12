package dynagoose

type Dynamo struct{}

func (d *Dynamo) Select() string {
	return "select"
}

func (d *Dynamo) Update() string {
	return "update"
}

func (d *Dynamo) Create() string {
	return "create"
}

func (d *Dynamo) Delete() string {
	return "delete"
}

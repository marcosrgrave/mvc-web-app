package utils

const (
	CREATE = "CREATE"
	READ   = "READ"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

// var test := CRUDMessage("Product").Create().Error()
func CRUDMessage(model string) *CRUD {
	return &CRUD{
		Model:  model,
		method: "",
	}
}

func (c *CRUD) Create() *CRUD {
	return &CRUD{
		Model:  c.Model,
		method: CREATE,
	}
}

func (c *CRUD) Read() *CRUD {
	return &CRUD{
		Model:  c.Model,
		method: READ,
	}
}

func (c *CRUD) Update() *CRUD {
	return &CRUD{
		Model:  c.Model,
		method: UPDATE,
	}
}

func (c *CRUD) Delete() *CRUD {
	return &CRUD{
		Model:  c.Model,
		method: DELETE,
	}
}

type CRUD struct {
	Model  string
	method string
}

func (c *CRUD) Error() string {
	// var msg string = string("%s %s the %s", status, method, model)

	if c.method == CREATE {
		return "Error creating the " + c.Model
	}
	if c.method == READ {
		return "Error reading the " + c.Model
	}
	if c.method == UPDATE {
		return "Error updating the " + c.Model
	}
	if c.method == DELETE {
		return "Error deleting the " + c.Model
	}
	return ""
}

func (c *CRUD) Success() string {
	return ""
}

func (c *CRUD) NotFound() string {
	return ""
}

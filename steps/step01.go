package GoZero

type Variable struct {
	Data int
}

func NewVariable(data int) *Variable {
	variable := new(Variable)
	variable.Data = data
	
	return variable
}
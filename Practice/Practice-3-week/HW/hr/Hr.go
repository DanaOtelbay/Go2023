package hr

type Hr struct {
	position string
	salary   float64
	address  string
}

func NewHr(position string, salary float64, address string) *Hr {
	return &Hr{position: position, salary: salary, address: address}
}

func (h *Hr) GetPosition() string {
	return h.position
}

func (h *Hr) SetPosition(position string) {
	h.position = position
}

func (h *Hr) GetSalary() float64 {
	return h.salary
}

func (h *Hr) SetSalary(salary float64) {
	h.salary = salary
}

func (h *Hr) GetAddress() string {
	return h.address
}

func (h *Hr) SetAddress(address string) {
	h.address = address
}

func (h *Hr) TerminateEmployee(name string) string {
	return "Employee terminated: " + name
}

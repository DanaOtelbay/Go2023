package manager

type Manager struct {
	position    string
	salary      float64
	address     string
	department  string
	phoneNumber string
}

func NewManager(position string, salary float64, address string, department string, phoneNumber string) *Manager {
	return &Manager{
		position:    position,
		salary:      salary,
		address:     address,
		department:  department,
		phoneNumber: phoneNumber,
	}
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}

func (m *Manager) GetSalary() float64 {
	return m.salary
}

func (m *Manager) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Manager) GetAddress() string {
	return m.address
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}

func (m *Manager) GetDepartment() string {
	return m.department
}

func (m *Manager) SetDepartment(department string) {
	m.department = department
}

func (m *Manager) GetPhoneNumber() string {
	return m.phoneNumber
}

func (m *Manager) SetPhoneNumber(phoneNumber string) {
	m.phoneNumber = phoneNumber
}

func (m *Manager) OrganizeMeetings() string {
	return "Meetings organized by Manager"
}

func (m *Manager) SetTargets() string {
	return "Targets set by Manager"
}

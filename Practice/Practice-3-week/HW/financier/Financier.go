package financier

type Financier struct {
	position string
	salary   float64
	address  string
}

func NewFinancier(position string, salary float64, address string) *Financier {
	return &Financier{position: position, salary: salary, address: address}
}

func (f *Financier) GetPosition() string {
	return f.position
}

func (f *Financier) SetPosition(position string) {
	f.position = position
}

func (f *Financier) GetSalary() float64 {
	return f.salary
}

func (f *Financier) SetSalary(salary float64) {
	f.salary = salary
}

func (f *Financier) GetAddress() string {
	return f.address
}

func (f *Financier) SetAddress(address string) {
	f.address = address
}

func (f *Financier) AnalyzeFinancialData() string {
	return "Financial data analyzed by Financier"
}

func (f *Financier) MakeInvestmentDecisions() string {
	return "Investment decisions made by Financier"
}

package main

import (
	"employee/consultant"
	"employee/developer"
	"employee/financier"
	"employee/hr"
	"employee/manager"
	"fmt"
)

func main() {
	managerObj := manager.NewManager("Manager", 50000.0, "123 Main St", "HR Department", "+1234567890")
	hrObj := hr.NewHr("HR Specialist", 45000.0, "456 Elm St")
	consultantObj := consultant.NewConsultant("Consultant", 55000.0, "789 Oak St")
	developerObj := developer.NewDeveloper("Developer", 60000.0, "101 Pine St")
	financierObj := financier.NewFinancier("Financier", 70000.0, "222 Maple St")

	fmt.Println("Manager:", managerObj.GetPosition(), "Salary:", managerObj.GetSalary(), "Address:", managerObj.GetAddress())
	fmt.Println("Hr:", hrObj.GetPosition(), "Salary:", hrObj.GetSalary(), "Address:", hrObj.GetAddress())
	fmt.Println("Consultant:", consultantObj.GetPosition(), "Salary:", consultantObj.GetSalary(), "Address:", consultantObj.GetAddress())
	fmt.Println("Developer:", developerObj.GetPosition(), "Salary:", developerObj.GetSalary(), "Address:", developerObj.GetAddress())
	fmt.Println("Financier:", financierObj.GetPosition(), "Salary:", financierObj.GetSalary(), "Address:", financierObj.GetAddress())

	fmt.Println(managerObj.OrganizeMeetings())
	fmt.Println(hrObj.TerminateEmployee("John"))
	fmt.Println(consultantObj.PerformConsultation())
	fmt.Println(developerObj.WriteCode())
	fmt.Println(financierObj.AnalyzeFinancialData())
}

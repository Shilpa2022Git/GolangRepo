package Entity

//import "time"

type Employee struct{
	EmpId  string `json:"EmpId"`
	EmpName string `json:"EmpName"`
	EmpAddress string `json:"EmpAddress"`
	EmpBirthDate string `json:"EmpBirthDate"`
	EmpGender string `json:"EmpGender"`
	EmpSalary string `json:"EmpSalary"`
	DeptNum string `json:"DeptNum"`
}
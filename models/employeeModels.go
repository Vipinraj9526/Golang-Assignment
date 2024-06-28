package models

type Employee struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	County      string `json:"county"`
	Postal      string `json:"postal"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Web         string `json:"web"`
}

type GetEmployeeDetails struct {
	Email string `json:"email" validate:"required,email"`
}

type UpdateEmployeeDetails struct {
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	County      string `json:"county"`
	Postal      string `json:"postal"`
	Phone       string `json:"phone"`
	Web         string `json:"web"`
}

type DeleteEmployeeDetails struct {
	Email string `json:"email" validate:"required,email"`
}

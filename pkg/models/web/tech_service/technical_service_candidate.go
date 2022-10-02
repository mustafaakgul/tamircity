package tech_service

type TechnicalServiceCandidateRequest struct {
	ServiceName         string `json:"service_name"`
	BusinessType        string `json:"business_type"`
	Address             string `json:"address"`
	NumberofBranches    int    `json:"number_of_branches"`
	NumberofTechnicians int    `json:"number_of_technicians"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	Email               string `json:"email"`
	PhoneNumber         string `json:"phone_number"`
}

type EmailRequest struct {
	Email string `json:"email"`
}

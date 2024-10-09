package handler

import "fmt"

type CreateOpeningRequest struct {
	Role        string  `json:"role"`
	Company     string  `json:"company"`
	Location    string  `json:"location"`
	Remote      *bool   `json:"remote"`
	Link        string  `json:"link"`
	Description string  `json:"description"`
	Salary      float64 `json:"salary"`
}

func errorParamIsRequired(name, typ string) error {
	return fmt.Errorf("parameter [%s] of type [%s] is required", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Description == "" && r.Salary == 0 && r.Remote == nil {
		return fmt.Errorf("request body is empty or invalid")
	}

	if r.Role == "" {
		return errorParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errorParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errorParamIsRequired("location", "string")
	}

	if r.Link == "" {
		return errorParamIsRequired("link", "string")
	}

	if r.Description == "" {
		return errorParamIsRequired("description", "string")
	}

	if r.Salary <= 0 {
		return errorParamIsRequired("salary", "float64")
	}

	if r.Remote == nil {
		return errorParamIsRequired("remote", "bool")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role        string  `json:"role"`
	Company     string  `json:"company"`
	Location    string  `json:"location"`
	Remote      *bool   `json:"remote"`
	Link        string  `json:"link"`
	Description string  `json:"description"`
	Salary      float64 `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Description != "" || r.Salary > 0 || r.Remote != nil {
		return nil
	}

	return fmt.Errorf("at least one parameter is required")
}

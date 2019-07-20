package request

// CreateUserRequest struct
type CreateUserRequest struct {
	Fullname    string `json:"fullname,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Role        string `json:"role,omitempty"`
	Active      bool   `json:"active,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
	Email       string `json:"email,omitempty"`
}

// UserLogInRequest struct
type UserLogInRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// GetUserRequest struct
type GetUserRequest struct {
	ParamUserID string `json:"user_id,omitempty"`
}

package request

// CreateUserRequest struct
type CreateUserRequest struct {
	Fullname    string `json:"fullname,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Role        string `json:"role,omitempty"`
	Active      bool   `json:"active,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
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
	ParamUserID string `json:"userID,omitempty"`
}

// GetUsersRequest struct
type GetUsersRequest struct {
	Fullname    string `json:"fullname,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	AlwaysPhone string `json:"alwaysPhone,omitempty"`
	Skip        string `json:"skip,omitempty"`
	Limit       string `json:"limit,omitempty"`
}

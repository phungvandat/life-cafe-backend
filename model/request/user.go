package request

// CreateUserRequest struct
type CreateUserRequest struct {
	User struct {
		Fullname string `json:"fullname,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Role     string `json:"role,omitempty"`
		Active   bool   `json:"active,omitempty"`
	} `json:"user,omitempty"`
}

// UserLogInRequest struct
type UserLogInRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

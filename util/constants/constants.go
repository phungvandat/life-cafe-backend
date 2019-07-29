package constants

// UserRole val
var UserRole = map[string]string{
	"master": "master",
	"admin":  "admin",
	"user":   "user",
}

// OrderType val
var OrderType = map[string]string{
	"import": "import",
	"export": "export",
}

// OrderStatus val
var OrderStatus = map[string]string{
	"created":    "created",
	"delivering": "delivering",
	"done":       "done",
}

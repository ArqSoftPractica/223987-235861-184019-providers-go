package Authorization

type RolePermissions struct {
	Roles map[string]string
}

var rolePermissions = RolePermissions{
	Roles: map[string]string{
		"master":      "MASTER",
		"admin":       "ADMIN",
		"providers":   "PROVIDERS",
		"maintenance": "MAINTENANCE",
		"employee":    "EMPLOYEE",
		"test":        "TEST",
	},
}

type EventMessageType struct {
	EventMessageType map[string]string
}

var eventMessageType = EventMessageType{
	EventMessageType: map[string]string{
		"create": "Create",
		"edit":   "Edit",
		"delete": "Delete",
	},
}

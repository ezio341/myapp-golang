package request

type EditUser struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	RoleID   uint   `json:"role_id"`
}

func (editUser *EditUser) IsValid() bool {
	return editUser.Username != ""
}

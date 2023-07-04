package response

type UserResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Address  string `json:"address"`
}

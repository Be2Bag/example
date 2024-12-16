package dto

// RegisterResponse คือโครงสร้างข้อมูลตอบกลับหลังการลงทะเบียน
type RegisterResponse struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

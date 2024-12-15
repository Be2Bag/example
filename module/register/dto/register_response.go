package dto

// RegisterResponse คือโครงสร้างข้อมูลตอบกลับหลังการลงทะเบียน
type RegisterResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

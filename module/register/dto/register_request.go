package dto

// RegisterRequest คือโครงสร้างข้อมูลคำขอลงทะเบียนผู้ใช้ใหม่
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"` // ชื่อผู้ใช้
	Email    string `json:"email" validate:"required,email"`           // อีเมล
	Password string `json:"password" validate:"required,min=6"`        // รหัสผ่าน
}
package model

import (
	"time"
)

// User คือโครงสร้างข้อมูลผู้ใช้
type User struct {
	UserID    string    `bson:"user_id" json:"user_id"`                 // รหัสผู้ใช้
	Username  string    `bson:"username" json:"username"`               // ชื่อผู้ใช้
	Email     string    `bson:"email" json:"email"`                     // อีเมล
	Password  string    `bson:"password" json:"-"`                      // รหัสผ่าน (ไม่เปิดเผยใน JSON)
	FirstName string    `bson:"first_name" json:"first_name"`           // ชื่อจริง
	LastName  string    `bson:"last_name" json:"last_name"`             // นามสกุล
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"` // วันที่สร้าง
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"` // วันที่แก้ไขล่าสุด
}

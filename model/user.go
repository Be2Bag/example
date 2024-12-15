package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User คือโครงสร้างข้อมูลผู้ใช้
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`                // รหัสผู้ใช้
	Username  string             `bson:"username" json:"username"`               // ชื่อผู้ใช้
	Email     string             `bson:"email" json:"email"`                     // อีเมล
	Password  string             `bson:"password" json:"-"`                      // รหัสผ่าน (ไม่เปิดเผยใน JSON)
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"` // วันที่สร้าง
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"` // วันที่แก้ไขล่าสุด
}

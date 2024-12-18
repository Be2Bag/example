```
GO-EXAMPLE/
├── app/                    # เริ่มการทำงานของแอปพลิเคชัน
├── cmd/                    # Entry points ของแอปพลิเคชัน
│   └── main.go             # Main file ที่ใช้สำหรับรันโปรเจกต์
├── config/                 # เก็บไฟล์การตั้งค่า เช่น config.yaml
├── middleware/             # Middleware ที่ใช้ร่วมกันในแอปพลิเคชัน
├── model/                  # โครงสร้างข้อมูล (struct) ของแอปพลิเคชัน
├── module/
│   ├── register/           # โมดูลสำหรับฟังก์ชันการลงทะเบียน
│   │   ├── dto/            # Data Transfer Objects สำหรับส่งและรับข้อมูล
│   │   ├── handler/        # Handlers สำหรับการจัดการ HTTP request
│   │   ├── middleware/     # Middleware เฉพาะของโมดูล register
│   │   ├── ports/          # Interfaces ของบริการที่ต้องใช้ในโมดูลนี้
│   │   └── services/       # Business logic ของโมดูล register
│   └── session/            # โมดูลสำหรับฟังก์ชันการจัดการ session
│       ├── dto/            # Data Transfer Objects สำหรับส่งและรับข้อมูล
│       ├── handler/        # Handlers สำหรับการจัดการ HTTP request
│       ├── middleware/     # Middleware เฉพาะของโมดูล session
│       ├── ports/          # Interfaces ของบริการที่ต้องใช้ในโมดูลนี้
│       └── services/       # Business logic ของโมดูล session
└── pkg/                    # Utilities, helper functions, และ library ที่ใช้ร่วมกัน
```

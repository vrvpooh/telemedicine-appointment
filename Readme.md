# 🏥 Telemedicine Appointment System

ระบบสำหรับจัดการนัดหมายแพทย์ (Telemedicine) พัฒนาโดยใช้ Go + SQLite พร้อมรองรับการรันผ่าน Docker

---

## 📌 Features

- ระบบจัดการผู้ใช้
- ระบบค้นหาและข้อมูลแพทย์
- ระบบจัดการตารางเวลา
- การจองและออกตั๋ว
- ระบบบันทึกผลตรวจ
- ระบบแจ้งเตือนและประเมินผล

---

## 🛠 Tech Stack

- **Backend:** Go (Golang)
- **Framework:** Gin
- **Database:** SQLite3
- **API Testing:** Postman
- **Container:** Docker
- **Version Control:** Git / GitHub

---

## 📂 Project Structure
```bash
telemedicine-appointment/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── config/
│   └── database.go              # Database connection setup
├── handler/                     # HTTP request/response handlers
├── middleware/                  # Authentication middleware
├── model/                       # Data models / structs
├── repository/                  # Database access layer
├── route/                       # API route definitions
├── service/                     # Business logic layer
├── postman_collection/          # Postman API test collection
├── Dockerfile                   # Docker image configuration
├── docker-compose.yml           # Docker Compose configuration
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ วิธีติดตั้งและรันแบบ Local

### 1. Clone โปรเจกต์

```bash
git clone https://github.com/vrvpooh/telemedicine-appointment.git
cd telemedicine-appointment
```

### 2. ติดตั้ง dependencies
```bash
go mod tidy
```

### 3. รันโปรแกรม
```bash
go run cmd/server/main.go
```

### 4. เข้าใช้งานผ่าน API
```bash
http://localhost:8080
```

---

## 🐳 วิธีรันด้วย Docker

### สำหรับ Windows
```bash
docker run -p 8080:8080 voravee/telemedicine-api
```

### สำหรับ Mac
```bash
docker run --platform linux/amd64 -p 8080:8080 voravee/telemedicine-api
```

--- 

## 🐳 วิธีรันด้วย Docker Compose

```bash
docker compose up --build
```

เมื่อต้องการหยุด container
```bash
docker compose down
```

--- 

## 🧪 วิธีการทดสอบ Unit Test

### รัน Unit Test ทั้งหมด
```bash
go test ./...
```

### Run Unit Test เฉพาะส่วน service (Business logic)
```bash
go test ./service/... -cover
```

### สร้างไฟล์ Coverage
```bash
go test ./... -coverprofile=coverage.out
```

### แสดงผล Coverage ผ่านเว็บเบราว์เซอร์
```bash
go tool cover -html=coverage.out
```

---

## 👥 สมาชิกและความรับผิดชอบ

### 1) นาย กษิดิศ คงประพันธ์ (6609650129) 
- ทำเรื่อง: ระบบแจ้งเตือนเเละประเมินผล

```bash
GET   /api/notification
POST  /api/feedback
PATCH /api/admin/verify-doctor/{id}
```

### 2) นาย กฤตชญา กลิ่นเดช (6609650160)
- ทำเรื่อง: ระบบจัดการผู้ใช้

```bash
POST /api/auth/register
POST /api/auth/login
GET  /api/users/me
```

### 3) นาย ณัฐนันท์ ชูดวง (6609650368)
- ทำเรื่อง: ระบบบันทึกผลตรวจ

```bash
POST /api/records
GET  /api/records/patient/me
GET  /api/records/{id}
```

### 4) นาย นนท์นริฐ รามณีย์กุลธวัช (6609650459)
- ทำเรื่อง: การจองและออกตั๋ว

```bash
POST  /api/appointments
GET   /api/appointments/{id}/zoom-token
PATCH /api/appointments/{id}/status
```

### 5) นาย วรวีร์ นุชธิสาร (6609650640)
- ทำเรื่อง: ระบบค้นหาและข้อมูลแพทย์

```bash
GET /api/doctors
GET /api/doctors/{id}
GET /api/specialties
```

### 6) นาย ศุภกฤต ธรรมางกูร (6609650657)
- ทำเรื่อง: ระบบจัดการตารางเวลา

```bash
POST   /api/doctor/{id}/slots
GET    /api/doctor/{id}/slots
DELETE /api/slots/{id}
```

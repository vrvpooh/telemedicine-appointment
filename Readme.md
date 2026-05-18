# 🏥 Telemedicine Appointment System

ระบบสำหรับจัดการนัดหมายแพทย์ (Telemedicine) พัฒนาโดยใช้ Go + SQLite พร้อมรองรับการรันผ่าน Docker

---

## 📌 Features

- ระบบจัดการผู้ใช้
- ระบบค้นหาเเละข้อมูลเเพทย์
- ระบบจัดการตารางเวลา
- การจองเเละออกตั๋ว
- ระบบบันทึกผลตรวจ
- ระบบเเจ้งเตือนเเละประเมินผล



---

## 🛠 Tech Stack

- **Backend:** Go (Golang)
- **Database:** SQLite3
- **Container:** Docker

---

## 📂 Project Structure
```bash
telemedicine-appointment/
│── cmd/server
    └── main.go
│── database/
│── repository/
│── handler/
│── model/
│── Dockerfile
│── go.mod
│── README.md
```

---

## ⚙️ วิธีติดตั้งและรัน (Local)

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

### 4. เข้าใช้งาน
```bash
http://localhost:8080
```

### 🐳 วิธีใช้งาน Docker
## 1. Build Docker Image
```bash
docker build -t telemedicine-api .
```
## 2. รัน Container

👉 สำหรับ PowerShell (Windows)
```bash
docker run -p 8080:8080 -v ${PWD}/data:/root/data telemedicine-api
```

👉 สำหรับ Mac / Linux
```bash
docker run -p 8080:8080 -v $(pwd)/data:/root/data telemedicine-api
```


package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() {
	db, err := sql.Open("sqlite3", "./telemedicine.db")
	if err != nil {
		log.Fatal("เปิด database ไม่ได้:", err)
	}

	// 	ให้เพื่อนๆเพิ่ม database ต่อตรงนี้เลยนะ ว่าตัวเองต้องใช้อะไรกันบ้าง
	schema := `
	CREATE TABLE IF NOT EXISTS slots (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		doctor_id  INTEGER NOT NULL,
		start_time TEXT    NOT NULL,
		end_time   TEXT    NOT NULL,
		is_booked  INTEGER NOT NULL DEFAULT 0
	);
	`

	if _, err := db.Exec(schema); err != nil {
		log.Fatal("สร้างตาราง slots ไม่สำเร็จ:", err)
	}

	DB = db
	log.Println("เชื่อมต่อ SQLite สำเร็จ -> telemedicine.db")
}

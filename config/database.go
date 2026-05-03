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

	schemaDoctors := `
	CREATE TABLE IF NOT EXISTS doctors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		specialty TEXT NOT NULL
	);
	`

	if _, err := db.Exec(schemaDoctors); err != nil {
		log.Fatal("สร้างตาราง doctors ไม่สำเร็จ:", err)
	}

	insertMockDoctors := `
	INSERT INTO doctors (name, specialty)
	SELECT 'Dr. Smith', 'Cardiology'
	WHERE NOT EXISTS (SELECT 1 FROM doctors);

	INSERT INTO doctors (name, specialty)
	SELECT 'Dr. John', 'Dermatology'
	WHERE NOT EXISTS (SELECT 1 FROM doctors WHERE name='Dr. John');

	INSERT INTO doctors (name, specialty)
	SELECT 'Dr. Lee', 'Neurology'
	WHERE NOT EXISTS (SELECT 1 FROM doctors WHERE name='Dr. Lee');
	`

	if _, err := db.Exec(insertMockDoctors); err != nil {
		log.Println("insert mock data error:", err)
	}

	DB = db
	log.Println("เชื่อมต่อ SQLite สำเร็จ -> telemedicine.db")
}

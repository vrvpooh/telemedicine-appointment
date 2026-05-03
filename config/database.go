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
		specialty TEXT NOT NULL,
		education TEXT,
		experience INTEGER,
		rating REAL
	);
	`

	if _, err := db.Exec(schemaDoctors); err != nil {
		log.Fatal("สร้างตาราง doctors ไม่สำเร็จ:", err)
	}

	insertMockDoctors := `
	INSERT INTO doctors (name, specialty, education, experience, rating)
	SELECT 'Dr. Smith', 'Cardiology', 'Harvard Medical School', 10, 4.8
	WHERE NOT EXISTS (SELECT 1 FROM doctors WHERE name='Dr. Smith');

	INSERT INTO doctors (name, specialty, education, experience, rating)
	SELECT 'Dr. John', 'Dermatology', 'Stanford University', 5, 4.5
	WHERE NOT EXISTS (SELECT 1 FROM doctors WHERE name='Dr. John');

	INSERT INTO doctors (name, specialty, education, experience, rating)
	SELECT 'Dr. Lee', 'Psychiatry', 'Chulalongkorn University', 8, 4.7
	WHERE NOT EXISTS (SELECT 1 FROM doctors WHERE name='Dr. Lee');
	`

	if _, err := db.Exec(insertMockDoctors); err != nil {
		log.Println("insert mock data error:", err)
	}

	DB = db
	log.Println("เชื่อมต่อ SQLite สำเร็จ -> telemedicine.db")
}

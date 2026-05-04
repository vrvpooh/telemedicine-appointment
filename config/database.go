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

	// ให้เพื่อนๆเพิ่ม database ต่อตรงนี้เลยนะ ว่าตัวเองต้องใช้อะไรกันบ้าง

  // slot (ten)
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

  // doctor (pooh)
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
	// Notification & Feedback (Tum)
    schemaExtra := `
    CREATE TABLE IF NOT EXISTS notifications (
        id          INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id     INTEGER NOT NULL,
        message     TEXT NOT NULL,
        is_read     INTEGER DEFAULT 0, -- 0=ยังไม่อ่าน, 1=อ่านแล้ว
        created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS feedbacks (
        id          INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id     INTEGER NOT NULL,
        doctor_id   INTEGER NOT NULL,
        rating      REAL NOT NULL,     -- คะแนน 1.0 - 5.0
        comment     TEXT,
        created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `
	if _, err := db.Exec(schemaExtra); err != nil {
        log.Fatal("สร้างตาราง Notification/Feedback ไม่สำเร็จ:", err)
    }

    _, _ = db.Exec("ALTER TABLE doctors ADD COLUMN is_verified INTEGER DEFAULT 0;")

    insertMockExtra := `
    INSERT INTO notifications (user_id, message) 
    SELECT 1, 'คุณมีการนัดหมายกับ Dr. Smith ในอีก 15 นาที'
    WHERE NOT EXISTS (SELECT 1 FROM notifications WHERE message LIKE '%Dr. Smith%');
    `
    if _, err := db.Exec(insertMockExtra); err != nil {
        log.Println("insert mock extra data error:", err)
    }

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
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

	// record (Pook)
	schemaRecords := `
	CREATE TABLE IF NOT EXISTS medical_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		appointment_id INTEGER NOT NULL,
		patient_id INTEGER NOT NULL,
		doctor_id INTEGER NOT NULL,
		symptoms TEXT,
		diagnosis TEXT,
		prescription TEXT,
		notes TEXT,
		created_at TEXT NOT NULL
	);
	`
	if _, err := db.Exec(schemaRecords); err != nil {
		log.Fatal("สร้างตาราง medical_records ไม่สำเร็จ:", err)
	}

	// appointment (Money)
	schemaAppointments := `
    CREATE TABLE IF NOT EXISTS appointments (
        id          INTEGER PRIMARY KEY AUTOINCREMENT,
        patient_id  INTEGER NOT NULL,
        slot_id     INTEGER NOT NULL UNIQUE,
        status      TEXT DEFAULT 'pending',
        zoom_token  TEXT,
        created_at  TEXT NOT NULL,
        updated_at  TEXT NOT NULL
    );
    `
	if _, err := db.Exec(schemaAppointments); err != nil {
		log.Fatal("สร้างตาราง appointments ไม่สำเร็จ:", err)
	}

	// สร้าง Unique Index เพื่อป้องกันการจองซ้ำสำหรับตารางที่มีอยู่แล้ว
	_, _ = db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_appointments_slot_id ON appointments(slot_id);")

	// Ensure updated_at exists (in case table was created before the schema update)
	if _, err := db.Exec("ALTER TABLE appointments ADD COLUMN updated_at DATETIME DEFAULT CURRENT_TIMESTAMP;"); err != nil {
		log.Println("Note: ALTER TABLE appointments (updated_at) might have already run or failed:", err)
	} else {
		log.Println("Successfully added updated_at column to appointments table.")
	}

	// users Authentication (Beer)
	schemaUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(schemaUsers); err != nil {
		log.Fatal("สร้างตาราง users ไม่สำเร็จ:", err)
	}

	insertMockRecords := `
	INSERT INTO medical_records (appointment_id, patient_id, doctor_id, symptoms, diagnosis, prescription, notes, created_at)
	SELECT 1, 101, 1, 'ปวดหัวตึบๆ มีไข้ 38 องศา', 'ไข้หวัดทั่วไป', 'Paracetamol 500mg กินทุก 4-6 ชั่วโมง', 'ดื่มน้ำอุ่นเยอะๆ และพักผ่อนให้เพียงพอ', '2026-05-01 10:00:00'
	WHERE NOT EXISTS (SELECT 1 FROM medical_records WHERE appointment_id=1);

	INSERT INTO medical_records (appointment_id, patient_id, doctor_id, symptoms, diagnosis, prescription, notes, created_at)
	SELECT 2, 101, 2, 'มีผื่นแดงที่แขนและคันมาก', 'ภูมิแพ้ผิวหนัง', 'Loratadine 10mg วันละ 1 เม็ด, ยาทา TA Cream', 'หลีกเลี่ยงการเกาและงดอาหารทะเลชั่วคราว', '2026-05-03 14:30:00'
	WHERE NOT EXISTS (SELECT 1 FROM medical_records WHERE appointment_id=2);

	INSERT INTO medical_records (appointment_id, patient_id, doctor_id, symptoms, diagnosis, prescription, notes, created_at)
	SELECT 3, 102, 1, 'เจ็บหน้าอกเวลาหายใจลึกๆ', 'กล้ามเนื้อหน้าอกอักเสบ', 'Ibuprofen 400mg หลังอาหารทันที', 'งดการยกของหนัก 1 สัปดาห์', '2026-05-04 09:15:00'
	WHERE NOT EXISTS (SELECT 1 FROM medical_records WHERE appointment_id=3);
	`
	if _, err := db.Exec(insertMockRecords); err != nil {
		log.Println("insert mock data (records) error:", err)
	}

	DB = db
	log.Println("เชื่อมต่อ SQLite สำเร็จ -> telemedicine.db")
}

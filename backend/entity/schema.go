package entity

import (
	"time"

	"gorm.io/gorm"
)

// แพทย์ เดี๋ยวจะต้องไปใช้ Enntity ของเพื่อน
type Doctor struct {
	gorm.Model
	Name      string
	Expertise string
	Email     string `gorm:"uniqueIndex"`
	Password  string `json:"-"`
	//แพทย์ 1 คน สามารถบันทึกข้อมูลการรักษาได้หลายคน
	Treatments []Treatment `gorm:"foreignKey:DoctorID"`
}

// โรค Disease
type Disease struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//1 โรค มีผู้ป่วยหลายคน
	Patients []Patient `gorm:"foreignKey:DiseaseID"`
}

// สถานะการรักษา Status
type Status struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//1 สถานะ มีผู้ป่วยหลายคน
	Patients []Patient `gorm:"foreignKey:StatusID"`
}

// สถานะติดตามผล Status
type Track struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//1 สถานะติดตามผล มีผู้ป่วยหลายคน
	Patients []Patient `gorm:"foreignKey:TrackID"`
}

// ผู้ป่วย Patient เดี๋ยวจะต้องไปใช้ Enntity ของเพื่อน
type Patient struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//Email string `gorm:"uniqueIndex"`
	// TeacherID ทำหน้าที่เป็น FK
	DiseaseID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Disease Disease `gorm:"references:id"`

	StatusID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Status Status `gorm:"references:id"`

	TrackID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Track Track `gorm:"references:id"`
}

// การรักษา
type Treatment struct {
	gorm.Model
	TREATMENT_ID string
	TREATMENT    string
	DATE         time.Time
	APPOINTMENT  string
	CONCLUSION   string
	GUIDANCE     string

	DoctorID *uint
	Doctor   Doctor `gorm:"references:id"`

	StatusID *uint
	Status   Status `gorm:"references:id"`

	TrackID *uint
	Track   Track `gorm:"references:id"`

	PatientID *uint
	Patient   Patient `gorm:"references:id"`

	DiseaseID *uint
	Disease   Disease `gorm:"references:id"`
}

//ระบบการเบิกอุปกรณ์เครื่องมือแลป Gg
//Medicine เทคนิคการแพทย์ เดี๋ยวจะต้องไปใช้ Enntity ของเพื่อน
type Medicine struct {
	gorm.Model
	Name      string
	Expertise string
	Email     string `gorm:"uniqueIndex"`
	Password  string `json:"-"`
	//เทคนิคการแพทย์ 1 คน สามารถบันทึกข้อมูลการเบิกได้หลายคำร้อง
	Requests []Request `gorm:"foreignKey:MedicineID"`
}
//อุปกรณ์
type Equipment struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//1 อุปกรณ์ มีหลายการเบิก
	Requests []Request `gorm:"foreignKey:EquipmentID"`
}
//สถานที่
type Location struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//1 สถสนที่ มีหลายการเบิก
	Requests []Request `gorm:"foreignKey:LocationID"`
}
//ตารางหลัก การเบิก
type Request struct {
	gorm.Model
	R_ID string
	R_NAME string
	QUANTITY    string //int
	TIME         time.Time
	
	MedicineID *uint
	Medicine   Medicine `gorm:"references:id"`

	EquipmentID *uint
	Equipment   Equipment `gorm:"references:id"`

	LocationID *uint
	Location   Location `gorm:"references:id"`
}


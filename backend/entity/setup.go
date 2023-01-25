package entity

import (
	//"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("se-65-g08.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrete the schema
	database.AutoMigrate(
		&Doctor{},
		&Disease{},
		&Status{},
		&Track{},
		&Patient{},
		&Treatment{},
		//
		&Location{},
		&Medicine{},
		&Equipment{},
		&Request{},
	)

	db = database

	//แพทย์
	//password, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	//password2, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	//เทคนิคการแพทย์
	//password3, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	//password4, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)

	db.Model(&Doctor{}).Create(&Doctor{
		Name:     "Khunjira",
		Email:    "khunjira@gmail.com",
		Password: "1234",
	})
	db.Model(&Doctor{}).Create(&Doctor{
		Name:     "Gg",
		Email:    "gg@gmail.com",
		Password:  "1234",
	})

	//
	db.Model(&Medicine{}).Create(&Medicine{
		Name:     "Nati",
		Email:    "nati@gmail.com",
		Password:  "1234",
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Name:     "non",
		Email:    "non@gmail.com",
		Password:  "1234",
	})

	var Khunjira Doctor
	var Gg Doctor
	db.Raw("SELECT * FROM users WHERE email = ?", "khunjira@gmail.com").Scan(&Khunjira)
	db.Raw("SELECT * FROM users WHERE email = ?", "gg@gmail.com").Scan(&Gg)
	//
	var Nati Doctor
	var non Doctor
	db.Raw("SELECT * FROM users WHERE email = ?", "nati@gmail.com").Scan(&Nati)
	db.Raw("SELECT * FROM users WHERE email = ?", "non@gmail.com").Scan(&non)

	// Disease Data
	D1 := Disease{
		Name: "โรคความดันโลหิตสูง",
	}
	db.Model(&Disease{}).Create(&D1)

	D2 := Disease{
		Name: "ออฟฟิศซินโดรม",
	}
	db.Model(&Disease{}).Create(&D2)

	D3 := Disease{
		Name: "โรคเครียดลงกระเพาะ",
	}
	db.Model(&Disease{}).Create(&D3)

	D4 := Disease{
		Name: "โรคหัวใจ",
	}
	db.Model(&Disease{}).Create(&D4)

	D5 := Disease{
		Name: "โรคเบาหวาน",
	}
	db.Model(&Disease{}).Create(&D5)

	D6 := Disease{
		Name: "โรคหลอดเลือดหัวใจตีบตัน",
	}
	db.Model(&Disease{}).Create(&D6)

	D7 := Disease{
		Name: "โรคหลอดเลือดสมอง",
	}
	db.Model(&Disease{}).Create(&D7)

	//  Status  Data
	S1 := Status{
		Name: "เข้ารับการรักษา",
	}
	db.Model(&Status{}).Create(&S1)

	S2 := Status{
		Name: "ส่งตรวจแลป",
	}
	db.Model(&Status{}).Create(&S2)

	S3 := Status{
		Name: "นอนโรงพยาบาล ติดตามผลการรักษา",
	}
	db.Model(&Status{}).Create(&S3)

	S4 := Status{
		Name: "รักษาหายแล้ว",
	}
	db.Model(&Status{}).Create(&S4)

	S5 := Status{
		Name: "ปฏิเสธการรักษา",
	}
	db.Model(&Status{}).Create(&S5)

	//  Track  Data
	T1 := Track{
		Name: "นัด",
	}
	db.Model(&Track{}).Create(&T1)

	T2 := Track{
		Name: "จ่ายยา",
	}
	db.Model(&Track{}).Create(&T2)

	T3 := Track{
		Name: "นัด และจ่ายยา",
	}
	db.Model(&Track{}).Create(&T3)

	T4 := Track{
		Name: "ไม่นัด",
	}
	db.Model(&Track{}).Create(&T4)

	T5 := Track{
		Name: "ไม่จ่ายยา",
	}
	db.Model(&Track{}).Create(&T5)

	// Patient Data
	P001 := Patient{
		Name:    "สมชาย ทันสมัย",
		Disease: D5,
	}
	db.Model(&Patient{}).Create(&P001)

	P002 := Patient{
		Name:    "สมหญิง จุลทล",
		Disease: D1,
	}
	db.Model(&Patient{}).Create(&P002)

	P003 := Patient{
		Name:    "มาสาย ประจำ",
		Disease: D6,
	}
	db.Model(&Patient{}).Create(&P003)

	////ระบบเบิกอุปกรณ์
	//  Equipment  Data
	E1 := Equipment{
		Name: "กระจกสไลด์แบบใส",
	}
	db.Model(&Equipment{}).Create(&E1)

	E2 := Equipment{
		Name: "กระจกสไลด์แบขุ่น",
	}
	db.Model(&Equipment{}).Create(&E2)

	E3 := Equipment{
		Name: "CoverGlass",
	}
	db.Model(&Equipment{}).Create(&E3)

	E4 := Equipment{
		Name: "กระปุกเก็บสิ่งตรวจ 2oz",
	}
	db.Model(&Equipment{}).Create(&E4)

	E5 := Equipment{
		Name: "ขวดแก้วรีเอเย่น",
	}
	db.Model(&Equipment{}).Create(&E5)
	E6 := Equipment{
		Name: "หลอดหยดยา",
	}
	db.Model(&Equipment{}).Create(&E6)

	E7 := Equipment{
		Name: "หน้ากาก",
	}
	db.Model(&Equipment{}).Create(&E7)

	E8 := Equipment{
		Name: "หมวกคลุมผม",
	}
	db.Model(&Equipment{}).Create(&E8)

	E9 := Equipment{
		Name: "หลอดเก็บเลือด",
	}
	db.Model(&Equipment{}).Create(&E9)

	E10 := Equipment{
		Name: "ถุงคลุมรองเท้า",
	}
	db.Model(&Equipment{}).Create(&E10)

	E11 := Equipment{
		Name: "ขวดแก้ว",
	}
	db.Model(&Equipment{}).Create(&E11)

	E12 := Equipment{
		Name: "ขี้ผึ้งอุดหลอดเก็บเลือด",
	}
	db.Model(&Equipment{}).Create(&E12)
	//สถานที่
	L1 := Location{
		Name: "อาคารอุบัติเหตุและเวชศาสตร์ฉุกเฉิน",
	}
	db.Model(&Location{}).Create(&L1)
	L2 := Location{
		Name: "อาคารเทคนิคการแทพย์",
	}
	db.Model(&Location{}).Create(&L2)
	L3 := Location{
		Name: "อาคารวิจัย",
	}
	db.Model(&Location{}).Create(&L3)
	L4 := Location{
		Name: "อาคารอายุรกรรม",
	}
	db.Model(&Location{}).Create(&L4)
	L5 := Location{
		Name: "อาคารเวชศาสตร์ฟื้นฟู",
	}
	db.Model(&Location{}).Create(&L5)

}

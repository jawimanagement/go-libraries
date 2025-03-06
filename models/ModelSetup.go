package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"os"

	"gorm.io/gorm/logger"

	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

var ActiveUser string
var DbConnection *gorm.DB

var OpenDB *gorm.DB

func mysqlConfig(configDbMaster string) *mysql.Config {

	config := &mysql.Config{
		DSN:                       configDbMaster,
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	return config
}

func JawiConnect() (*sql.DB, *gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("error loading .env file")
	}
	// mysql connection
	// dsn := "root:pass@tcp(127.0.0.1:3306)/dewan?charset=utf8mb3&parseTime=True&loc=Asia/Jakarta"
	configDbMaster := os.Getenv("jawiDsn")
	sqlDB, err := sql.Open("mysql", configDbMaster)
	if err != nil {
		return nil, nil, fmt.Errorf("%s", err)
	}

	config := mysqlConfig(configDbMaster)

	dbMaster, err := gorm.Open(mysql.New(*config), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		QueryFields: true,
		NowFunc: func() time.Time {
			loc, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(loc)
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("jawi database connection error : %s", err)
	}
	// prevent global updates
	dbMaster.Session(&gorm.Session{AllowGlobalUpdate: false})

	sqlDB.SetMaxIdleConns(1)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(3)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(15 * time.Second)
	// OpenDB = dbMaster
	return sqlDB, dbMaster, nil
}

func JawiLiveConnect() (*sql.DB, *gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("error loading .env file")
	}
	//mysql connection
	// dsn := "root:pass@tcp(127.0.0.1:3306)/dewan?charset=utf8mb3&parseTime=True&loc=Asia/Jakarta"
	configDbMaster := os.Getenv("jawiLiveDsn")
	sqlDB, err := sql.Open("mysql", configDbMaster)
	if err != nil {
		return nil, nil, fmt.Errorf("%s", err)
	}

	config := mysqlConfig(configDbMaster)

	dbMaster, err := gorm.Open(mysql.New(*config), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		QueryFields: true,
		NowFunc: func() time.Time {
			loc, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(loc)
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("jawi live database connection error : %s", err)
	}
	// prevent global updates
	dbMaster.Session(&gorm.Session{AllowGlobalUpdate: false})

	sqlDB.SetMaxIdleConns(1)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(3)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(15 * time.Second)
	// OpenDB = dbMaster

	return sqlDB, dbMaster, nil
}

// parse null string on model
type NullString struct {
	sql.NullString
}

// parse null string on model
func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return []byte(`null`), nil
}

func NullStringInput(s string) NullString {
	if len(s) == 0 {
		return NullString{sql.NullString{String: `null`, Valid: false}}
	}
	return NullString{sql.NullString{String: s, Valid: true}}
}

// parse null int on model
type NullInt64 struct {
	sql.NullInt64
}

// parse null int on model
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return []byte(`null`), nil
}

func NullInt64Input(s int64) NullInt64 {
	if s == 0 {
		return NullInt64{sql.NullInt64{Int64: 0, Valid: false}}
	}
	return NullInt64{sql.NullInt64{Int64: s, Valid: true}}
}

// parse null time on model
type NullDateTime struct {
	sql.NullTime
}

// parse null time on model
func (nt NullDateTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		t := nt.Time
		return json.Marshal(t.Format("2006-01-02 15:04:05"))
	}
	return []byte(`null`), nil
}

func NullDateTimeInput(t time.Time) NullDateTime {
	if t.IsZero() {
		return NullDateTime{sql.NullTime{Time: time.Time{}, Valid: false}}
	}
	return NullDateTime{sql.NullTime{Time: t, Valid: true}}
}

// date
type NullDate struct {
	sql.NullTime
}

func (nt NullDate) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		t := nt.Time
		return json.Marshal(t.Format("2006-01-02"))
	}
	return []byte(`null`), nil
}

func NullDateInput(t time.Time) NullDate {
	if t.IsZero() {
		return NullDate{sql.NullTime{Time: time.Time{}, Valid: false}}
	}
	return NullDate{sql.NullTime{Time: t, Valid: true}}
}

// get Columns Name from Model
func GetColumnName(db *gorm.DB, model []interface{}) *gorm.DB {
	for _, v := range model {
		fmt.Println(v)
	}
	return db
}

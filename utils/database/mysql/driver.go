package mysql

import (
	"fmt"
	"log"

	"aziz-wahyudin/technical-test-k-style-hub/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Member{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&ReviewProdcut{})
	db.AutoMigrate(&LikeReview{})
}

type Member struct {
	gorm.Model
	Username       string
	Gender         string
	Skintype       string
	Skincolor      string
	ReviewProdcuts []ReviewProdcut
	LikeReviews    []LikeReview
}

type Product struct {
	gorm.Model
	NameProduct    string
	Price          string
	ReviewProdcuts []ReviewProdcut
}

type ReviewProdcut struct {
	gorm.Model
	DescReview  string
	MemberID    uint
	LikeReviews []LikeReview
}

type LikeReview struct {
	gorm.Model
	Likes           int
	Dislikes        int
	MemberID        uint
	ReviewProdcutID uint
}

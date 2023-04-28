package repository

import (
	"aziz-wahyudin/technical-test-k-style-hub/features/product"

	"gorm.io/gorm"
)

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
	Likes           int           `gorm:"default:0"`
	Dislikes        int           `gorm:"default:0"`
	MemberID        uint          `gorm:"primaryKey"`
	ReviewProdcutID uint          `gorm:"primaryKey"`
	Review          ReviewProdcut `gorm:"foreignKey:ReviewID"`
	Member          Member        `gorm:"foreignKey:MemberID"`
}

func (dataModel *Product) ToCore() product.Product {
	return product.Product{}
}

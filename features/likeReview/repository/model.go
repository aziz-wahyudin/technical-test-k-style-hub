package repository

import (
	likereview "aziz-wahyudin/technical-test-k-style-hub/features/likeReview"

	"gorm.io/gorm"
)

type LikeReview struct {
	gorm.Model
	Likes           int           `gorm:"default:0"`
	Dislikes        int           `gorm:"default:0"`
	MemberID        uint          `gorm:"primaryKey"`
	ReviewProdcutID uint          `gorm:"primaryKey"`
	Review          ReviewProdcut `gorm:"foreignKey:ReviewID"`
	Member          Member        `gorm:"foreignKey:MemberID"`
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

type ReviewProdcut struct {
	gorm.Model
	DescReview  string
	MemberID    uint
	LikeReviews []LikeReview
}

func FromCore(dataCore likereview.LikeReview) LikeReview {
	likeReviewGorm := LikeReview{
		Likes:           dataCore.Likes,
		Dislikes:        dataCore.Dislikes,
		MemberID:        dataCore.MemberID,
		ReviewProdcutID: dataCore.ReviewProdcutID,
	}
	return likeReviewGorm
}

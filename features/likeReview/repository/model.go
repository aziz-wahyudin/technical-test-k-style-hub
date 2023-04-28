package repository

import (
	likereview "aziz-wahyudin/technical-test-k-style-hub/features/likeReview"

	"gorm.io/gorm"
)

type LikeReview struct {
	gorm.Model
	Likes           int `gorm:"default:0"`
	Dislikes        int `gorm:"default:0"`
	MemberID        uint
	ReviewProdcutID uint
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

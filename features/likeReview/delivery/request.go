package delivery

import (
	likereview "aziz-wahyudin/technical-test-k-style-hub/features/likeReview"
)

type LikeReviewReq struct {
	Likes           int  `json:"likes" form:"likes"`
	Dislikes        int  `json:"dislikes" form:"dislikes"`
	MemberID        uint `json:"member_id" form:"member_id"`
	ReviewProdcutID uint `json:"review_product_id" form:"review_product_id"`
}

func (data *LikeReviewReq) reqToCore() likereview.LikeReview {
	return likereview.LikeReview{
		Likes:           data.Likes,
		Dislikes:        data.Dislikes,
		MemberID:        data.MemberID,
		ReviewProdcutID: data.ReviewProdcutID,
	}
}

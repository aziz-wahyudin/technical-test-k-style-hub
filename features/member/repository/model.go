package repository

import (
	"aziz-wahyudin/technical-test-k-style-hub/features/member"

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

func FromCore(dataCore member.Member) Member {
	memberGorm := Member{
		Username:  dataCore.Username,
		Gender:    dataCore.Gender,
		Skintype:  dataCore.Skintype,
		Skincolor: dataCore.Skincolor,
	}
	return memberGorm
}

func (dataModel *Member) ToCore() member.Member {
	return member.Member{
		ID:        dataModel.ID,
		Username:  dataModel.Username,
		Gender:    dataModel.Gender,
		Skintype:  dataModel.Skintype,
		Skincolor: dataModel.Skincolor,
	}
}

func ToCoreList(dataModel []Member) []member.Member {
	var dataCore []member.Member
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ToCore())
	}
	return dataCore
}

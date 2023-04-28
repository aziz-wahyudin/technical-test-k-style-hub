package repository

import (
	likereview "aziz-wahyudin/technical-test-k-style-hub/features/likeReview"
	"errors"

	"gorm.io/gorm"
)

type likeReviewRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) likereview.RepositoryInterface {
	return &likeReviewRepository{
		db: db,
	}
}

// Create implements likereview.RepositoryInterface
func (r *likeReviewRepository) Create(input likereview.LikeReview) (row int, err error) {
	likeGorm := FromCore(input)
	tx := r.db.Create(&likeGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, error query")
	}
	return int(tx.RowsAffected), nil
}

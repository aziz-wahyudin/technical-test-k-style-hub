package repository

import (
	"aziz-wahyudin/technical-test-k-style-hub/features/member"
	"errors"

	"gorm.io/gorm"
)

type memberRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) member.RepositoryInterface {
	return &memberRepository{
		db: db,
	}
}

// Create implements member.RepositoryInterface
func (r *memberRepository) Create(input member.Member) (row int, err error) {
	memberGorm := FromCore(input)
	tx := r.db.Create(&memberGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed, error query")
	}
	return int(tx.RowsAffected), nil
}

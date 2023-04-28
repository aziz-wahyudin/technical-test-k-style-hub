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

// Update implements member.RepositoryInterface
func (r *memberRepository) Update(input member.Member, id int) (err error) {
	memberGorm := FromCore(input)
	tx := r.db.Where("id= ?", id).Updates(memberGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed, error query")
	}
	return nil
}

// Delete implements member.RepositoryInterface
func (r *memberRepository) Delete(id int) (row int, err error) {
	member := Member{}

	tx := r.db.Delete(&member, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete by id failed, error query")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements member.RepositoryInterface
func (r *memberRepository) GetAll() (data []member.Member, err error) {
	var members []Member

	tx := r.db.Find(&members)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := ToCoreList(members)
	return dataCore, nil
}

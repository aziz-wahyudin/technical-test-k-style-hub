package service

import (
	likereview "aziz-wahyudin/technical-test-k-style-hub/features/likeReview"
)

type likeReviewService struct {
	likeReviewRepository likereview.RepositoryInterface
}

func New(repo likereview.RepositoryInterface) likereview.ServiceInterface {
	return &likeReviewService{
		likeReviewRepository: repo,
	}
}

// Create implements likereview.ServiceInterface
func (s *likeReviewService) Create(input likereview.LikeReview) (err error) {
	_, err = s.likeReviewRepository.Create(input)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements likereview.ServiceInterface
func (s *likeReviewService) Delete(idMember int, idReview int) (err error) {
	_, err = s.likeReviewRepository.Delete(idMember, idReview)
	if err != nil {
		return err
	}
	return nil
}

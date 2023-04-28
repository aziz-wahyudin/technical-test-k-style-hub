package factory

import (
	memberDelivery "aziz-wahyudin/technical-test-k-style-hub/features/member/delivery"
	memberRepo "aziz-wahyudin/technical-test-k-style-hub/features/member/repository"
	memberService "aziz-wahyudin/technical-test-k-style-hub/features/member/service"

	likereviewDelivery "aziz-wahyudin/technical-test-k-style-hub/features/likeReview/delivery"
	likereviewRepo "aziz-wahyudin/technical-test-k-style-hub/features/likeReview/repository"
	likereviewService "aziz-wahyudin/technical-test-k-style-hub/features/likeReview/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	memberRepoFactory := memberRepo.New(db)
	memberServiceFactory := memberService.New(memberRepoFactory)
	memberDelivery.New(memberServiceFactory, e)

	likereviewRepoFactory := likereviewRepo.New(db)
	likereviewServiceFactory := likereviewService.New(likereviewRepoFactory)
	likereviewDelivery.New(likereviewServiceFactory, e)
}

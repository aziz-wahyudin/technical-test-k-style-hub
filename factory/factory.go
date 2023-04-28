package factory

import (
	memberDelivery "aziz-wahyudin/technical-test-k-style-hub/features/member/delivery"
	memberRepo "aziz-wahyudin/technical-test-k-style-hub/features/member/repository"
	memberService "aziz-wahyudin/technical-test-k-style-hub/features/member/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	memberRepoFactory := memberRepo.New(db)
	memberServiceFactory := memberService.New(memberRepoFactory)
	memberDelivery.New(memberServiceFactory, e)
}

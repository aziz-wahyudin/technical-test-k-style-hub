package delivery

import (
	"strconv"

	likereview "aziz-wahyudin/technical-test-k-style-hub/features/likeReview"
	"aziz-wahyudin/technical-test-k-style-hub/middlewares"
	"aziz-wahyudin/technical-test-k-style-hub/utils/helper"

	"net/http"

	"github.com/labstack/echo/v4"
)

type LikeReviewDelivery struct {
	LikeReviewService likereview.ServiceInterface
}

func New(service likereview.ServiceInterface, e *echo.Echo) {
	handler := &LikeReviewDelivery{
		LikeReviewService: service,
	}
	e.POST("/likereviews/:id_review", handler.CreateLikeReview)
	e.DELETE("/likereviews/:id_review", handler.DeleteLikeReview)
}

func (d *LikeReviewDelivery) CreateLikeReview(c echo.Context) error {
	idReview, _ := strconv.Atoi(c.Param("id_review"))
	idMember := middlewares.ExtractTokenUserId(c)
	inputData := LikeReviewReq{}
	inputData.MemberID = uint(idMember)
	inputData.ReviewProdcutID = uint(idReview)

	errBind := c.Bind(&inputData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed input data"+errBind.Error()))
	}

	dataCore := inputData.reqToCore()
	err := d.LikeReviewService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create like and review information"))
}

func (d *LikeReviewDelivery) DeleteLikeReview(c echo.Context) error {
	idReview, _ := strconv.Atoi(c.Param("id_review"))
	idMember := middlewares.ExtractTokenUserId(c)
	err := d.LikeReviewService.Delete(idMember, idReview)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete like & review information"))
}

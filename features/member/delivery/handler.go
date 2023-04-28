package delivery

import (
	// "strconv"

	"aziz-wahyudin/technical-test-k-style-hub/features/member"
	"aziz-wahyudin/technical-test-k-style-hub/utils/helper"

	"net/http"

	"github.com/labstack/echo/v4"
)

type MemberDelivery struct {
	MemberService member.ServiceInterface
}

func New(service member.ServiceInterface, e *echo.Echo) {
	handler := &MemberDelivery{
		MemberService: service,
	}
	e.POST("/members", handler.CreateMember)
}

func (d *MemberDelivery) CreateMember(c echo.Context) error {
	input := MemberReq{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding data "+errBind.Error()))
	}
	dataCore := input.reqToCore()
	err := d.MemberService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create member information"))
}

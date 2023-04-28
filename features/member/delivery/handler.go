package delivery

import (
	"strconv"

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
	e.PUT("/members/:id_member", handler.UpdateMember)
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

func (d *MemberDelivery) UpdateMember(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id_member"))
	inputData := MemberReq{}
	errBind := c.Bind(&inputData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed input data"+errBind.Error()))
	}

	dataUpdateCore := inputData.reqToCore()
	err := d.MemberService.Update(dataUpdateCore, idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed update data"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update member information"))

}

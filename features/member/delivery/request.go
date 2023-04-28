package delivery

import (
	"aziz-wahyudin/technical-test-k-style-hub/features/member"
)

type MemberReq struct {
	Username  string `json:"username" form:"username"`
	Gender    string `json:"gender" form:"gender"`
	Skintype  string `json:"skintype" form:"skintype"`
	Skincolor string `json:"skincolor" form:"skincolor"`
}

func (data *MemberReq) reqToCore() member.Member {
	return member.Member{
		Username:  data.Username,
		Gender:    data.Gender,
		Skintype:  data.Skintype,
		Skincolor: data.Skincolor,
	}
}

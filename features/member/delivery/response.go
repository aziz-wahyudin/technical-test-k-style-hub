package delivery

import "aziz-wahyudin/technical-test-k-style-hub/features/member"

type MemberResp struct {
	ID        uint   `json:"id"`
	Username  string `json:"username" form:"username"`
	Gender    string `json:"gender" form:"gender"`
	Skintype  string `json:"skintype" form:"skintype"`
	Skincolor string `json:"skincolor" form:"skincolor"`
}

func FromCore(dataCore member.Member) MemberResp {
	return MemberResp{
		ID:        dataCore.ID,
		Username:  dataCore.Username,
		Gender:    dataCore.Gender,
		Skintype:  dataCore.Skintype,
		Skincolor: dataCore.Skincolor,
	}
}

func FromCoreList(dataCore []member.Member) []MemberResp {
	var dataResp []MemberResp
	for _, v := range dataCore {
		dataResp = append(dataResp, FromCore(v))
	}
	return dataResp
}

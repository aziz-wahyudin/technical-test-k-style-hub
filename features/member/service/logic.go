package service

import (
	"aziz-wahyudin/technical-test-k-style-hub/features/member"
)

type memberService struct {
	memberRepository member.RepositoryInterface
}

func New(repo member.RepositoryInterface) member.ServiceInterface {
	return &memberService{
		memberRepository: repo,
	}
}

// Create implements member.ServiceInterface
func (s *memberService) Create(input member.Member) (err error) {
	_, err = s.memberRepository.Create(input)
	if err != nil {
		return err
	}
	return nil
}

// Update implements member.ServiceInterface
func (s *memberService) Update(input member.Member, id int) (err error) {
	err = s.memberRepository.Update(input, id)
	if err != nil {
		return err
	}
	return nil
}

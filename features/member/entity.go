package member

type Member struct {
	ID             uint
	Username       string
	Gender         string
	Skintype       string
	Skincolor      string
	ReviewProdcuts []ReviewProdcut
	LikeReviews    []LikeReview
}

type ReviewProdcut struct {
	ID          uint
	DescReview  string
	MemberID    uint
	LikeReviews []LikeReview
}

type LikeReview struct {
	ID              uint
	Likes           int
	Dislikes        int
	MemberID        uint
	ReviewProdcutID uint
}

type ServiceInterface interface {
	Create(input Member) (err error)
	Update(input Member, id int) (err error)
	Delete(id int) (err error)
	GetAll() (data []Member, err error)
}

type RepositoryInterface interface {
	Create(input Member) (row int, err error)
	Update(input Member, id int) (err error)
	Delete(id int) (row int, err error)
	GetAll() (data []Member, err error)
}

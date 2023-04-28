package likereview

type LikeReview struct {
	ID              uint
	Likes           int
	Dislikes        int
	MemberID        uint
	ReviewProdcutID uint
}

type ServiceInterface interface {
	Create(input LikeReview) (err error)
}

type RepositoryInterface interface {
	Create(input LikeReview) (row int, err error)
}

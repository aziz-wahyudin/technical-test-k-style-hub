package product

type Member struct {
	ID             uint
	Username       string
	Gender         string
	Skintype       string
	Skincolor      string
	ReviewProdcuts []ReviewProdcut
	LikeReviews    []LikeReview
}

type Product struct {
	ID             uint
	NameProduct    string
	Price          string
	ReviewProdcuts []ReviewProdcut
}

type ReviewProdcut struct {
	ID          uint
	DescReview  string
	MemberID    uint
	LikeReviews []LikeReview
}

type LikeReview struct {
	Likes           int
	Dislikes        int
	MemberID        uint
	ReviewProdcutID uint
	Review          ReviewProdcut
	Member          Member
}

type ServiceInterface interface {
	GetById(id int) (data Product, err error)
}

type RepositoryInterface interface {
	GetById(id int) (data Product, err error)
}

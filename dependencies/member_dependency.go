package dependencies

import (
	cont "be/controllers/member"
	implCont "be/controllers/member/impl"
	"be/models/repositories"
	"be/models/repositories/impl"
	memberRepo "be/models/repositories/member"
	implRepo "be/models/repositories/member/impl"
	"be/models/services/member"
	implServ "be/models/services/member/impl"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type MemberDependency struct {
	Db          *gorm.DB
	UserRepo    repositories.UserRepo
	UserServ    member.UserServMember
	UserCont    cont.UserContMember
	ProductRepo memberRepo.ProductRepoMember
	ProductServ member.ProductServMember
	ProductCont cont.ProductContMember
}

func NewMemberDependency(db *gorm.DB, validator *validator.Validate, jwtKey string) *MemberDependency {
	// Repositories
	userRepo := impl.NewUserRepoImpl()
	productRepo := implRepo.NewProductRepoMemberImpl()

	// Services
	userServ := implServ.NewUserServMemberImpl(userRepo, db, validator, jwtKey)
	productServ := implServ.NewProductServMemberImpl(db, validator, productRepo)

	// Controllers
	userCont := implCont.NewUserContMemberImpl(userServ)
	productCont := implCont.NewProductContMemberImpl(productServ)

	return &MemberDependency{Db: db, UserRepo: userRepo, UserServ: userServ, UserCont: userCont, ProductCont: productCont}
}

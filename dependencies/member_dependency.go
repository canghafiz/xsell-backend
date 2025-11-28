package dependencies

import (
	cont "be/controllers/member"
	implCont "be/controllers/member/impl"
	"be/models/domains"
	"be/models/repositories"
	"be/models/repositories/impl"
	memberRepo "be/models/repositories/member"
	implRepo "be/models/repositories/member/impl"
	"be/models/services"
	implGeneralServ "be/models/services/impl"
	"be/models/services/member"
	implServ "be/models/services/member/impl"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type MemberDependency struct {
	Db             *gorm.DB
	RedisServ      services.RedisService
	UserRepo       repositories.UserRepo
	UserServ       member.UserServMember
	UserCont       cont.UserContMember
	ProductRepo    memberRepo.ProductRepoMember
	ProductServ    member.ProductServMember
	ProductCont    cont.ProductContMember
	PageLayoutRepo repositories.PageLayoutRepo
	PageLayoutServ member.PageLayoutServMember
	PageLayoutCont cont.PageLayoutContMember
}

func NewMemberDependency(db *gorm.DB, validator *validator.Validate, redisConfig domains.RedisConfig, jwtKey string) *MemberDependency {
	// Repositories
	userRepo := impl.NewUserRepoImpl()
	productRepo := implRepo.NewProductRepoMemberImpl()
	pageLayoutRepo := impl.NewPageLayoutRepoImpl()

	// Services
	redisServ := implGeneralServ.NewRedisServiceImpl(redisConfig)
	userServ := implServ.NewUserServMemberImpl(userRepo, db, validator, jwtKey)
	productServ := implServ.NewProductServMemberImpl(db, validator, productRepo)
	pageLayoutServ := implServ.NewPageLayoutServMemberImpl(db, redisServ, pageLayoutRepo)

	// Controllers
	userCont := implCont.NewUserContMemberImpl(userServ)
	productCont := implCont.NewProductContMemberImpl(productServ)
	pageLayoutCont := implCont.NewPageLayoutContMemberImpl(pageLayoutServ)

	return &MemberDependency{Db: db, RedisServ: redisServ, UserRepo: userRepo, UserCont: userCont, ProductCont: productCont, PageLayoutCont: pageLayoutCont}
}

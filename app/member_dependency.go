package app

import (
	cont "be/controllers/member"
	implCont "be/controllers/member/impl"
	"be/models/repositories"
	"be/models/repositories/impl"
	"be/models/services/member"
	implServ "be/models/services/member/impl"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type MemberDependency struct {
	Db       *gorm.DB
	UserRepo repositories.UserRepo
	UserServ member.UserServMember
	UserCont cont.UserContMember
}

func NewMemberDependency(db *gorm.DB, validator *validator.Validate, jwtKey string) *MemberDependency {
	// Repositories
	userRepo := impl.NewUserRepoImpl()

	// Services
	userServ := implServ.NewUserServMemberImpl(userRepo, db, validator, jwtKey)

	// Controllers
	userCont := implCont.NewUserContMemberImpl(userServ)

	return &MemberDependency{Db: db, UserRepo: userRepo, UserServ: userServ, UserCont: userCont}
}

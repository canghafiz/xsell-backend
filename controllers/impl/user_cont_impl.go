package impl

import (
	"be/models/services"

	"github.com/gin-gonic/gin"
)

type UserContImpl struct {
	UserServ services.UserServ
}

func NewUserContImpl(userServ services.UserServ) *UserContImpl {
	return &UserContImpl{UserServ: userServ}
}

func (cont *UserContImpl) RegisterMember(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (cont *UserContImpl) LoginMember(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (cont *UserContImpl) Logout(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (cont *UserContImpl) UpdateData(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

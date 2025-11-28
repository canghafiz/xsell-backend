package dependencies

import (
	"be/controllers"
	contImpl "be/controllers/impl"
	"be/models/domains"
	repo "be/models/repositories/impl"
	"be/models/services/impl"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Dependency struct {
	FileCont controllers.FileCont
	OtpCont  controllers.OtpCont
}

func NewDependency(db *gorm.DB, validator *validator.Validate, smtp domains.Smtp, appName string) *Dependency {
	// Repo
	otpRepo := repo.NewOtpRepoImpl()

	// Serv
	fileServ := impl.NewFileServImpl()
	smtpServ := impl.NewSmtpServImpl(smtp, appName)
	otpServ := impl.NewOtpServImpl(db, validator, otpRepo, smtpServ)

	// Cont
	fileCont := contImpl.NewFileContImpl(fileServ)
	otpCont := contImpl.NewOtpContImpl(otpServ)

	return &Dependency{
		FileCont: fileCont,
		OtpCont:  otpCont,
	}
}

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
	FileCont     controllers.FileCont
	OtpCont      controllers.OtpCont
	BannerCont   controllers.BannerCont
	CategoryCont controllers.CategoryCont
	MetaSeoCont  controllers.MetaSeoCont
}

func NewDependency(db *gorm.DB, validator *validator.Validate, smtp domains.Smtp, appName string) *Dependency {
	// Repo
	otpRepo := repo.NewOtpRepoImpl()
	bannerRepo := repo.NewBannerRepoImpl()
	categoryRepo := repo.NewCategoryRepoImpl()
	metaRepo := repo.NewMetaSeoRepoImpl()

	// Serv
	fileServ := impl.NewFileServImpl()
	smtpServ := impl.NewSmtpServImpl(smtp, appName)
	otpServ := impl.NewOtpServImpl(db, validator, otpRepo, smtpServ)
	bannerServ := impl.NewBannerServImpl(db, bannerRepo)
	categoryServ := impl.NewCategoryServImpl(db, categoryRepo)
	metaServ := impl.NewMetaSeoServImpl(db, metaRepo)

	// Cont
	fileCont := contImpl.NewFileContImpl(fileServ)
	otpCont := contImpl.NewOtpContImpl(otpServ)
	bannerCont := contImpl.NewBannerContImpl(bannerServ)
	categoryCont := contImpl.NewCategoryContImpl(categoryServ)
	metaCont := contImpl.NewMetaSeoContImpl(metaServ)

	return &Dependency{
		FileCont:     fileCont,
		OtpCont:      otpCont,
		BannerCont:   bannerCont,
		CategoryCont: categoryCont,
		MetaSeoCont:  metaCont,
	}
}

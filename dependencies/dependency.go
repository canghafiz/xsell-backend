package dependencies

import (
	"be/controllers"
	contImpl "be/controllers/impl"
	"be/models/services/impl"
)

type Dependency struct {
	FileCont controllers.FileCont
}

func NewDependency() *Dependency {
	// Serv
	fileServ := impl.NewFileServImpl()

	// Cont
	fileCont := contImpl.NewFileContImpl(fileServ)

	return &Dependency{
		FileCont: fileCont,
	}
}

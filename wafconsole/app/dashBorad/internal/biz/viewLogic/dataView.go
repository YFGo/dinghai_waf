package viewLogic

type DataViewRepo interface {
}

type DataViewUsecase struct {
	repo DataViewRepo
}

func NewDataViewUsecase(repo DataViewRepo) *DataViewUsecase {
	return &DataViewUsecase{repo: repo}
}

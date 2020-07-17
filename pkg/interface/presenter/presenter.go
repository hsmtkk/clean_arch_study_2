package presenter

type Presenter interface {
}

type presenterImpl struct{}

func New() Presenter {
	return &presenterImpl{}
}

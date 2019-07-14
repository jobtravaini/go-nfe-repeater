package nfe

type INfeRepository interface {
	FindByAccessKey(accessKey string) (Nfe, error)
	CreateOrUpdate(nfe Nfe) error
}

type INfeService interface {
	GetNfe(accessKey string) (Nfe, error)
}
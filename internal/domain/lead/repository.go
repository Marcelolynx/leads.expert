package lead

type Repository interface {
	Salvar(lead *Lead) error
}

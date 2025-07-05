package storage

type PodInfo struct {
	Name      string
	Namespace string
	Labels    map[string]string
}

type Store interface {
	Save(pod PodInfo) error
	GetAll() ([]PodInfo, error)
}

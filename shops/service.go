package shops

import (
	"github.com/deliriumproducts/aumo"
)

type service struct {
	store aumo.ShopStore
}

//New returns an instance of `aumo.ShopService`
func New(store aumo.ShopStore) aumo.ShopService {
	return &service{
		store: store,
	}
}

func (ss *service) Shop(id uint) (*aumo.Shop, error) {
	return ss.store.FindByID(nil, id, false)
}

func (ss *service) Shops() ([]aumo.Shop, error) {
	return ss.store.FindAll(nil)
}

func (ss *service) Create(s *aumo.Shop) error {
	return ss.store.Save(nil, s)
}

func (ss *service) Update(id uint, s *aumo.Shop) error {
	return ss.store.Update(nil, id, s)
}

func (ss *service) Delete(id uint) error {
	return ss.store.Delete(nil, id)
}

func (ss *service) Owners() error {
	return ss.store.FindByID(nil, id, true)
}
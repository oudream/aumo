package rest

import (
	"errors"
	"net/http"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/auth"
)

func (rest *Rest) shopGetAll(w http.ResponseWriter, r *http.Request) {
	shops, err := rest.shopService.Shops()
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, shops, http.StatusOK)
}

func (rest *Rest) shopGet(w http.ResponseWriter, r *http.Request) {
	sID := rest.ParamNumber(w, r, "id")
	user, err := auth.CurrentUser(r.Context())
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	withOwners := false
	if user.Role == aumo.ShopOwner {
		for _, shop := range user.Shops {
			// If the user owns the given shop, they can get the owners
			if shop.ID == sID {
				withOwners = true
				break
			}
		}
	}

	if user.Role == aumo.Admin {
		withOwners = true
	}

	shop, err := rest.shopService.Shop(sID, user.Role != aumo.Customer && withOwners) // only get the owners if the user is not a customer
	switch {
	case err == nil:
		break
	case errors.Is(err, aumo.ErrShopNotFound):
		rest.JSONError(w, err, http.StatusNotFound)
		return
	default:
		rest.JSONError(w, err, http.StatusInternalServerError)
		return

	}

	rest.JSON(w, shop, http.StatusOK)
}

func (rest *Rest) shopCreate(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name  string `form:"name" validate:"required" json:"name"`
		Image string `form:"image" validate:"required,url" json:"image"`
	}

	var um request
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	shop := aumo.NewShop(um.Name, um.Image)

	err := rest.shopService.Create(shop)

	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, shop, http.StatusOK)
}

func (rest *Rest) shopEdit(w http.ResponseWriter, r *http.Request) {
	sID := rest.ParamNumber(w, r, "id")

	type request struct {
		Name  string `form:"name" validate:"required" json:"name"`
		Image string `form:"image" validate:"required,url" json:"image"`
	}

	var um request
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	shop := aumo.NewShop(um.Name, um.Image)
	err := rest.shopService.Update(sID, shop)
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, shop, http.StatusOK)

}

func (rest *Rest) shopAddOwner(w http.ResponseWriter, r *http.Request) {
	sID := rest.ParamNumber(w, r, "id")

	type request struct {
		UserID string `form:"user_id" validate:"required" json:"user_id"`
	}

	var um request
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	so := aumo.NewShopOwners(sID, um.UserID)
	err := rest.shopService.AddOwner(so)

	switch {
	case err == nil:
		break
	case errors.Is(err, aumo.ErrShopOwnerUserNotFound):
		rest.JSONError(w, err, http.StatusNotFound)
		return
	default:
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, Message{"Owner successfully added!"}, http.StatusOK)
}

func (rest *Rest) shopRemoveOwner(w http.ResponseWriter, r *http.Request) {
	sID := rest.ParamNumber(w, r, "id")

	type request struct {
		UserID string `form:"user_id" validate:"required" json:"user_id"`
	}

	var um request
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	so := aumo.NewShopOwners(sID, um.UserID)
	err := rest.shopService.RemoveOwner(so)

	switch {
	case err == nil:
		break
	case errors.Is(err, aumo.ErrShopOwnerUserNotFound):
		rest.JSONError(w, err, http.StatusNotFound)
		return
	default:
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, Message{"Owner successfully removed!"}, http.StatusOK)
}

package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var (
	ErrBadTypeAssertion = errors.New("web: failed to assert type")
)

const (
	CookieStoreKey = "aumo"
	UserContextKey = "user"
	UserSessionKey = "user"
)

type UserForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

func (wb *Web) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var um UserForm
	if err := json.NewDecoder(r.Body).Decode(&um); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := wb.CreateUser(um.Name, um.Email, um.Password, um.Avatar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (wb *Web) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var um UserForm

	if err := json.NewDecoder(r.Body).Decode(&um); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session, err := wb.store.Get(r, CookieStoreKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := wb.GetUserByEmail(um.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if !user.ValidatePassword(um.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
	}

	session.Values[UserSessionKey] = &user

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wb *Web) ClaimReceiptHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, "User unauthorized", http.StatusUnauthorized)
		return
	}

	receipt, err := wb.GetReceiptByID(uint(id))
	if err != nil {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	if err := wb.Aumo.SetReceiptUserID(user, receipt); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

type BuyForm struct {
	Quantity uint `json:"quantity"`
}

func (wb *Web) BuyHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var bf BuyForm
	if err := json.NewDecoder(r.Body).Decode(&bf); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, "User unauthorized", http.StatusUnauthorized)
		return
	}

	si, err := wb.GetShopItemByID(uint(id))
	if err != nil {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	if err := wb.Aumo.BuyUserShopItem(user, si, bf.Quantity); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (wb *Web) MeHandler(w http.ResponseWriter, r *http.Request) {
	u, err := GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, "User unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewEncoder(w).Encode(u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
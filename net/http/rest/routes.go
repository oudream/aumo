package rest

import (
	"github.com/deliriumproducts/aumo"
	"github.com/go-chi/chi"
)

func (rest *Rest) routes() {
	rest.router.Post("/register", rest.userRegister)
	rest.router.Post("/login", rest.userLogin)
	rest.router.Group(func(r chi.Router) {
		r.Use(rest.WithAuth())
		r.Get("/logout", rest.userLogout)
		r.Get("/me", rest.userGetCurrent)
	})

	rest.router.Route("/receipts", func(r chi.Router) {
		r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.receiptCreate)
		r.With(rest.WithAuth(aumo.Customer)).Get("/{id}", rest.receiptClaim)
	})

	rest.router.Route("/products", func(r chi.Router) {
		r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.productCreate)
		r.Get("/", rest.productGetAll)
		r.Get("/{id}", rest.productGet)
		r.Put("/{id}", rest.productEdit)
	})

	rest.router.Route("/orders", func(r chi.Router) {
		r.With(rest.WithAuth(aumo.Customer)).Post("/", rest.orderCreate)
		r.With(rest.WithAuth(aumo.Admin)).Get("/", rest.orderGetAll)
		r.With(rest.WithAuth(aumo.Admin)).Get("/{id}", rest.orderGet)
	})
}

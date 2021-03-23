package router

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/controllers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/controllers/auth"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Route struct defining all of this project routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Public      bool
}

// Routes slice of Route
type Routes []Route

// NewRouter registers public routes
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	authenticatedRouter := router.PathPrefix("/").Subrouter()

	for _, route := range routes {
		appRouter := authenticatedRouter
		if route.Public {
			appRouter = router
		}
		appRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	authenticatedRouter.Use(loggingMiddleware)

	return router
}

var routes = Routes{
	Route{
		Name:        "Home",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controllers.RenderHome,
	},
	Route{
		Name:        "Get Order",
		Method:      "GET",
		Pattern:     "/order/{id}",
		HandlerFunc: controllers.GetOrder,
		Public:      true,
	},
	Route{
		Name:        "Create Order",
		Method:      "POST",
		Pattern:     "/order/new",
		HandlerFunc: controllers.CreateOrder,
		Public:      false,
	},
	Route{
		Name:        "Update Order",
		Method:      "POST",
		Pattern:     "/order/{id}/update",
		HandlerFunc: controllers.UpdateOrder,
		Public:      false,
	},
	Route{
		Name:        "Add Order Line",
		Method:      "POST",
		Pattern:     "/order/{id}/orderLine/new",
		HandlerFunc: controllers.AddOrderLine,
		Public:      false,
	},
	Route{
		Name:        "Update Order Line quantity",
		Method:      "POST",
		Pattern:     "/orderLine/{id}/updateQuantity",
		HandlerFunc: controllers.UpdateOrderLineQuantity,
		Public:      false,
	},
	Route{
		Name:        "Delete Order Line",
		Method:      "POST",
		Pattern:     "/orderLine/{id}/delete",
		HandlerFunc: controllers.DeleteOrderLine,
		Public:      false,
	},
	//Auth
	Route{
		Name:        "Sign In",
		Method:      "POST",
		Pattern:     "/signin",
		HandlerFunc: auth.Signin,
		Public:      true,
	},
	Route{
		Name:        "Sign Up",
		Method:      "POST",
		Pattern:     "/signup",
		HandlerFunc: auth.SignUp,
		Public:      true,
	},
	Route{
		Name:        "Refresh",
		Method:      "GET",
		HandlerFunc: auth.Refresh,
		Public:      false,
	},
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				helpers.WriteErrorJSON(w, http.StatusUnauthorized, "auth is not logged in")
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tknStr := c.Value
		claims := &auth.GoQRSClaims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				helpers.WriteErrorJSON(w, http.StatusUnauthorized, "invalid signature")
				return
			}
			helpers.WriteErrorJSON(w, http.StatusBadRequest, "Bad request")
			return
		}
		if !tkn.Valid {
			helpers.WriteErrorJSON(w, http.StatusUnauthorized, "invalid token")
			return
		}

		auth.Refresh(w, r)

		next.ServeHTTP(w, r)
	})
}

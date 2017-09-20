package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/goadesign/oauth2/design"
)

var OAuth2Sec = OAuth2("/oauth2/authorize", "/oauth2/token", func() {
	Scope("api:read")
	Scope("api:write")
})

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "POST")
	})
	Files("/css/*filepath", "public/css")
	Files("/js/*filepath", "public/js")
})

var _ = Resource("authUI", func() {
	BasePath("/auth")
	Action("promptAuthorization", func() {
		Description("Prompt the user for client authorization")
		Routing(GET("/authorize-client"))
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)

	})
	Action("confirmAuthorization", func() {
		Description("Confirm the authorization of the client")
		Routing(POST("/confirm-authorization"))
		Params(func() {
			Param("confirmed", Boolean, "Is the authorization confirmed.")
		})
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("login", func() {
	BasePath("/login")
	Action("showLogin", func() {
		Description("Shows a login screen")
		Routing(GET(""))
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(Unauthorized, ErrorMedia)
	})
})
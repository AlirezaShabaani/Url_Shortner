package handleHttp

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"url_shortner/internal/core/port"
)

type HttpHandler struct {
	urlServices port.UrlServices
}

func New(urlServices port.UrlServices) HttpHandler {
	return HttpHandler{urlServices: urlServices}
}

func (h *HttpHandler) Save(ctx echo.Context) error {
	var ourl string
	if ourl = ctx.QueryParam("ourl"); ourl == "" {
		return ctx.JSON(http.StatusBadRequest, "url param can't be empty")
	}
	surl, err := h.urlServices.Save(ourl)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "shortened url : localhost:10000/"+surl)
}


func (h *HttpHandler) Redirect(ctx echo.Context) error {
	var surl string
	if surl = ctx.Param("surl");surl == ""{
		return ctx.JSON(http.StatusBadRequest, "url param can't be empty")
	}
	ourl, err := h.urlServices.Read(surl)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.Redirect(301,ourl)
}

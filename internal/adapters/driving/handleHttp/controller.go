package handleHttp

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"url_shortner/internal/adapters/driving/handleHttp/model"
	"url_shortner/internal/core/port"
)

type HttpHandler struct {
	urlServices port.UrlServices
}

func New(urlServices port.UrlServices) HttpHandler {
	return HttpHandler{urlServices: urlServices}
}

func (h *HttpHandler) Save(ctx echo.Context) error {
	m := new(model.GenerateUrlReq)
	if err := ctx.Bind(m); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	surl, err := h.urlServices.Save(m.Ourl)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "shortened url :"+surl)
}


func (h *HttpHandler) Redirect(ctx echo.Context) error {
	var surl string
	if surl = ctx.QueryParam("surl"); surl == "" {
		return ctx.JSON(http.StatusBadRequest, "url param can't be empty")
	}
	ourl, err := h.urlServices.Read(surl)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.Redirect(http.StatusPermanentRedirect,ourl)
	//if err != nil {
	//	return ctx.JSON(http.StatusBadRequest, err.Error())
	//}
	//return ctx.JSON(http.StatusOK,ourl)
}

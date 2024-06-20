package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autorizationheader = "Authorization"
	userCtx            = "user_id"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(autorizationheader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		fmt.Println(header)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		fmt.Println(header)
		return
	}

	userId, err := h.services.Auth.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		fmt.Println(header)
		return
	}

	c.Set(userCtx, userId)
}

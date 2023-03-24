package delivery

import (
	"net/http"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *viewHandler) RenderHome(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)

	var name string

	if userId != uuid.Nil {
		res := handler.usecase.GetUserById(model.GetUserByIdRequest{ID: userId})

		if res.Status == http.StatusOK {
			user, ok := res.Data["user"].(model.User)

			if ok {
				name = user.Name
			}
		}
	}

	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"Name": name,
	})
}

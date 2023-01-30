package landing

import (
  "net/http"

  "github.com/gin-gonic/gin"
)


func Render(ctx *gin.Context) {
  ctx.HTML(http.StatusOK, "landing.html", gin.H{})
}

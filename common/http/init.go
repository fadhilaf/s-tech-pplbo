package http

import ( 
  "github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Router *gin.Engine
}

func NewHTTPServer() HTTPServer {

	router := gin.Default()

  router.LoadHTMLGlob("internal/view/*.html")

	return HTTPServer{
		Router: router,
	}
}

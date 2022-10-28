package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/devfest-kutaisi-2022/internal/model"
)

type WineRepository interface {
	GetWineList(ctx context.Context) (map[string]model.Wine, error)
	GetWine(ctx context.Context, id string) (*model.Wine, error)
	AddWine(ctx context.Context, id string, w model.Wine) error
}

type Server struct {
	repo WineRepository
	srv  *http.Server
}

func New(address string, repo WineRepository) *Server {

	srv := &http.Server{
		Addr: address,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	s := &Server{
		repo: repo,
		srv:  srv,
	}

	router.GET("/", hello)
	router.POST("/wine", s.AddWine)
	router.GET("/wine/:id", s.GetWine)
	router.GET("/wine", s.ListWine)

	srv.Handler = router

	return s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

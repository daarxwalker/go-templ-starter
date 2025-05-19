package templ_service

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type TemplService struct {
	c *gin.Context
}

func New(c *gin.Context) *TemplService {
	return &TemplService{
		c: c,
	}
}

func (s *TemplService) Render(status int, component templ.Component) error {
	s.c.Header("Content-Type", "text/html; charset=utf-8")
	s.c.Status(status)
	return component.Render(s.c, s.c.Writer)
}

func (s *TemplService) MustRender(status int, component templ.Component) {
	if err := s.Render(status, component); err != nil {
		s.c.Error(fmt.Errorf("templ component render failed: %v", err))
	}
}

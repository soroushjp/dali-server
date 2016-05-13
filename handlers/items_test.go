package handlers

import (
	"os"
	"testing"

	"github.com/albrow/fipple"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/soroushjp/dali-server/context"
	"github.com/stretchr/testify/suite"
)

type ItemsHandlerSuite struct {
	suite.Suite

	app *context.AppContext
	rec *fipple.Recorder
}

func TestItemsHandlerSuite(t *testing.T) {
	suite.Run(t, &ItemsHandlerSuite{})
}

func (s *ItemsHandlerSuite) SetupSuite() {
	// get DB connection
	db, err := sqlx.Connect(
		"postgres",
		"user="+os.Getenv("DALI_TEST_DB_USER")+
			" dbname="+os.Getenv("DALI_TEST_DB_NAME")+
			" sslmode=disable",
	)
	s.Require().NoError(err)
	// create test app context
	app := &context.AppContext{
		DB: db,
	}
	s.Require().NoError(err)
	s.app = app
	// create gin engine, set test mode and initialize fipple recorder with engine
	gin.SetMode(gin.TestMode)
	eng := NewEngine(s.app)
	s.rec = fipple.NewRecorder(s.T(), eng)
}

func (s *ItemsHandlerSuite) SetupTest() {
	_, err := s.app.DB.Exec("DELETE FROM items WHERE true;")
	s.Require().NoError(err)
}

func (s *ItemsHandlerSuite) TestItemsHandler_Index() {
	res := s.rec.Get("/items")
	res.ExpectOk()
	res.ExpectBodyContains("[]")
}

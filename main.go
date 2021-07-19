package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ae-tech-behind/turbo-dollop/controller"
	"github.com/ae-tech-behind/turbo-dollop/router"
	"github.com/ae-tech-behind/turbo-dollop/service"
	"github.com/ae-tech-behind/turbo-dollop/store"
	"github.com/ae-tech-behind/turbo-dollop/usecase"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/oauth2/v1"
)

const (
	// AppName application name
	AppName = "go-service-demo"
)

var httpClient = &http.Client{}

func main() {

	logger := logrus.New()

	//postgres://postgres:Alma2097@localhost:5432/Library?sslmode=disable

	db, err := store.New("user=postgres password=Alma2097 dbname=Library sslmode=disable")
	if err != nil {
		logger.Fatal("Something went wrong")
	}
	// Service init
	someService := service.NewSomeService("something")

	// UseCase init
	somethingUC := usecase.NewSomething(someService)
	statusUC := usecase.NewStatus(AppName)
	userUC := usecase.NewUsers(db)
	bookUC := usecase.NewBooks(db)
	loanUC := usecase.NewLoans(db)

	// Controller init
	somethingC := controller.NewSomething(somethingUC)
	statusC := controller.NewStatus(statusUC)
	userC := controller.NewUsers(userUC) // se inicializa el user controller
	bookC := controller.NewBooks(bookUC)
	loanC := controller.NewLoans(loanUC)
	// Create router
	r := router.New(somethingC, statusC, userC, bookC, loanC)

	r.Use(TokenValidator())

	// Define stop signal for the end of excecution
	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	go func() {
		address := ":8080"
		if err := r.Start(address); err != nil {
			logger.Fatal("something went wrong")
		}
	}()

	<-stop

	logger.Info("shutting down server...")
}

func TokenValidator() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := c.Response()
			req := c.Request()

			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Access-Control-Allow-Headers:", "*")
			res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			res.Header().Set("Access-Control-Expose-Headers", "Accept, Accept-Lango gguage, Content-Type, token, X-Auth-Token")

			token := req.Header.Get("Authorization")
			_, err := verifyIdToken(token)
			if err != nil {
				err = fmt.Errorf("Invalid Acces token")
				c.String(http.StatusUnauthorized, err.Error())
				return err
			}

			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

func verifyIdToken(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

//cd "C:\Program Files (x86)\Google\Chrome\Application"
//chrome.exe --user-data-dir="C:/Chrome dev session" --disable-web-security

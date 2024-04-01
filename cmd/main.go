package main

import (
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"learn-ddd/db"
	"learn-ddd/gen/api/v1/apiv1connect"
	"learn-ddd/internal/handler"
	dbrepo "learn-ddd/internal/infrastructure/db"
	//memrepo "learn-ddd/internal/infrastructure/memory"
	"learn-ddd/internal/interceptor"
	"learn-ddd/internal/usecase"
	"learn-ddd/lib/errctrl"
)

func init() {
	time.Local = time.FixedZone("Local", 9*60*60)
}

func main() {
	errctrl.MustExec(db.Connect())
	interceptors := connect.WithInterceptors(interceptor.NewErrorResponseInterceptor(), interceptor.NewTransactionInterceptor())
	validator := errctrl.Must(protovalidate.New())

	mux := http.NewServeMux()
	mux.Handle(apiv1connect.NewUserServiceHandler(handler.NewUserServiceHandler(usecase.NewUserUseCase(dbrepo.NewUserRepository()), validator), interceptors))
	//mux.Handle(apiv1connect.NewUserServiceHandler(handler.NewUserServiceHandler(usecase.NewUserUseCase(memrepo.NewUserMemoryRepository()), validator), interceptors))
	mux.Handle(apiv1connect.NewTaskServiceHandler(handler.NewTaskServiceHandler(usecase.NewTaskUseCase(dbrepo.NewTaskRepository(), dbrepo.NewUserRepository()), validator), interceptors))

	log.Println("Server is starting...")
	errctrl.MustExec(http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{})))
}

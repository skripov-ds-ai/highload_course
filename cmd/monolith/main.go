package main

import (
	"github.com/skripov-ds-ai/highload_course/internal/app/monolith"
)

func main() {
	//fx.New(monolith.CreateApp()).Run()

	monolith.NewApp().Run()
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//app := fx.New(monolith.CreateApp())
	//
	//err := app.Start(ctx)
	//if err != nil {
	//	// TODO: zap in new vs zap here
	//}
	//
	//<-ctx.Done() // ожидаем завершения контекста в случае ошибки или получения сигнала
	//
	//err = app.Stop(context.Background())
	//if err != nil {
	//	// TODO: zap in new vs zap here
	//}
}

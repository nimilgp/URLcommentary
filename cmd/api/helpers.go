package main

import (
	"fmt"
	"net/http"
)

func (app *application) newEmailData() map[string]any {
	data := map[string]any{
		"BaseURL": app.config.baseURL,
	}

	return data
}

func (app *application) backgroundTask(r *http.Request, fn func() error) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			err := recover()
			if err != nil {
				app.reportServerError(r, fmt.Errorf("%s", err))
			}
		}()

		err := fn()
		if err != nil {
			app.reportServerError(r, err)
		}
	}()
}

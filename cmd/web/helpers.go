package main

import (
    "html/template"
    "net/http"
)

func (app *application) render(
    w http.ResponseWriter,
    page string,
    data any,
) {
    files := []string{
        "./ui/html/base.tmpl.html",
        "./ui/html/pages/" + page,
        "./ui/html/partials/nav.tmpl.html",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        http.Error(
            w,
            err.Error(),
            http.StatusInternalServerError,
        )
        return
    }

    err = ts.ExecuteTemplate(
        w,
        "base",
        data,
    )

    if err != nil {
        http.Error(
            w,
            err.Error(),
            http.StatusInternalServerError,
        )
    }
}
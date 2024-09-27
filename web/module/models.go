package module

import "github.com/a-h/templ"

type Module struct {
    ID string
    Url string
    Title string
    Icon func(string) templ.Component
}

func New(id, url, title string, icon func(string) templ.Component) Module {
    return Module{
        ID: id,
        Url: url,
        Title: title,
        Icon: icon,
    }
}

func (module Module) Equals(other Module) bool {
    return module.ID == other.ID
}

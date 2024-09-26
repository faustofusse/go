package module

import "github.com/a-h/templ"

type Module struct {
    ID string
    Title string
    Icon func(string) templ.Component
}

func New(id, title string, icon func(string) templ.Component) Module {
    return Module{
        ID: id,
        Title: title,
        Icon: icon,
    }
}

func (module Module) Equals(other Module) bool {
    return module.ID == other.ID
}

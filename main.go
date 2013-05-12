var Routes []*goku.Route = []*goku.Route{
    &goku.Route{
        Name:     "static",
        IsStatic: true,
        Pattern:  "/public/(.*)",
    },
    &goku.Route{
        Name:       "edit",
        Pattern:    "/{controller}/{id}/{action}",
        Default:    map[string]string{"action": "edit"},
        Constraint: map[string]string{"id": "\\d+"},
    },
    &goku.Route{
        Name:    "default",
        Pattern: "/{controller}/{action}",
        Default: map[string]string{"controller": "todo", "action": "index"},
    },
}

// a home controller builder
var cb *goku.ControllerBuilder = goku.Controller("home")
// add a "index" action to "home" controller for http get
cb.Get("index", func(ctx *goku.HttpContext) goku.ActionResulter {
    return ctx.Html("Hello World")
})j

// create server with the rules
rt := &goku.RouteTable{Routes: todo.Routes}
s := goku.CreateServer(rt, nil, nil)
log.Fatal(s.ListenAndServe())

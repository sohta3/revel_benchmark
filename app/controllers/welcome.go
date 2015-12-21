package controllers

import "github.com/revel/revel"

type Welcome struct {
	*revel.Controller
}

func (c Welcome) Index() revel.Result {
	return c.RenderJson(`[{"name": "Chris McCord"},  {"name": "Matt Sears"},  {"name": "David Stump"},  {"name": "Ricardo Thompson"}]`)
}

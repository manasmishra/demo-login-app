package controllers

import (
	"TaskManager/common"

	"gopkg.in/mgo.v2"
)

// Struct used for maintaining HTTP Request context
type Context struct {
	MongoSession *mgo.Session
	UserName     string
}

// close mongo session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// Return Mongo collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

// Create a new context for each HTTP request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
		UserName:     "",
	}
	return context
}

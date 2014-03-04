package goflow

type Application interface {
	Name() string
	Id() AppId
}

type AppId uint64

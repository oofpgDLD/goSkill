package desk

type IFactoryDesk interface {
	Create() Desk
}

type Desk interface {
	Open()
	Close()
}


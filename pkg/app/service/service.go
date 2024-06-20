package service

type IService[T any] interface {
	IAdd
	IUpdate
	IUpdateField
	IList[T]
	IListPage[T]
	IContent[T]
	IStatus
	IDelete
}

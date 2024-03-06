package gform

// IOrmExecute ...
type IOrmExecute interface {
	GetForce() bool
	// insert,insertGetId
	Insert(data ...interface{}) (int64, error)
	InsertGetId(data ...interface{}) (int64, error)
	Update(data ...interface{}) (int64, error)
	Save(data ...interface{}) (int64, error)
	// UpdateOrInsert

	// increment,decrement
	// 在操作过程中你还可以指定额外的列进行更新：
	Increment(args ...interface{}) (int64, error)
	Decrement(args ...interface{}) (int64, error)
	// delete
	Delete() (int64, error)
	//LastInsertId() int64
	Force() IOrm
}

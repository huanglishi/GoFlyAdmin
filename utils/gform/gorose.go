package gform

// gform_IMG ...
const gform_IMG = `gofly为您护航`

const (
	// VERSION_TEXT ...
	VERSION_TEXT = "\ngolang orm of gform's version : "
	// VERSION_NO ...
	VERSION_NO = "v2.1.10"
	// VERSION ...
	VERSION = VERSION_TEXT + VERSION_NO + gform_IMG
)

// Open ...
func Open(conf ...interface{}) (engin *Engin, err error) {
	// 驱动engin
	engin, err = NewEngin(conf...)
	if err != nil {
		if engin.GetLogger().EnableErrorLog() {
			engin.GetLogger().Error(err.Error())
		}
		return
	}

	return
}

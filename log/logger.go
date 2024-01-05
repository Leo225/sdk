package log

type Logger interface {
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
}

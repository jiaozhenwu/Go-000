学习笔记
1.1 Error
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
    Error() string
}

1.2 Sentinel Error

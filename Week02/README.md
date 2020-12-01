学习笔记
1.1 Error
    // The error built-in interface type is the conventional interface for
    // representing an error condition, with the nil value representing no error.
    `type error interface {
        Error() string
    }`

1.2 Error type 
    1.2.1 sentinel error // 哨兵错误
        通过判断值相等。像 io.EOF
        Go1.3 之后可以使用 if errors.Is(err,io.EOF){ // do something}
    1.2.2 type error // 自定义错误
        Go1.3 之后可以使用 if errors.As(err,MyError){ // do something}
        和上面Is大体上是一样的，区别在于Is是严格判断相等，即两个 error是否相等，
        而As判断类型是否相同，并提取第一个符合目标类型的错误，用来统一处理某一类错误
1.3 wrap error
   1.3.1
        在你的应用代码中使用 errors.New 或者 errors.Errorf返回错误
        `if len(name)<3 {
            return errors.Errorf("name too short")
        }`
        如果调用其他包内的函数，通常简单的直接返回
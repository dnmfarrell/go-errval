errval
------
errval is a go package that uses generics to return a single value which contains either an error or a value. The `Catch` method is used to handle the error or return the value:

	if ok, val := foo().Catch(logger); ok {
		fmt.Println(val)
		// Output: bar
	}

This requires callers to explictly handle potential errors. However whilst experimenting with errval I found it didn't help shorten or de-obfuscate code with sequential calls that must return the first error found ([example](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md)).

The root problem of golang's mundane error handling is its strict differentiation between expressions and statements: since only statements control program execution, do not compose (and `goto` is neutured), it requires new statement types to be added to the language to provide try/catch style exception handling.

The current [Go2 proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md) describes a check/handle mechanism which is similar to try/catch.

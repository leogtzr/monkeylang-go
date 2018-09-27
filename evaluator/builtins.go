package evaluator

import "chango/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},

	"max": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			max, ok := args[0].(*object.Integer)
			if !ok {
				return newError("arguments to `max` not supported, got %s", args[0].Type())
			}

			for _, arg := range args {
				n, ok := arg.(*object.Integer)
				if !ok {
					return newError("argument to `max` not supported, got %s", arg.Type())
				}
				if n.Value > max.Value {
					max = n
				}
			}
			return &object.Integer{Value: max.Value}

			// switch arg := args[0].(type) {
			// case *object.Integer:
			// 	return &object.Integer{Value: int64()}
			// default:
			// 	return newError("argument to `len` not supported, got %s",
			// 		args[0].Type())
			// }
		},
	},
}

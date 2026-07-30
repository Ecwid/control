package lazy

import (
	"github.com/ecwid/control"
)

type Locator func() (*control.Node, error)

func (l Locator) Resolve() (*control.Node, error) {
	return l()
}

func (l Locator) Query(css string) Locator {
	return func() (*control.Node, error) {
		node, err := l()
		if err != nil {
			return nil, err
		}
		return node.Query(css).Unwrap()
	}
}

func (l Locator) InnerText() Callable1[string] {
	return caller1(l, func(n *control.Node) (string, error) {
		return n.GetText().Unwrap()
	})
}

func (l Locator) CallFunctionOn(function string, args ...any) Callable1[any] {
	return caller1(l, func(n *control.Node) (any, error) {
		return n.CallFunctionOn(function, args...).Unwrap()
	})
}

func (l Locator) Click() Callable {
	return caller(l, func(n *control.Node) error {
		return n.Click()
	})
}

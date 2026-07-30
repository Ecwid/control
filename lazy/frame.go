package lazy

import (
	"github.com/ecwid/control"
)

type Frame func() (*control.Frame, error)

func (f Frame) Resolve() (*control.Frame, error) {
	return f()
}

func New(f *control.Frame) Frame {
	return func() (*control.Frame, error) {
		return f, nil
	}
}

func (f Frame) Navigate(url string) Callable {
	return caller(f, func(f *control.Frame) error {
		return f.Navigate(url)
	})
}

func (f Frame) Reload(ignoreCache bool, scriptToEvaluateOnLoad string) Callable {
	return caller(f, func(f *control.Frame) error {
		return f.Reload(ignoreCache, scriptToEvaluateOnLoad)
	})
}

func (f Frame) Query(css string) Locator {
	return func() (*control.Node, error) {
		frame, err := f.Resolve()
		if err != nil {
			return nil, err
		}
		doc, err := frame.Document().Unwrap()
		if err != nil {
			return nil, err
		}
		return doc.Query(css).Unwrap()
	}
}

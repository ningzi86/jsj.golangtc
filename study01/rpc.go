package study01

import "errors"

type DemoService struct {
}

type Args struct {
	Number1 int
	Number2 int
}

func (DemoService) Add(args Args, result *int) error {
	*result = args.Number1 + args.Number2
	return nil
}

func (DemoService) Div(args Args, result *float32) error {
	if args.Number2 == 0 {
		return errors.New("division by zero")
	}
	*result = float32(args.Number1) / float32(args.Number2)
	return nil
}


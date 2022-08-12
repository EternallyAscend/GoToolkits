package command

import "bytes"

type Result struct {
	err    error
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func (that *Result) GetOutputAsString() string {
	return that.stdout.String()
}

func (that *Result) GetErr() error {
	return that.err
}

func (that *Result) GetErrorAsString() string {
	return that.stderr.String()
}

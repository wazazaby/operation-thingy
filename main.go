package thingy

import "context"

type Operation func(map[string]string)

type Thingy struct {
	opCh     chan Operation
	resultCh chan map[string]string
}

func New() Thingy {
	return Thingy{
		opCh:     make(chan Operation),
		resultCh: make(chan map[string]string),
	}
}

func (t Thingy) Run(ctx context.Context) {
	subject := make(map[string]string)
	for {
		select {
		case <-ctx.Done():
			t.resultCh <- subject
			return
		case op := <-t.opCh:
			op(subject)
		}
	}
}

func (t Thingy) Upsert(key, value string) {
	t.opCh <- func(m map[string]string) {
		m[key] = value
	}
}

func (t Thingy) Delete(key string) {
	t.opCh <- func(m map[string]string) {
		delete(m, key)
	}
}

func (t Thingy) Done() map[string]string {
	return <-t.resultCh
}

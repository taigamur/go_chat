package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer

	tracer := New(&buf)

	if tracer == nil {
		t.Error("Newからも戻り値がnil")
	} else {
		tracer.Trace("こんにちは、tracerパッケージ")
		if buf.String() != "こんにちは、tracerパッケージ\n" {
			t.Errorf("'%s'という誤った文字列が出力", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("データ")
}

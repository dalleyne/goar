package goar
import (
	"testing"
	. "github.com/onsi/ginkgo"
	"github.com/mailgun/godebug/lib"
	. "github.com/onsi/gomega"
)
var active_record_suite_test_go_scope = godebug.EnteringNewScope(active_record_suite_test_go_contents)
func TestGoar(t *testing.T) {
	ctx, ok := godebug.EnterFunc(func() {
		TestGoar(t)
	})
	if !ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	scope := active_record_suite_test_go_scope.EnteringNewChildScope()
	scope.Declare("t", &t)
	godebug.Line(ctx, scope, 11)
	RegisterFailHandler(Fail)
	godebug.Line(ctx, scope, 12)
	RunSpecs(t, "ActiveRecord Suite")
}

var active_record_suite_test_go_contents = `package goar

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ActiveRecord Suite")
}
`

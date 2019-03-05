package arquitetura

import (
	"runtime"
	"testing"
)

//TestDependente verifica qual a arquitera da maquina
func TestDependente(t *testing.T) {

	t.Parallel()

	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	t.Fail()
}

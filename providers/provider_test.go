package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	lenBefore := len(Providers)
	Register("tester", Provider{Name: "tester"})

	assert.Equal(t, lenBefore+1, len(Providers), "Length of Providers should be one more than before.")
	assert.Equal(t, lenBefore+1, len(Names), "Length of Names should be one more than before.")
	assert.Equal(t, "tester", Providers["tester"].Name, "Provider name should be 'tester'.")

	Register("tester", Provider{Name: "tester-2"})

	assert.Equal(t, lenBefore+1, len(Providers), "Duplicate provider registration should not increment Provider length.")
	assert.Equal(t, lenBefore+1, len(Names), "Duplicate provider registration should not increment Names length.")
	assert.Equal(t, "tester-2", Providers["tester"].Name, "Provider name should be updated to 'tester-2'.")
}

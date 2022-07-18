package hub

import (
	"testing"

	"github.com/sgettys/bellman/pkg/criers"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestRegisterAddsCrierToRegistry(t *testing.T) {
	m := criers.MockCrier{}
	r := SingleMessageRegistry{}
	exp := map[string]criers.Crier{
		"mock": &m,
	}
	r.Register("mock", &m)
	assert.Equal(t, exp, r.criers)
}

func TestSingleMessageRegistrySendOnlySendsToSpecifiedCrier(t *testing.T) {
	dat := "test"
	m1 := criers.MockCrier{}
	m1.On("Send", mock.Anything, []byte(dat)).Return(nil)
	m2 := criers.MockCrier{}
	r := SingleMessageRegistry{}
	r.Register("m1", &m1)
	r.Register("m2", &m2)
	expReg := map[string]criers.Crier{
		"m1": &m1,
		"m2": &m2,
	}
	assert.Equal(t, expReg, r.criers)
	r.Send("m1", []byte(dat))
	m1.AssertCalled(t, "Send", mock.Anything, []byte(dat))
	m2.AssertNotCalled(t, "Send", mock.Anything, []byte(dat))
}
func TestSingleMessageRegistrySendAllSendsToAllRegisteredCriers(t *testing.T) {
	dat := "test"
	m1 := criers.MockCrier{}
	m1.On("Send", mock.Anything, []byte(dat)).Return(nil)
	m2 := criers.MockCrier{}
	m2.On("Send", mock.Anything, []byte(dat)).Return(nil)
	r := SingleMessageRegistry{}
	r.Register("m1", &m1)
	r.Register("m2", &m2)
	expReg := map[string]criers.Crier{
		"m1": &m1,
		"m2": &m2,
	}
	assert.Equal(t, expReg, r.criers)
	r.SendAll([]byte(dat))
	m1.AssertCalled(t, "Send", mock.Anything, []byte(dat))
	m2.AssertCalled(t, "Send", mock.Anything, []byte(dat))
}

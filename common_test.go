package checkmate

import (
	"testing"

	. "gopkg.in/check.v1"
)

type CheckmateSuite struct{}

var _ = Suite(&CheckmateSuite{})

func Test(t *testing.T) { TestingT(t) }
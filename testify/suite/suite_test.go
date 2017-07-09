package suite

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	log.Println("SetupTest")
	suite.VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TearDownTest() {
	log.Println("TearDownTest")
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	log.Println("TestExample")
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	// 运行ExampleTestSuite中Test开始的函数
	suite.Run(t, new(ExampleTestSuite))
}

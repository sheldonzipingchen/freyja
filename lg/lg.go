// lg short for log
package lg

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()
var stdFormatter *prefixed.TextFormatter
var fileFormatter *prefixed.TextFormatter

func Init() {
}

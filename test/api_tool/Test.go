package apitool_test

import (
	"testing"
)

//#TestFileCheck
//Purpose: This test checks the file read, validating json format and
//         required parts present to generate a unit test
//JSON name: refrence-01.json
//Format:
// {  "package_methodname" : {
//                   "signature" : {
//                           "1_<type>" : [from, this, to, this],
//                           "2_<type>" : [from, to]...
//                            },
//                  "return_type" : {"1_type" : [from, this, to, this],
//                                   "2_type" : [ ... ]
//                             },
//                  "exposed_variables" : {
//                                 "E1" : [from, this, to, this],
//                                   ...
//
//                                },
//                  "behaviour" : {
//                               "db" : {},
//                               "io" : {},
//                               "net" : {},
//                               "object" : {}
//                               },
//                  "usedby" : {}
//                  }
//
// }
//
func TestFileCheck(t *testing.T) {

	t.Error() // to indicate test failed
}

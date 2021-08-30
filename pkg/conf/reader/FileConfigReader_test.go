package reader

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	TEST_PRIORITY = 10
	tempFCR,err       = NewFileConfigReader([]string{os.Getenv("HOME")}, false, "base.yaml", "yaml", TEST_PRIORITY)
)

func TestValidateAndAbsolutePaths(t *testing.T){
	tests := []struct {
		name   string
		input  FileConfigReader
		output interface{}
	}{
		{
			name:   "Validation Success",
			input:  tempFCR,
			output: nil,
		},
		
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			err := test.input.validateAndAbsolutePaths()
			assert.Equal(t, test.output, err)
		})
	}
}

func TestFileConfigReader(t *testing.T) {
	tests := []struct {
		name   string
		input  FileConfigReader
		output interface{}
		methodName string
		args  []interface{}
	}{
		//TODO fixed below commented test case which is failing due to reflection(nil error return by called method)
		/*{
			name:   "ValidateAndAbsolutePaths Success",
			input:  tempFCR,
			output: nil,
			methodName: "validateAndAbsolutePaths",
		},*/
		{
			name:   "Get Priority Success",
			input:  tempFCR,
			output: TEST_PRIORITY,
			methodName: "GetPriority",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			//err := test.input.validateAndAbsolutePaths()
			output := invoke(test.input,test.methodName, test.args...)
			assert.Equal(t, test.output, output[0])
		})
	}
}

func invoke(anyStruct interface{}, methodName string, args ...interface{}) []interface{} {
	methodInputs := make([]reflect.Value, len(args))
	var returnValues []interface{}
	for i, _ := range args {
		methodInputs[i] = reflect.ValueOf(args[i])
	}

	values := reflect.ValueOf(anyStruct).MethodByName(methodName).Call(methodInputs)

	for _,v := range values{
		 returnValues = append(returnValues, v.Interface())
	}
	return returnValues
}

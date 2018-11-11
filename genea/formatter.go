package genea

import (
	"fmt"
	"io"

	"github.com/json-iterator/go"
)

func (tree *Tree) Format(out io.Writer, version Version) error {
	var data interface{}
	switch version {
	case VersionV1:
		// TODO: formatting for V1
	case VersionV2:
		// TODO: formatting for V2
	default:
		return fmt.Errorf("unsupported version `%v`", version)
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	n, err := out.Write(content)
	if err != nil {
		return err
	}
	if n != len(content) {
		return fmt.Errorf("incomplete write")
	}
	return nil
}

package stats

import (
	"io"

	"github.com/LouisBrunner/regenea/core"
	"github.com/LouisBrunner/regenea/genea"
	"github.com/LouisBrunner/regenea/report/procs"
)

func Process(tree *genea.Tree, out io.Writer, marshaller func(v interface{}) ([]byte, error)) error {
	processors := []procs.Processor{
		&procs.Counter{},
		&procs.Oldest{},
		&procs.Youngest{},
		&procs.Weird{},
	}

	processors2 := make([]core.Processor, len(processors), len(processors))
	for i := range processors {
		processors2[i] = core.Processor(processors[i])
	}

	core.ProcessTree(tree, processors2)

	result := map[string][]interface{}{}
	for _, proc := range processors {
		category, content := proc.Output()
		if _, ok := result[category]; !ok {
			result[category] = []interface{}{}
		}
		result[category] = append(result[category], content)
	}
	output, err := marshaller(result)
	if err != nil {
		return err
	}
	_, err = out.Write(output)
	return err
}

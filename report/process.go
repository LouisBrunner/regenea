package report

import (
	"io"

	"github.com/LouisBrunner/regenea/core"
	"github.com/LouisBrunner/regenea/genea"
	"github.com/LouisBrunner/regenea/report/procs"
)

// TODO: add "lived through <historic event>" processor
// TODO: add chronological timeline

func Process(tree *genea.Tree, out io.Writer, marshaller func(v interface{}) ([]byte, error)) error {
	processors := []procs.Processor{
		&procs.Counter{},
		&procs.Ages{},
		&procs.Weird{},
	}

	processors2 := make([]core.Processor, len(processors), len(processors))
	for i := range processors {
		processors2[i] = core.Processor(processors[i])
	}

	core.ProcessTree(tree, processors2)

	result := map[string]procs.StringMap{}
	for _, proc := range processors {
		category, content := proc.Output()
		if _, ok := result[category]; !ok {
			result[category] = procs.StringMap{}
		}
		for k, v := range content {
			result[category][k] = v
		}
	}
	output, err := marshaller(result)
	if err != nil {
		return err
	}
	_, err = out.Write(output)
	return err
}

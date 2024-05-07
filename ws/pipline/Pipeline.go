package pipline

import "architecture/ws/controlunit"

type Pipeline struct {
	ControlUnit *controlunit.ControlUnit
	Stages      []PipelineStage
}

type PipelineStage struct {
	Name string
	Func func()
}

func NewPipeline(controlunit *controlunit.ControlUnit) *Pipeline {
	pipline := &Pipeline{
		ControlUnit: controlunit,
		Stages: []PipelineStage{
			{Name: "Fetch", Func: controlunit.FetchIntstruction},
			{Name: "Decode", Func: controlunit.DecodeInstruction},
			{Name: "Execute", Func: controlunit.Execute},
		},
	}
	return pipline
}

func (p *Pipeline) Run() {
	for {
		p.ExecuteStages()
	}
}

func (p *Pipeline) ExecuteStages() {
	for _, stage := range p.Stages {
		stage.Func()
	}
}

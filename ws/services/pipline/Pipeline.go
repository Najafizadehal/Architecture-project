package pipeline

// import "architecture/ws/services/controlunit"

// // Pipeline represents the CPU pipeline with multiple stages.
// type Pipeline struct {
// 	ControlUnit *controlunit.ControlUnit
// 	Stages      []PipelineStage
// }

// // PipelineStage represents a single stage in the pipeline.
// type PipelineStage struct {
// 	Name string
// 	Func func()
// }

// // NewPipeline creates a new Pipeline with the given Control Unit.
// func NewPipeline(controlUnit *controlunit.ControlUnit) *Pipeline {
// 	return &Pipeline{
// 		ControlUnit: controlUnit,
// 		Stages: []PipelineStage{
// 			{Name: "Fetch", Func: controlUnit.Fetch},
// 			{Name: "Decode", Func: controlUnit.DecodeInstruction},
// 			{Name: "Execute", Func: controlUnit.Execute},
// 		},
// 	}
// }

// // Run executes the pipeline stages in a loop.
// func (p *Pipeline) Run() {
// 	for {
// 		p.ExecuteStages()
// 	}
// }

// // ExecuteStages runs each stage of the pipeline.
// func (p *Pipeline) ExecuteStages() {
// 	for _, stage := range p.Stages {
// 		stage.Func()
// 	}
// }

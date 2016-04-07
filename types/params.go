package types

type FuncRunCode func(p RunParams) RunResults

type RunParams struct {
	Language string `json:"lang" binding:"required"`
	Source   string `json:"source" binding:"required"`
	Spec     string `json:"spec" binding:"required"`
}

type JsonResult map[string]interface{}

type RunResults struct {
	Output interface{} `json:"output"`
}

type RunContext struct {
	Params  RunParams
	Results RunResults
	Path    string
}

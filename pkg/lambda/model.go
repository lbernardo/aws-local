package lambda

type ParamsLambda struct {
	Bin     string
	Env     map[string]string
	Runtime string
	Volume  string
	Body    string
}

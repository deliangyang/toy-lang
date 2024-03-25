package object

type Enviroment struct {
	envs map[string]Object
}

func NewEnviroment() *Enviroment {
	return &Enviroment{
		envs: make(map[string]Object),
	}
}

func (e *Enviroment) Get(name string) (Object, bool) {
	obj, ok := e.envs[name]
	return obj, ok
}

func (e *Enviroment) Set(name string, obj Object) Object {
	e.envs[name] = obj
	return obj
}

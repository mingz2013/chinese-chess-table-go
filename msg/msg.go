package msg

//type Params map[string]interface{}
//type Results map[string]interface{}
//
//type Msg struct {
//	Cmd string
//	Params Params
//	Results Results
//}

type Msg map[string]interface{}

func (m Msg) GetCmd() string {
	return m["cmd"].(string)
}

func (m Msg) SetCmd(cmd string) {
	m["cmd"] = cmd
}

func (m Msg) GetParams() map[string]interface{} {
	return m["params"].(map[string]interface{})
}

func (m Msg) SetParams(params map[string]interface{}) {
	m["params"] = params
}

func (m Msg) GetResults() map[string]interface{} {
	return m["results"].(map[string]interface{})
}

func (m Msg) SetResults(results map[string]interface{}) {
	m["results"] = results
}

func (m Msg) SetKey(key string, value interface{}) {
	m[key] = value
}

func (m Msg) GetKey(key string) interface{} {
	return m[key]
}

func NewMsg() Msg {
	return Msg{}
}

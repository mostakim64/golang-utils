package msgutil

type Data map[string]interface{}

type Msg struct {
	Data Data
}

func NewMessage() Msg {
	return Msg{
		Data: make(Data),
	}
}

func (m Msg) Set(key string, value interface{}) Msg {
	m.Data[key] = value
	return m
}

func (m Msg) Done() Data {
	return m.Data
}

func RequestBodyParseErrorResponseMsg() Data {
	return NewMessage().Set("message", "Failed to parse request body").Done()
}

func InvalidEmailMsg() Data {
	return NewMessage().Set("message", "Invalid email").Done()
}

func TooManyRequestErrorResponseMsg() Data {
	return NewMessage().Set("message", "Too many requests").Done()
}

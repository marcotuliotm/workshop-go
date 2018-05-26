package domain

// User usario
type User struct {
	Name       string                 `json:"name,omitempty"`
	Age        int                    `json:"age,omitempty"`
	Phone      []string               `json:"phone,omitempty"`
	Parentesco map[string]interface{} `json:"parentesco,omitempty"`
}

//GetName retorna o nome
func (me *User) GetName() string {
	return me.Name
}

//GetAge retorna a idade
func (me *User) GetAge() int {
	return me.Age
}

//AddParentes add parentes
func (me *User) AddParentes(key string, value interface{}) {
	me.Parentesco[key] = value
}

func (me *User) String() string {
	return me.Name
}

package gomap

type gmap interface {
	getKeys() []string
}

type StringBoolMap map[string]bool

func (m StringBoolMap) GetKeys() []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

type IntBoolMap map[int]bool

func (m IntBoolMap) GetKeys() []int {
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

type StringStringMap map[string]string

func (m StringStringMap) GetKeys() []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

type StringIntMap map[string]int

func (m StringIntMap) GetKeys() []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

type IntStringMap map[int]string

func (m IntStringMap) GetKeys() []int {
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

type IntIntMap map[int]int

func (m IntIntMap) GetKeys() []int {
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}
# gomap
This is a customized type based on map[K]V, adding function GetKeys(). Since Go 1.x doesn't support generics type, only 6 six common pairs of types are supported by the module.

Import it in your program as:
```go
      import "github.com/pochard/gomap"
```

## API
### type StringBoolMap
	func GetKeys() ([]string) 
### type StringStringMap
	func GetKeys() ([]string) 
### type StringIntMap
	func GetKeys() ([]string) 
### type IntStringMap
	func GetKeys() ([]int) 
### type IntIntMap
	func GetKeys() ([]int) 
### type IntBoolMap
	func GetKeys() ([]int) 

## Example
``` go
func main(){
	sbmap := make(gomap.StringBoolMap)
	sbmap["China"] = true
	sbmap["USA"] = true
 	sbmap["Japan"] = true

	fmt.Printf("StringBoolMap keys:\n")
	for _,key := range sbmap.GetKeys(){
 		fmt.Printf("%s\n", key)
 	}

	ssmap := make(gomap.StringStringMap)
	ssmap["China"] = "Chinese"
	ssmap["USA"] = "English"
 	ssmap["Japan"] = "Japanese"

	fmt.Printf("StringStringMap keys:\n")
	for _,key := range ssmap.GetKeys(){
 		fmt.Printf("%s\n", key)
 	}

	simap := make(gomap.StringIntMap)
	simap["China"] = 145
	simap["USA"] = 6756
	simap["Japan"] = 987

	fmt.Printf("StringIntMap keys:\n")
	for _,key := range simap.GetKeys(){
		fmt.Printf("%s\n", key)
	}

	ismap := make(gomap.IntStringMap)
	ismap[145] = "China"
	ismap[6756] = "USA"
	ismap[987] = "Japan"

	fmt.Printf("IntStringMap keys:\n")
	for _,key := range ismap.GetKeys(){
		fmt.Printf("%d\n", key)
	}

	iimap := make(gomap.IntIntMap)
	iimap[145] = 12
	iimap[6756] = 24
	iimap[987] = 32

	fmt.Printf("IntIntMap keys:\n")
	for _,key := range iimap.GetKeys(){
		fmt.Printf("%d\n", key)
	}

	ibmap := make(gomap.IntBoolMap)
	ibmap[145] = true
	ibmap[6756] = true
	ibmap[987] = true

	fmt.Printf("IntBoolMap keys:\n")
	for _,key := range ibmap.GetKeys(){
		fmt.Printf("%d\n", key)
	}
}

```
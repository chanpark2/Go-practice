package main

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		println("int", t)
	case string:
		println("string", t)
	case bool:
		println("bool", t)
	default:
		println("unknown")
	}
}
func main() {

}

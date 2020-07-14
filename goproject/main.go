package main

func isValid(s string) bool {
	slice := []byte{}
	ret := true
	mapList := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for i := 0; i < len(s); i++ {
		value := s[i]
		if value == '(' || value == '{' || value == '[' {
			slice = append(slice, value)
		}
		if value == ')' || value == '}' || value == ']' {
			if 0 == len(slice) {
				ret = false
				break
			}
			topKey := slice[len(slice)-1]
			topValue, exists := mapList[topKey]
			if !exists || topValue != value {
				ret = false
				break
			}
			slice = slice[:len(slice)-1]
		}
	}
	return ret
}

func main() {

}

package lookup

func Lookup(hash string) (string, int, string) {
	position := 0
	stingz := ""
	file := "abs"
	hashmap := map[string][]byte{}
	for key, val := range hashmap {
		if hash == key {
			stingz = string(val)
		}
	}
	return file, position, stingz
}

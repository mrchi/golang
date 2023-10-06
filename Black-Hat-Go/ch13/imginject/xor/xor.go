package xor

func xor(input []byte, key []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}
	return output
}

func XorEncode(input []byte, key string) []byte {
	return xor(input, []byte(key))
}

func XorDecode(input []byte, key string) []byte {
	return xor(input, []byte(key))
}

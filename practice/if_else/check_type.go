package checktype

func CheckAnimal(sound string) string {
	if sound == "maa" {
		return "cow"
	} else if sound == "meow" {
		return "cat"
	} else {
		return "nothing"
	}
}
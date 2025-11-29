package project

func Ponctuations(slicee []string) string {
	tempstr := ""
	tempslice := []string{}

	for i, element := range slicee {
		if i != len(slicee)-1 {
			tempstr += element + " "
		} else {
			tempstr += element
		}
	}

	for _, ch := range tempstr {
		tempslice = append(tempslice, string(ch))
	}

	for i := 0; i < len(tempslice); i++ {
		ch := tempslice[i]
		if ch != "," && ch != "." && ch != ":" && ch != ";" && ch != "!" && ch != "?" {
			continue
		}
		if i > 0 && tempslice[i-1] == " " {
			tempslice = append(tempslice[:i-1], tempslice[i:]...)
			i--
		}
		if i+1 < len(tempslice) && tempslice[i+1] != " " {
			tempslice = append(tempslice[:i+1], append([]string{" "}, tempslice[i+1:]...)...)
		}
	}

	result := ""
	for _, ch := range tempslice {
		result += ch
	}
	return result
}

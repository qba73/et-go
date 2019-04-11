package naming

func Color(name string) string {
	switch name {
	case "blue":
		return "#0000FF"
	case "white":
		return "#FFFFFF"
	default:
		return "#000000"
	}
}

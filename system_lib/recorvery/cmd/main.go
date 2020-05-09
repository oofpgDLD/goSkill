package main


func main() {
}

func toInt32(o interface{}) int32{
	switch t := o.(type) {
	case int32:
		return int32(t)
	default:
		return int32(t)
	}
}

package _2_adapter

type AC interface {
	OutputAC() int
}

//中国220V
type AC220 struct {
}

func (t *AC220)OutputAC() int{
	return 220
}

//日本110V
type AC110 struct {
}

func (t *AC110)OutputAC() int{
	return 110
}

//被适配这目标
type DC5Adapter interface {
	OutputDC5(ac AC) int
}

type ChinaPowerAdapter struct {
}

func (t *ChinaPowerAdapter) OutputDC5(ac AC) int{
	return ac.OutputAC() / 44
}


type JapanPowerAdapter struct {
}

func (t *JapanPowerAdapter) OutputDC5(ac AC) int{
	return ac.OutputAC() / 22
}

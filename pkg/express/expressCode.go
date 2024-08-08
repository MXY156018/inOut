package express

type CpCode string

const SF1 CpCode = "01"  //顺丰
const SF2 CpCode = "30"  //顺丰
const SF3 CpCode = "10"  //顺丰
const SF4 CpCode = "26"  //顺丰跨境
const SF5 CpCode = "-9"  //顺丰
const SF6 CpCode = "-14" //顺丰快运
const SF7 CpCode = "-16" //顺丰快运
const YT1 CpCode = "05"  //圆通
const YT2 CpCode = "34"  //圆通
const YT3 CpCode = "58"  //圆通承诺达
const YT4 CpCode = "-5"  //圆通承诺达

func (c CpCode) String() string {
	switch c {
	case SF1:
		return "顺丰快递"
	case SF2:
		return "顺丰快递"
	case SF3:
		return "顺丰快递"
	case SF4:
		return "顺丰跨境"
	case SF5:
		return "顺丰快递"
	case SF6:
		return "顺丰快运"
	case SF7:
		return "顺丰快运"
	case YT1:
		return "圆通快递"
	case YT2:
		return "圆通快递"
	case YT3:
		return "圆通承诺达"
	case YT4:
		return "圆通承诺达"
	}
	return c.String()
}

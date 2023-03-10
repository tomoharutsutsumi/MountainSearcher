package domain

type Weather struct {
  Daily struct {
	Weathercode []int
  }
}

func GetWeatherName(code int) string {
  codeMap := map[int]string{
	0: "晴天",
	1: "ほぼ晴天",
	2: "所々曇り",
	3: "曇り",
	45: "霧",
	48: "深い霧",
	51: "軽い小雨",
	53: "少し強い小雨",
	55: "強い小雨",
	56: "軽い凍る小雨",
	57: "強い凍る小雨",
	61: "軽い雨",
	63: "少し強い雨",
	65: "大雨",
	66: "軽い凍る雨",
	67: "凍る大雨",
	71: "弱い雪",
	73: "少し強い雪",
	75: "強い雪",
	77: "ひょう",
	80: "豪雨",
	81: "強い豪雨",
	82: "大豪雨",
	85: "豪雪",
	86: "大豪雪",
	95: "台風",
	96: "弱いひょうのある台風",
	99: "強いひょうのある台風",
  }
  return codeMap[code]
}
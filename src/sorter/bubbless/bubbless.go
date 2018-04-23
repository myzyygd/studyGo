package bubbless

type SorterData []int

func (a *SorterData) Bubbless() []int {
	flag := true
	tmpData := *a
	for i := 0; i < len(tmpData)-1; i++ {
		flag = true
		for j := 0; j < len(tmpData)-i-1; j++ {
			if tmpData[j] > tmpData[j+1] {
				tmpData[j], tmpData[j+1] = tmpData[j+1], tmpData[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
	return tmpData
}

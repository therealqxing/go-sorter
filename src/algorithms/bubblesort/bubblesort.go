package bubblesort

func BubbleSort(values []int) {
	flag := true

	for i := 0; i < len(values) - 1; i++ {
		flag = true
		for j := 0; j < len(values) - i - 1; j++ {
			if values[j] > values [j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}

		if flag == true {
			//如果已经顺序对了，就不用继续冒泡排序了
			break
		}
	}
}
package main

// Ex は
// 与えられた組み合わせ数とリストと結果数から、
// 部分和問題の解の結果と解のリストを返します。
func Ex(listNum int, list []int, resolve int) (success bool, sumList []int) {
	for i := 0; i < len(list); i++ {
		if list[i] == resolve {
			success = true
			sumList = append(sumList, list[i])
			return success, sumList
		}
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == resolve {
				success = true
				sumList = append(sumList, list[i])
				sumList = append(sumList, list[j])
				return success, sumList
			}
			for k := j + 1; k < len(list); k++ {
				if list[i]+list[j]+list[k] == resolve {
					success = true
					sumList = append(sumList, list[i])
					sumList = append(sumList, list[j])
					sumList = append(sumList, list[k])
					return success, sumList
				}
			}
		}
	}
	return false, sumList
}

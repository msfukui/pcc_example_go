package main

// Ex は
// 与えられた組み合わせ数とリストと結果数から、
// 部分和問題の解の結果と解のリストを返します。
func Ex(listNum int, list []int, resolve int) (success bool, sumList []int) {
	iList := list
	for i := 0; i < len(iList); i++ {
		if iList[i] == resolve {
			success = true
			sumList = append(sumList, iList[i])
			return success, sumList
		}
		jList := deleteAt(iList, i)
		for j := 0; j < len(jList); j++ {
			if iList[i]+jList[j] == resolve {
				success = true
				sumList = append(sumList, iList[i])
				sumList = append(sumList, jList[j])
				return success, sumList
			}
			kList := deleteAt(jList, j)
			for k := 0; k < len(kList); k++ {
				if iList[i]+jList[j]+kList[k] == resolve {
					success = true
					sumList = append(sumList, iList[i])
					sumList = append(sumList, jList[j])
					sumList = append(sumList, kList[k])
					return success, sumList
				}
			}
		}
	}
	return false, sumList
}

// delete_at は
// 与えられたリストとインデックス番号から、
// インデックス番号の要素を削除した結果のリストを返します。
func deleteAt(list []int, i int) (deleteList []int) {
	if i > len(list) || i < 0 {
		return list
	}
	deleted := make([]int, len(list)-1)
	if i == 0 {
		copy(deleted, list[i+1:len(list)])
		deleteList := deleted
		return deleteList
	}
	copy(deleted, list[0:i-1])
	for j := i + 1; j < len(list); j++ {
		deleted = append(deleted, list[j])
	}
	deleteList = deleted
	return deleteList
}

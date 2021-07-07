package leetcode

func DisplayTable(orders [][]string) [][]string {
	strHeard := make([]string, 0)
	strHeards := make([][]string, 0)
	mapOrders := make(map[string]int, 0)
	strHeard = append(strHeard, "Table")
	for _, val := range orders {
		strHeard = append(strHeard, val[2])
		if mapOrders[val[2]] != 0 {
			mapOrders[val[2]] = mapOrders[val[2]] + 1
		} else {
			mapOrders[val[2]] = 1
		}
	}
	strHeards = append(strHeards, strHeard)

	return nil
}

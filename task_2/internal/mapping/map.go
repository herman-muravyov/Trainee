package mapping

func Mapping(slice []int) []int {
    if len(slice) == 0 {
        return nil
    }

    setMap := make(map[int]struct{}, len(slice))
    for _, value := range slice {
        setMap[value] = struct{}{}
    }

    set := make([]int, 0, len(setMap))
    for key := range setMap {
        set = append(set, key)
    }

    return set
}

/* Best solution without Map */
func RemoveDuplicates(slice []int) int {
    if len(slice) == 0 {
        return 0
    }

    k := 0

    for i, _ := range(slice) {
        if slice[i] != slice[k] {
            k++
            slice[k] = slice[i]
        }
    }

    return k + 1
}
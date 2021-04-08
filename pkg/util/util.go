package util

import "github.com/spf13/cast"

/*
*@Author Administrator
*@Date 7/4/2021 11:22
*@desc
 */
func IntInlice(i int, arr []int) bool {
	for _, v := range arr {

		if v == i {
			return true
		}
	}
	return false
}
func StringInSlice(s string, sl []string) bool {
	for _, s2 := range sl {
		if s2 == s {
			return true
		}
	}
	return false

}

/*
*@Author Administrator
*@Date 7/4/2021 11:27
*@desc
 */
func IntInStringSlice(i int, s []string) bool {
	v := cast.ToString(i)
	for _, s2 := range s {
		if v == s2 {
			return true
		}
	}
	return false
}

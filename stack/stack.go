package stack

type Stack struct {
	s []int
}

func (st *Stack) Append(ele int) []int {
	st.s = append(st.s, ele)
	return st.s
}

func (st *Stack) Top() int {
	return st.s[len(st.s)-1]
}

func (st *Stack) Pop() []int {
	st.s = st.s[:len(st.s)-1]
	return st.s
}

//保持其他元素序列, 后半截前移覆盖
func (st *Stack) Remove(i int) []int {
	copy(st.s[i:], st.s[i+1:])
	return st.s[:len(st.s)-1]
}

//不保存其它元素序列, 末位前移占位
func (st *Stack) RemoveFast(i int) []int {
	st.s[i] = st.s[len(st.s)-1]
	return st.s[:len(st.s)-1]
}

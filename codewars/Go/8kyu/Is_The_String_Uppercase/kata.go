package kata

import "fmt"

func main() {
	s := make([]string, 3)
	/*
		배열과 다르게 슬라이스는 원소의 갯수가 아닌 포함된 원소들로만 작성됨.
		길이가 0인 빈 배열을 만들기 위해선 내장 함수 make를 사용
		위 코드는 길이가 3인 문자열 배열을 만듦
	*/
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	/*
		len은 알다시피 슬라이스의 길이를 반환함
		이러한 기본 연산과 함께, 슬라이스는 배열보다 풍부한 기능들을
		지원함. 그 중 하나는 새로운 값이 추가된 슬라이스를 반환하는
		append 함수임. 주의할건 새로운 슬라이스를 사용하기위해선
		append 로부터 반환되는 값을 사용해야함
	*/

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	/*
		슬라이스는 복사가 가능함. 위 코드는 s와 같은 길이를 갖는
		빈 슬라이스 c를 생성하고, s를 c로 복사함
	*/

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	/*
		한 줄로 슬라이스 선언 및 초기화 하는 것이 가능함
	*/

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	/*
		슬라이스는 다차원 데이터를 구성할 수도 있음.
		다차원 배열과 다르게 다중 슬라이스의 내부 슬라이스들은 가변적인 길이를
		가질 수 있음
	*/
}

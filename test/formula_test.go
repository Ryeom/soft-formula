package test

import (
	"fmt"
	"regexp"
	"strconv"
	"test/calculator"
	"testing"
)

/*
	formula : 공식
	1. 수식분석기 : 토큰 분리 및 타입 판별
		1) 수식을 토큰으로 분리
		2) 토큰 타입 판별
	2. 계산기 엔진 : 분석된 토큰을 바탕으로 계산 수행 (스택,후위연산)
		1) 사용자로부터 입력받은 수식을 수식분석기에 전달
		2) 수식분석기에서는 중위 표기법으로 표현된 수식을 후위 표기법으로 변환하여 계산하기 쉽게 준비
		3) 준비된 후위 표기법 수식을 스택을 이용하여 계산 스택에는 피연산자들과 연산자들이 차례로 저장
		4) 수식을 한 번 순회할 때마다, 연산자를 만나면 스택에서 피연산자를 두 개 꺼내서 연산을 수행하고, 그 결과를 다시 스택에 넣음.
		5) 수식을 모두 순회한 후, 스택에는 최종 결과값이 남게 됩니다. 이 값을 출력해주면 계산기의 계산 끝.
	3. 유틸리티 함수 : 수식의 유효성검사 및 분수연산등의 추가적인 기능

	<연산자 우선순위>
	1. 괄호: ()
	2. 단항 연산자: +, -
	3. 지수 연산자: ^
	4. 곱셈/나눗셈 연산자: *, /, %
	5. 덧셈/뺄셈 연산자: +, -
*/

func RunOp(sign string, a, b float64) float64 {
	s := byte(sign[0])
	switch s {
	case byte(43):
		return a + b
	case byte(45):
		return a - b
	case byte(47):
		return a / b
	case byte(42):
		return a * b
	default:
		return 0
	}
}

func TestFormula(t *testing.T) {
	//a := &Calculation{InfixFormula: "(a+b)*c+d", Parameters: []interface{}{1, 2.3, 2, 2}}
	//
	//fmt.Println(a.Init())
	//b := &Calculation{PostfixFormula: "12+3/*4+", Parameters: []interface{}{4, 5.1, 7, 4}}
	//
	//fmt.Println(b.Init())
	//a, aerr := calculator.NewCalculation("postfix", "(a+b)*Z+d", []interface{}{1, 2.3, 2, 2})
	b, berr := calculator.NewCalculation("infix", "ab+z*d+", []interface{}{4, 5.1, 7, 4})

	//fmt.Println("a : ", a, aerr)
	fmt.Println("b : ", b, berr)

	//a.Purpose = []string{"a", "b"}
	b.Purpose = []string{"c", "d"}

	fmt.Println("------------------")
	//fmt.Println("a : ", a, aerr)
	fmt.Println("b : ", b, berr)

}

//func (c *Calculation) Init() *Calculation {
//	if c.PostfixFormula == "" {
//		p, err := infixToPostfix(c.PostfixFormula)
//		if err != nil {
//			fmt.Println(p, err)
//		}
//		c.InfixFormula = p
//	} else if c.InfixFormula == "" {
//		i, err := postfixToInfix(c.InfixFormula)
//		if err != nil {
//			fmt.Println(i, err)
//		}
//		c.PostfixFormula = i
//	}
//	return c
//}

func avg(list []float64) float64 {
	var total float64
	for _, v := range list {
		total = total + v
	}
	return total / float64(len(list))
}

func add[T calculator.Number](a, b T) T {
	return a + b
}

func TestTransfer(t *testing.T) {
	//s := "(mother_age+father_age+sister_age)/baby*2.0"
	s := "(a2+5b+m_7)/x*2.0"
	//ss := "a2 5b + 64 * 2 8 + /"

	list := calculator.PlanTextToInfixFormula(s)
	i := infixToPostfix(list)
	fmt.Println("결과 i : ", i)
	infix := []string{"(", "62", "+", "144", ")", "*", "4", "/", "(", "120", "+", "80", ")"}
	//postfix := []string{"A", "B", "+", "C", "*", "D", "E", "+", "/"}
	i = infixToPostfix(infix)
	fmt.Println("결과 i : ", i)
	fmt.Println("-0-------")

	postfix := []string{"21", "5", "+", "64", "*", "2", "8", "+", `/`}
	p := postfixToInfix(postfix)
	fmt.Println("결과 p : ", p)
	//
	//var nr []interface{}
	//for _, w := range p {
	//	if len(w) == 0 {
	//		w1 := rune(w[0])
	//		if calculator.IsOperator(w1) || calculator.IsBracket(w1) {
	//			nr = append(nr, w)
	//			continue
	//		}
	//		v, errpf := strconv.ParseFloat(w, 64)
	//		if errpf != nil {
	//			nr = append(nr, w)
	//			continue
	//		}
	//		nr = append(nr, v)
	//	} else {
	//		v, errpf := strconv.ParseFloat(w, 64)
	//		if errpf != nil {
	//			nr = append(nr, w)
	//			continue
	//		}
	//		nr = append(nr, v)
	//	}
	//}
	v, _ := calculator.EvaluatePostfix(i)

	fmt.Println("result !!!!!", v)

}

//후위를 중위로~
func postfixToInfix(postfix []string) []string {
	stack := make([]string, 0)
	for _, token := range postfix {
		if isOperator(token) {
			// 연산자인 경우
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			infix := "(" + operand1 + " " + token + " " + operand2 + ")"
			stack = append(stack, infix)
		} else {
			// 피연산자인 경우
			stack = append(stack, token)
		}
	}
	return stack
}

func isOperator(token string) bool {
	if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
		return true
	}
	return false
}
func TestRegexp(t *testing.T) {
	matched, _ := regexp.MatchString("[a-zA-Z0-9_]+", "a2")
	fmt.Println(matched)

}

//중위를 후위로~
func infixToPostfix(tokens []string) []string {
	var stack []string
	var postfix []string
	operators := map[string]int{
		"+": 1, "-": 1,
		"*": 2, "/": 2, "%": 2,
		"^": 3,
	}

	for _, token := range tokens {
		matched, _ := regexp.MatchString("[a-zA-Z0-9_]+", token)
		fmt.Println(matched)
		if isNumber(token) || matched { // 혹은 변수인지 ?
			postfix = append(postfix, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil // 괄호가 맞지 않는 경우 nil 반환
			}
			// 왼쪽 괄호 제거
			stack = stack[:len(stack)-1]
		} else if _, ok := operators[token]; ok {
			for len(stack) > 0 && stack[len(stack)-1] != "(" &&
				operators[stack[len(stack)-1]] >= operators[token] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else {
			return nil // 올바른 토큰이 아닌 경우 nil 반환
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil // 괄호가 맞지 않는 경우 nil 반환
		}
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

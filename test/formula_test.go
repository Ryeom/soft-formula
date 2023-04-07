package test

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
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

type Number interface {
	constraints.Integer | constraints.Float
}
type Calculation struct {
	InfixFormula   string // 연산자 중위 표기법
	PostfixFormula string // 연산자 후위 표기법
	Parameters     []interface{}
	result         float64
}

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
	a := &Calculation{InfixFormula: "(a+b)*c+d", Parameters: []interface{}{1, 2.3, 2, 2}}

	fmt.Println(a.Init())
	b := &Calculation{PostfixFormula: "12+3/*4+", Parameters: []interface{}{4, 5.1, 7, 4}}

	fmt.Println(b.Init())
}
func (c *Calculation) Init() *Calculation {
	if c.PostfixFormula == "" {
		p, err := infixToPostfix(c.PostfixFormula)
		if err != nil {

		}
		c.InfixFormula = p
	} else if c.InfixFormula == "" {
		i, err := postfixToInfix(c.InfixFormula)
		if err != nil {

		}
		c.PostfixFormula = i
	}
	return c
}

func postfixToInfix(postfix string) (string, error) {
	var stack []string

	for _, char := range postfix {
		if 65 <= char && char <= 122 {
			stack = append(stack, string(char))
		} else {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid postfix expression: %s", postfix)
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("(%s %c %s)", left, char, right))
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid postfix expression: %s", postfix)
	}

	return stack[0], nil
}

func infixToPostfix(input string) (string, error) {
	fmt.Println([]rune("azAZ"))
	// 연산자 우선순위 정의
	precedence := map[rune]int{'*': 3, '/': 3, '+': 2, '-': 2}

	var output string // 결과를 담을 변수
	var stack []rune  // 스택
	fmt.Println()
	for _, char := range input {
		fmt.Println(reflect.TypeOf(char))
		switch {
		//case unicode.IsDigit(char):
		case 65 <= char && char <= 122:
			output += string(char) // 숫자이면 결과에 추가
		case char == '(':
			stack = append(stack, char) // 여는 괄호는 스택에 추가
		case char == ')':
			// 닫는 괄호를 만날 때까지 스택에서 연산자를 pop하고 결과에 추가
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				output += string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return "", errors.New("mismatched parentheses")
			}
			stack = stack[:len(stack)-1] // 여는 괄호 pop
		case isOperator(char):
			// 스택에서 우선순위가 높은 연산자를 pop하고 결과에 추가
			for len(stack) > 0 && stack[len(stack)-1] != '(' &&
				precedence[char] <= precedence[stack[len(stack)-1]] {
				output += string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, char) // 현재 연산자를 스택에 추가
		default:
			return "", fmt.Errorf("unknown character '%v'", string(char))
		}
	}

	// 스택에 남아있는 연산자를 모두 pop하여 결과에 추가
	for len(stack) > 0 {
		if stack[len(stack)-1] == '(' {
			return "", errors.New("mismatched parentheses")
		}
		output += string(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return output, nil
}

// 주어진 문자가 연산자인지 확인하는 함수
func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func (c Calculation) calc() float64 {
	r := 0.0

	c.result = r
	return r
}

func avg(list []float64) float64 {
	var total float64
	for _, v := range list {
		total = total + v
	}
	return total / float64(len(list))
}

func add[T Number](a, b T) T {
	return a + b
}

package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStackCalc(t *testing.T) {
	var stack []int

	for {
		var input string
		fmt.Scan(&input)

		// 입력이 "q"일 경우 프로그램 종료
		if input == "q" {
			break
		}

		// 입력이 숫자일 경우 스택에 추가
		number, err := strconv.Atoi(input)
		if err == nil {
			stack = append(stack, number)
			continue
		}

		// 입력이 연산자일 경우 스택에서 값을 꺼내서 계산
		if len(stack) < 2 {
			fmt.Println("스택에는 최소한 2개의 숫자가 필요합니다.")
			continue
		}

		var result int

		switch input {
		case "+":
			result = stack[len(stack)-2] + stack[len(stack)-1]
		case "-":
			result = stack[len(stack)-2] - stack[len(stack)-1]
		case "*":
			result = stack[len(stack)-2] * stack[len(stack)-1]
		case "/":
			result = stack[len(stack)-2] / stack[len(stack)-1]
		default:
			fmt.Println("잘못된 입력입니다.")
			continue
		}

		// 계산 결과를 스택에 추가
		stack = stack[:len(stack)-2]
		stack = append(stack, result)

		fmt.Println(result)
	}
}

func precedence(s string) int {
	switch s {
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	default:
		return 0
	}
}

// 중위 표기법을 후위 표기법으로 변환하는 함수
func InfixToPostfix(infix string) string {
	// 우선순위 설정
	precedence := map[rune]int{'*': 3, '/': 3, '+': 2, '-': 2, '(': 1}
	stack := []rune{}
	postfix := []string{}

	// 중위 표기식 문자열을 반복해서 탐색
	for _, char := range infix {
		switch {
		case char >= '0' && char <= '9': // 피연산자인 경우
			postfix = append(postfix, string(char))
		case char == '(':
			stack = append(stack, char)
		case char == ')':
			// '('를 만날 때까지 스택에서 pop한 연산자를 후위 표기식에 추가
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				postfix = append(postfix, string(top))
			}
			// '('를 버림
			stack = stack[:len(stack)-1]
		default: // 연산자인 경우
			for len(stack) > 0 && precedence[char] <= precedence[stack[len(stack)-1]] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				postfix = append(postfix, string(top))
			}
			stack = append(stack, char)
		}
	}
	// 스택에 남아있는 연산자들을 후위 표기식에 추가
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		postfix = append(postfix, string(top))
	}

	// 후위 표기식 문자열을 반환
	return strings.Join(postfix, " ")
}

// 중위에서 후위로 계산식 표기법 변경
func TestMiddleToAfter(t *testing.T) {
	// 중위 표기식을 후위 표기식으로 변환하는 함수 테스트
	testCases := []struct {
		infix    string
		expected string
	}{
		{"1+2", "1 2 +"},
		{"1+2*3", "1 2 3 * +"},
		{"1*2+3", "1 2 * 3 +"},
		{"(1+2)*3", "1 2 + 3 *"},
		{"1+(2+3)*4", "1 2 3 + 4 * +"},
		{"(1+2)*(3+4)", "1 2 + 3 4 + *"},
		{"1+2-3*4/5", "1 2 + 3 4 * 5 / -"},
		{"(1+2)*(3-4)/5", "1 2 + 3 4 - * 5 /"},
		{"(1+2*3)/(4-5*6)", "1 2 3 * + 4 5 6 * - /"},
		{"1", "1"},
		{"1+2", "1 2 +"},
		{"(1+2)*3", "1 2 + 3 *"},
		{"1*2+3", "1 2 * 3 +"},
		{"1+2*3", "1 2 3 * +"},
		{"1+(2+3)*4", "1 2 3 + 4 * +"},
		{"(1+2)*(3+4)", "1 2 + 3 4 + *"},
		{"1+2-3*4/5", "1 2 + 3 4 * 5 / -"},
		{"(1+2)*(3-4)/5", "1 2 + 3 4 - * 5 /"},
		{"1+2*3-4/5", "1 2 3 * + 4 5 / -"},
		{"(1+2)*3-4*5/6", "1 2 + 3 * 4 5 * 6 / -"},
	}

	for i, tc := range testCases {
		output := InfixToPostfix(tc.infix)
		if output != tc.expected {
			fmt.Printf("Test case %d failed: expected %s but got %s\n", i+1, tc.expected, output)
		}
		fmt.Println(output)
	}

	return
	// 후위 표기식을 중위 표기식으로 변환하는 함수 테스트
	testCases2 := []struct {
		postfix string
		infix   string
		err     bool
	}{
		{"23+", "(2 + 3)", false},
		{"23*5+", "(2 * 3 + 5)", false},
		{"23*5-4+", "(2 * 3 - 5 + 4)", false},
		{"23+5*4+", "(2 + 3 * 5 + 4)", false},
		{"23+5*4", "", true},
		{"23+", "", true},
		{"+", "", true},
		{"23+", "(2 + 3)", false},
		{"23*5+", "(2 * 3 + 5)", false},
		{"23*5-4+", "(2 * 3 - 5 + 4)", false},
		{"23+5*4+", "(2 + 3 * 5 + 4)", false},
		{"23+5*4", "", true},
		{"23+", "", true},
		{"+", "", true},
		{"23+56+*", "((2 + 3) * (5 + 6))", false},
		{"23+56*+", "(2 + 3) + (5 * 6)", false},
		{"123+*4+", "((1 * 2) + 3) + 4", false},
		{"123*+4+", "(1 + (2 * 3)) + 4", false},
		{"12+34+*", "((1 + 2) * 3) + 4", false},
		{"12+3*4+", "(1 + 2 * 3) + 4", false},
		{"1234-*/", "", true},
		{"12+3-4+", "", true},
		{"12+3/*4+", "", true},
		{"12+3*4*/", "", true},
		{"12+3*4-", "", true},
		{"+23", "", true},
		{"23+", "(2 + 3)", false},
		{"23+45+*", "(2 + 3) * (4 + 5)", false},
		{"23+45+*6+", "(2 + 3) * (4 + 5) + 6", false},
	}

	for i, tc := range testCases2 {
		infix, err := PostfixToInfix(tc.postfix)
		if (err != nil) != tc.err {
			fmt.Printf("Test case %d failed: expected error %v, but got %v\n", i+1, tc.err, err)
			continue
		}
		if infix != tc.infix {
			fmt.Printf("Test case %d failed: expected %s, but got %s\n", i+1, tc.infix, infix)
			continue
		}
		fmt.Printf("Test case %d passed\n", i+1)
	}
}

func PostfixToInfix(postfix string) (string, error) {
	var stack []string

	for _, c := range postfix {
		if c >= '0' && c <= '9' {
			stack = append(stack, string(c))
		} else {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid postfix expression: %s", postfix)
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("(%s %c %s)", left, c, right))
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid postfix expression: %s", postfix)
	}

	return stack[0], nil
}

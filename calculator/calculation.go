package calculator

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
)

type Number interface {
	constraints.Integer | constraints.Float
}
type Calculation struct {
	InfixFormula   string // 연산자 중위 표기법
	PostfixFormula string // 연산자 후위 표기법
	Parameters     []interface{}
	Purpose        []string
	Result         float64
}

func NewCalculation(formulaType, formula string, parameters []interface{}) (*Calculation, error) {
	if formulaType == "postfix" {
		p, err := infixToPostfix(formula)
		if err != nil {
			fmt.Println(p, err)
			return nil, nil
		}
		return &Calculation{InfixFormula: formula, PostfixFormula: p, Parameters: parameters}, nil
	} else if formulaType == "infix" {
		i, err := postfixToInfix(formula)
		if err != nil {
			fmt.Println(i, err)
			return nil, nil
		}
		return &Calculation{InfixFormula: i, PostfixFormula: formula, Parameters: parameters}, nil
	}
	return nil, nil
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
			fmt.Println(left, " / ", char, " / ", right)
			stack = append(stack, fmt.Sprintf("(%s %c %s)", left, char, right))
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid postfix expression: %s", postfix)
	}
	fmt.Println("postfixToInfix stack : ", stack, len(stack))
	return stack[0], nil
}

func infixToPostfix(input string) (string, error) {
	//fmt.Println([]rune("azAZ"))
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
		case IsOperator(char):
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

	fmt.Println("infixToPostfix : ", stack)
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
func IsOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func IsBracket(c rune) bool {
	return c == ')' || c == '('
}

func PlanTextToInfixFormula(s string) []string {
	var list []string
	var word string
	for i, char := range s {
		if !IsOperator(char) && !IsBracket(char) {
			word = word + string(char)
			if len(s) <= i+1 {
				list = append(list, word)
				break
			}
			if len(s) > i && (IsOperator(rune(s[i+1])) || IsBracket(rune(s[i+1]))) {
				list = append(list, word)
				word = ""
			}
		} else {
			list = append(list, string(char))
		}
	}
	return list
}
func EvaluatePostfix(expr []interface{}) (float64, error) {
	stack := make([]float64, 0)
	for _, token := range expr {
		switch t := token.(type) {
		case float64:
			stack = append(stack, t)
		case string:
			switch t {
			case "+":
				if len(stack) < 2 {
					return 0, errors.New("invalid expression")
				}
				op1 := stack[len(stack)-2]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, op1+op2)
			case "-":
				if len(stack) < 2 {
					return 0, errors.New("invalid expression")
				}
				op1 := stack[len(stack)-2]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, op1-op2)
			case "*":
				if len(stack) < 2 {
					return 0, errors.New("invalid expression")
				}
				op1 := stack[len(stack)-2]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, op1*op2)
			case "/":
				if len(stack) < 2 {
					return 0, errors.New("invalid expression")
				}
				op1 := stack[len(stack)-2]
				op2 := stack[len(stack)-1]
				if op2 == 0 {
					return 0, errors.New("cannot be divided by zero")
				}
				stack = stack[:len(stack)-2]
				stack = append(stack, op1/op2)
			default:
				return 0, errors.New("unknown operator")
			}
		default:
			return 0, errors.New("unknown token")
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}
	return stack[0], nil
}

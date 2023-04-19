package calculator

type Transaction struct { // 공식을 묶어 하나의 유의미한 결과값으로 만듦
	Id       string
	Sequence map[string]Formula // 결과값의 명칭 : 공식
	View     []string           // 보이게 어떻게하지?
}

func TransactionToFormula() {

}

package main

import (
	"fmt"
	//"github.com/skorobogatov/input"
	"strconv"
)

type Lexem struct {
	Tag
	Image string
}

type Tag int

type Variable struct{
	Val int
	Name string
}


const (
	ERROR  Tag = 1 << iota // Неправильная лексема
	NUMBER                 // Целое число
	VAR                    // Имя переменной
	PLUS                   // Знак +
	MINUS                  // Знак -
	MUL                    // Знак *
	DIV                    // Знак /
	LPAREN                 // Левая круглая скобка
	RPAREN                 // Правая круглая скобка
)

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() int{
	if s.IsEmpty() {
		return 0
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element
	}
}

var stack Stack

var (
	lexems []Lexem
	lexemL int
	lenexpr int
	lexemI int
	token []string
	variables []Variable
	error bool
)

func getNumber(term string, i int)  int {
	for  i < lenexpr && term[i] > 47 && term[i] < 58 { i++ }
	return i
}

func getVariable(term string, i int) int {
	for i < lenexpr && ((term[i] > 47 && term[i] < 58) || (term[i] > 64 && term[i] < 91) || (term[i] > 96 && term[i] < 123)) { i++ }
	return i
}

func lexer(expr string) {
	var lx Lexem
	for i := 0; i < lenexpr; i++ {
		lexemL++
		switch expr[i] {
		case 10:
			continue
		case 32:
			continue
		case 43:
			lx.Tag = PLUS
			lx.Image = "+"
			lexems = append(lexems, lx)
			break
		case 45:
			lx.Tag = MINUS
			lx.Image = "-"
			lexems = append(lexems, lx)
			break
		case 42:
			lx.Tag = MUL
			lx.Image = "*"
			lexems = append(lexems, lx)
			break
		case 47:
			lx.Tag = DIV
			lx.Image = "/"
			lexems = append(lexems, lx)
			break
		case 40:
			lx.Tag = LPAREN
			lx.Image = "("
			lexems = append(lexems, lx)
			break
		case 41:
			lx.Tag = RPAREN
			lx.Image = ")"
			lexems = append(lexems, lx)
			break
		default:
			if expr[i] > 47 && expr[i] < 58 {
				lx.Tag = NUMBER
				j := getNumber(expr, i)
				lx.Image = expr[i:j]
				lexems = append(lexems, lx)
				i = j - 1
			} else if (expr[i] > 64 && expr[i] < 91) || (expr[i] > 96 && expr[i] < 123) {
				j := getVariable(expr, i)
				lx.Tag = VAR
				lx.Image = expr[i:j]
				lexems = append(lexems, lx)
				i = j - 1
			} else {
				lx.Tag = ERROR
				lexems = append(lexems, lx)
				error = true
			}
			break
		}
	}
}

func E() {
	T()
	Ec()
}

func T() {
	F()
	Tc()
}

func Ec() {
	if lexemL > lexemI {
		lx := lexems[lexemI]
		if lx.Tag & (PLUS|MINUS) != 0 {
			lexemI++
			T()
			token = append(token, lx.Image)
			Ec()
		} else if lx.Tag & (VAR|NUMBER) != 0 { error = true }
	}
}

func Tc() {
	if lexemL > lexemI {
		lx := lexems[lexemI]
		if lx.Tag & (DIV|MUL) != 0 {
			lexemI++
			F()
			token = append(token, lx.Image)
			Tc()
		}
	}
}

func F() {
	if lexemL > lexemI {
		lx := lexems[lexemI]
		if lx.Tag & (NUMBER|VAR) != 0 {
			lexemI++
			token = append(token, lx.Image)
		} else if lx.Tag & MINUS != 0 {
			lexemI++
			token = append(token, "-1")
			F()
			token = append(token, "*")
		} else if lx.Tag & LPAREN != 0 {
			lexemI++
			E()
			if lexemL > lexemI {
				lx = lexems[lexemI]
				lexemI++
				if lx.Tag & RPAREN == 0 { error = true }
			} else { error = true }
		} else { error = true }
	} else { error = true }
}

func getVal(str string) int {
	for _, v := range variables {
		if v.Name == str { return v.Val }
	}
	return 0
}

func Contains(x string) bool {
	for _, n := range variables {
		if x == n.Name { return true }
	}
	return false
}

func eval(){
	for _, t := range token {
		if (t[0] > 47 && t[0] < 58) || (t[0] == '-' && len(t) > 1){
			i, _ := strconv.Atoi(t)
			stack.Push(i)
		} else if (t[0] > 64 &&  t[0] < 91) || (t[0] > 96 && t[0] < 123){
			val := getVal(t)
			stack.Push(val)
		} else {
			switch t {
			case "+":
				stack.Push(stack.Pop() + stack.Pop())
				break
			case "-":
				tmp := stack.Pop()
				stack.Push(stack.Pop() - tmp)
				break
			case "*":
				stack.Push(stack.Pop() * stack.Pop())
				break
			case "/":
				tmp := stack.Pop()
				stack.Push(stack.Pop() / tmp)
			}
		}
	}
}

func main() {
	var (
		expr string
		x int
		v Variable
	)
	//expr = input.Gets()
	_, _ = fmt.Scanf("%s", &expr)
	lenexpr = len(expr)
	lexer(expr)
	E()
	lexemI = len(token) - 1
	if error { fmt.Println("error")
	} else {
		for _, lx := range lexems{
			if lx.Tag & VAR != 0 && !Contains(lx.Image){
				//input.Scanf("%d", &x)
				_, _ = fmt.Scanf("%d", &x)
				v.Name = lx.Image
				v.Val = x
				variables = append(variables, v)
			}
		}
		eval()
		fmt.Println(stack.Pop())
	}
}
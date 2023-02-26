package tasks

import (
	"fmt"
	"os"
	"strings"
)

// В самолете n рядов и по три кресла слева и справа в каждом ряду.
// Крайние кресла (A и F) находятся у окна, центральные (C и D) – у прохода.
// На регистрацию приходят группы из одного, двух или трех пассажиров.
// Они желают сидеть рядом, то есть на одном ряду и на одной стороне: левой или правой.
// Например, группа из двух пассажиров может сесть на кресла B и C,
// но не может сесть на кресла C и D, потому что они разделены проходом,
// а также не может сесть на кресла A и C, потому что тогда они окажутся не рядом.
// Кроме того, один из пассажиров каждой группы очень требовательный – он хочет сесть либо у окна, либо у прохода.
// Конечно же, каждая группа из пассажиров хочет занять места в ряду с как можно меньшим номером,
// Ведь тогда они скорее выйдут из самолета после посадки. Для каждой группы пассажиров определите,
// есть ли места в самолете, подходящие для них.

// Input:
// Первая строка содержит число n(1 ≤ n ≤ 100) – количество рядов в самолете. Далее в n
// строках вводится изначальная рассадка в самолете по рядам (от первого до n-го), где символами .(точка)
// обозначены свободные места, символами # (решетка) обозначены занятые места,
// а символами _ (нижнее подчеркивание) обозначен проход между креслами C и D каждого ряда.
// Следующая строка содержит число m(1 ≤ m ≤ 100) – количество групп пассажиров.
// Далее в m строках содержатся описания групп пассажиров. Формат описания такой:
// num side position , где num – количество пассажиров (число 1, 2 или 3),
// side – желаемая сторона самолета (строка left или right),
// position – желаемое место требовательного пассажира (строка aisle или window).

// Output:
// Если группа может сесть на места, удовлетворяющие ее требованиям,
// то выведите строку Passengers can take seats: и список их мест в формате
// row letter, упорядоченный по возрастанию буквы места. Затем выведите в n
// строках получившуюся рассадку в самолете, в формате, описанном выше, причем места,
// занятые текущей группой пассажиров, должны быть обозначены символом X.
// Если группа не может найти места, удовлетворяющие ее требованиям,
// то выведите строку Cannot fulfill passengers requirements.
// Ответ сравнивается с правильным посимвольно,
// поэтому ваше решение не должно выводить никаких лишних символов,
// в том числе лишних переводов строк или пробельных символов в концах строк.
// В конце каждой строки (включая последнюю) должен быть выведен символ перевода строки.

func PlaneSeats() {
	f, _ := os.Open("./tasks/input.txt")
	defer f.Close()
	oldstdin := os.Stdin
	defer func() {
		os.Stdin = oldstdin
	}()
	os.Stdin = f
	var rows, num, group int
	var pos, side, position string
	fmt.Scan(&rows)
	seats := make([]string, rows)
	for idx := range seats {
		fmt.Scan(&pos)
		seats[idx] = pos
	}
	fmt.Scan(&group)
	for idx := 0; idx < group; idx++ {
		// Attention: on other devices \n in the beggining of format string in scanf might not be needed
		fmt.Scanf("\n%d %s %s", &num, &side, &position)
		if res, ok, idx := managePassangers(num, side, position, seats); !ok {
			fmt.Println("Cannot fulfill passengers requirements")
		} else {
			fmt.Printf("Passengers can take seats: %s\n", res)
			for _, str := range seats {
				fmt.Println(str)
			}
			seats[idx] = strings.ReplaceAll(seats[idx], "X", "#")
		}
	}

}

func managePassangers(num int, side, position string, seats []string) (string, bool, int) {
	for idx, row := range seats {
		switch num {
		case 1:
			if side == "left" {
				if position == "window" && row[0] == '.' {
					seats[idx] = "X" + row[1:]
					return fmt.Sprintf("%d%s", idx+1, "A"), true, idx
				}
				if position == "aisle" && row[2] == '.' {
					seats[idx] = row[:2] + "X" + row[3:]
					return fmt.Sprintf("%d%s", idx+1, "C"), true, idx
				}
			} else {
				if position == "window" && row[6] == '.' {
					seats[idx] = row[:6] + "X"
					return fmt.Sprintf("%d%s", idx+1, "F"), true, idx
				}
				if position == "aisle" && row[4] == '.' {
					seats[idx] = row[:4] + "X" + row[5:]
					return fmt.Sprintf("%d%s", idx+1, "D"), true, idx
				}
			}
		case 2:
			if side == "left" {
				if position == "window" && row[0:2] == ".." {
					seats[idx] = "XX" + row[2:]
					return fmt.Sprintf("%d%s %d%s", idx+1, "A", idx+1, "B"), true, idx
				}
				if position == "aisle" && row[1:3] == ".." {
					seats[idx] = row[:1] + "XX" + row[3:]
					return fmt.Sprintf("%d%s %d%s", idx+1, "B", idx+1, "C"), true, idx
				}
			} else {
				if position == "aisle" && row [4:6] == ".." {
					seats[idx] = row[:4] + "XX" + row[6:]
					return fmt.Sprintf("%d%s %d%s", idx+1, "D", idx+1, "E"), true, idx
				}
				if position == "window" && row[5:] == ".." {
					seats[idx] = row[:5] + "XX"
					return fmt.Sprintf("%d%s %d%s", idx+1, "E", idx+1, "F"), true, idx
				}
			}
		case 3:
			if side == "left" && row[0:3] == "..." {
				seats[idx] = "XXX" + row[3:]
				return fmt.Sprintf("%dA %dB %dC", idx+1, idx+1, idx+1), true, idx
			}
			if side == "right" && row[4:] == "..." {
				seats[idx] = row[0:4] + "XXX"
				return fmt.Sprintf("%dD %dE %dF", idx+1, idx+1, idx+1), true, idx
			}
		}
	}
	return "", false, 0
}

// P.S though code is kinda big for an easy task like this, it works fast
// Amount of lines can be reduced and code can be optimized, but for this specific task 
// i feel like this is not that bad

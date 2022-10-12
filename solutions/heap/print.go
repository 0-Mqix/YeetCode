package heap

import "fmt"

func (h *Heap) Print() {
	lines := []string{"", "", "", ""}
	for i, v := range h.data {
		if i == 0 {
			lines[0] = fmt.Sprintf("          %d", v)
		} else if i < 3 {
			if i%2 != 0 {
				lines[1] += fmt.Sprintf("     %d         ", v)
			} else {
				lines[1] += fmt.Sprintf("%d", v)
			}
		} else if i < 7 {
			if i%2 != 0 {
				lines[2] += fmt.Sprintf("   %d   ", v)
			} else {
				lines[2] += fmt.Sprintf("%d  ", v)
			}
		} else if i < 15 {
			if i%2 != 0 {
				lines[3] += fmt.Sprintf(" %d ", v)
			} else {
				lines[3] += fmt.Sprintf("%d ", v)
			}
		}
	}

	fmt.Println("--------------------------------------")
	for _, v := range lines {
		println(v)
	}
	fmt.Printf("--------------------------------------\n\n")
}

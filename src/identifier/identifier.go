// 标识符
package identifier

import (
	"fmt"
)

func complexDemo() {
	var score complex64 = complex(1, 2)
	var number complex128 = complex(23.23, 11.11)
	fmt.Print("Score = ", score, "\nNumber = ", number, "\n")
	fmt.Print("Real Score = ", real(score), ", Image Score = ", imag(score), "\n")
	fmt.Print("Real Number = ", real(number), ", Image Number = ", imag(number))
}

func Demo() {
	words := `if-else，switch-case-fallthrough-default，for-range-break-continue，var-const-type，import-package，go-defer，goto-return，func-interface，struct-map，chan-select`
	ids := `bool-true-false，byte，int-int8-int16-int32-int64，uint-uint8-uint16-uint32-uint64，float32-float64，complex-imag-real，complex64-complex128，string-cap-len，itoa，nil-upintptr，new-make，panic-recover，print-println，close，copy，append`
	fmt.Println("25个保留关键字：", words)
	fmt.Println("36个预定义标识符：", ids)
	complexDemo()
}

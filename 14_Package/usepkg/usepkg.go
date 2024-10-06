package main

import (
	"14_Package/usepkg/custompkg" // 모듈 내 패키지
	"fmt"                         // 표준 패키지

	// 외부 저장소 패키지
	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo2/ch14/expkg"
	// 다운 받은 외부 패키지들은 GOPATH/pkg/mod 폴더에 버전별로 저장되어 있음
	// 따라서 이미 다운받은 패키지들은 다른 모듈에서 사용되더라도 같은 버전이라면 다시 다운 로드 하지 않고 사용함
)

// go mod tidy
//	- go 모듈에 필요한 패키지를 찾아서 다운로드 해주고,
// 	- 필요한 패키지 정보를 go.mod 파일과 go.sum 파일에 명시

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

package analyze_ipa

import (
	"fmt"
	"testing"
)

func TestCC(t *testing.T) {
	info, _ := Parser("./Link_Live_1.0.ipa")
	fmt.Println(info)
	//fmt.Println("data:image/png;base64,"+base64.RawStdEncoding.EncodeToString(info.IconBody))
}

package font

import (
	_ "embed"
	"fmt"
	"sync"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

//go:embed AlibabaPuHuiTi-3-65-Medium.ttf
var AlibabaPuHuiTi []byte

var (
	once       sync.Once
	collection []font.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		face, err := opentype.Parse(AlibabaPuHuiTi)
		if err != nil {
			panic(fmt.Errorf("failed to parse AlibabaPuHuiTi font: %s", err))
		}

		fc := font.FontFace{
			Font: font.Font{
				Typeface: "阿里巴巴普惠体 3.0",
			},
			Face: face,
		}

		collection = []font.FontFace{fc}
	})
	return collection
}

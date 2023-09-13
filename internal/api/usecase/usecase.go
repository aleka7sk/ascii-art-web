package usecase

import "github.com/aleka7sk/ascii-art-web/pkg"

func ConvertToAscii(text, font string) string {
	return pkg.Run(text, font)
}

package usecase

import "ascii-art-web/pkg"

func ConvertToAscii(text, font string) string {
	return pkg.Run(text, font)
}

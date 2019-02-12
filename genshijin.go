package genshijin

import (
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

// Shaberu ヘンスウ s イレ ル オレ ゲンシジン カエ ス
func Shaberu(s string) string {
	t := tokenizer.New()
	text := strings.TrimSpace(s)
	if text == "" {
		return ""
	}
	tokens := t.Tokenize(text)
	var items []string
	for _, token := range tokens {
		features := token.Features()
		if len(features) == 0 || token.Surface == "BOS" || token.Surface == "EOS" {
			continue
		}
		if features[0] == "助詞" || features[6] == "*" {
			continue
		}
		items = append(items, features[7])
	}
	return strings.Join(items, " ")
}

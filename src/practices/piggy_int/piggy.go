package piggy_int

import (
	"fmt"
	"math/rand"
)

// ======================================================================================================
// LESSON7
// 練習問題
// 新しい貯金箱プログラムを書きましょう。
// 今度は整数型を使い、ドルではなくセントで金額を管理します。
// 空の貯金箱に5セント、10セント、25セントのコインを、ランダムに入れることを貯金が20ドルに達するまで繰り返します。
// コインを投入するたびに貯金箱の残高をドル単位で（例えば$1.05のように）表示してください。
// ======================================================================================================

var savingCosts = [...]int{5, 10, 25}
var bank = 0
var maxSavingCost = 2000

func Answer() {
	savingCostsLength := len(savingCosts)
	for bank <= maxSavingCost {
		savingCost := savingCosts[rand.Intn(savingCostsLength)]
		bank += savingCost

		fmt.Printf("貯金額：$%d.%02d（$0.%02d 貯金しました！）\n", bank/100, bank%100, savingCost)
	}
	fmt.Printf("最終貯金額：$%d.%02d\n", bank/100, bank%100)
}

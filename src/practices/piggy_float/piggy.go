package piggy_float

import (
	"fmt"
	"math/rand"
)

// ======================================================================================================
// LESSON6
// 練習問題
// 友達に贈り物をするため、貯金をします。
// 空の貯金箱が少なくとも20ドル（$20.00）に達するまで5セント（$0.05）か10セント（$0.10）か25セント（$0.25）を
// ランダムに選んで貯金するプログラムを書きましょう。貯金するたびに、貯金箱の残高を表示します。
// その時、適切な幅と精度で整形するようにしてください。
// ======================================================================================================

var savingCosts = [...]float64{0.05, 0.1, 0.25}
var bank = 0.0
var maxSavingCost = 20.0

func Answer() {
	savingCostsLength := len(savingCosts)
	for bank <= maxSavingCost {
		savingCost := savingCosts[rand.Intn(savingCostsLength)]
		bank += savingCost

		fmt.Printf("貯金額：%.2f（%.2fセント貯金しました！）\n", bank, savingCost)
	}

	fmt.Printf("最終貯金額：%.2f\n", bank)
}

package seeds

import (
	"Ewallet-infotecs/internal/fakes"
	"database/sql"
	"github.com/sirupsen/logrus"
)

func WalletSeed(db *sql.DB) {
	checkQuery := `SELECT id, address, balance from wallet`
	rows, _ := db.Query(checkQuery)
	defer rows.Close()
	if !rows.Next() {
		for i := 0; i < 10; i++ {
			fakeAddress := fakes.FakeWallet()
			stmt, _ := db.Prepare(`INSERT INTO wallet(address, balance) VALUES (?,?)`)
			_, err := stmt.Exec(fakeAddress, 100.00)
			if err != nil {
				panic(err)
			}
			logrus.Printf("Добавлен кошелек с адресом: %s", fakeAddress)
		}
		logrus.Print("Добавлено 10 новых кошельков в базу данных")
	}
}

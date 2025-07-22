package main

import (
	"fmt"

	"gorm.io/gorm"
)

// 第二题 指针SQL语句练习
// 题目2：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
// （包含字段 id 主键，from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
// 如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
// 如果余额不足，则回滚事务。

type accounts struct {
	gorm.Model
	// ID      uint `gorm:"primarykey"`
	Name    string
	Balance float64
}

type transactions struct {
	// ID              uint `gorm:"primarykey"`
	gorm.Model
	From_account_id uint
	To_account_id   uint
	Amount          float64
}

// func (account *accounts) BeforeUpdate(tx *gorm.DB) (err error) {

// 	myValue, _ := tx.Get("Name")
// 	fmt.Println("更新前:", myValue)
// 	// fmt.Println("更新前:", account.Balance)

// 	if account.Balance < 0 {
// 		return errors.New(account.Name + "余额不足!")
// 	}
// 	return
// }

func Question_2(db *gorm.DB) {
	//1 创建表
	db.AutoMigrate(&accounts{}, &transactions{})

	//2 插入数据
	// accounts_1 := []accounts{{Name: "A", Balance: 80}, {Name: "B", Balance: 50}}
	// result := db.Create(&accounts_1) // 通过数据的指针来创建插入,否则不会插入
	// fmt.Println(result.RowsAffected)

	//开余额划转事务

	//划转资金
	change_amount := 20.0

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询账户A
	var accountA_1 accounts
	db.Debug().First(&accountA_1, "name = ?", "A")
	fmt.Println(accountA_1)

	var accountB_1 accounts
	db.Debug().First(&accountB_1, "name = ?", "B")
	fmt.Println(accountB_1)

	// 如果账户A余额不足 , 则更新失败, 回滚
	if accountA_1.Balance-change_amount < 0 {
		fmt.Println("账户A余额不足")
		tx.Rollback()
		return
	}

	// 账户A余额充足, 则更新账户A余额
	if err := tx.Model(&accounts{}).Where("id = ?", accountA_1.ID).
		Update("Balance", gorm.Expr("Balance - ?", change_amount)).Error; err != nil {
		fmt.Println("扣款失败:", err)
		tx.Rollback()
		return
	}

	// 如果更新账户B余额失败, 回滚
	if err := tx.Model(&accounts{}).Where("id = ?", accountB_1.ID).
		Update("Balance", accountB_1.Balance+change_amount).Error; err != nil {
		fmt.Println("入账失败:", err)
		tx.Rollback()
		return
	}

	// 插入交易记录
	transaction := transactions{
		From_account_id: accountA_1.ID,
		To_account_id:   accountB_1.ID,
		Amount:          change_amount,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		fmt.Println("创建交易记录失败:", err)
		return
	}

	tx.Commit()

	fmt.Println("入账成功")

}

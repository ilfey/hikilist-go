package main

import (
	"context"
	"fmt"

	"github.com/ilfey/hikilist-go/config"
	"github.com/ilfey/hikilist-go/data/database"
	userModels "github.com/ilfey/hikilist-go/data/models/user"
	"github.com/ilfey/hikilist-go/internal/logger"
)

func main() {
	logger.SetLevel(logger.LevelTrace)

	config.LoadEnvironment()

	config := config.New()

	database.New(config.Database)

	// dm := userModels.DetailModel{}

	// err := dm.Get(context.Background(), "id = 9")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("dm: %+v\n", dm)

	lm := userModels.ListModel{}

	err := lm.Paginate(context.Background(), &userModels.Paginate{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("lm.Results[0]: %v\n", lm.Results[0])

	// var data3 collectionModels.DetailModel

	// sel3 := orm.Select(&data3, "collections")
	// sel3.Resolve("User", func(ctx context.Context, dm *collectionModels.DetailModel) error {
	// 	sel := orm.Select(&userModels.ListItemModel{}, "users")

	// 	sel.Where(fmt.Sprintf("id = %d", dm.UserID))

	// 	item, err := sel.QueryRow(ctx, db)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	dm.User = item

	// 	return nil
	// })

	// dm2, err := sel3.QueryRow(context.Background(), db)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("dm2: %+v\n", dm2)
	// fmt.Printf("*dm2.User: %+v\n", *dm2.User)
}

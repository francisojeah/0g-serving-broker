package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/ethereum/go-ethereum/common"

	"github.com/0glabs/0g-serving-agent/common/config"
	"github.com/0glabs/0g-serving-agent/common/contract"
	database "github.com/0glabs/0g-serving-agent/provider/internal/db"
	"github.com/0glabs/0g-serving-agent/provider/internal/handler"
	"github.com/0glabs/0g-serving-agent/provider/internal/proxy"
)

func Main() {
	config := config.GetConfig()

	db, err := gorm.Open(mysql.Open(config.Database.Provider), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	if err := database.Migrate(db); err != nil {
		panic(err)
	}

	client := contract.MustNewWeb3(config.ChainUrl, config.PrivateKey)
	defer client.Close()
	c, err := contract.NewServingContract(common.HexToAddress(config.ContractAddress), client, config.CustomGasPrice, config.CustomGasLimit)
	if err != nil {
		panic(err)
	}

	r := gin.New()
	p := proxy.New(db, r, c, config.Address, config.PrivateKey)
	if err := p.Start(); err != nil {
		panic(err)
	}

	h := handler.New(db, p, c, config.ServingUrl, config.PrivateKey)
	h.Register(r)

	// Listen and Serve, config port with PORT=X
	if err := r.Run(); err != nil {
		panic(err)
	}
}

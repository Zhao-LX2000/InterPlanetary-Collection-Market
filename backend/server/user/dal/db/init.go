// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"github.com/IPAM/pkg/consts"
	//sql2 "github.com/seata/seata-go/pkg/datasource/sql"
	"time"

	//"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

//var aDB *sql.DB
//
//// AInit init the native DB
//func AInit() {
//
//	dsn := "root:1119@(127.0.0.1:3306)/seata_client1?multiStatements=true&interpolateParams=true"
//	dbAt, err := sql.Open(sql2.SeataATMySQLDriver, dsn)
//	if err != nil {
//		panic("init seata at mysql driver error")
//	}
//
//	aDB = dbAt
//}

// Init init DB
func Init() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	//DB, err = gorm.Open(mysql.New(mysql.Config{
	//	DriverName: sql2.SeataATMySQLDriver,
	//	DSN:        consts.MySQLDefaultDSN,
	//}),
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      gormlogrus,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
}

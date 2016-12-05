package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-ini/ini"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var MasterDB *xorm.Engine

var dns string
var ConfigFile *ini.File

func init() {
	ConfigFile, _ = ini.Load("config/env.ini")
	mysqlConfig := ConfigFile.Section("mysql").KeysHash()
	//if err != nil {
	//		fmt.Println("get mysql config error:", err)
	//		return
	//	}

	fillDns(mysqlConfig)

	// 启动时就打开数据库连接
	err := initEngine()
	if err != nil {
		panic(err)
	}
}

var (
	ConnectDBErr = errors.New("connect db error")
	UseDBErr     = errors.New("use db error")
)

// TestDB 测试数据库
func TestDB() error {

	mysqlConfig := ConfigFile.Section("mysql").KeysHash()

	tmpDns := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=True&loc=Local",
		mysqlConfig["user"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["charset"])
	egnine, err := xorm.NewEngine("mysql", tmpDns)
	if err != nil {
		fmt.Println("new engine error:", err)
		return err
	}
	defer egnine.Close()

	// 测试数据库连接是否 OK
	if err = egnine.Ping(); err != nil {
		fmt.Println("ping db error:", err)
		return ConnectDBErr
	}

	_, err = egnine.Exec("use " + mysqlConfig["dbname"])
	if err != nil {
		fmt.Println("use db error:", err)
		_, err = egnine.Exec("CREATE DATABASE " + mysqlConfig["dbname"] + " DEFAULT CHARACTER SET " + mysqlConfig["charset"])
		if err != nil {
			fmt.Println("create database error:", err)

			return UseDBErr
		}

		fmt.Println("create database successfully!")
	}

	// 初始化 MasterDB
	Init()

	return nil
}

func Init() error {
	mysqlConfig := ConfigFile.Section("mysql").KeysHash()

	fillDns(mysqlConfig)

	// 启动时就打开数据库连接
	err := initEngine()
	if err != nil {
		fmt.Println("mysql is not open:", err)
		return err
	}

	return nil
}

func fillDns(mysqlConfig map[string]string) {
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConfig["user"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["dbname"],
		mysqlConfig["charset"])
}

func initEngine() error {
	var err error

	MasterDB, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		return err
	}

	maxIdle := ConfigFile.Section("mysql").Key("max_idle").MustInt()
	maxConn := ConfigFile.Section("mysql").Key("max_conn").MustInt()

	MasterDB.SetMaxIdleConns(maxIdle)
	MasterDB.SetMaxOpenConns(maxConn)

	showSQL := ConfigFile.Section("xorm").Key("show_sql").MustBool(false)
	logLevel := ConfigFile.Section("xorm").Key("log_level").MustInt()

	MasterDB.ShowSQL(showSQL)
	MasterDB.Logger().SetLevel(core.LogLevel(logLevel))

	// 启用缓存
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// MasterDB.SetDefaultCacher(cacher)

	return nil
}

func StdMasterDB() *sql.DB {
	return MasterDB.DB().DB
}

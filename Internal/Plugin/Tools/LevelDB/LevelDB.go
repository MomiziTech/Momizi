/*
 * @Author: NyanCatda
 * @Date: 2022-03-24 14:04:35
 * @LastEditTime: 2022-03-24 15:09:05
 * @LastEditors: NyanCatda
 * @Description: LevelDB封装
 * @FilePath: \Momizi\Internal\Plugin\Tools\LevelDB\LevelDB.go
 */
package LevelDB

import (
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type DB struct {
	ldb  *leveldb.DB
	Path string
}

/**
 * @description: 打开数据库连接
 * @param {string} Path 数据库路径
 * @return {*DB} 数据库连接
 */
func Open(Path string) *DB {
	ldb, err := leveldb.OpenFile(Path, nil)
	if err != nil {
		Log.Error("LevelDB", err)
		return nil
	}
	return &DB{
		ldb:  ldb,
		Path: Path,
	}
}

/**
 * @description: 写入值
 * @param {string} Key 键
 * @param {string} Value 值
 * @return {error} 错误信息
 */
func (DB DB) Set(Key, Value string) error {
	LDB := DB.ldb
	return LDB.Put([]byte(Key), []byte(Value), nil)
}

/**
 * @description: 读取值
 * @param {string} Key 键
 * @return {string} 值
 * @return {error} 错误信息
 */
func (DB DB) Get(Key string) (string, error) {
	LDB := DB.ldb
	data, err := LDB.Get([]byte(Key), nil)
	return string(data), err
}

/**
 * @description: 删除值
 * @param {string} Key 键
 * @return {error} 错误信息
 */
func (DB DB) Delete(Key string) error {
	LDB := DB.ldb
	return LDB.Delete([]byte(Key), nil)
}

/**
 * @description: 获取键列表
 * @param {Prefix} Prefix 前缀
 * @return {[]string} 键
 * @return {error} 错误信息
 */
func (DB DB) Key(Prefix, Suffix string) ([]string, error) {
	LDB := DB.ldb

	var iter iterator.Iterator
	if Prefix == "" && Suffix == "" {
		iter = LDB.NewIterator(nil, nil)
	}

	if Prefix != "" && Suffix == "" {
		iter = LDB.NewIterator(util.BytesPrefix([]byte(Prefix)), nil)
	}

	if Prefix == "" && Suffix != "" {
		iter = LDB.NewIterator(&util.Range{Start: []byte{}, Limit: []byte(Suffix)}, nil)
	}

	if Prefix != "" && Suffix != "" {
		iter = LDB.NewIterator(&util.Range{Start: []byte(Prefix), Limit: []byte(Suffix)}, nil)
	}

	var keys []string
	for iter.Next() {
		keys = append(keys, string(iter.Key()))
	}
	return keys, nil
}

/**
 * @description: 关闭连接
 * @param {*}
 * @return {*}
 */
func (DB DB) Close() error {
	return DB.ldb.Close()
}

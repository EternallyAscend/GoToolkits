package goRedis

// go-redis

// errorNotFound 中数据查询为空时的返回值。
const errorNotFound = "redis: nil"

// 自定义对外暴露的错误信息

// ErrorClientIsNil 当对应当Client指针为空时抛出的错误。
const ErrorClientIsNil = "Redis Client is nil. "

// ErrorDuplicateKey 当尝试创建某键值对时，如已存在对应键名抛出的错误。
const ErrorDuplicateKey = "Duplicate key. "

// ErrorWIP 调用正在开发的函数将会抛出此错误。
const ErrorWIP = "This function is Working in Progress. "

// StatusDeleted 正常删除的返回值。
const StatusDeleted = 1

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAPPRESPONSECODESCAT = "APP_RESPONSE_CODES_CAT"

// APPRESPONSECODESCAT mapped from table <APP_RESPONSE_CODES_CAT>
type APPRESPONSECODESCAT struct {
	SEQID int64  `gorm:"column:SEQ_ID;type:int;primaryKey" json:"SEQ_ID"`
	NAME  string `gorm:"column:NAME;type:varchar;not null" json:"NAME"`
}

// TableName APPRESPONSECODESCAT's table name
func (*APPRESPONSECODESCAT) TableName() string {
	return TableNameAPPRESPONSECODESCAT
}
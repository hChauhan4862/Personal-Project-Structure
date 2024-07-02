// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameVMSDISPLAY = "VMS_DISPLAYS"

// VMSDISPLAY mapped from table <VMS_DISPLAYS>
type VMSDISPLAY struct {
	DOORDISPLAYID              int64   `gorm:"column:DOOR_DISPLAY_ID;type:int;primaryKey" json:"DOOR_DISPLAY_ID"`
	DOORDISPLAYNO              string  `gorm:"column:DOOR_DISPLAY_NO;type:nvarchar;not null;uniqueIndex:WN_DOOR_DISPLAY_DOOR_DISPLAY_NO_key,priority:1" json:"DOOR_DISPLAY_NO"`
	DOORDISPLAYNAME            string  `gorm:"column:DOOR_DISPLAY_NAME;type:nvarchar;not null" json:"DOOR_DISPLAY_NAME"`
	DOORDISPLAYLOCATION        *string `gorm:"column:DOOR_DISPLAY_LOCATION;type:nvarchar" json:"DOOR_DISPLAY_LOCATION"`
	OFFICEID                   int64   `gorm:"column:OFFICE_ID;type:int;not null" json:"OFFICE_ID"`
	FLOORID                    *int64  `gorm:"column:FLOOR_ID;type:int" json:"FLOOR_ID"`
	ROOMID                     *int64  `gorm:"column:ROOM_ID;type:int" json:"ROOM_ID"`
	PANTRYID                   *int64  `gorm:"column:PANTRY_ID;type:int" json:"PANTRY_ID"`
	MACHINEID                  string  `gorm:"column:MACHINE_ID;type:nvarchar;not null" json:"MACHINE_ID"`
	IPADDRESS                  *string `gorm:"column:IP_ADDRESS;type:nvarchar" json:"IP_ADDRESS"`
	SUBNETMASK                 *string `gorm:"column:SUBNET_MASK;type:nvarchar" json:"SUBNET_MASK"`
	MACADDRESS                 *string `gorm:"column:MAC_ADDRESS;type:nvarchar" json:"MAC_ADDRESS"`
	ISALLOWNOTICE              *bool   `gorm:"column:IS_ALLOW_NOTICE;type:bit;not null;default:1" json:"IS_ALLOW_NOTICE"`
	ISALLOWBOOKING             *bool   `gorm:"column:IS_ALLOW_BOOKING;type:bit;not null;default:1" json:"IS_ALLOW_BOOKING"`
	ISALLOWAMENITIES           *bool   `gorm:"column:IS_ALLOW_AMENITIES;type:bit;not null;default:1" json:"IS_ALLOW_AMENITIES"`
	ISALLOWHOSTINFO            *bool   `gorm:"column:IS_ALLOW_HOST_INFO;type:bit;not null;default:1" json:"IS_ALLOW_HOST_INFO"`
	ISALLOWNEXTMEETING         *bool   `gorm:"column:IS_ALLOW_NEXT_MEETING;type:bit;not null;default:1" json:"IS_ALLOW_NEXT_MEETING"`
	ISALLOWADDINTERNALUSER     *bool   `gorm:"column:IS_ALLOW_ADD_INTERNAL_USER;type:bit;not null;default:1" json:"IS_ALLOW_ADD_INTERNAL_USER"`
	ISALLOWADDEXTERNALUSER     *bool   `gorm:"column:IS_ALLOW_ADD_EXTERNAL_USER;type:bit;not null;default:1" json:"IS_ALLOW_ADD_EXTERNAL_USER"`
	ISALLOWBOOKINGCONFIRMATION *bool   `gorm:"column:IS_ALLOW_BOOKING_CONFIRMATION;type:bit;not null;default:1" json:"IS_ALLOW_BOOKING_CONFIRMATION"`
	LICENCEKEY                 *string `gorm:"column:LICENCE_KEY;type:nvarchar" json:"LICENCE_KEY"`
	LICENCEACTIVATIONDATE      *string `gorm:"column:LICENCE_ACTIVATION_DATE;type:char" json:"LICENCE_ACTIVATION_DATE"`
	LICENCEEXPIRYDATE          *string `gorm:"column:LICENCE_EXPIRY_DATE;type:char" json:"LICENCE_EXPIRY_DATE"`
	ISACTIVE                   *bool   `gorm:"column:IS_ACTIVE;type:bit;not null;default:1" json:"IS_ACTIVE"`
	CREATEDAT                  string  `gorm:"column:CREATED_AT;type:char;not null" json:"CREATED_AT"`
	UPDATEDAT                  *string `gorm:"column:UPDATED_AT;type:char" json:"UPDATED_AT"`
}

// TableName VMSDISPLAY's table name
func (*VMSDISPLAY) TableName() string {
	return TableNameVMSDISPLAY
}
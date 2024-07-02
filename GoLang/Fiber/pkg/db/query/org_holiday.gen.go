// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"SMART_OFFICE_APP/pkg/db/model"
)

func newORGHOLIDAY(db *gorm.DB, opts ...gen.DOOption) oRGHOLIDAY {
	_oRGHOLIDAY := oRGHOLIDAY{}

	_oRGHOLIDAY.oRGHOLIDAYDo.UseDB(db, opts...)
	_oRGHOLIDAY.oRGHOLIDAYDo.UseModel(&model.ORGHOLIDAY{})

	tableName := _oRGHOLIDAY.oRGHOLIDAYDo.TableName()
	_oRGHOLIDAY.ALL = field.NewAsterisk(tableName)
	_oRGHOLIDAY.HOLIDAYID = field.NewInt64(tableName, "HOLIDAY_ID")
	_oRGHOLIDAY.HOLIDAYDATE = field.NewString(tableName, "HOLIDAY_DATE")
	_oRGHOLIDAY.HOLIDAYNAME = field.NewString(tableName, "HOLIDAY_NAME")
	_oRGHOLIDAY.HOLIDAYREASON = field.NewString(tableName, "HOLIDAY_REASON")
	_oRGHOLIDAY.COMPANYHOLIDAY = field.NewBool(tableName, "COMPANY_HOLIDAY")
	_oRGHOLIDAY.OFFICEID = field.NewInt64(tableName, "OFFICE_ID")

	_oRGHOLIDAY.fillFieldMap()

	return _oRGHOLIDAY
}

type oRGHOLIDAY struct {
	oRGHOLIDAYDo

	ALL            field.Asterisk
	HOLIDAYID      field.Int64
	HOLIDAYDATE    field.String
	HOLIDAYNAME    field.String
	HOLIDAYREASON  field.String
	COMPANYHOLIDAY field.Bool
	OFFICEID       field.Int64

	fieldMap map[string]field.Expr
}

func (o oRGHOLIDAY) Table(newTableName string) *oRGHOLIDAY {
	o.oRGHOLIDAYDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o oRGHOLIDAY) As(alias string) *oRGHOLIDAY {
	o.oRGHOLIDAYDo.DO = *(o.oRGHOLIDAYDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *oRGHOLIDAY) updateTableName(table string) *oRGHOLIDAY {
	o.ALL = field.NewAsterisk(table)
	o.HOLIDAYID = field.NewInt64(table, "HOLIDAY_ID")
	o.HOLIDAYDATE = field.NewString(table, "HOLIDAY_DATE")
	o.HOLIDAYNAME = field.NewString(table, "HOLIDAY_NAME")
	o.HOLIDAYREASON = field.NewString(table, "HOLIDAY_REASON")
	o.COMPANYHOLIDAY = field.NewBool(table, "COMPANY_HOLIDAY")
	o.OFFICEID = field.NewInt64(table, "OFFICE_ID")

	o.fillFieldMap()

	return o
}

func (o *oRGHOLIDAY) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *oRGHOLIDAY) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["HOLIDAY_ID"] = o.HOLIDAYID
	o.fieldMap["HOLIDAY_DATE"] = o.HOLIDAYDATE
	o.fieldMap["HOLIDAY_NAME"] = o.HOLIDAYNAME
	o.fieldMap["HOLIDAY_REASON"] = o.HOLIDAYREASON
	o.fieldMap["COMPANY_HOLIDAY"] = o.COMPANYHOLIDAY
	o.fieldMap["OFFICE_ID"] = o.OFFICEID
}

func (o oRGHOLIDAY) clone(db *gorm.DB) oRGHOLIDAY {
	o.oRGHOLIDAYDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o oRGHOLIDAY) replaceDB(db *gorm.DB) oRGHOLIDAY {
	o.oRGHOLIDAYDo.ReplaceDB(db)
	return o
}

type oRGHOLIDAYDo struct{ gen.DO }

func (o oRGHOLIDAYDo) Debug() *oRGHOLIDAYDo {
	return o.withDO(o.DO.Debug())
}

func (o oRGHOLIDAYDo) WithContext(ctx context.Context) *oRGHOLIDAYDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o oRGHOLIDAYDo) ReadDB() *oRGHOLIDAYDo {
	return o.Clauses(dbresolver.Read)
}

func (o oRGHOLIDAYDo) WriteDB() *oRGHOLIDAYDo {
	return o.Clauses(dbresolver.Write)
}

func (o oRGHOLIDAYDo) Session(config *gorm.Session) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Session(config))
}

func (o oRGHOLIDAYDo) Clauses(conds ...clause.Expression) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o oRGHOLIDAYDo) Returning(value interface{}, columns ...string) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o oRGHOLIDAYDo) Not(conds ...gen.Condition) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o oRGHOLIDAYDo) Or(conds ...gen.Condition) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o oRGHOLIDAYDo) Select(conds ...field.Expr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o oRGHOLIDAYDo) Where(conds ...gen.Condition) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o oRGHOLIDAYDo) Order(conds ...field.Expr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o oRGHOLIDAYDo) Distinct(cols ...field.Expr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o oRGHOLIDAYDo) Omit(cols ...field.Expr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o oRGHOLIDAYDo) Join(table schema.Tabler, on ...field.Expr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o oRGHOLIDAYDo) LeftJoin(table schema.Tabler, on ...field.Expr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o oRGHOLIDAYDo) RightJoin(table schema.Tabler, on ...field.Expr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o oRGHOLIDAYDo) Group(cols ...field.Expr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o oRGHOLIDAYDo) Having(conds ...gen.Condition) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o oRGHOLIDAYDo) Limit(limit int) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o oRGHOLIDAYDo) Offset(offset int) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o oRGHOLIDAYDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o oRGHOLIDAYDo) Unscoped() *oRGHOLIDAYDo {
	return o.withDO(o.DO.Unscoped())
}

func (o oRGHOLIDAYDo) Create(values ...*model.ORGHOLIDAY) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o oRGHOLIDAYDo) CreateInBatches(values []*model.ORGHOLIDAY, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o oRGHOLIDAYDo) Save(values ...*model.ORGHOLIDAY) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o oRGHOLIDAYDo) First() (*model.ORGHOLIDAY, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGHOLIDAY), nil
	}
}

func (o oRGHOLIDAYDo) Take() (*model.ORGHOLIDAY, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGHOLIDAY), nil
	}
}

func (o oRGHOLIDAYDo) Last() (*model.ORGHOLIDAY, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGHOLIDAY), nil
	}
}

func (o oRGHOLIDAYDo) Find() ([]*model.ORGHOLIDAY, error) {
	result, err := o.DO.Find()
	return result.([]*model.ORGHOLIDAY), err
}

func (o oRGHOLIDAYDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ORGHOLIDAY, err error) {
	buf := make([]*model.ORGHOLIDAY, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o oRGHOLIDAYDo) FindInBatches(result *[]*model.ORGHOLIDAY, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o oRGHOLIDAYDo) Attrs(attrs ...field.AssignExpr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o oRGHOLIDAYDo) Assign(attrs ...field.AssignExpr) *oRGHOLIDAYDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o oRGHOLIDAYDo) Joins(fields ...field.RelationField) *oRGHOLIDAYDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o oRGHOLIDAYDo) Preload(fields ...field.RelationField) *oRGHOLIDAYDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o oRGHOLIDAYDo) FirstOrInit() (*model.ORGHOLIDAY, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGHOLIDAY), nil
	}
}

func (o oRGHOLIDAYDo) FirstOrCreate() (*model.ORGHOLIDAY, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ORGHOLIDAY), nil
	}
}

func (o oRGHOLIDAYDo) FindByPage(offset int, limit int) (result []*model.ORGHOLIDAY, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o oRGHOLIDAYDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o oRGHOLIDAYDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o oRGHOLIDAYDo) Delete(models ...*model.ORGHOLIDAY) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *oRGHOLIDAYDo) withDO(do gen.Dao) *oRGHOLIDAYDo {
	o.DO = *do.(*gen.DO)
	return o
}
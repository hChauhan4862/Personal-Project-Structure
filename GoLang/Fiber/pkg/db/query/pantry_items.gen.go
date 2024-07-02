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

func newPANTRYITEM(db *gorm.DB, opts ...gen.DOOption) pANTRYITEM {
	_pANTRYITEM := pANTRYITEM{}

	_pANTRYITEM.pANTRYITEMDo.UseDB(db, opts...)
	_pANTRYITEM.pANTRYITEMDo.UseModel(&model.PANTRYITEM{})

	tableName := _pANTRYITEM.pANTRYITEMDo.TableName()
	_pANTRYITEM.ALL = field.NewAsterisk(tableName)
	_pANTRYITEM.SEQ = field.NewInt64(tableName, "SEQ")
	_pANTRYITEM.ITEMCODE = field.NewString(tableName, "ITEM_CODE")
	_pANTRYITEM.ITEMICON = field.NewString(tableName, "ITEM_ICON")
	_pANTRYITEM.ITEMNAME = field.NewString(tableName, "ITEM_NAME")
	_pANTRYITEM.CATSEQ = field.NewInt64(tableName, "CAT_SEQ")

	_pANTRYITEM.fillFieldMap()

	return _pANTRYITEM
}

type pANTRYITEM struct {
	pANTRYITEMDo

	ALL      field.Asterisk
	SEQ      field.Int64
	ITEMCODE field.String
	ITEMICON field.String
	ITEMNAME field.String
	CATSEQ   field.Int64

	fieldMap map[string]field.Expr
}

func (p pANTRYITEM) Table(newTableName string) *pANTRYITEM {
	p.pANTRYITEMDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pANTRYITEM) As(alias string) *pANTRYITEM {
	p.pANTRYITEMDo.DO = *(p.pANTRYITEMDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pANTRYITEM) updateTableName(table string) *pANTRYITEM {
	p.ALL = field.NewAsterisk(table)
	p.SEQ = field.NewInt64(table, "SEQ")
	p.ITEMCODE = field.NewString(table, "ITEM_CODE")
	p.ITEMICON = field.NewString(table, "ITEM_ICON")
	p.ITEMNAME = field.NewString(table, "ITEM_NAME")
	p.CATSEQ = field.NewInt64(table, "CAT_SEQ")

	p.fillFieldMap()

	return p
}

func (p *pANTRYITEM) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pANTRYITEM) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 5)
	p.fieldMap["SEQ"] = p.SEQ
	p.fieldMap["ITEM_CODE"] = p.ITEMCODE
	p.fieldMap["ITEM_ICON"] = p.ITEMICON
	p.fieldMap["ITEM_NAME"] = p.ITEMNAME
	p.fieldMap["CAT_SEQ"] = p.CATSEQ
}

func (p pANTRYITEM) clone(db *gorm.DB) pANTRYITEM {
	p.pANTRYITEMDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pANTRYITEM) replaceDB(db *gorm.DB) pANTRYITEM {
	p.pANTRYITEMDo.ReplaceDB(db)
	return p
}

type pANTRYITEMDo struct{ gen.DO }

func (p pANTRYITEMDo) Debug() *pANTRYITEMDo {
	return p.withDO(p.DO.Debug())
}

func (p pANTRYITEMDo) WithContext(ctx context.Context) *pANTRYITEMDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pANTRYITEMDo) ReadDB() *pANTRYITEMDo {
	return p.Clauses(dbresolver.Read)
}

func (p pANTRYITEMDo) WriteDB() *pANTRYITEMDo {
	return p.Clauses(dbresolver.Write)
}

func (p pANTRYITEMDo) Session(config *gorm.Session) *pANTRYITEMDo {
	return p.withDO(p.DO.Session(config))
}

func (p pANTRYITEMDo) Clauses(conds ...clause.Expression) *pANTRYITEMDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pANTRYITEMDo) Returning(value interface{}, columns ...string) *pANTRYITEMDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pANTRYITEMDo) Not(conds ...gen.Condition) *pANTRYITEMDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pANTRYITEMDo) Or(conds ...gen.Condition) *pANTRYITEMDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pANTRYITEMDo) Select(conds ...field.Expr) *pANTRYITEMDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pANTRYITEMDo) Where(conds ...gen.Condition) *pANTRYITEMDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pANTRYITEMDo) Order(conds ...field.Expr) *pANTRYITEMDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pANTRYITEMDo) Distinct(cols ...field.Expr) *pANTRYITEMDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pANTRYITEMDo) Omit(cols ...field.Expr) *pANTRYITEMDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pANTRYITEMDo) Join(table schema.Tabler, on ...field.Expr) *pANTRYITEMDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pANTRYITEMDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pANTRYITEMDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pANTRYITEMDo) RightJoin(table schema.Tabler, on ...field.Expr) *pANTRYITEMDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pANTRYITEMDo) Group(cols ...field.Expr) *pANTRYITEMDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pANTRYITEMDo) Having(conds ...gen.Condition) *pANTRYITEMDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pANTRYITEMDo) Limit(limit int) *pANTRYITEMDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pANTRYITEMDo) Offset(offset int) *pANTRYITEMDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pANTRYITEMDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pANTRYITEMDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pANTRYITEMDo) Unscoped() *pANTRYITEMDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pANTRYITEMDo) Create(values ...*model.PANTRYITEM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pANTRYITEMDo) CreateInBatches(values []*model.PANTRYITEM, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pANTRYITEMDo) Save(values ...*model.PANTRYITEM) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pANTRYITEMDo) First() (*model.PANTRYITEM, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEM), nil
	}
}

func (p pANTRYITEMDo) Take() (*model.PANTRYITEM, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEM), nil
	}
}

func (p pANTRYITEMDo) Last() (*model.PANTRYITEM, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEM), nil
	}
}

func (p pANTRYITEMDo) Find() ([]*model.PANTRYITEM, error) {
	result, err := p.DO.Find()
	return result.([]*model.PANTRYITEM), err
}

func (p pANTRYITEMDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PANTRYITEM, err error) {
	buf := make([]*model.PANTRYITEM, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pANTRYITEMDo) FindInBatches(result *[]*model.PANTRYITEM, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pANTRYITEMDo) Attrs(attrs ...field.AssignExpr) *pANTRYITEMDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pANTRYITEMDo) Assign(attrs ...field.AssignExpr) *pANTRYITEMDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pANTRYITEMDo) Joins(fields ...field.RelationField) *pANTRYITEMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pANTRYITEMDo) Preload(fields ...field.RelationField) *pANTRYITEMDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pANTRYITEMDo) FirstOrInit() (*model.PANTRYITEM, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEM), nil
	}
}

func (p pANTRYITEMDo) FirstOrCreate() (*model.PANTRYITEM, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PANTRYITEM), nil
	}
}

func (p pANTRYITEMDo) FindByPage(offset int, limit int) (result []*model.PANTRYITEM, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pANTRYITEMDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pANTRYITEMDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pANTRYITEMDo) Delete(models ...*model.PANTRYITEM) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pANTRYITEMDo) withDO(do gen.Dao) *pANTRYITEMDo {
	p.DO = *do.(*gen.DO)
	return p
}
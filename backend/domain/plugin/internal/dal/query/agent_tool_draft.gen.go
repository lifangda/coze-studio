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

	"github.com/coze-dev/coze-studio/backend/domain/plugin/internal/dal/model"
)

func newAgentToolDraft(db *gorm.DB, opts ...gen.DOOption) agentToolDraft {
	_agentToolDraft := agentToolDraft{}

	_agentToolDraft.agentToolDraftDo.UseDB(db, opts...)
	_agentToolDraft.agentToolDraftDo.UseModel(&model.AgentToolDraft{})

	tableName := _agentToolDraft.agentToolDraftDo.TableName()
	_agentToolDraft.ALL = field.NewAsterisk(tableName)
	_agentToolDraft.ID = field.NewInt64(tableName, "id")
	_agentToolDraft.AgentID = field.NewInt64(tableName, "agent_id")
	_agentToolDraft.PluginID = field.NewInt64(tableName, "plugin_id")
	_agentToolDraft.ToolID = field.NewInt64(tableName, "tool_id")
	_agentToolDraft.CreatedAt = field.NewInt64(tableName, "created_at")
	_agentToolDraft.SubURL = field.NewString(tableName, "sub_url")
	_agentToolDraft.Method = field.NewString(tableName, "method")
	_agentToolDraft.ToolName = field.NewString(tableName, "tool_name")
	_agentToolDraft.ToolVersion = field.NewString(tableName, "tool_version")
	_agentToolDraft.Operation = field.NewField(tableName, "operation")

	_agentToolDraft.fillFieldMap()

	return _agentToolDraft
}

// agentToolDraft Draft Agent Tool
type agentToolDraft struct {
	agentToolDraftDo

	ALL         field.Asterisk
	ID          field.Int64  // Primary Key ID
	AgentID     field.Int64  // Agent ID
	PluginID    field.Int64  // Plugin ID
	ToolID      field.Int64  // Tool ID
	CreatedAt   field.Int64  // Create Time in Milliseconds
	SubURL      field.String // Sub URL Path
	Method      field.String // HTTP Request Method
	ToolName    field.String // Tool Name
	ToolVersion field.String // Tool Version, e.g. v1.0.0
	Operation   field.Field  // Tool Openapi Operation Schema

	fieldMap map[string]field.Expr
}

func (a agentToolDraft) Table(newTableName string) *agentToolDraft {
	a.agentToolDraftDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a agentToolDraft) As(alias string) *agentToolDraft {
	a.agentToolDraftDo.DO = *(a.agentToolDraftDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *agentToolDraft) updateTableName(table string) *agentToolDraft {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.AgentID = field.NewInt64(table, "agent_id")
	a.PluginID = field.NewInt64(table, "plugin_id")
	a.ToolID = field.NewInt64(table, "tool_id")
	a.CreatedAt = field.NewInt64(table, "created_at")
	a.SubURL = field.NewString(table, "sub_url")
	a.Method = field.NewString(table, "method")
	a.ToolName = field.NewString(table, "tool_name")
	a.ToolVersion = field.NewString(table, "tool_version")
	a.Operation = field.NewField(table, "operation")

	a.fillFieldMap()

	return a
}

func (a *agentToolDraft) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *agentToolDraft) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["agent_id"] = a.AgentID
	a.fieldMap["plugin_id"] = a.PluginID
	a.fieldMap["tool_id"] = a.ToolID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["sub_url"] = a.SubURL
	a.fieldMap["method"] = a.Method
	a.fieldMap["tool_name"] = a.ToolName
	a.fieldMap["tool_version"] = a.ToolVersion
	a.fieldMap["operation"] = a.Operation
}

func (a agentToolDraft) clone(db *gorm.DB) agentToolDraft {
	a.agentToolDraftDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a agentToolDraft) replaceDB(db *gorm.DB) agentToolDraft {
	a.agentToolDraftDo.ReplaceDB(db)
	return a
}

type agentToolDraftDo struct{ gen.DO }

type IAgentToolDraftDo interface {
	gen.SubQuery
	Debug() IAgentToolDraftDo
	WithContext(ctx context.Context) IAgentToolDraftDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAgentToolDraftDo
	WriteDB() IAgentToolDraftDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAgentToolDraftDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAgentToolDraftDo
	Not(conds ...gen.Condition) IAgentToolDraftDo
	Or(conds ...gen.Condition) IAgentToolDraftDo
	Select(conds ...field.Expr) IAgentToolDraftDo
	Where(conds ...gen.Condition) IAgentToolDraftDo
	Order(conds ...field.Expr) IAgentToolDraftDo
	Distinct(cols ...field.Expr) IAgentToolDraftDo
	Omit(cols ...field.Expr) IAgentToolDraftDo
	Join(table schema.Tabler, on ...field.Expr) IAgentToolDraftDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAgentToolDraftDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAgentToolDraftDo
	Group(cols ...field.Expr) IAgentToolDraftDo
	Having(conds ...gen.Condition) IAgentToolDraftDo
	Limit(limit int) IAgentToolDraftDo
	Offset(offset int) IAgentToolDraftDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAgentToolDraftDo
	Unscoped() IAgentToolDraftDo
	Create(values ...*model.AgentToolDraft) error
	CreateInBatches(values []*model.AgentToolDraft, batchSize int) error
	Save(values ...*model.AgentToolDraft) error
	First() (*model.AgentToolDraft, error)
	Take() (*model.AgentToolDraft, error)
	Last() (*model.AgentToolDraft, error)
	Find() ([]*model.AgentToolDraft, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AgentToolDraft, err error)
	FindInBatches(result *[]*model.AgentToolDraft, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AgentToolDraft) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAgentToolDraftDo
	Assign(attrs ...field.AssignExpr) IAgentToolDraftDo
	Joins(fields ...field.RelationField) IAgentToolDraftDo
	Preload(fields ...field.RelationField) IAgentToolDraftDo
	FirstOrInit() (*model.AgentToolDraft, error)
	FirstOrCreate() (*model.AgentToolDraft, error)
	FindByPage(offset int, limit int) (result []*model.AgentToolDraft, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAgentToolDraftDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a agentToolDraftDo) Debug() IAgentToolDraftDo {
	return a.withDO(a.DO.Debug())
}

func (a agentToolDraftDo) WithContext(ctx context.Context) IAgentToolDraftDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a agentToolDraftDo) ReadDB() IAgentToolDraftDo {
	return a.Clauses(dbresolver.Read)
}

func (a agentToolDraftDo) WriteDB() IAgentToolDraftDo {
	return a.Clauses(dbresolver.Write)
}

func (a agentToolDraftDo) Session(config *gorm.Session) IAgentToolDraftDo {
	return a.withDO(a.DO.Session(config))
}

func (a agentToolDraftDo) Clauses(conds ...clause.Expression) IAgentToolDraftDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a agentToolDraftDo) Returning(value interface{}, columns ...string) IAgentToolDraftDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a agentToolDraftDo) Not(conds ...gen.Condition) IAgentToolDraftDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a agentToolDraftDo) Or(conds ...gen.Condition) IAgentToolDraftDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a agentToolDraftDo) Select(conds ...field.Expr) IAgentToolDraftDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a agentToolDraftDo) Where(conds ...gen.Condition) IAgentToolDraftDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a agentToolDraftDo) Order(conds ...field.Expr) IAgentToolDraftDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a agentToolDraftDo) Distinct(cols ...field.Expr) IAgentToolDraftDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a agentToolDraftDo) Omit(cols ...field.Expr) IAgentToolDraftDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a agentToolDraftDo) Join(table schema.Tabler, on ...field.Expr) IAgentToolDraftDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a agentToolDraftDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAgentToolDraftDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a agentToolDraftDo) RightJoin(table schema.Tabler, on ...field.Expr) IAgentToolDraftDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a agentToolDraftDo) Group(cols ...field.Expr) IAgentToolDraftDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a agentToolDraftDo) Having(conds ...gen.Condition) IAgentToolDraftDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a agentToolDraftDo) Limit(limit int) IAgentToolDraftDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a agentToolDraftDo) Offset(offset int) IAgentToolDraftDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a agentToolDraftDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAgentToolDraftDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a agentToolDraftDo) Unscoped() IAgentToolDraftDo {
	return a.withDO(a.DO.Unscoped())
}

func (a agentToolDraftDo) Create(values ...*model.AgentToolDraft) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a agentToolDraftDo) CreateInBatches(values []*model.AgentToolDraft, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a agentToolDraftDo) Save(values ...*model.AgentToolDraft) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a agentToolDraftDo) First() (*model.AgentToolDraft, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentToolDraft), nil
	}
}

func (a agentToolDraftDo) Take() (*model.AgentToolDraft, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentToolDraft), nil
	}
}

func (a agentToolDraftDo) Last() (*model.AgentToolDraft, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentToolDraft), nil
	}
}

func (a agentToolDraftDo) Find() ([]*model.AgentToolDraft, error) {
	result, err := a.DO.Find()
	return result.([]*model.AgentToolDraft), err
}

func (a agentToolDraftDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AgentToolDraft, err error) {
	buf := make([]*model.AgentToolDraft, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a agentToolDraftDo) FindInBatches(result *[]*model.AgentToolDraft, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a agentToolDraftDo) Attrs(attrs ...field.AssignExpr) IAgentToolDraftDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a agentToolDraftDo) Assign(attrs ...field.AssignExpr) IAgentToolDraftDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a agentToolDraftDo) Joins(fields ...field.RelationField) IAgentToolDraftDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a agentToolDraftDo) Preload(fields ...field.RelationField) IAgentToolDraftDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a agentToolDraftDo) FirstOrInit() (*model.AgentToolDraft, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentToolDraft), nil
	}
}

func (a agentToolDraftDo) FirstOrCreate() (*model.AgentToolDraft, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentToolDraft), nil
	}
}

func (a agentToolDraftDo) FindByPage(offset int, limit int) (result []*model.AgentToolDraft, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a agentToolDraftDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a agentToolDraftDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a agentToolDraftDo) Delete(models ...*model.AgentToolDraft) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *agentToolDraftDo) withDO(do gen.Dao) *agentToolDraftDo {
	a.DO = *do.(*gen.DO)
	return a
}

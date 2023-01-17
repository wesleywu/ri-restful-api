package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/WesleyWu/ri-restful-api/app/model"
	"github.com/WesleyWu/ri-restful-api/app/service/internal/dao"
	"github.com/WesleyWu/ri-restful-api/util/errors"
	"github.com/WesleyWu/ri-restful-api/util/orm"
	"github.com/gogf/gf/v2/frame/g"
)

type IVideoCollectionTwoPks interface {
	Count(ctx context.Context, req *model.VideoCollectionTwoPksCountReq) (*model.VideoCollectionTwoPksCountRes, error)
	One(ctx context.Context, req *model.VideoCollectionTwoPksOneReq) (*model.VideoCollectionTwoPksOneRes, error)
	List(ctx context.Context, req *model.VideoCollectionTwoPksListReq) (*model.VideoCollectionTwoPksListRes, error)
}

type IVideoCollectionTwoPksRepo interface {
	Create(ctx context.Context, req *model.VideoCollectionTwoPksCreateReq) (*model.VideoCollectionTwoPksCreateRes, error)
	Update(ctx context.Context, req *model.VideoCollectionTwoPksUpdateReq) (*model.VideoCollectionTwoPksUpdateRes, error)
	Upsert(ctx context.Context, req *model.VideoCollectionTwoPksUpsertReq) (*model.VideoCollectionTwoPksCreateRes, error)
	Delete(ctx context.Context, req *model.VideoCollectionTwoPksDeleteReq) (*model.VideoCollectionTwoPksDeleteRes, error)
}

type VideoCollectionTwoPksImpl struct {
}

type VideoCollectionTwoPksRepoImpl struct {
}

var (
	VideoCollectionTwoPks     IVideoCollectionTwoPks     = new(VideoCollectionTwoPksImpl)
	VideoCollectionTwoPksRepo IVideoCollectionTwoPksRepo = new(VideoCollectionTwoPksRepoImpl)
)

// Count 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionTwoPksImpl) Count(ctx context.Context, req *model.VideoCollectionTwoPksCountReq) (*model.VideoCollectionTwoPksCountRes, error) {
	var err error
	m := dao.VideoCollectionTwoPks.Ctx(ctx).WithAll()
	m, err = orm.ParseConditions(ctx, req, dao.VideoCollectionTwoPks.ColumnMap, m)
	if err != nil {
		return nil, err
	}
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	return &model.VideoCollectionTwoPksCountRes{Total: count}, err
}

// List 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionTwoPksImpl) List(ctx context.Context, req *model.VideoCollectionTwoPksListReq) (*model.VideoCollectionTwoPksListRes, error) {
	var (
		total int
		page  int
		order string
		list  []*model.VideoCollectionTwoPks
		err   error
	)
	m := dao.VideoCollectionTwoPks.Ctx(ctx).WithAll()
	m, err = orm.ParseConditions(ctx, req, dao.VideoCollectionTwoPks.ColumnMap, m)
	if err != nil {
		return nil, err
	}
	total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		err = errors.WrapServiceErrorf(err, req, "获取数据总记录数失败")
		return nil, err
	}
	if req.Page == 0 {
		req.Page = 1
	}
	page = int(req.Page)
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	if !g.IsEmpty(req.OrderBy) {
		order = req.OrderBy
	}
	list = []*model.VideoCollectionTwoPks{}
	err = m.Fields(model.VideoCollectionTwoPks{}).Page(page, int(req.PageSize)).Order(order).Scan(&list)
	if err != nil {
		g.Log().Error(ctx, err)
		err = errors.WrapServiceErrorf(err, req, "获取数据列表失败")
		return nil, err
	}
	return &model.VideoCollectionTwoPksListRes{
		Total:   uint64(total),
		Current: uint32(page),
		Items:   list,
	}, nil
}

// One 根据req指定的查询条件获取单条数据
// 支持排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionTwoPksImpl) One(ctx context.Context, req *model.VideoCollectionTwoPksOneReq) (*model.VideoCollectionTwoPksOneRes, error) {
	var (
		list  []*model.VideoCollectionTwoPks
		order string
		err   error
	)
	m := dao.VideoCollectionTwoPks.Ctx(ctx).WithAll()
	m, err = orm.ParseConditions(ctx, req, dao.VideoCollectionTwoPks.ColumnMap, m)
	if err != nil {
		return nil, err
	}
	if !g.IsEmpty(req.OrderBy) {
		order = req.OrderBy
	}
	err = m.Fields(model.VideoCollectionTwoPks{}).Order(order).Limit(1).Scan(&list)
	if err != nil {
		g.Log().Error(ctx, err)
		err = errors.WrapServiceErrorf(err, req, "获取单条数据记录失败")
		return nil, err
	}
	if g.IsEmpty(list) || len(list) == 0 {
		return nil, errors.NewNotFoundErrorf(req, "找不到要获取的数据")
	}
	return &model.VideoCollectionTwoPksOneRes{
		VideoCollectionTwoPks: *list[0],
	}, nil
}

// Create 插入记录
// 包括表中所有字段，支持字段类型自动转换，支持对非主键且可为空字段不赋值
// 未赋值或赋值为nil的字段将被更新为 NULL 或数据库表指定的DEFAULT
func (s *VideoCollectionTwoPksRepoImpl) Create(ctx context.Context, req *model.VideoCollectionTwoPksCreateReq) (*model.VideoCollectionTwoPksCreateRes, error) {
	var (
		result       sql.Result
		lastInsertId int64
		rowsAffected int64
		err          error
	)
	result, err = dao.VideoCollectionTwoPks.Ctx(ctx).Insert(req)
	if err != nil {
		if reqErr, ok := errors.DbErrorToRequestError(req, err, dao.VideoCollectionTwoPksDbType); ok {
			return nil, reqErr
		}
		err = errors.WrapServiceErrorf(err, req, "插入失败")
		g.Log().Error(ctx, err)
		return nil, err
	}
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入记录主键键值出错")
		g.Log().Error(ctx, err)
		return nil, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入记录条数出错")
		g.Log().Error(ctx, err)
		return nil, err
	}
	message := "插入成功"
	if rowsAffected == 0 {
		message = "未插入任何记录" // should not happen
	}
	return &model.VideoCollectionTwoPksCreateRes{
		Message:      message,
		LastInsertId: lastInsertId,
		RowsAffected: rowsAffected,
	}, nil
}

// Update 根据主键更新对应记录
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新（即不会修改原记录的字段值）
func (s *VideoCollectionTwoPksRepoImpl) Update(ctx context.Context, req *model.VideoCollectionTwoPksUpdateReq) (*model.VideoCollectionTwoPksUpdateRes, error) {
	var (
		result       sql.Result
		rowsAffected int64
		err          error
	)
	result, err = dao.VideoCollectionTwoPks.Ctx(ctx).
		FieldsEx(dao.VideoCollectionTwoPks.Columns.Id1,
			dao.VideoCollectionTwoPks.Columns.Id2,
			dao.VideoCollectionTwoPks.Columns.CreatedAt).
		WherePri(g.Map{
			dao.VideoCollectionTwoPks.Columns.Id1: req.Id1,
			dao.VideoCollectionTwoPks.Columns.Id2: req.Id2,
		}).
		Update(req)
	if err != nil {
		if reqErr, ok := errors.DbErrorToRequestError(req, err, dao.VideoCollectionTwoPksDbType); ok {
			return nil, reqErr
		}
		err = errors.WrapServiceErrorf(err, req, "更新失败")
		g.Log().Error(ctx, err)
		return nil, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入记录条数出错")
		return nil, err
	}
	message := "更新成功"
	if rowsAffected == 0 {
		return nil, errors.NewNotFoundErrorf(req, "不存在要更新的记录")
	}
	return &model.VideoCollectionTwoPksUpdateRes{
		Message:      message,
		RowsAffected: rowsAffected,
	}, nil
}

// Upsert 根据主键（或唯一索引）是否存在且已在req中赋值，更新或插入对应记录。
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新/插入（即更新时不会修改原记录的字段值）
func (s *VideoCollectionTwoPksRepoImpl) Upsert(ctx context.Context, req *model.VideoCollectionTwoPksUpsertReq) (*model.VideoCollectionTwoPksCreateRes, error) {
	var (
		result       sql.Result
		lastInsertId int64
		rowsAffected int64
		err          error
	)
	result, err = dao.VideoCollectionTwoPks.Ctx(ctx).FieldsEx(dao.VideoCollectionTwoPks.Columns.CreatedAt).Data(req).Save()
	if err != nil {
		if reqErr, ok := errors.DbErrorToRequestError(req, err, dao.VideoCollectionTwoPksDbType); ok {
			return nil, reqErr
		}
		err = errors.WrapServiceErrorf(err, req, "插入/更新失败")
		g.Log().Error(ctx, err)
		return nil, err
	}
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入记录主键键值出错")
		g.Log().Error(ctx, err)
		return nil, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入/更新记录条数出错")
		g.Log().Error(ctx, err)
		return nil, err
	}
	message := "更新成功"
	if rowsAffected == 0 {
		message = "未插入/更新任何记录" // should not happen
	} else if rowsAffected == 1 {
		message = "插入成功"
	}
	return &model.VideoCollectionTwoPksCreateRes{
		Message:      message,
		LastInsertId: lastInsertId,
		RowsAffected: rowsAffected,
	}, nil
}

// Delete 根据req指定的条件删除表中记录
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionTwoPksRepoImpl) Delete(ctx context.Context, req *model.VideoCollectionTwoPksDeleteReq) (*model.VideoCollectionTwoPksDeleteRes, error) {
	var (
		result       sql.Result
		rowsAffected int64
		err          error
	)
	m := dao.VideoCollectionTwoPks.Ctx(ctx).WithAll()
	if req.IsEmpty() {
		return nil, errors.NewBadRequestErrorf(req, "必须指定删除条件")
	}
	m, err = orm.ParseConditions(ctx, req, dao.VideoCollectionTwoPks.ColumnMap, m)
	if err != nil {
		return nil, err
	}
	result, err = m.Delete()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "删除失败")
		g.Log().Error(ctx, err)
		return nil, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取删除记录条数出错")
		return nil, err
	}
	message := fmt.Sprintf("已删除%d条记录", rowsAffected)
	if rowsAffected == 0 {
		message = "不存在要删除的记录"
	}

	return &model.VideoCollectionTwoPksDeleteRes{
		Message:      message,
		RowsAffected: rowsAffected,
	}, nil
}

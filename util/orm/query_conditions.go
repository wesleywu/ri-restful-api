package orm

import (
	"context"
	"github.com/WesleyWu/ri-restful-api/util/errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
	"strings"
)

const (
	ConditionQueryPrefix = "condition{"
	ConditionQuerySuffix = "}"
	TagNameMulti         = "multi"
	TagNameWildcard      = "wildcard"
)

type Condition struct {
	Operator OperatorType `json:"operator"`
	Multi    MultiType    `json:"multi"`
	Wildcard WildcardType `json:"wildcard"`
	Value    interface{}  `json:"value"`
}

// ParseConditions 根据传入 query 结构指针的值来设定 Where 条件
// ctx         context
// queryPtr    传入 query 结构指针的值
// columnMap   表的字段定义map，key为GoField，value为表字段名
// m           gdb.Model
func ParseConditions(ctx context.Context, req interface{}, columnMap map[string]string, m *gdb.Model) (*gdb.Model, error) {
	var err error
	p := reflect.TypeOf(req)
	if p.Kind() != reflect.Ptr { // 要求传入值必须是个指针
		return m, errors.NewBadRequestErrorf(req, "服务函数的输入参数必须是结构体指针")
	}
	t := p.Elem()
	g.Log().Debugf(ctx, "kind of input parameter is %s", t.Name())

	queryValue := reflect.ValueOf(req).Elem()

	// 循环结构体的字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == "Meta" {
			continue
		}
		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct { // 内置结构体类似 XxxQuery
			structValue := queryValue.Field(i)
			g.Log().Debugf(ctx, "kind of field %s is %s", field.Name, field.Type.Kind().String())
			g.Log().Debugf(ctx, "value of field %s is %x", field.Name, structValue)
			for si := 0; si < fieldType.NumField(); si++ {
				innerField := fieldType.Field(si)
				if innerField.Type.Kind() != reflect.Interface { // 仅处理类型为 interface{} 的字段
					continue
				}
				columnName, exists := columnMap[innerField.Name] // 仅处理在表字段定义中有的字段
				if !exists {
					continue
				}
				fieldValue := structValue.Field(si).Interface()
				if fieldValue == nil { // 不出来值为nil的字段
					continue
				}
				g.Log().Debugf(ctx, "inner field %s kind:%si, column:%s, value:%s", innerField.Name, innerField.Type.Kind().String(), columnName, fieldValue)
				m, err = parseField(ctx, req, columnName, innerField.Tag, fieldValue, m)
				if err != nil {
					return nil, err
				}
			}
		} else if fieldType.Kind() == reflect.Interface { // 普通字段，类型为 interfact{}
			columnName, exists := columnMap[field.Name]
			if !exists {
				continue
			}
			fieldValue := queryValue.Field(i).Interface()
			if fieldValue == nil {
				continue
			}
			m, err = parseField(ctx, req, columnName, field.Tag, fieldValue, m)
			if err != nil {
				return nil, err
			}
		}
	}
	return m, nil
}

func parseField(ctx context.Context, req interface{}, columnName string, tag reflect.StructTag, value interface{}, m *gdb.Model) (*gdb.Model, error) {
	if value == nil {
		return m, nil
	}
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Ptr:
		if t.Elem() == reflect.TypeOf(Condition{}) {
			return AddCondition(ctx, columnName, value.(*Condition), m)
		}
		return m.Where(columnName, value), nil
	case reflect.Slice, reflect.Array:
		valueSlice := gconv.SliceAny(value)
		multiTag, ok := tag.Lookup(TagNameMulti)
		if ok {
			multi, err := ParseMultiType(multiTag)
			if err != nil {
				return nil, err
			}
			switch len(valueSlice) {
			case 0:
				return m, nil
			case 1:
				return m.Where(columnName, valueSlice[0]), nil
			case 2:
				if multi == Between {
					return m.WhereBetween(columnName, valueSlice[0], valueSlice[1]), nil
				}
			default:
				return m.WhereIn(columnName, valueSlice), nil
			}
		} else {
			switch len(valueSlice) {
			case 0:
				return m, nil
			case 1:
				return m.Where(columnName, valueSlice[0]), nil
			default:
				return m.WhereIn(columnName, valueSlice), nil
			}
		}
	case reflect.Struct, reflect.Func, reflect.Map, reflect.Chan:
		g.Log().Warningf(ctx, "Query field type struct is not supported")
		return m, nil
	case reflect.String:
		valueString := value.(string)
		if g.IsEmpty(valueString) {
			return m, nil
		}
		if strings.HasPrefix(valueString, ConditionQueryPrefix) && strings.HasSuffix(valueString, ConditionQuerySuffix) {
			var condition *Condition
			err := gjson.DecodeTo(valueString[9:], &condition)
			if err != nil {
				return nil, errors.NewBadRequestErrorf(req, err.Error())
			} else {
				g.Log().Debugf(ctx, "Query field type is orm.Condition: %s", gjson.MustEncodeString(condition))
				return AddCondition(ctx, columnName, condition, m)
			}
		}
		wildcardString, ok := tag.Lookup(TagNameWildcard)
		if ok {
			wildcard, err := ParseWildcardType(wildcardString)
			if err != nil {
				return nil, err
			}
			switch wildcard {
			case Contains:
				return m.WhereLike(columnName, "%"+valueString+"%"), nil
			case StartsWith:
				return m.WhereLike(columnName, valueString+"%"), nil
			case EndsWith:
				return m.WhereLike(columnName, "%"+valueString), nil
			default:
				return m.WhereIn(columnName, valueString), nil
			}
		} else {
			return m.Where(columnName, valueString), nil
		}
	default:
		return m.Where(columnName, value), nil
	}
	return m, nil
}

func AddCondition(_ context.Context, columnName string, condition *Condition, m *gdb.Model) (*gdb.Model, error) {
	if condition == nil {
		return m, nil
	}
	if condition.Value == nil {
		return m, nil
	}
	switch condition.Operator {
	case EQ:
		switch condition.Multi {
		case Exact:
			return m.Where(columnName, condition.Value), nil
		case Between:
			valueSlice := gconv.SliceAny(condition.Value)
			valueLen := len(valueSlice)
			if valueLen == 0 {
				return m, nil
			} else if valueLen == 1 {
				return m.Where(columnName, valueSlice[0]), nil
			} else if valueLen == 2 {
				return m.WhereBetween(columnName, valueSlice[0], valueSlice[1]), nil
			} else {
				return m, errors.NewBadRequestErrorf("column %s requires between query but given %d values", columnName, valueLen)
			}
		case NotBetween:
			valueSlice := gconv.SliceAny(condition.Value)
			valueLen := len(valueSlice)
			if valueLen == 0 {
				return m, nil
			} else if valueLen == 1 {
				return m.WhereNot(columnName, valueSlice[0]), nil
			} else if valueLen == 2 {
				return m.WhereNotBetween(columnName, valueSlice[0], valueSlice[1]), nil
			} else {
				return m, errors.NewBadRequestErrorf("column %s requires between query but given %d values", columnName, valueLen)
			}
		case In:
			valueSlice := gconv.SliceAny(condition.Value)
			valueLen := len(valueSlice)
			if valueLen == 0 {
				return m, nil
			} else if valueLen == 1 {
				return m.Where(columnName, valueSlice[0]), nil
			} else {
				return m.WhereIn(columnName, valueSlice), nil
			}
		case NotIn:
			valueSlice := gconv.SliceAny(condition.Value)
			valueLen := len(valueSlice)
			if valueLen == 0 {
				return m, nil
			} else if valueLen == 1 {
				return m.WhereNot(columnName, valueSlice[0]), nil
			} else {
				return m.WhereNotIn(columnName, valueSlice), nil
			}
		}
	case NE:
		return m.WhereNot(columnName, condition.Value), nil
	case GT:
		return m.WhereGT(columnName, condition.Value), nil
	case GTE:
		return m.WhereGTE(columnName, condition.Value), nil
	case LT:
		return m.WhereLT(columnName, condition.Value), nil
	case LTE:
		return m.WhereLTE(columnName, condition.Value), nil
	case Like:
		valueStr := gconv.String(condition.Value)
		if g.IsEmpty(valueStr) {
			return m, nil
		}
		switch condition.Wildcard {
		case Contains:
			valueStr = "%" + valueStr + "%"
		case StartsWith:
			valueStr = valueStr + "%"
		case EndsWith:
			valueStr = "%" + valueStr
		}
		return m.WhereLike(columnName, valueStr), nil
	case NotLike:
		valueStr := gconv.String(condition.Value)
		if g.IsEmpty(valueStr) {
			return m, nil
		}
		switch condition.Wildcard {
		case Contains:
			valueStr = "%" + valueStr + "%"
		case StartsWith:
			valueStr = valueStr + "%"
		case EndsWith:
			valueStr = "%" + valueStr
		}
		return m.WhereNotLike(columnName, valueStr), nil
	case Null:
		return m.WhereNot(columnName, condition.Value), nil
	case NotNull:
		return m.WhereNot(columnName, condition.Value), nil
	}
	return m, nil
}

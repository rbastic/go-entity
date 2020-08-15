package entity

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Entity struct {
	Type     string
	KV       map[string]interface{}
	Children map[string]Array
	Dirty    bool
}

func (o *Entity) Clone() (*Entity, error) {
	dest := New(o.Type)

	// 'deep' cloning is implemented using JSON
	// marshal/unmarshal ... for now.
	data, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func New(entType string) *Entity {
	return &Entity{
		Type:     entType,
		KV:       makeEmptyMap(),
		Children: makeEmptyChildrenMap(),
		Dirty:    true,
	}
}

func makeEmptyMap() map[string]interface{} {
	return make(map[string]interface{})
}

func (o *Entity) Get(k string) (interface{}, bool) {
	v, ok := o.KV[k]
	return v, ok
}

func (o *Entity) GetInt64(k string) (int64, error) {
	v, ok := o.KV[k]
	if !ok {
		return 0, ErrKeyIsMissing
	}

	switch v.(type) {
	case int:
		fl := v.(int)
		return int64(fl), nil
	case int32:
		fl := v.(int32)
		return int64(fl), nil
	case int64:
		fl := v.(int64)
		return fl, nil
	case uint32:
		fl := v.(uint32)
		return int64(fl), nil
	case uint64:
		fl := v.(uint64)
		return int64(fl), nil
	case float32:
		fl := v.(float32)
		return int64(fl), nil
	case float64:
		fl := v.(float64)
		return int64(fl), nil
	case string:
		fl := v.(string)
		val, err := strconv.ParseInt(fl, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("%s: key:'%s' v:[%v]", err.Error(), k, v)

		}
		return val, nil
	case sql.NullString:
		fl := v.(sql.NullString)
		if !fl.Valid {
			return 0, ErrValueIsNil
		}
		return strconv.ParseInt(fl.String, 10, 64)
	case bool:
		return 0, ErrRefusingCast
	case nil:
		return 0, ErrValueIsNil
	default:
		return 0, fmt.Errorf("GetInt64: unrecognized type %v", reflect.TypeOf(v))
	}
}

func (o *Entity) GetString(k string) (string, error) {
	v, ok := o.KV[k]
	if !ok {
		return "", ErrKeyIsMissing
	}

	switch v.(type) {
	case []uint8:
		fl := v.([]uint8)
		return string(fl), nil
	case float32:
		fl := v.(float32)
		return strconv.FormatFloat(float64(fl), 'f', 6, 64), nil
	case float64:
		fl := v.(float64)
		return strconv.FormatFloat(fl, 'f', 6, 64), nil
	case int:
		fl := v.(int)
		return strconv.Itoa(fl), nil
	case int64:
		fl := v.(int64)
		return strconv.FormatInt(fl, 10), nil
	case uint64:
		fl := v.(uint64)
		return strconv.FormatUint(fl, 10), nil
	case string:
		fl := v.(string)
		return fl, nil
	case *string:
		fl := v.(*string)
		return *fl, nil
	case sql.NullString:
		fl := v.(sql.NullString)
		if !fl.Valid {
			return "", ErrValueIsNil
		}
		return fl.String, nil
	case *sql.NullString:
		fl := v.(*sql.NullString)
		if !fl.Valid {
			return "", ErrValueIsNil
		}
		return fl.String, nil
	case bool:
		fl := v.(bool)
		if fl {
			return "true", nil
		} else {
			return "false", nil
		}
	case nil:
		return "", ErrValueIsNil
	default:
		return "", fmt.Errorf("GetString: unrecognized type %v", reflect.TypeOf(v))
	}
}

func (o *Entity) GetTime(k string) (*time.Time, error) {
	v, ok := o.KV[k]
	if !ok {
		return nil, ErrKeyIsMissing
	}

	switch v.(type) {
	case time.Time:
		fl := v.(time.Time)
		return &fl, nil
	case *time.Time:
		fl := v.(*time.Time)
		return fl, nil
	case nil:
		return nil, ErrValueIsNil
	default:
		return nil, fmt.Errorf("GetTime: unrecognized type %v", reflect.TypeOf(v))
	}
}

func (o *Entity) GetFloat64(k string) (float64, error) {
	v, ok := o.KV[k]
	if !ok {
		return 0, ErrKeyIsMissing
	}

	switch v.(type) {
	case float32:
		fl := v.(float32)
		return float64(fl), nil
	case float64:
		fl := v.(float64)
		return fl, nil
	case int:
		fl := v.(int)
		return float64(fl), nil
	case int64:
		fl := v.(int64)
		return float64(fl), nil
	case uint64:
		fl := v.(uint64)
		return float64(fl), nil
	case string:
		fl := v.(string)
		return strconv.ParseFloat(fl, 64)
	case bool:
		return 0, ErrRefusingCast
	case nil:
		return 0, ErrValueIsNil
	default:
		return 0, fmt.Errorf("GetFloat64: unrecognized type %v", reflect.TypeOf(v))
	}
}

func (o *Entity) GetUint64(k string) (uint64, error) {
	v, ok := o.KV[k]
	if !ok {
		return 0, ErrKeyIsMissing
	}

	switch v.(type) {
	case int:
		fl := v.(int)
		return uint64(fl), nil
	case int32:
		fl := v.(int32)
		return uint64(fl), nil
	case int64:
		fl := v.(int64)
		return uint64(fl), nil
	case uint:
		fl := v.(uint)
		return uint64(fl), nil
	case uint32:
		fl := v.(uint32)
		return uint64(fl), nil
	case uint64:
		fl := v.(uint64)
		return fl, nil
	case float32:
		fl := v.(float32)
		return uint64(fl), nil
	case float64:
		fl := v.(float64)
		return uint64(fl), nil
	case string:
		fl := v.(string)
		return strconv.ParseUint(fl, 10, 64)
	case bool:
		return 0, ErrRefusingCast
	case nil:
		return 0, ErrValueIsNil
	default:
		return 0, fmt.Errorf("GetUint64: unrecognized type %v", reflect.TypeOf(v))
	}
}

func (o *Entity) Set(k string, v interface{}) {
	o.KV[k] = v
}

func (o *Entity) MarkDirty(status bool) {
	o.Dirty = status
}

func (o *Entity) IsDirty() bool {
	return o.Dirty
}

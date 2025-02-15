package utils

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
)

type BuildTreeOptions struct {
	CurKey      string
	ParentKey   string
	ChildrenKey string
}

func CheckFieldExists[T any](fieldName string) bool {
	var instance T
	instanceType := reflect.TypeOf(instance)
	_, found := instanceType.FieldByName(fieldName)
	return found
}
func CheckFieldsExists[T any](fieldNames ...string) bool {
	for _, fieldName := range fieldNames {
		if !CheckFieldExists[T](fieldName) {
			return false
		}
	}
	return true
}

func BuildTree[T any](data []T, options BuildTreeOptions) ([]*T, error) {
	curKey := DefaultIfEmpty(options.CurKey, "id")
	parentKey := DefaultIfEmpty(options.ParentKey, "parentId")
	childrenKey := DefaultIfEmpty(options.ChildrenKey, "children")
	if !CheckFieldsExists[T](curKey, parentKey, childrenKey) {
		return nil, errors.New("curKey, parentKey, childrenKey must be exists!")
	}
	nodeMap := make(map[any]*T)
	for _, item := range data {
		id := reflect.ValueOf(item).FieldByName(curKey).Interface()
		itemCopy := item
		nodeMap[id] = &itemCopy
	}
	var result []*T
	for _, item := range data {
		parentId := reflect.ValueOf(item).FieldByName(parentKey)
		curId := reflect.ValueOf(item).FieldByName(curKey).Interface()
		if parentId.IsNil() {
			result = append(result, nodeMap[curId])
			continue
		}
		var parent *T
		var exists bool
		pidInterface := parentId.Interface()
		switch reflect.TypeOf(pidInterface).String() {
		case "*uint":
			pid := pidInterface.(*uint)
			parent, exists = nodeMap[*pid]
		case "*string":
			pid := pidInterface.(*string)
			parent, exists = nodeMap[*pid]
		case "*int":
			pid := pidInterface.(*int)
			parent, exists = nodeMap[*pid]
		case "uuid.UUID":
			pid := pidInterface.(uuid.UUID)
			parent, exists = nodeMap[pid]
		default:
			return nil, errors.New("parentId must be uint or string")
		}
		if exists {
			children := reflect.ValueOf(parent).Elem().FieldByName(childrenKey)
			children.Set(reflect.Append(children, reflect.ValueOf(item)))
		}
	}
	return result, nil
}

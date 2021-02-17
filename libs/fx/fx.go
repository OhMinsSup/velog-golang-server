package fx

import "reflect"

func Filter(arr interface{}, cond func(interface{}) bool) interface{} {
	contentType := reflect.TypeOf(arr)
	contentValue := reflect.ValueOf(arr)

	newContent := reflect.MakeSlice(contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); cond(content.Interface()) {
			newContent = reflect.Append(newContent, content)
		}
	}
	return newContent.Interface()
}

func Contain(iters []string, selector string) (string, bool) {
	for _, iter := range iters {
		if iter == selector {
			return iter, true
		}
	}
	return "", false
}

func ContainSelector(iters []string, selector string) (string, bool) {
	for _, iter := range iters {
		if iter == selector {
			return selector, true
		}
	}
	return "", false
}

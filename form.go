package ghttp

import (
	"io"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Http表单文件数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type FormFiles []FormFile
type FormFile  struct {
		FieldName string
		FileName  string
		Datas     io.Reader
	}


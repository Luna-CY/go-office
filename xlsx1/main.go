package xlsx1

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/Luna-CY/go-office/xlsx1/sheet"
	"github.com/Luna-CY/go-office/xlsx1/workbook"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

// New 新建一个xlsx文档
func New() (*Document, error) {
	u, err := user.Current()
	if nil != err {
		return nil, errors.New(fmt.Sprintf("获取系统用户信息失败: %v", err))
	}

	document := new(Document)
	document.contentType = NewContentType()
	document.rootRelationship = NewRootRelationship()
	document.docPropsApp = NewDocPropsApp()
	document.docPropsCore = NewDocPropsCore()
	document.docPropsCore.SetUsername(u.Username)

	document.workbookRelationship = workbook.NewBookRelationship()
	document.workbook = workbook.NewWorkbook()

	return document, nil
}

// Document xlsx文档
type Document struct {
	// contentType [Content_Types].xml
	contentType *ContentType

	// rootRelationship /_rels/.rels
	rootRelationship *RootRelationship

	// docPropsApp /docProps/app.xml
	docPropsApp *DocPropsApp

	// docPropsCore /docProps/core.xml
	docPropsCore *DocPropsCore

	// workbookRelationship /xl/_rels/workbook.xml.rels
	workbookRelationship *workbook.BookRelationship

	workbook *workbook.Workbook
}

func (d *Document) SetDocumentUser(user string) *Document {
	d.docPropsCore.SetUsername(user)

	return d
}

func (d *Document) GetWorkbook() *workbook.Workbook {
	return d.workbook
}

// NewSheet 代理 Workbook.NewSheet 方法
func (d *Document) NewSheet() *sheet.Sheet {
	return d.workbook.NewSheet(d.workbookRelationship.NextRId())
}

// NewSheetWithName 代理 Workbook.NewSheetWithName 方法
func (d *Document) NewSheetWithName(name string) *sheet.Sheet {
	return d.workbook.NewSheetWithName(name, d.workbookRelationship.NextRId())
}

func (d *Document) Save(path string) error {
	// 最少保存一个工作表
	if 0 == d.workbook.SheetNum() {
		d.workbook.NewSheet(d.workbookRelationship.NextRId())
	}

	d.docPropsCore.SetTime(time.Now())

	d.contentType.AddDefault(Default{Extension: "rels", ContentType: "application/vnd.openxmlformats-package.relationships+xml"})
	d.contentType.AddDefault(Default{Extension: "xml", ContentType: "application/xml"})
	d.contentType.AddOverride(Override{PartName: fmt.Sprintf("/%v", d.docPropsCore.Filepath()), ContentType: CoreContentType})
	d.contentType.AddOverride(Override{PartName: fmt.Sprintf("/%v", d.docPropsApp.Filepath()), ContentType: AppContentType})
	d.contentType.AddOverride(Override{PartName: fmt.Sprintf("/%v", d.workbook.Filepath()), ContentType: workbook.ContentType})

	d.rootRelationship.AddRelationship(Relationship{Id: d.rootRelationship.NextRId(), Type: AppRelationshipType, Target: d.docPropsApp.Filepath()})
	d.rootRelationship.AddRelationship(Relationship{Id: d.rootRelationship.NextRId(), Type: CoreRelationshipType, Target: d.docPropsCore.Filepath()})
	d.rootRelationship.AddRelationship(Relationship{Id: d.rootRelationship.NextRId(), Type: workbook.RelationshipType, Target: d.workbook.Filepath()})

	for _, s := range d.workbook.Sheets.GetSheets() {
		d.contentType.AddOverride(Override{PartName: fmt.Sprintf("/%v", s.Filepath()), ContentType: sheet.ContentType})

		tokens := strings.Split(s.Filepath(), "/")
		d.workbookRelationship.AddRelationship(workbook.Relationship{Id: s.RelationshipId(), Type: sheet.RelationShipType, Target: filepath.Join(tokens[1:]...)})
	}

	return d.save(path)
}

func (d *Document) save(path string) error {
	file, err := os.Create(path)
	if nil != err {
		return errors.New(fmt.Sprintf("创建文件失败: %v %v", path, err))
	}
	defer file.Close()

	book := zip.NewWriter(file)
	book.RegisterCompressor(zip.Store, nil)
	defer book.Close()

	files := []XmlFile{d.contentType, d.rootRelationship, d.docPropsCore, d.docPropsApp, d.workbookRelationship, d.workbook}
	for _, f := range d.workbook.Sheets.GetSheets() {
		files = append(files, f)
	}

	for _, f := range files {
		writer, err := book.Create(f.Filepath())
		if nil != err {
			return errors.New(fmt.Sprintf("创建压缩文件路径失败: %v", err))
		}

		content, err := f.Marshal()
		if nil != err {
			return err
		}

		if _, err := writer.Write(content); nil != err {
			return errors.New(fmt.Sprintf("写入文件失败: %v", err))
		}
	}

	return nil
}

type XmlEntry interface {
	// Marshal 序列化文件的xml内容
	Marshal() ([]byte, error)
}

// XmlFile xml文件结构接口
type XmlFile interface {
	XmlEntry

	// Filepath 返回文件路径
	Filepath() string
}

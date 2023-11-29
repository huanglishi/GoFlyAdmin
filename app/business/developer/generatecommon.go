package developer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gofly/global"
	"gofly/utils"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gohouse/gorose/v2"
)

/**
* 代码安装工具
 */
/********************************后端*****************************************/
// tablename, tablenamecate, fields string
func MarkeGoCode(file_path, filename, packageName string, parameter map[string]interface{}) {
	//变量参数
	tablename := utils.InterfaceTostring(parameter["tablename"])
	tablenamecate := utils.InterfaceTostring(parameter["cate_tablename"])
	fields := utils.InterfaceTostring(parameter["fields"])
	// 创建go文件
	filePath := filepath.Join(file_path, filename+".go")
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			os.Create(filePath)
		}
	}
	//复制go文件模板到新创建文件
	copyfile := "list"
	if parameter["tpl_type"] != "" {
		copyfile = utils.InterfaceTostring(parameter["tpl_type"])
		if parameter["tpl_type"] == "contentcatelist" {
			filename_cate := filename + "cate"
			filePath_cate := filepath.Join(file_path, filename_cate+".go")
			MarkeBelongCate(filePath_cate, filename_cate, packageName, tablenamecate, fields)
		}
	}
	err := CopyFileContents(filepath.Join("resource/staticfile/codetpl/go/", copyfile+".gos"), filePath)
	if err != nil {
		panic(err)
	}
	//打开新键go文件内容-并替换
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "Replace") {
			datestr := strings.ReplaceAll(string(a), "Replace", utils.FirstUpper(filename))
			result += datestr + "\n"
		} else if strings.Contains(string(a), "packageName") {
			datestr := strings.ReplaceAll(string(a), "packageName", packageName)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{tablename}") {
			datestr := strings.ReplaceAll(string(a), "{tablename}", tablename)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{tablenamecate}") {
			datestr := strings.ReplaceAll(string(a), "{tablenamecate}", tablenamecate)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{fields}") {
			datestr := strings.ReplaceAll(string(a), "{fields}", fields)
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 创建数据关联的分类
func MarkeBelongCate(filePath, filename, packageName, tablename, fields string) {
	err := CopyFileContents(filepath.Join("resource/staticfile/codetpl/go/contentcate.gos"), filePath)
	if err != nil {
		panic(err)
	}
	//打开新键go文件内容-并替换
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "Replace") {
			datestr := strings.ReplaceAll(string(a), "Replace", utils.FirstUpper(filename))
			result += datestr + "\n"
		} else if strings.Contains(string(a), "packageName") {
			datestr := strings.ReplaceAll(string(a), "packageName", packageName)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{tablename}") {
			datestr := strings.ReplaceAll(string(a), "{tablename}", tablename)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{fields}") {
			datestr := strings.ReplaceAll(string(a), "{fields}", fields)
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

/**************************前端处理**********************************/
// 1修改api.ts
//packageName=包名，filename文件名
func ApitsReplay(filePath, packageName, filename string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "modname/filename") {
			datestr := strings.ReplaceAll(string(a), "modname/filename", fmt.Sprintf("%s/%s", packageName, filename))
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 1.1、修改data.ts字段
// file_path文件路径，tablefieldname 字段
func UpFieldData(file_path string, tablefieldname interface{}) {
	f, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "{},") {
			relaystr := ""
			for _, Val := range tablefieldname.([]interface{}) {
				webb, _ := json.Marshal(Val)
				var webjson map[string]interface{}
				_ = json.Unmarshal(webb, &webjson)
				if webjson["value"].(string) == "createtime" || webjson["value"].(string) == "updatetime" {
					relaystr += fmt.Sprintf("     {\n       title:  '%v',\n       dataIndex: '%v',\n       slotName: 'createtime',\n       align:'center'\n     },\n", webjson["label"].(string), webjson["value"].(string))
				} else if webjson["value"].(string) == "id" {
					relaystr += fmt.Sprintf("     {\n       title:  '%v',\n       dataIndex: '%v',\n       width: 76,\n       align:'center'\n     },\n", webjson["label"].(string), webjson["value"].(string))
				} else if webjson["value"].(string) == "image" {
					relaystr += fmt.Sprintf("     {\n       title:  '%v',\n       dataIndex: '%v',\n       slotName: 'image',\n       align:'center'\n     },\n", webjson["label"].(string), webjson["value"].(string))
				} else if webjson["value"].(string) == "status" {
					relaystr += fmt.Sprintf("     {\n       title:  '%v',\n       dataIndex: '%v',\n       slotName: 'status',\n       align:'center'\n     },\n", webjson["label"].(string), webjson["value"].(string))
				} else if webjson["value"].(string) == "cid" {
					relaystr += fmt.Sprintf("     {\n       title:  '%v',\n       dataIndex: 'catename',\n       align:'center'\n     },\n", webjson["label"].(string))
				} else if webjson["value"].(string) == "content" {
				} else {
					relaystr += fmt.Sprintf("     {\n       title:  '%v',\n       dataIndex: '%v',\n       align:'center'\n     },\n", webjson["label"].(string), webjson["value"].(string))
				}
			}
			datestr := strings.ReplaceAll(string(a), "{},", relaystr)
			result += datestr
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 2.1、修改AddForm.vue字段
// file_path文件路径，tablefieldname 字段
func UpFieldAddForm(file_path, fields string, tablefieldname interface{}) {
	f, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	//处理数据
	FieldData := ""    //数据字段初始
	relayhtml := ""    //HTML模板
	replaceFile := ""  //替换附件字段
	replaceimage := "" //替换图片字段
	for _, val := range tablefieldname.([]interface{}) {
		webb, _ := json.Marshal(val)
		var webjson map[string]interface{}
		_ = json.Unmarshal(webb, &webjson)
		value_str := webjson["value"].(string)
		label_str := webjson["label"].(string)
		type_str := webjson["type"].(string)
		if strings.Contains(value_str, "file") {
			replaceFile = value_str
		}
		if strings.Contains(value_str, "image") {
			replaceimage = value_str
		}
		if value_str != "id" {
			defval := "\"\""
			if type_str == "tinyint" || type_str == "int" || type_str == "decimal" || type_str == "double" {
				defval = "0"
			}
			FieldData += fmt.Sprintf("            %v:%v,\n", value_str, defval)
		}
		//处理html模版
		if value_str != "content" && value_str != "id" && value_str != "createtime" && value_str != "updatetime" {
			if type_str == "varchar" && (value_str == "des" || value_str == "remark") {
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"12\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" :rules=\"%v\" >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-textarea v-model=\"formData.%v\" placeholder=\"请填%v\" :auto-size=\"{minRows:3,maxRows:5}\"/>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", value_str, label_str, "[{required:true,message:'请填写"+label_str+"'}]", value_str, label_str)
			} else if type_str == "int" && value_str != "cid" {
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"12\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" :rules=\"%v\" >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-input-number v-model=\"formData.%v\" placeholder=\"请填%v\" />\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", value_str, label_str, "[{required:true,message:'请填写"+label_str+"'}]", value_str, label_str)
			} else if type_str == "tinyint" && value_str == "status" {
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"12\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-radio-group v-model=\"formData.%v\" :options=\"SHoptions\" />\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", value_str, label_str, value_str)
			} else if type_str == "varchar" && value_str == "image" {
				relayhtml += "\t\t\t\t\t\t\t\t\t\t" + `<a-col :span="24">
											<a-form-item field="image" label="` + label_str + `" style="margin-bottom:15px;">
												<div class="upimagebox">
													<div class="imagebtn">
														<div
														class="upload-show-picture"
														v-if="formData.` + value_str + `"
														>
														<a-image
															:src="formData.` + value_str + `"
															height="90"
															:preview-visible="visibleimage"
															@preview-visible-change="() => { visibleimage= false }"
															/>
														<div class="upload-show-picture-mask">
															<a-space><icon-eye @click="()=>visibleimage=true" class="opbtn"/> <IconEdit  @click="UpImage" class="opbtn"/></a-space>
														</div>
														</div>
														<div class="upload-picture-card" v-else @click="UpImage">
														<div class="upload-picture-card-text">
														<IconPlus />
														<div style="margin-top: 10px; font-weight: 600">上传图片</div>
														</div>
														</div>
													  </div>
												</div>
											</a-form-item>
										</a-col>` + "\n"
			} else if type_str == "varchar" && value_str == "file" {
				relayhtml += "\t\t\t\t\t\t\t\t\t\t" + `<a-col :span="12">
										<a-form-item field="file_link" label="` + label_str + `" style="margin-bottom:15px;">
										<a-upload
											accept=".zip,.rar"
											:show-file-list="false"
											:custom-request="customupFile"
											>
											<template #upload-button>
												<div class="upfilezip" >
												<div class="upbtn">
													<a-button type="primary" >
															<template #icon>
															<icon-plus />
															</template>
															<template #default>选择附件</template>
													</a-button>
												</div>
												<div class="showfile">
													<a v-if="formData.` + value_str + `" href="formData.` + value_str + `" download="源码文件"><icon-upload /> 已上传代码包</a>
													<span v-else>未上传</span>
												</div>
												</div>
											</template>
											</a-upload>
										</a-form-item>
									</a-col>` + "\n"
			} else if type_str == "varchar" { //文本输入框
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"12\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" :rules=\"%v\" >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-input v-model=\"formData.%v\" placeholder=\"请填%v\" />\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", value_str, label_str, "[{required:true,message:'请填写"+label_str+"'}]", value_str, label_str)
			}
		}
	}
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}

		fields_arr := strings.Split(fields, `,`)
		if strings.Contains(string(a), "isEditor=ref(false)") && utils.IsContainStrin(fields_arr, "content") {
			datestr := strings.ReplaceAll(string(a), "isEditor=ref(false)", "isEditor=ref(true)")
			result += datestr + "\n"
		} else if strings.Contains(string(a), "replaceField:null") {
			datestr := strings.ReplaceAll(string(a), "replaceField:null", FieldData)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "['replaceFile']") && replaceFile != "" {
			datestr := strings.ReplaceAll(string(a), "['replaceFile']", "."+replaceFile)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "['replaceimage']") && replaceimage != "" {
			datestr := strings.ReplaceAll(string(a), "['replaceimage']", "."+replaceimage)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "<!--replaceTpl-->") {
			datestr := strings.ReplaceAll(string(a), "<!--replaceTpl-->", relayhtml)
			result += datestr
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

/********************************************卸载前后端********************************/
// 卸载/删除文件
func UnInstallCodeFile(data gorose.Data) {
	//1.删除后端代码
	//go文件目录
	file_path_go_root := filepath.Join("app/", utils.InterfaceTostring(data["api_path"]))
	//go文件
	filego_path := filepath.Join(file_path_go_root, utils.InterfaceTostring(data["api_filename"]))
	if _, err := os.Stat(filego_path); err == nil {
		//1.文件存在删除文件
		os.Remove(filego_path)
		if utils.InterfaceTostring(data["tpl_type"]) == "contentcatelist" {
			filename_arr := strings.Split(utils.InterfaceTostring(data["api_filename"]), `.`)
			filecatego_path := filepath.Join(file_path_go_root, filename_arr[0]+"cate.go")
			os.Remove(filecatego_path)
		}
		//2.删除文件夹
		dir, _ := os.ReadDir(file_path_go_root)
		if len(dir) == 0 {
			os.RemoveAll(file_path_go_root)
			//3.移除路由
			packgename_arr := strings.Split(utils.InterfaceTostring(data["api_path"]), `/`)
			modelname := "business" //模块名称
			if len(packgename_arr) > 0 {
				modelname = packgename_arr[0]
			}
			CheckApiRemoveController(modelname, utils.InterfaceTostring(data["api_path"]))
		}
	}
	//2.2 删除views下代码
	vue_component := utils.InterfaceTostring(data["component"])
	component_arr := strings.Split(vue_component, `/`)
	if data["component"] != nil {
		componentpah_arr := strings.Split(data["component"].(string), (component_arr[len(component_arr)-1]))
		vue_path := filepath.Join(global.App.Config.App.Vueobjroot, "/src/views/", componentpah_arr[0]) //前端文件路径
		if _, err := os.Stat(vue_path); err == nil {
			os.RemoveAll(vue_path)
			//2.3.模块目录文件夹
			vue_model_path := filepath.Join(global.App.Config.App.Vueobjroot, "/src/views/", component_arr[0])
			dirs, _ := os.ReadDir(vue_model_path)
			if len(dirs) == 0 {
				os.RemoveAll(vue_model_path)
			}
		}
	}

}

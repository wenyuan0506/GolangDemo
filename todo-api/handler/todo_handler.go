package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"todo-api/service"
	"todo-api/util"
)

// 統一回傳格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := service.GetAllTodos()
	util.JSON(w, http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    todos,
	})
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.JSON(w, http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "ID 必須是數字",
		})
		return
	}

	todo, found := service.GetTodoByID(id)
	if !found {
		util.JSON(w, http.StatusNotFound, Response{
			Code:    http.StatusNotFound,
			Message: "找不到該筆待辦",
		})
		return
	}

	util.JSON(w, http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    todo,
	})
}

func MssqlTest(w http.ResponseWriter, r *http.Request) {
	msg := service.MssqlTest()
	util.JSON(w, http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    map[string]string{"message": msg},
	})
}

func GetConnString(w http.ResponseWriter, r *http.Request) {
	conn := service.ConnString()
	util.JSON(w, http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    map[string]string{"connString": conn},
	})
}

func GetAllTableNames(w http.ResponseWriter, r *http.Request) {
	names, err := service.GetAllTableNames()
	if err != nil {
		util.JSON(w, http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "無法獲取所有表名",
		})
		return
	}
	util.JSON(w, http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    names,
	})
}

// 獲取表格數據 (POST /table)
func GetTableData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.JSON(w, http.StatusMethodNotAllowed, Response{
			Code:    http.StatusMethodNotAllowed,
			Message: "只支援 POST 方法",
		})
		return
	}

	// 解析 body
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.JSON(w, http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "解析 JSON 失敗: " + err.Error(),
		})
		return
	}
	defer r.Body.Close()

	name := strings.TrimSpace(req.Name)
	if name == "" {
		util.JSON(w, http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "請提供表格名稱",
		})
		return
	}

	data, err := service.GetTableData(name)
	if err != nil {
		util.JSON(w, http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "無法獲取表格數據: " + err.Error(),
		})
		return
	}

	util.JSON(w, http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    data,
	})
}

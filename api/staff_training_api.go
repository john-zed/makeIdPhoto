package api

import (
	. "../../staff_remote/model"
	. "./liantuofu"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddStaffTrainingApi(c *gin.Context) {

	staffName := c.Request.FormValue("staff_name")
	gender := c.Request.FormValue("gender")
	age, _ := strconv.Atoi(c.Request.FormValue("age"))
	phone := c.Request.FormValue("phone")
	address := c.Request.FormValue("address")
	idCardNo := c.Request.FormValue("id_card_no")
	idCardFront := c.Request.FormValue("id_card_front")
	idCardBack := c.Request.FormValue("id_card_back")
	avatar := c.Request.FormValue("avatar")
	isTrained, _ := strconv.ParseBool(c.Request.FormValue("is_trained"))
	workingState := c.Request.FormValue("working_state")

	staffTraining := StaffTraining{StaffName: staffName, Gender: gender, Age: age, Phone: phone, Address: address, IdCardNo: idCardNo, IdCardFront: idCardFront, IdCardBack: idCardBack, Avatar: avatar, IsTrained: isTrained, WorkingState: workingState}

	_, err := staffTraining.AddStaffTraining()

	if err != nil {
		renderError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    CodeOk,
		"message": "",
		"data":    staffTraining,
	})
}

func QueryStaffTrainingsApi(c *gin.Context) {
	staffTrainings, _ := GetStaffTrainings()

	c.JSON(http.StatusOK, gin.H{
		"code":           CodeOk,
		"message":        "",
		"staffTrainings": staffTrainings,
	})
}

func UploadFile(c *gin.Context) {

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    CodeError,
			"message": fmt.Sprintf("get form error: %s", err.Error()),
		})
		return
	}

	filename := xid.New().String() + path.Ext(filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, "./upload/"+filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    CodeError,
			"message": fmt.Sprintf("upload file err: %s", err.Error()),
		})
		return
	}

	imagePath := Http + Host + "/upload/" + filename

	c.JSON(http.StatusOK, gin.H{
		"code":     CodeOk,
		"message":  "",
		"filepath": imagePath,
	})
}

func MakePhoto(c *gin.Context) {

	// 保存原图
	file, err := c.FormFile("file")
	if err != nil {
		renderError(c, err)
		return
	}

	filename := xid.New().String() + path.Ext(filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, "./upload/"+filename); err != nil {
		renderError(c, err)
		return
	}

	// 请求处理证件照API
	originalImage, err := file.Open()
	if err != nil {
		renderError(c, err)
		return
	}
	originalImageBytes, _ := ioutil.ReadAll(originalImage)
	photoBase64 := base64.StdEncoding.EncodeToString(originalImageBytes)
	spec := c.Request.FormValue("spec")
	res, err := MakeIdPhotoFree(photoBase64, spec)

	if err != nil {
		renderError(c, err)
		return
	}

	// 处理res
	// 去掉空格回车
	res = strings.ReplaceAll(res, " ", "")
	res = strings.ReplaceAll(res, "\n", "")

	// 取出图片地址信息
	imagePath := gjson.Get(res, "result.file_name_wm.0").String()

	// 保存结果图
	resImage, err := http.Get(imagePath)
	if err != nil {
		renderError(c, err)
		return
	}
	defer resImage.Body.Close()
	resBody, err := ioutil.ReadAll(resImage.Body)
	if err != nil {
		renderError(c, err)
		return
	}
	// 随机文件名 文件名后缀与原图相同
	resImagePath := "./upload/"
	// 随机文件名 文件名后缀与原图相同
	resFileName := xid.New().String() + path.Ext(filepath.Base(file.Filename))
	resImagePath += resFileName
	resImageFile, err := os.Create(resImagePath)
	if err != nil {
		renderError(c, err)
		return
	}
	// 写入
	defer resImageFile.Close()
	resImageFile.Write(resBody)
	// 图片地址带有&需要使用PureJSON
	c.PureJSON(http.StatusOK, gin.H{
		"code":    CodeOk,
		"message": "",
		"data":    Http + Host + "/upload/" + resFileName,
	})
}

func LiantuofuPrecreate(c *gin.Context) {

	totalAmount, _ := strconv.ParseFloat(c.Request.FormValue("total_amount"), 64)

	if totalAmount != TotalAmount {
		c.JSON(http.StatusOK, gin.H{
			"code":    CodeError,
			"message": "订单金额不正确",
		})
		return
	}

	outTradeNo := MakeTradeNo()

	reqParas := GetPrecreateReqParas(totalAmount, outTradeNo)

	resp := Post(UrlQIPrecreate, reqParas, 20)

	payUrl := gjson.Get(resp, "qrCode").String()

	var resData struct {
		PayUrl  string `json:"payUrl""`
		TradeNo string `json:"tradeNo""`
	}
	resData.PayUrl = payUrl
	resData.TradeNo = outTradeNo

	c.JSON(http.StatusOK, gin.H{
		"code":    CodeOk,
		"message": "",
		"data":    resData,
	})
}

func LiantuofuOpenPay(c *gin.Context) {

	/*totalAmount, _ := strconv.ParseFloat(c.Request.FormValue("total_amount"), 64)

	if totalAmount != TotalAmount {
		c.JSON(http.StatusOK, gin.H{
			"code":    CodeError,
			"message": "订单金额不正确",
		})
		return
	}*/

	openId := c.Request.FormValue("open_id")

	if openId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    CodeError,
			"message": "open_id 不能为空",
		})
		return
	}

	outTradeNo := MakeTradeNo()

	reqParas := GetOpenPayReqParas(TotalAmount, outTradeNo, "WXPAY", openId)

	resp := Post(UrlQIPrecreate, reqParas, 20)

	resCode := gjson.Get(resp, "code").String()
	msg := gjson.Get(resp, "msg").String()

	if resCode == "SUCCESS" {
		c.JSON(http.StatusOK, gin.H{
			"code":       CodeOk,
			"message":    msg,
			"outTradeNo": outTradeNo,
			"appId":      gjson.Get(resp, "appId").String(),
			"timeStamp":  gjson.Get(resp, "timeStamp").String(),
			"nonceStr":   gjson.Get(resp, "nonceStr").String(),
			"signType":   gjson.Get(resp, "signType").String(),
			"paySign":    gjson.Get(resp, "paySign").String(),
			"payPackage": gjson.Get(resp, "payPackage").String(),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    CodeError,
			"message": msg,
		})
	}

}

func LiantuofuQueryOrder(c *gin.Context) {
	outTradeNo := c.Request.FormValue("trade_no")
	reqParas := GetOrderQueryReqParas(outTradeNo)
	resp := Post(UrlQIPayQuery, reqParas, 20)
	resCode := gjson.Get(resp, "code").String()
	msg := gjson.Get(resp, "msg").String()
	totalAmount := gjson.Get(resp, "totalAmount").Float()

	if resCode == "SUCCESS" && totalAmount == totalAmount {
		c.JSON(http.StatusOK, gin.H{
			"code":    CodeOk,
			"message": msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    CodeError,
			"message": msg,
		})
	}
}

func renderError(c *gin.Context, err error) {
	log.Println(err.Error())
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    CodeError,
		"message": err.Error(),
	})
}

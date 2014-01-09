package controllers

import (
	"big_brother/models"
	"github.com/astaxie/beego"
)

type ManageController struct {
	beego.Controller
}

func (this *ManageController) GetManagePage() {
	var machineList []*models.Register
	o.Using("default")
	_, err := o.QueryTable("register").Limit(-1).All(&machineList, "id", "ip", "host_name", "hardware_addr", "agent_version", "machine_role", "status")
	if err != nil {
		this.Data["machine_list"] = nil
	} else {
		this.Data["machine_list"] = machineList
	}
	this.Data["role_mapper"] = map[string]string{
		"ercifenjian": "二次分拣",
		"kaipiao":     "财务开票",
		"dabao":       "打包",
		"fenbo":       "分拨",
	}
	this.Data["status_mapper"] = map[int]string{
		0:  "已正常关机",
		1:  "正常运行中",
		-1: "运行异常",
		-2: "不再使用",
	}

	statusLabelMapper := map[int]string{
		0:  "label",
		1:  "label label-success",
		-1: "label label-important",
		-2: "label label-inverse",
	}
	machineStatusLabels := make(map[string]string)
	for _, machine := range machineList {
		machineStatusLabels[machine.Hardware_addr] = statusLabelMapper[machine.Status]
	}
	this.Data["status_labels"] = machineStatusLabels
	this.TplNames = "manage_machine.html"
}

func (this *ManageController) DelMachine() {
	id := this.GetString("id")
	if ip == "" || mac == "" {
		this.Data["json"] = map[string]string{
			"Status": "failure",
			"Msg": "参数不全，未能删除机器！",
		}
	}else{
		o.Using("default")
		num, err := o.Delete(&models.Register{Id: id})
	}
	this.ServeJson()
}

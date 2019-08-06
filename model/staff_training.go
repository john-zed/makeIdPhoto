package model

import (
	. "../../staff_remote/database"
	qb "github.com/didi/gendry/builder"
	"time"
)

type StaffTraining struct {
	Id           int    `json:"id"`
	StaffName    string `json:"staffName"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	IdCardNo     string `json:"idCardNo"`
	IdCardFront  string `json:"idCardFront"`
	IdCardBack   string `json:"idCardBack"`
	Avatar       string `json:"avatar"`
	IsTrained    bool   `json:"isTrained"`
	WorkingState string `json:"workingState"`
	Status       string `json:"status"`
	CreateTime   string `json:"createTime"`
}

func (st *StaffTraining) AddStaffTraining() (id int64, err error) {
	t := time.Now().Format("2006-01-02 15:04:05")
	rs, err := SqlDB.Exec("INSERT INTO staff_training(staff_name, gender, age, phone, address, id_card_no,  id_card_front, id_card_back, avatar, is_trained, working_state, status, create_time ) VALUES (?, ?, ?, ?, ?, ? ,? ,? ,? ,? ,? ,? ,? )", st.StaffName, st.Gender, st.Age, st.Phone, st.Address, st.IdCardNo, st.IdCardFront, st.IdCardBack, st.Avatar, st.IsTrained, st.WorkingState, Applied, t)
	if err != nil {
		return 0, err
	}

	id, err = rs.LastInsertId()
	return
}

func GetStaffTrainings() (staffTrainings []StaffTraining, err error) {

	staffTrainings = make([]StaffTraining, 0)

	where := map[string]interface{}{
		"_orderby": "create_time desc",
	}
	cond, vals, err := qb.BuildSelect("staff_training", where, []string{"id", "staff_name", "gender", "age", "phone", "address", "id_card_no", "id_card_front", "id_card_back", "avatar", "is_trained", "working_state", "status", "create_time"})
	if nil != err {
		panic(err)
	}
	rows, err := SqlDB.Query(cond, vals...)
	if nil != err {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var staffTraining StaffTraining
		rows.Scan(&staffTraining.Id, &staffTraining.StaffName, &staffTraining.Gender, &staffTraining.Age, &staffTraining.Phone, &staffTraining.Address, &staffTraining.IdCardNo, &staffTraining.IdCardFront, &staffTraining.IdCardBack, &staffTraining.Avatar, &staffTraining.IsTrained, &staffTraining.WorkingState, &staffTraining.Status, &staffTraining.CreateTime)
		staffTrainings = append(staffTrainings, staffTraining)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
	return
}

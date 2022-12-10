package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func InsertAbnormalState(state entity.AbnormalState) int64 {
	var sql = "insert into abnormal_state(name,introduce,icon) values (?,?,?)"
	var icon = ""
	if state.Icon.Valid {
		icon = state.Icon.String
	}
	result, _ := utils.Database.Exec(sql, state.Name, state.Introduce, icon)
	id, _ := result.LastInsertId()
	return id
}

func GetAllAbnormalState() []entity.AbnormalStateResponse {
	var sql = "select id,name,introduce,icon from abnormal_state"
	row, _ := utils.Database.Query(sql)
	defer row.Close()
	var list = make([]entity.AbnormalStateResponse, 0)
	for row.Next() {
		var state entity.AbnormalState
		row.Scan(
			&state.Id,
			&state.Name,
			&state.Introduce,
			&state.Icon,
		)
		var response = entity.AbnormalStateResponse{
			Id:        state.Id,
			Name:      state.Name,
			Introduce: state.Introduce,
		}
		response.Icon = ""
		if state.Icon.Valid {
			response.Icon = state.Icon.String
		}
		list = append(list, response)
	}
	return list
}

func UpdateAbnormalState(state entity.AbnormalStateResponse) int64 {
	var sql = "update abnormal_state set name=?,introduce=?,icon=? where id=?"
	row, _ := utils.Database.Exec(
		sql,
		state.Name,
		state.Introduce,
		state.Icon,
		state.Id,
	)
	stateId, _ := row.RowsAffected()
	return stateId
}

package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func GetSkillByName(page int, amount int, name string) ([]entity.Skill, int) {
	var skills = make([]entity.Skill, 0)
	var sql = "select s.id, s.name, s.description, s.value, s.amount, s.speed, s.is_genetic, s.additional_effects, s.is_be, st.id, st.name, g.id, g.name from skill s left join skill_type st on s.skill_type_id = st.id left join genius_attributes g on s.attributes_id = g.id where s.name like ? order by s.id desc limit ? offset ?"
	row, err := utils.Database.Query(sql, "%"+name+"%", amount,
		(page-1)*amount)
	if err != nil {
		panic("failed")
	}
	defer row.Close()
	for row.Next() {
		var skill entity.Skill
		row.Scan(
			&skill.Id,
			&skill.Name,
			&skill.Description,
			&skill.Value,
			&skill.Amount,
			&skill.Speed,
			&skill.IsGenetic,
			&skill.AdditionalEffects,
			&skill.IsBe,
			&skill.SkillType.Id,
			&skill.SkillType.Name,
			&skill.Attributes.Id,
			&skill.Attributes.Name,
		)
		skills = append(skills, skill)
	}
	row, err = utils.Database.Query("select COUNT(*) from skill s left join skill_type st on s.skill_type_id = st.id left join genius_attributes g on s.attributes_id = g.id where s.name like ?",
		"%"+name+"%")
	if err != nil {
		panic("failed")
	}
	defer row.Close()
	var total int
	for row.Next() {
		row.Scan(&total)
	}
	return skills, total
}

func InsertSkill(skill *entity.Skill) int64 {
	sql := "insert into skill(name,skill_type_id,attributes_id,value,amount,speed,is_genetic,is_be,additional_effects,description) value(?,?,?,?,?,?,?,?,?,?)"
	result, _ := utils.Database.Exec(sql,
		skill.Name,
		skill.SkillType.Id,
		skill.Attributes.Id,
		skill.Value,
		skill.Amount,
		skill.Speed,
		skill.IsGenetic,
		skill.IsBe,
		skill.AdditionalEffects,
		skill.Description)
	id, _ := result.LastInsertId()
	return id
}

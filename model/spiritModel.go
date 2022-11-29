package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
)

func GetSpiritList(page int, amount int, keywords string) ([]entity.SpiritListItem, int) {
	var sql = "select g.id,g.number,g.avatar,g.name,att.id,att.name ,att2.id,att2.name from `genius` g  left join `genius_attributes` att on g.primary_attributes_id = att.id left join `genius_attributes` att2 on g.secondary_attributes_id = att2.id where g.name like ?  order by g.id desc limit ? offset ?"
	var sqlAmount = "select COUNT(*) from `genius` g  left join `genius_attributes` att on g.primary_attributes_id = att.id left join `genius_attributes` att2 on g.secondary_attributes_id = att2.id where g.name like ?"
	row, err := utils.Database.Query(sql, "%"+keywords+"%", amount, (page-1)*amount)
	if err != nil {
		return *new([]entity.SpiritListItem), 0
	}
	defer row.Close()
	var list = make([]entity.SpiritListItem, 0)
	for row.Next() {
		var spirit entity.SpiritListItem
		row.Scan(&spirit.ID,
			&spirit.Number,
			&spirit.Avatar,
			&spirit.Name,
			&spirit.PrimaryAttributes.Id,
			&spirit.PrimaryAttributes.Name,
			&spirit.SecondaryAttributes.Id,
			&spirit.SecondaryAttributes.Name)
		list = append(list, spirit)
	}
	row, err = utils.Database.Query(sqlAmount, "%"+keywords+"%")
	if err != nil {
		return *new([]entity.SpiritListItem), 0
	}
	defer row.Close()
	var total int
	for row.Next() {
		row.Scan(&total)
	}
	return list, total
}

func InsertSpirit(spirit *entity.Spirit) int64 {
	var sql = "insert into genius(avatar,number,name,description,primary_attributes_id,secondary_attributes_id,race_power,race_attack,race_defense,race_magic_attack,race_magic_defense,race_speed,group_id,height,weight,hobby) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	row, _ := utils.Database.Exec(
		sql,
		spirit.Avatar,
		spirit.Number,
		spirit.Name,
		spirit.Description,
		spirit.PrimaryAttributes.Id,
		spirit.SecondaryAttributes.Id,
		spirit.RacePower,
		spirit.RaceAttack,
		spirit.RaceDefense,
		spirit.RaceMagicAttack,
		spirit.RaceMagicDefense,
		spirit.RaceSpeed,
		spirit.Group.Id,
		spirit.Height,
		spirit.Weight,
		spirit.Hobby,
	)
	spiritSqlId, _ := row.LastInsertId()
	for _, v := range spirit.Skills {
		row, _ = utils.Database.Exec("insert into genius_skill(genius_id,skill_id) values (?,?)",
			spiritSqlId,
			v.Id)
		_, _ = row.LastInsertId() // 操作影响的行数
	}
	return spiritSqlId
}

func GetSpiritDetailsById(id int) (*entity.Spirit, error) {
	row, err := utils.Database.Query("select gen.avatar,gen.number,gen.name,gen.description,gen.race_power,gen.race_attack,gen.race_defense,gen.race_magic_attack,gen.race_magic_defense,gen.race_speed,gen.height,gen.weight,gen.hobby,att.id,att.name,att2.id,att2.name,gro.id,gro.name from `genius` gen left join `genius_attributes` att on gen.primary_attributes_id = att.id left join `genius_attributes` att2 on gen.secondary_attributes_id = att2.id left join `group_table` gro on gen.group_id = gro.id where gen.number = ?",
		id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var details entity.Spirit
	if row.Next() {
		err = row.Scan(
			&details.Avatar,
			&details.Number,
			&details.Name,
			&details.Description,
			&details.RacePower,
			&details.RaceAttack,
			&details.RaceDefense,
			&details.RaceMagicAttack,
			&details.RaceMagicDefense,
			&details.RaceSpeed,
			&details.Height,
			&details.Weight,
			&details.Hobby,
			&details.PrimaryAttributes.Id,
			&details.PrimaryAttributes.Name,
			&details.SecondaryAttributes.Id,
			&details.SecondaryAttributes.Name,
			&details.Group.Id,
			&details.Group.Name)
		if err != nil {
			return nil, err
		}
	}
	row, err = utils.Database.Query("select s.Id, s.name, s.description, s.value, s.amount, s.speed, s.is_genetic, s.additional_effects, s.is_be, att3.id, att3.name, st.id, st.name from `genius_skill` gs left join  `skill` s on gs.skill_id = s.id left join `genius_attributes` att3 on s.attributes_id = att3.id  left join `skill_type` st on s.skill_type_id = st.id  where gs.genius_id = ?",
		id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	details.Skills = make([]entity.Skill, 0)
	for row.Next() {
		var skill entity.Skill
		err = row.Scan(
			&skill.Id,
			&skill.Name,
			&skill.Description,
			&skill.Value,
			&skill.Amount,
			&skill.Speed,
			&skill.IsGenetic,
			&skill.AdditionalEffects,
			&skill.IsBe,
			&skill.Attributes.Id,
			&skill.Attributes.Name,
			&skill.SkillType.Id,
			&skill.SkillType.Name)
		details.Skills = append(details.Skills, skill)
		if err != nil {
			return nil, err
		}
	}
	return &details, nil
}

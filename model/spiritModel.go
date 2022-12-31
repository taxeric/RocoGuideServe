package model

import (
	"RocoGuide/entity"
	"RocoGuide/utils"
	"fmt"
)

func GetSpiritList(page int, amount int, keywords string, seriesId int) ([]entity.SpiritListItem, int) {
	var sql = "select g.id,g.number,g.avatar,g.name,att.id,att.name ,att2.id,att2.name from `genius` g  left join `genius_attributes` att on g.primary_attributes_id = att.id left join `genius_attributes` att2 on g.secondary_attributes_id = att2.id where g.name like ? and g.series = ? order by g.id desc limit ? offset ?"
	var sqlAmount = "select COUNT(*) from `genius` g  left join `genius_attributes` att on g.primary_attributes_id = att.id left join `genius_attributes` att2 on g.secondary_attributes_id = att2.id where g.name like ? and g.series = ?"
	row, err := utils.Database.Query(sql, "%"+keywords+"%", seriesId, amount, (page-1)*amount)
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
	row, err = utils.Database.Query(sqlAmount, "%"+keywords+"%", seriesId)
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
	var sql = "insert into genius(avatar,number,name,description,primary_attributes_id,secondary_attributes_id,race_power,race_attack,race_defense,race_magic_attack,race_magic_defense,race_speed,group_id,series,lineage,height,weight,hobby) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	row, err := utils.Database.Exec(
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
		spirit.Series.Id,
		spirit.Lineage.Id,
		spirit.Height,
		spirit.Weight,
		spirit.Hobby,
	)
	if err != nil {
		fmt.Println(err)
	}
	spiritSqlId, err := row.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	if len(spirit.Skills) != 0 {
		for _, v := range spirit.Skills {
			fmt.Printf("data -> %v", v)
			row, _ = utils.Database.Exec("insert into genius_skill(genius_id,skill_id) values (?,?)",
				spiritSqlId,
				v.Id)
			_, _ = row.LastInsertId() // 操作影响的行数
		}
	}
	return spiritSqlId
}

func GetSpiritDetailsById(id int) (*entity.Spirit, error) {
	row, err := utils.Database.Query("select gen.id,gen.avatar,gen.number,gen.name,gen.description,gen.race_power,gen.race_attack,gen.race_defense,gen.race_magic_attack,gen.race_magic_defense,gen.race_speed,gen.height,gen.weight,gen.hobby,seri.id,seri.name,line.id,line.name,line.introduce,line.icon,att.id,att.name,att2.id,att2.name,gro.id,gro.name from `genius` gen left join `genius_attributes` att on gen.primary_attributes_id = att.id left join `genius_attributes` att2 on gen.secondary_attributes_id = att2.id left join `group_table` gro on gen.group_id = gro.id left join `series` seri on gen.series=seri.id left join `lineage` line on gen.lineage=line.id where gen.number = ?",
		id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var details entity.Spirit
	if row.Next() {
		err = row.Scan(
			&details.Id,
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
			&details.Series.Id,
			&details.Series.Name,
			&details.Lineage.Id,
			&details.Lineage.Name,
			&details.Lineage.Introduce,
			&details.Lineage.Icon,
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

func UpdateSpirit(spirit *entity.Spirit) int64 {
	sql := "update genius set avatar=?,name=?,description=?,primary_attributes_id=?,secondary_attributes_id=?,race_power=?,race_attack=?,race_defense=?,race_magic_attack=?,race_magic_defense=?,race_speed=?,group_id=?,lineage=?,series=?,height=?,weight=?,hobby=? where number = ?"
	row, _ := utils.Database.Exec(sql,
		spirit.Avatar,
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
		spirit.Lineage.Id,
		spirit.Series.Id,
		spirit.Height,
		spirit.Weight,
		spirit.Hobby,
		spirit.Number,
	)
	spiritId, _ := row.RowsAffected()
	return spiritId
}

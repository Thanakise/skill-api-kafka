package skill

import (
	"fmt"

	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/lib/pq"
	"github.com/thanakize/skill-api-kafka/consumer/database"
)

type Handler struct {
	DB database.Db
}

func InitHandler(db database.Db) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h Handler) ActiveHandler(msg *sarama.ConsumerMessage) {
	key := string(msg.Key)
	switch key {
	case "insert":
		var skill Skill
		err := json.Unmarshal(msg.Value, &skill)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		_, err = h.InsertSkill(skill)
		if err != nil {
			fmt.Println(err.Error())
			break

		}

	case "put":
		var skill Skill
		key := msg.Headers[0].Value
		err := json.Unmarshal(msg.Value, &skill)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		newSkill, err := h.UpdateSkill(skill, string(key))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(newSkill)

	default:
		fmt.Println("unknow type")
	}
}

func (h Handler) InsertSkill(skill Skill) (Skill, error) {
	q := "INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key, name, description, logo, tags"
	row := h.DB.DB.QueryRow(q, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
	var newSkill Skill
	err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))
	if err != nil {

		return Skill{}, err
	}
	return newSkill, nil
}
func (h Handler) UpdateSkill(skill Skill, key string) (Skill, error) {
	q := "update skill set name = $2, description = $3, logo = $4, tags = $5 where key = $1 RETURNING key, name, description, logo, tags"
	// q := "INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id"
	row := h.DB.DB.QueryRow(q, key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
	var newSkill Skill
	err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))

	if err != nil {
		return Skill{}, err
	}
	return newSkill, nil
}

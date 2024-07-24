package skill

import (
	"errors"
	"fmt"
	"reflect"

	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/lib/pq"
	"github.com/thanakize/skill-api-kafka/consumer/database"
)

type Handler struct {
	database.Db
}

type HandlerInterface interface {
	// bark() string
	ActiveHandler(msg *sarama.ConsumerMessage)
	// InsertSkill(skill Skill) (Skill, error)
	// UpdateSkill(skill Skill, key string) (Skill, error)
	// DeleteSkill(key string) error
	// PatchSkill(skill Skill, key string) (Skill, error)
}

func InitHandler(db database.Db) *Handler {
	return &Handler{
		db,
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
		_, err = h.UpdateSkill(skill, string(key))
		if err != nil {
			fmt.Println(err.Error())
			break
		}

	case "delete":
		key := msg.Headers[0].Value

		err := h.DeleteSkill(string(key))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	case "patch":
		var skill Skill
		key := msg.Headers[0].Value
		err := json.Unmarshal(msg.Value, &skill)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		_, err = h.PatchSkill(skill, string(key))
		if err != nil {
			fmt.Println(err.Error())
			break
		}

	default:
		fmt.Println("unknow type")
	}
}

func (h Handler) InsertSkill(skill Skill) (Skill, error) {
	q := "INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key, name, description, logo, tags"
	row := h.DB.QueryRow(q, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
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
	row := h.DB.QueryRow(q, key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
	var newSkill Skill
	err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))

	if err != nil {
		return Skill{}, err
	}
	return newSkill, nil
}
func (h Handler) DeleteSkill(key string) error {
	row, err := h.DB.Exec("DELETE FROM skill WHERE key = $1", key)
	if err != nil {
		return err
	}
	affect, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if affect == 0 {
		return errors.New("afect row is 0")
	}
	return nil
}
func (h Handler) PatchSkill(skill Skill, key string) (Skill, error) {
	v := reflect.ValueOf(skill)
	typeOfS := v.Type()
	var updateKey string
	var updateValue any
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		if value != "" {
			updateKey = typeOfS.Field(i).Name
			updateValue = value
			if reflect.TypeOf(updateValue) == reflect.TypeOf([]string{}) {
				updateValue = pq.Array(updateValue)
			}
			break
		}
	}
	result := fmt.Sprintf("update skill set %v = $1 where key = $2 RETURNING key, name, description, logo, tags", updateKey)
	row := h.DB.QueryRow(result, updateValue, key)
	var newSkill Skill
	err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))

	if err != nil {

		return Skill{}, err
	}

	return newSkill, nil
}

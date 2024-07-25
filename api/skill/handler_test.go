package skill_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thanakize/skill-api-kafka/api/skill"
)

type MockSkillReposiory struct {
	skill skill.Skill
	err   error
}

func InitMockSkillRepo(skill skill.Skill, err error) *MockSkillReposiory {
	return &MockSkillReposiory{
		skill: skill,
		err:   err,
	}
}

func (repo MockSkillReposiory) GetSkills() ([]skill.Skill, error) {
	return []skill.Skill{repo.skill}, repo.err
}
func (repo MockSkillReposiory) GetSkill(key string) (skill.Skill, error) {
	return repo.skill, repo.err
}

type MockProducer struct{}

func CreateMockProducer() *MockProducer {
	return &MockProducer{}
}

func (producer MockProducer) ProduceMessage(key string, skill skill.Skill, skillKey string) (skill.Skill, error) {
	return skill, nil
}

func TestGetHandler(t *testing.T) {
	mockSkill := skill.Skill{
		Key:         "test",
		Name:        "test",
		Logo:        "test",
		Description: "test",
		Tags:        []string{"test"},
	}
	mockProducer := CreateMockProducer()
	t.Run("Success", func(t *testing.T) {
		mockRepo := InitMockSkillRepo(mockSkill, nil)
		handler := skill.NewHandler(mockRepo, mockProducer)

		// Create a response recorder
		w := httptest.NewRecorder()
		// Create a new gin context
		c, _ := gin.CreateTestContext(w)
		// Create a new request
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/skills", nil)
		c.Request = req

		handler.GetSkillsHandler(c)

		// Check the status code
		assert.Equal(t, http.StatusOK, w.Code)
		// Check the response body
		expectedResponse := `{"status":"success","data":[{"key":"test","name":"test","description":"test","logo":"test","tags":["test"]}]}`
		assert.JSONEq(t, expectedResponse, w.Body.String())

	})
	t.Run("Error", func(t *testing.T) {
		mockRepo := InitMockSkillRepo(skill.Skill{}, errors.New("error"))
		handler := skill.NewHandler(mockRepo, mockProducer)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/skills", nil)
		c.Request = req

		handler.GetSkillsHandler(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedResponse := `{"status":"error","message":"error"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())

	})
}

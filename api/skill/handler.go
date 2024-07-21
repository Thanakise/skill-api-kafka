package skill

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
	skillRepo SkillReposiory
	producer  *Producer
}

func NewHandler(SkillRepo SkillReposiory, producer *Producer) *SkillHandler {
	return &SkillHandler{
		skillRepo: SkillRepo,
		producer:  producer,
	}
}

func (h SkillHandler) GetSkillsHandler(ctx *gin.Context) {
	skill, err := h.skillRepo.GetSkills()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, SkillsResponse{
		Status: "success",
		Data:   skill,
	})

}
func (h SkillHandler) GetSkillHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	skill, err := h.skillRepo.GetSkill(key)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: "Skill not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, SkillResponse{
		Status: "success",
		Data:   skill,
	})
}
func (h SkillHandler) PostSkillHandler(ctx *gin.Context) {
	var skill Skill
	if err := ctx.ShouldBindJSON(&skill); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	newSkill, err := h.producer.ProduceMessage("insert", skill, "")
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: "Skill already exists",
		})
		return
	}
	ctx.JSON(http.StatusOK, SkillResponse{
		Status: "success",
		Data:   newSkill,
	})
}
func (h SkillHandler) PutSkillHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	var skill Skill
	if err := ctx.ShouldBindJSON(&skill); err != nil {
		log.Println(err.Error())

		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	newSkill, err := h.producer.ProduceMessage("put", skill, key)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: "not be able to update skill",
		})
		return
	}
	ctx.JSON(http.StatusOK, SkillResponse{
		Status: "success",
		Data:   newSkill,
	})
}

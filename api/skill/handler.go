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
func (h SkillHandler) DeleteSkillHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	_, err := h.producer.ProduceMessage("delete", Skill{}, key)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: "not be able to delete skill",
		})
		return
	}
	ctx.JSON(http.StatusOK, DeleteResponse{
		Status:  "success",
		Message: "Skill deleted",
	})
}

func (h SkillHandler) PatchNameSkillHandler(ctx *gin.Context) {
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
	newSkill, err := h.producer.ProduceMessage("patch", Skill{Name: skill.Name}, key)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: "not be able to update skill name",
		})
		return
	}
	ctx.JSON(http.StatusOK, SkillResponse{
		Status: "success",
		Data:   newSkill,
	})
}
func (h SkillHandler) PatchDescriptionSkillHandler(ctx *gin.Context) {
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
	newSkill, err := h.producer.ProduceMessage("patch", Skill{Description: skill.Description}, key)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: "not be able to update skill description",
		})
		return
	}
	ctx.JSON(http.StatusOK, SkillResponse{
		Status: "success",
		Data:   newSkill,
	})
}
func (h SkillHandler) PatchLogoSkillHandler(ctx *gin.Context) {
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
	newSkill, err := h.producer.ProduceMessage("patch", Skill{Logo: skill.Logo}, key)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: "not be able to update skill logo",
		})
		return
	}
	ctx.JSON(http.StatusOK, SkillResponse{
		Status: "success",
		Data:   newSkill,
	})
}
func (h SkillHandler) PatchTagsSkillHandler(ctx *gin.Context) {
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

	newSkill, err := h.producer.ProduceMessage("patch", Skill{Tags: skill.Tags}, key)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, ErrorResponse{
			Status:  "error",
			Message: "not be able to update skill tags",
		})
		return
	}
	ctx.JSON(http.StatusOK, SkillResponse{
		Status: "success",
		Data:   newSkill,
	})

}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thanakize/skill-api-kafka/api/skill"
)

func InitRoute(r *gin.Engine, skillHandler *skill.SkillHandler) {

	router := r.Group("/api/v1/skills")

	router.GET("/", skillHandler.GetSkillsHandler)
	router.GET("/:key", skillHandler.GetSkillHandler)
	router.POST("/", skillHandler.PostSkillHandler)
	router.PUT("/:key", skillHandler.PutSkillHandler)
	router.DELETE("/:key", skillHandler.DeleteSkillHandler)
	router.PATCH("/:key/actions/name", skillHandler.PatchNameSkillHandler)
	router.PATCH("/:key/actions/description", skillHandler.PatchDescriptionSkillHandler)
	router.PATCH("/:key/actions/logo", skillHandler.PatchLogoSkillHandler)
	router.PATCH("/:key/actions/tags", skillHandler.PatchTagsSkillHandler)
}

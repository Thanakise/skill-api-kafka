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
	// router.DELETE("/:key", controllers.DeleteController(usecase))
	// router.PATCH("/:key/actions/name",controllers.PatchNameController(usecase))
	// router.PATCH("/:key/actions/description",controllers.PatchDescriptionController(usecase))
	// router.PATCH("/:key/actions/logo",controllers.PatchLogoController(usecase))
	// router.PATCH("/:key/actions/tags",controllers.PatchTagsController(usecase))
}

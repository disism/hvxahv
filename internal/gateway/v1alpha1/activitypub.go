package v1alpha1

import (
	"github.com/disism/hvxahv/internal/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func V1ActivityPub(r *gin.Engine) {
	// HTTP API for public query of ActivityPub.
	// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	r.GET("/.well-known/webfinger", handlers.WebFingerHandler)

	// https://www.w3.org/TR/activitypub/#actor-objects
	// Get the actors in the activityPub protocol.
	r.GET("/u/:actor", handlers.GetActorHandler)

	// https://www.w3.org/TR/activitypub/#inbox
	// Inbox
	r.POST("/u/:actor/inbox", handlers.InboxHandler)

}

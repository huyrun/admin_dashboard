package tables

import (
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"gorm.io/gorm"
)

func NewGenerators(db *gorm.DB, conn db.Connection) (map[string]table.Generator, error) {
	user, err := NewUser(db, conn)
	if err != nil {
		return nil, err
	}
	entity, err := NewEntity(db, conn)
	if err != nil {
		return nil, err
	}
	activity, err := NewActivity(user, entity, db, conn)
	if err != nil {
		return nil, err
	}
	entityComments, err := NewEntityComments(user, entity, db, conn)
	if err != nil {
		return nil, err
	}
	likedEntities, err := NewLikedEntities(user, entity, db, conn)
	if err != nil {
		return nil, err
	}
	savedEntities, err := NewSavedEntities(user, entity, db, conn)
	if err != nil {
		return nil, err
	}
	userRelationships, err := NewUserRelationships(user, entity, db, conn)
	if err != nil {
		return nil, err
	}
	tag, err := NewTag(db, conn)
	if err != nil {
		return nil, err
	}
	entityTag, err := NewEntityTags(entity, tag)
	if err != nil {
		return nil, err
	}
	category, err := NewCategory(db, conn)
	if err != nil {
		return nil, err
	}
	wishStory, err := NewWishStory(entity, db, conn)
	if err != nil {
		return nil, err
	}
	wish, err := NewWish(user, entity, category, db, conn)
	if err != nil {
		return nil, err
	}

	return map[string]table.Generator{
		"entity_comments":    entityComments.GetEntityCommentsTable,
		"entity_tags":        entityTag.GetEntityTagTable,
		"categories":         category.GetCategoriesTable,
		"entities":           entity.GetEntitiesTable,
		"activities":         activity.GetActivitiesTable,
		"goose_db_version":   GetGooseDbVersionTable,
		"liked_entities":     likedEntities.GetLikedEntitiesTable,
		"user_relationships": userRelationships.GetUserRelationshipsTable,
		"saved_entities":     savedEntities.GetSavedEntitiesTable,
		"wishes":             wish.GetWishTable,
		"wish_stories":       wishStory.GetWishStoryTable,
		"tags":               tag.GetTagsTable,
		"users":              user.GetUsersTable,

		// generators end
	}, nil
}

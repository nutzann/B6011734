package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edge"
)

// WasherEntity holds the schema definition for the WasherEntity entity.
type WasherEntity struct {
	ent.Schema
}

// Fields of the WasherEntity.
func (WasherEntity) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("url").NotEmpty(),
	}
}

// Edges of the WasherEntity.
func (WasherEntity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("WasherEntity").Unique(),
		edge.To("playlist_videos", Playlist_Video.Type).StorageKey(edge.Column("video_id")),
	}
}

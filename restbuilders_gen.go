package disgord

// Code generated - This file has been automatically generated by generate/restbuilders/main.go - DO NOT EDIT.
// Warning: This file is overwritten by the "go generate" command
// This file holds all the basic RESTBuilder methods a builder is expected to.

// CreateGuildEmojiBuilder is the interface for the builder.
type CreateGuildEmojiBuilder interface {
	IgnoreCache() CreateGuildEmojiBuilder
	CancelOnRatelimit() CreateGuildEmojiBuilder
	URLParam(name string, v interface{}) CreateGuildEmojiBuilder
	Set(name string, v interface{}) CreateGuildEmojiBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *createGuildEmojiBuilder) IgnoreCache() CreateGuildEmojiBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *createGuildEmojiBuilder) CancelOnRatelimit() CreateGuildEmojiBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *createGuildEmojiBuilder) URLParam(name string, v interface{}) CreateGuildEmojiBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *createGuildEmojiBuilder) Set(name string, v interface{}) CreateGuildEmojiBuilder {
	b.r.body[name] = v
	return b
}

// UpdateChannelBuilder is the interface for the builder.
type UpdateChannelBuilder interface {
	IgnoreCache() UpdateChannelBuilder
	CancelOnRatelimit() UpdateChannelBuilder
	URLParam(name string, v interface{}) UpdateChannelBuilder
	Set(name string, v interface{}) UpdateChannelBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateChannelBuilder) IgnoreCache() UpdateChannelBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateChannelBuilder) CancelOnRatelimit() UpdateChannelBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateChannelBuilder) URLParam(name string, v interface{}) UpdateChannelBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateChannelBuilder) Set(name string, v interface{}) UpdateChannelBuilder {
	b.r.body[name] = v
	return b
}

// UpdateCurrentUserBuilder is the interface for the builder.
type UpdateCurrentUserBuilder interface {
	IgnoreCache() UpdateCurrentUserBuilder
	CancelOnRatelimit() UpdateCurrentUserBuilder
	URLParam(name string, v interface{}) UpdateCurrentUserBuilder
	Set(name string, v interface{}) UpdateCurrentUserBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateCurrentUserBuilder) IgnoreCache() UpdateCurrentUserBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateCurrentUserBuilder) CancelOnRatelimit() UpdateCurrentUserBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateCurrentUserBuilder) URLParam(name string, v interface{}) UpdateCurrentUserBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateCurrentUserBuilder) Set(name string, v interface{}) UpdateCurrentUserBuilder {
	b.r.body[name] = v
	return b
}

// UpdateGuildBuilder is the interface for the builder.
type UpdateGuildBuilder interface {
	IgnoreCache() UpdateGuildBuilder
	CancelOnRatelimit() UpdateGuildBuilder
	URLParam(name string, v interface{}) UpdateGuildBuilder
	Set(name string, v interface{}) UpdateGuildBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildBuilder) IgnoreCache() UpdateGuildBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildBuilder) CancelOnRatelimit() UpdateGuildBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildBuilder) URLParam(name string, v interface{}) UpdateGuildBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildBuilder) Set(name string, v interface{}) UpdateGuildBuilder {
	b.r.body[name] = v
	return b
}

// UpdateGuildEmbedBuilder is the interface for the builder.
type UpdateGuildEmbedBuilder interface {
	IgnoreCache() UpdateGuildEmbedBuilder
	CancelOnRatelimit() UpdateGuildEmbedBuilder
	URLParam(name string, v interface{}) UpdateGuildEmbedBuilder
	Set(name string, v interface{}) UpdateGuildEmbedBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildEmbedBuilder) IgnoreCache() UpdateGuildEmbedBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildEmbedBuilder) CancelOnRatelimit() UpdateGuildEmbedBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildEmbedBuilder) URLParam(name string, v interface{}) UpdateGuildEmbedBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildEmbedBuilder) Set(name string, v interface{}) UpdateGuildEmbedBuilder {
	b.r.body[name] = v
	return b
}

// UpdateGuildEmojiBuilder is the interface for the builder.
type UpdateGuildEmojiBuilder interface {
	IgnoreCache() UpdateGuildEmojiBuilder
	CancelOnRatelimit() UpdateGuildEmojiBuilder
	URLParam(name string, v interface{}) UpdateGuildEmojiBuilder
	Set(name string, v interface{}) UpdateGuildEmojiBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildEmojiBuilder) IgnoreCache() UpdateGuildEmojiBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildEmojiBuilder) CancelOnRatelimit() UpdateGuildEmojiBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildEmojiBuilder) URLParam(name string, v interface{}) UpdateGuildEmojiBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildEmojiBuilder) Set(name string, v interface{}) UpdateGuildEmojiBuilder {
	b.r.body[name] = v
	return b
}

// UpdateGuildMemberBuilder is the interface for the builder.
type UpdateGuildMemberBuilder interface {
	IgnoreCache() UpdateGuildMemberBuilder
	CancelOnRatelimit() UpdateGuildMemberBuilder
	URLParam(name string, v interface{}) UpdateGuildMemberBuilder
	Set(name string, v interface{}) UpdateGuildMemberBuilder

	KickFromVoice() UpdateGuildMemberBuilder
	DeleteNick() UpdateGuildMemberBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildMemberBuilder) IgnoreCache() UpdateGuildMemberBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildMemberBuilder) CancelOnRatelimit() UpdateGuildMemberBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildMemberBuilder) URLParam(name string, v interface{}) UpdateGuildMemberBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildMemberBuilder) Set(name string, v interface{}) UpdateGuildMemberBuilder {
	b.r.body[name] = v
	return b
}

// UpdateGuildRoleBuilder is the interface for the builder.
type UpdateGuildRoleBuilder interface {
	IgnoreCache() UpdateGuildRoleBuilder
	CancelOnRatelimit() UpdateGuildRoleBuilder
	URLParam(name string, v interface{}) UpdateGuildRoleBuilder
	Set(name string, v interface{}) UpdateGuildRoleBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateGuildRoleBuilder) IgnoreCache() UpdateGuildRoleBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateGuildRoleBuilder) CancelOnRatelimit() UpdateGuildRoleBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateGuildRoleBuilder) URLParam(name string, v interface{}) UpdateGuildRoleBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateGuildRoleBuilder) Set(name string, v interface{}) UpdateGuildRoleBuilder {
	b.r.body[name] = v
	return b
}

// UpdateMessageBuilder is the interface for the builder.
type UpdateMessageBuilder interface {
	IgnoreCache() UpdateMessageBuilder
	CancelOnRatelimit() UpdateMessageBuilder
	URLParam(name string, v interface{}) UpdateMessageBuilder
	Set(name string, v interface{}) UpdateMessageBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateMessageBuilder) IgnoreCache() UpdateMessageBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateMessageBuilder) CancelOnRatelimit() UpdateMessageBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateMessageBuilder) URLParam(name string, v interface{}) UpdateMessageBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateMessageBuilder) Set(name string, v interface{}) UpdateMessageBuilder {
	b.r.body[name] = v
	return b
}

// UpdateWebhookBuilder is the interface for the builder.
type UpdateWebhookBuilder interface {
	IgnoreCache() UpdateWebhookBuilder
	CancelOnRatelimit() UpdateWebhookBuilder
	URLParam(name string, v interface{}) UpdateWebhookBuilder
	Set(name string, v interface{}) UpdateWebhookBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *updateWebhookBuilder) IgnoreCache() UpdateWebhookBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *updateWebhookBuilder) CancelOnRatelimit() UpdateWebhookBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *updateWebhookBuilder) URLParam(name string, v interface{}) UpdateWebhookBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *updateWebhookBuilder) Set(name string, v interface{}) UpdateWebhookBuilder {
	b.r.body[name] = v
	return b
}

// BasicBuilder is the interface for the builder.
type BasicBuilder interface {
	IgnoreCache() BasicBuilder
	CancelOnRatelimit() BasicBuilder
	URLParam(name string, v interface{}) BasicBuilder
	Set(name string, v interface{}) BasicBuilder
}

// IgnoreCache will not fetch the data from the cache if available, and always execute a
// a REST request. However, the response will always update the cache to keep it synced.
func (b *basicBuilder) IgnoreCache() BasicBuilder {
	b.r.IgnoreCache()
	return b
}

// CancelOnRatelimit will disable waiting if the request is rate limited by Discord.
func (b *basicBuilder) CancelOnRatelimit() BasicBuilder {
	b.r.CancelOnRatelimit()
	return b
}

// URLParam adds or updates an existing URL parameter.
// eg. URLParam("age", 34) will cause the URL `/test` to become `/test?age=34`
func (b *basicBuilder) URLParam(name string, v interface{}) BasicBuilder {
	b.r.queryParam(name, v)
	return b
}

// Set adds or updates an existing a body parameter
// eg. Set("age", 34) will cause the body `{}` to become `{"age":34}`
func (b *basicBuilder) Set(name string, v interface{}) BasicBuilder {
	b.r.body[name] = v
	return b
}

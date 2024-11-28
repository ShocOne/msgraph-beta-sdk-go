package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type ProductRevision struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewProductRevision instantiates a new ProductRevision and sets the default values.
func NewProductRevision()(*ProductRevision) {
    m := &ProductRevision{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateProductRevisionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductRevisionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProductRevision(), nil
}
// GetCatalogEntry gets the catalogEntry property value. The catalogEntry property
// returns a CatalogEntryable when successful
func (m *ProductRevision) GetCatalogEntry()(CatalogEntryable) {
    val, err := m.GetBackingStore().Get("catalogEntry")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CatalogEntryable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the content. Read-only.
// returns a *string when successful
func (m *ProductRevision) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProductRevision) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalogEntry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCatalogEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogEntry(val.(CatalogEntryable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isHotpatchUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHotpatchUpdate(val)
        }
        return nil
    }
    res["knowledgeBaseArticle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKnowledgeBaseArticleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKnowledgeBaseArticle(val.(KnowledgeBaseArticleable))
        }
        return nil
    }
    res["osBuild"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBuildVersionDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuild(val.(BuildVersionDetailsable))
        }
        return nil
    }
    res["product"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProduct(val)
        }
        return nil
    }
    res["releaseDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateTime(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetIsHotpatchUpdate gets the isHotpatchUpdate property value. True indicates that the content is hotpatchable; otherwise, false. For more information, see Deploy a hotpatch quality update using Windows Autopatch. Read-only.
// returns a *bool when successful
func (m *ProductRevision) GetIsHotpatchUpdate()(*bool) {
    val, err := m.GetBackingStore().Get("isHotpatchUpdate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKnowledgeBaseArticle gets the knowledgeBaseArticle property value. The knowledge base article associated with the product revision.
// returns a KnowledgeBaseArticleable when successful
func (m *ProductRevision) GetKnowledgeBaseArticle()(KnowledgeBaseArticleable) {
    val, err := m.GetBackingStore().Get("knowledgeBaseArticle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(KnowledgeBaseArticleable)
    }
    return nil
}
// GetOsBuild gets the osBuild property value. The osBuild property
// returns a BuildVersionDetailsable when successful
func (m *ProductRevision) GetOsBuild()(BuildVersionDetailsable) {
    val, err := m.GetBackingStore().Get("osBuild")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BuildVersionDetailsable)
    }
    return nil
}
// GetProduct gets the product property value. The product of the revision. Possible values are: Windows 10, Windows 11. Read-only.
// returns a *string when successful
func (m *ProductRevision) GetProduct()(*string) {
    val, err := m.GetBackingStore().Get("product")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReleaseDateTime gets the releaseDateTime property value. The release date for the content. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *ProductRevision) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("releaseDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetVersion gets the version property value. The version of the feature update. Read-only.
// returns a *string when successful
func (m *ProductRevision) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ProductRevision) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("catalogEntry", m.GetCatalogEntry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHotpatchUpdate", m.GetIsHotpatchUpdate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("knowledgeBaseArticle", m.GetKnowledgeBaseArticle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("osBuild", m.GetOsBuild())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("product", m.GetProduct())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalogEntry sets the catalogEntry property value. The catalogEntry property
func (m *ProductRevision) SetCatalogEntry(value CatalogEntryable)() {
    err := m.GetBackingStore().Set("catalogEntry", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the content. Read-only.
func (m *ProductRevision) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHotpatchUpdate sets the isHotpatchUpdate property value. True indicates that the content is hotpatchable; otherwise, false. For more information, see Deploy a hotpatch quality update using Windows Autopatch. Read-only.
func (m *ProductRevision) SetIsHotpatchUpdate(value *bool)() {
    err := m.GetBackingStore().Set("isHotpatchUpdate", value)
    if err != nil {
        panic(err)
    }
}
// SetKnowledgeBaseArticle sets the knowledgeBaseArticle property value. The knowledge base article associated with the product revision.
func (m *ProductRevision) SetKnowledgeBaseArticle(value KnowledgeBaseArticleable)() {
    err := m.GetBackingStore().Set("knowledgeBaseArticle", value)
    if err != nil {
        panic(err)
    }
}
// SetOsBuild sets the osBuild property value. The osBuild property
func (m *ProductRevision) SetOsBuild(value BuildVersionDetailsable)() {
    err := m.GetBackingStore().Set("osBuild", value)
    if err != nil {
        panic(err)
    }
}
// SetProduct sets the product property value. The product of the revision. Possible values are: Windows 10, Windows 11. Read-only.
func (m *ProductRevision) SetProduct(value *string)() {
    err := m.GetBackingStore().Set("product", value)
    if err != nil {
        panic(err)
    }
}
// SetReleaseDateTime sets the releaseDateTime property value. The release date for the content. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ProductRevision) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("releaseDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. The version of the feature update. Read-only.
func (m *ProductRevision) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
type ProductRevisionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCatalogEntry()(CatalogEntryable)
    GetDisplayName()(*string)
    GetIsHotpatchUpdate()(*bool)
    GetKnowledgeBaseArticle()(KnowledgeBaseArticleable)
    GetOsBuild()(BuildVersionDetailsable)
    GetProduct()(*string)
    GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersion()(*string)
    SetCatalogEntry(value CatalogEntryable)()
    SetDisplayName(value *string)()
    SetIsHotpatchUpdate(value *bool)()
    SetKnowledgeBaseArticle(value KnowledgeBaseArticleable)()
    SetOsBuild(value BuildVersionDetailsable)()
    SetProduct(value *string)()
    SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersion(value *string)()
}

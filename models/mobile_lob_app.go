package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileLobApp 
type MobileLobApp struct {
    MobileApp
    // The internal committed content version.
    committedContentVersion *string
    // The list of content versions for this app.
    contentVersions []MobileAppContentable
    // The name of the main Lob application file.
    fileName *string
    // The total size, including all uploaded files.
    size *int64
}
// NewMobileLobApp instantiates a new MobileLobApp and sets the default values.
func NewMobileLobApp()(*MobileLobApp) {
    m := &MobileLobApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.mobileLobApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMobileLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidLobApp":
                        return NewAndroidLobApp(), nil
                    case "#microsoft.graph.iosLobApp":
                        return NewIosLobApp(), nil
                    case "#microsoft.graph.macOSDmgApp":
                        return NewMacOSDmgApp(), nil
                    case "#microsoft.graph.macOSLobApp":
                        return NewMacOSLobApp(), nil
                    case "#microsoft.graph.win32LobApp":
                        return NewWin32LobApp(), nil
                    case "#microsoft.graph.windowsAppX":
                        return NewWindowsAppX(), nil
                    case "#microsoft.graph.windowsMobileMSI":
                        return NewWindowsMobileMSI(), nil
                    case "#microsoft.graph.windowsPhone81AppX":
                        return NewWindowsPhone81AppX(), nil
                    case "#microsoft.graph.windowsPhone81AppXBundle":
                        return NewWindowsPhone81AppXBundle(), nil
                    case "#microsoft.graph.windowsPhoneXAP":
                        return NewWindowsPhoneXAP(), nil
                    case "#microsoft.graph.windowsUniversalAppX":
                        return NewWindowsUniversalAppX(), nil
                }
            }
        }
    }
    return NewMobileLobApp(), nil
}
// GetCommittedContentVersion gets the committedContentVersion property value. The internal committed content version.
func (m *MobileLobApp) GetCommittedContentVersion()(*string) {
    return m.committedContentVersion
}
// GetContentVersions gets the contentVersions property value. The list of content versions for this app.
func (m *MobileLobApp) GetContentVersions()([]MobileAppContentable) {
    return m.contentVersions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["committedContentVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCommittedContentVersion)
    res["contentVersions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppContentFromDiscriminatorValue , m.SetContentVersions)
    res["fileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileName)
    res["size"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetSize)
    return res
}
// GetFileName gets the fileName property value. The name of the main Lob application file.
func (m *MobileLobApp) GetFileName()(*string) {
    return m.fileName
}
// GetSize gets the size property value. The total size, including all uploaded files.
func (m *MobileLobApp) GetSize()(*int64) {
    return m.size
}
// Serialize serializes information the current object
func (m *MobileLobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("committedContentVersion", m.GetCommittedContentVersion())
        if err != nil {
            return err
        }
    }
    if m.GetContentVersions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetContentVersions())
        err = writer.WriteCollectionOfObjectValues("contentVersions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCommittedContentVersion sets the committedContentVersion property value. The internal committed content version.
func (m *MobileLobApp) SetCommittedContentVersion(value *string)() {
    m.committedContentVersion = value
}
// SetContentVersions sets the contentVersions property value. The list of content versions for this app.
func (m *MobileLobApp) SetContentVersions(value []MobileAppContentable)() {
    m.contentVersions = value
}
// SetFileName sets the fileName property value. The name of the main Lob application file.
func (m *MobileLobApp) SetFileName(value *string)() {
    m.fileName = value
}
// SetSize sets the size property value. The total size, including all uploaded files.
func (m *MobileLobApp) SetSize(value *int64)() {
    m.size = value
}

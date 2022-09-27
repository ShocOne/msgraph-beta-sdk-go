package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GeoCoordinates 
type GeoCoordinates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
    altitude *float64
    // Optional. The latitude, in decimal, for the item. Writable on OneDrive Personal.
    latitude *float64
    // Optional. The longitude, in decimal, for the item. Writable on OneDrive Personal.
    longitude *float64
    // The OdataType property
    odataType *string
}
// NewGeoCoordinates instantiates a new geoCoordinates and sets the default values.
func NewGeoCoordinates()(*GeoCoordinates) {
    m := &GeoCoordinates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.geoCoordinates";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGeoCoordinatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGeoCoordinatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGeoCoordinates(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GeoCoordinates) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAltitude gets the altitude property value. Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
func (m *GeoCoordinates) GetAltitude()(*float64) {
    return m.altitude
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GeoCoordinates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["altitude"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetAltitude)
    res["latitude"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetLatitude)
    res["longitude"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetLongitude)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetLatitude gets the latitude property value. Optional. The latitude, in decimal, for the item. Writable on OneDrive Personal.
func (m *GeoCoordinates) GetLatitude()(*float64) {
    return m.latitude
}
// GetLongitude gets the longitude property value. Optional. The longitude, in decimal, for the item. Writable on OneDrive Personal.
func (m *GeoCoordinates) GetLongitude()(*float64) {
    return m.longitude
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GeoCoordinates) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GeoCoordinates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("altitude", m.GetAltitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("latitude", m.GetLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("longitude", m.GetLongitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GeoCoordinates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAltitude sets the altitude property value. Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
func (m *GeoCoordinates) SetAltitude(value *float64)() {
    m.altitude = value
}
// SetLatitude sets the latitude property value. Optional. The latitude, in decimal, for the item. Writable on OneDrive Personal.
func (m *GeoCoordinates) SetLatitude(value *float64)() {
    m.latitude = value
}
// SetLongitude sets the longitude property value. Optional. The longitude, in decimal, for the item. Writable on OneDrive Personal.
func (m *GeoCoordinates) SetLongitude(value *float64)() {
    m.longitude = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GeoCoordinates) SetOdataType(value *string)() {
    m.odataType = value
}

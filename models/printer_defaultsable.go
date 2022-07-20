package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterDefaultsable 
type PrinterDefaultsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColorMode()(*PrintColorMode)
    GetContentType()(*string)
    GetCopiesPerJob()(*int32)
    GetDocumentMimeType()(*string)
    GetDpi()(*int32)
    GetDuplexConfiguration()(*PrintDuplexConfiguration)
    GetDuplexMode()(*PrintDuplexMode)
    GetFinishings()([]string)
    GetFitPdfToPage()(*bool)
    GetInputBin()(*string)
    GetMediaColor()(*string)
    GetMediaSize()(*string)
    GetMediaType()(*string)
    GetMultipageLayout()(*PrintMultipageLayout)
    GetOdataType()(*string)
    GetOrientation()(*PrintOrientation)
    GetOutputBin()(*string)
    GetPagesPerSheet()(*int32)
    GetPdfFitToPage()(*bool)
    GetPresentationDirection()(*PrintPresentationDirection)
    GetPrintColorConfiguration()(*PrintColorConfiguration)
    GetPrintQuality()(*PrintQuality)
    GetQuality()(*PrintQuality)
    GetScaling()(*PrintScaling)
    SetColorMode(value *PrintColorMode)()
    SetContentType(value *string)()
    SetCopiesPerJob(value *int32)()
    SetDocumentMimeType(value *string)()
    SetDpi(value *int32)()
    SetDuplexConfiguration(value *PrintDuplexConfiguration)()
    SetDuplexMode(value *PrintDuplexMode)()
    SetFinishings(value []string)()
    SetFitPdfToPage(value *bool)()
    SetInputBin(value *string)()
    SetMediaColor(value *string)()
    SetMediaSize(value *string)()
    SetMediaType(value *string)()
    SetMultipageLayout(value *PrintMultipageLayout)()
    SetOdataType(value *string)()
    SetOrientation(value *PrintOrientation)()
    SetOutputBin(value *string)()
    SetPagesPerSheet(value *int32)()
    SetPdfFitToPage(value *bool)()
    SetPresentationDirection(value *PrintPresentationDirection)()
    SetPrintColorConfiguration(value *PrintColorConfiguration)()
    SetPrintQuality(value *PrintQuality)()
    SetQuality(value *PrintQuality)()
    SetScaling(value *PrintScaling)()
}

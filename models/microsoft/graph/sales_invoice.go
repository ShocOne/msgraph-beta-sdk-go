package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SalesInvoice 
type SalesInvoice struct {
    Entity
    // 
    billingPostalAddress *PostalAddressType;
    // 
    billToCustomerId *string;
    // 
    billToCustomerNumber *string;
    // 
    billToName *string;
    // 
    currency *Currency;
    // 
    currencyCode *string;
    // 
    currencyId *string;
    // 
    customer *Customer;
    // 
    customerId *string;
    // 
    customerName *string;
    // 
    customerNumber *string;
    // 
    customerPurchaseOrderReference *string;
    // 
    discountAmount *float64;
    // 
    discountAppliedBeforeTax *bool;
    // 
    dueDate *string;
    // 
    email *string;
    // 
    externalDocumentNumber *string;
    // 
    invoiceDate *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    number *string;
    // 
    orderId *string;
    // 
    orderNumber *string;
    // 
    paymentTerm *PaymentTerm;
    // 
    paymentTermsId *string;
    // 
    phoneNumber *string;
    // 
    pricesIncludeTax *bool;
    // 
    salesInvoiceLines []SalesInvoiceLine;
    // 
    salesperson *string;
    // 
    sellingPostalAddress *PostalAddressType;
    // 
    shipmentMethod *ShipmentMethod;
    // 
    shipmentMethodId *string;
    // 
    shippingPostalAddress *PostalAddressType;
    // 
    shipToContact *string;
    // 
    shipToName *string;
    // 
    status *string;
    // 
    totalAmountExcludingTax *float64;
    // 
    totalAmountIncludingTax *float64;
    // 
    totalTaxAmount *float64;
}
// NewSalesInvoice instantiates a new salesInvoice and sets the default values.
func NewSalesInvoice()(*SalesInvoice) {
    m := &SalesInvoice{
        Entity: *NewEntity(),
    }
    return m
}
// GetBillingPostalAddress gets the billingPostalAddress property value. 
func (m *SalesInvoice) GetBillingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.billingPostalAddress
    }
}
// GetBillToCustomerId gets the billToCustomerId property value. 
func (m *SalesInvoice) GetBillToCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerId
    }
}
// GetBillToCustomerNumber gets the billToCustomerNumber property value. 
func (m *SalesInvoice) GetBillToCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToCustomerNumber
    }
}
// GetBillToName gets the billToName property value. 
func (m *SalesInvoice) GetBillToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.billToName
    }
}
// GetCurrency gets the currency property value. 
func (m *SalesInvoice) GetCurrency()(*Currency) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// GetCurrencyCode gets the currencyCode property value. 
func (m *SalesInvoice) GetCurrencyCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyCode
    }
}
// GetCurrencyId gets the currencyId property value. 
func (m *SalesInvoice) GetCurrencyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currencyId
    }
}
// GetCustomer gets the customer property value. 
func (m *SalesInvoice) GetCustomer()(*Customer) {
    if m == nil {
        return nil
    } else {
        return m.customer
    }
}
// GetCustomerId gets the customerId property value. 
func (m *SalesInvoice) GetCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerId
    }
}
// GetCustomerName gets the customerName property value. 
func (m *SalesInvoice) GetCustomerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerName
    }
}
// GetCustomerNumber gets the customerNumber property value. 
func (m *SalesInvoice) GetCustomerNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerNumber
    }
}
// GetCustomerPurchaseOrderReference gets the customerPurchaseOrderReference property value. 
func (m *SalesInvoice) GetCustomerPurchaseOrderReference()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerPurchaseOrderReference
    }
}
// GetDiscountAmount gets the discountAmount property value. 
func (m *SalesInvoice) GetDiscountAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.discountAmount
    }
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. 
func (m *SalesInvoice) GetDiscountAppliedBeforeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discountAppliedBeforeTax
    }
}
// GetDueDate gets the dueDate property value. 
func (m *SalesInvoice) GetDueDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueDate
    }
}
// GetEmail gets the email property value. 
func (m *SalesInvoice) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetExternalDocumentNumber gets the externalDocumentNumber property value. 
func (m *SalesInvoice) GetExternalDocumentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalDocumentNumber
    }
}
// GetInvoiceDate gets the invoiceDate property value. 
func (m *SalesInvoice) GetInvoiceDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invoiceDate
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *SalesInvoice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNumber gets the number property value. 
func (m *SalesInvoice) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// GetOrderId gets the orderId property value. 
func (m *SalesInvoice) GetOrderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderId
    }
}
// GetOrderNumber gets the orderNumber property value. 
func (m *SalesInvoice) GetOrderNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderNumber
    }
}
// GetPaymentTerm gets the paymentTerm property value. 
func (m *SalesInvoice) GetPaymentTerm()(*PaymentTerm) {
    if m == nil {
        return nil
    } else {
        return m.paymentTerm
    }
}
// GetPaymentTermsId gets the paymentTermsId property value. 
func (m *SalesInvoice) GetPaymentTermsId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.paymentTermsId
    }
}
// GetPhoneNumber gets the phoneNumber property value. 
func (m *SalesInvoice) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// GetPricesIncludeTax gets the pricesIncludeTax property value. 
func (m *SalesInvoice) GetPricesIncludeTax()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pricesIncludeTax
    }
}
// GetSalesInvoiceLines gets the salesInvoiceLines property value. 
func (m *SalesInvoice) GetSalesInvoiceLines()([]SalesInvoiceLine) {
    if m == nil {
        return nil
    } else {
        return m.salesInvoiceLines
    }
}
// GetSalesperson gets the salesperson property value. 
func (m *SalesInvoice) GetSalesperson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.salesperson
    }
}
// GetSellingPostalAddress gets the sellingPostalAddress property value. 
func (m *SalesInvoice) GetSellingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.sellingPostalAddress
    }
}
// GetShipmentMethod gets the shipmentMethod property value. 
func (m *SalesInvoice) GetShipmentMethod()(*ShipmentMethod) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethod
    }
}
// GetShipmentMethodId gets the shipmentMethodId property value. 
func (m *SalesInvoice) GetShipmentMethodId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipmentMethodId
    }
}
// GetShippingPostalAddress gets the shippingPostalAddress property value. 
func (m *SalesInvoice) GetShippingPostalAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.shippingPostalAddress
    }
}
// GetShipToContact gets the shipToContact property value. 
func (m *SalesInvoice) GetShipToContact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToContact
    }
}
// GetShipToName gets the shipToName property value. 
func (m *SalesInvoice) GetShipToName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shipToName
    }
}
// GetStatus gets the status property value. 
func (m *SalesInvoice) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTotalAmountExcludingTax gets the totalAmountExcludingTax property value. 
func (m *SalesInvoice) GetTotalAmountExcludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountExcludingTax
    }
}
// GetTotalAmountIncludingTax gets the totalAmountIncludingTax property value. 
func (m *SalesInvoice) GetTotalAmountIncludingTax()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalAmountIncludingTax
    }
}
// GetTotalTaxAmount gets the totalTaxAmount property value. 
func (m *SalesInvoice) GetTotalTaxAmount()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.totalTaxAmount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SalesInvoice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["billingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingPostalAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["billToCustomerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToCustomerId(val)
        }
        return nil
    }
    res["billToCustomerNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToCustomerNumber(val)
        }
        return nil
    }
    res["billToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillToName(val)
        }
        return nil
    }
    res["currency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrency() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val.(*Currency))
        }
        return nil
    }
    res["currencyCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyCode(val)
        }
        return nil
    }
    res["currencyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyId(val)
        }
        return nil
    }
    res["customer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomer() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomer(val.(*Customer))
        }
        return nil
    }
    res["customerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["customerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerName(val)
        }
        return nil
    }
    res["customerNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerNumber(val)
        }
        return nil
    }
    res["customerPurchaseOrderReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerPurchaseOrderReference(val)
        }
        return nil
    }
    res["discountAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAmount(val)
        }
        return nil
    }
    res["discountAppliedBeforeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAppliedBeforeTax(val)
        }
        return nil
    }
    res["dueDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDate(val)
        }
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["externalDocumentNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalDocumentNumber(val)
        }
        return nil
    }
    res["invoiceDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceDate(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["orderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderId(val)
        }
        return nil
    }
    res["orderNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderNumber(val)
        }
        return nil
    }
    res["paymentTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPaymentTerm() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaymentTerm(val.(*PaymentTerm))
        }
        return nil
    }
    res["paymentTermsId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaymentTermsId(val)
        }
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["pricesIncludeTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPricesIncludeTax(val)
        }
        return nil
    }
    res["salesInvoiceLines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSalesInvoiceLine() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SalesInvoiceLine, len(val))
            for i, v := range val {
                res[i] = *(v.(*SalesInvoiceLine))
            }
            m.SetSalesInvoiceLines(res)
        }
        return nil
    }
    res["salesperson"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSalesperson(val)
        }
        return nil
    }
    res["sellingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSellingPostalAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["shipmentMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShipmentMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipmentMethod(val.(*ShipmentMethod))
        }
        return nil
    }
    res["shipmentMethodId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipmentMethodId(val)
        }
        return nil
    }
    res["shippingPostalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShippingPostalAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["shipToContact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToContact(val)
        }
        return nil
    }
    res["shipToName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToName(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["totalAmountExcludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountExcludingTax(val)
        }
        return nil
    }
    res["totalAmountIncludingTax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountIncludingTax(val)
        }
        return nil
    }
    res["totalTaxAmount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTaxAmount(val)
        }
        return nil
    }
    return res
}
func (m *SalesInvoice) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SalesInvoice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("billingPostalAddress", m.GetBillingPostalAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("billToCustomerId", m.GetBillToCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("billToCustomerNumber", m.GetBillToCustomerNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("billToName", m.GetBillToName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("currencyCode", m.GetCurrencyCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("currencyId", m.GetCurrencyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customer", m.GetCustomer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerName", m.GetCustomerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerNumber", m.GetCustomerNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerPurchaseOrderReference", m.GetCustomerPurchaseOrderReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("discountAmount", m.GetDiscountAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("discountAppliedBeforeTax", m.GetDiscountAppliedBeforeTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dueDate", m.GetDueDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalDocumentNumber", m.GetExternalDocumentNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invoiceDate", m.GetInvoiceDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("orderId", m.GetOrderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("orderNumber", m.GetOrderNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("paymentTerm", m.GetPaymentTerm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("paymentTermsId", m.GetPaymentTermsId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pricesIncludeTax", m.GetPricesIncludeTax())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSalesInvoiceLines()))
        for i, v := range m.GetSalesInvoiceLines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("salesInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("salesperson", m.GetSalesperson())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sellingPostalAddress", m.GetSellingPostalAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shipmentMethod", m.GetShipmentMethod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shipmentMethodId", m.GetShipmentMethodId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shippingPostalAddress", m.GetShippingPostalAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shipToContact", m.GetShipToContact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shipToName", m.GetShipToName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalAmountExcludingTax", m.GetTotalAmountExcludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalAmountIncludingTax", m.GetTotalAmountIncludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalTaxAmount", m.GetTotalTaxAmount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBillingPostalAddress sets the billingPostalAddress property value. 
func (m *SalesInvoice) SetBillingPostalAddress(value *PostalAddressType)() {
    if m != nil {
        m.billingPostalAddress = value
    }
}
// SetBillToCustomerId sets the billToCustomerId property value. 
func (m *SalesInvoice) SetBillToCustomerId(value *string)() {
    if m != nil {
        m.billToCustomerId = value
    }
}
// SetBillToCustomerNumber sets the billToCustomerNumber property value. 
func (m *SalesInvoice) SetBillToCustomerNumber(value *string)() {
    if m != nil {
        m.billToCustomerNumber = value
    }
}
// SetBillToName sets the billToName property value. 
func (m *SalesInvoice) SetBillToName(value *string)() {
    if m != nil {
        m.billToName = value
    }
}
// SetCurrency sets the currency property value. 
func (m *SalesInvoice) SetCurrency(value *Currency)() {
    if m != nil {
        m.currency = value
    }
}
// SetCurrencyCode sets the currencyCode property value. 
func (m *SalesInvoice) SetCurrencyCode(value *string)() {
    if m != nil {
        m.currencyCode = value
    }
}
// SetCurrencyId sets the currencyId property value. 
func (m *SalesInvoice) SetCurrencyId(value *string)() {
    if m != nil {
        m.currencyId = value
    }
}
// SetCustomer sets the customer property value. 
func (m *SalesInvoice) SetCustomer(value *Customer)() {
    if m != nil {
        m.customer = value
    }
}
// SetCustomerId sets the customerId property value. 
func (m *SalesInvoice) SetCustomerId(value *string)() {
    if m != nil {
        m.customerId = value
    }
}
// SetCustomerName sets the customerName property value. 
func (m *SalesInvoice) SetCustomerName(value *string)() {
    if m != nil {
        m.customerName = value
    }
}
// SetCustomerNumber sets the customerNumber property value. 
func (m *SalesInvoice) SetCustomerNumber(value *string)() {
    if m != nil {
        m.customerNumber = value
    }
}
// SetCustomerPurchaseOrderReference sets the customerPurchaseOrderReference property value. 
func (m *SalesInvoice) SetCustomerPurchaseOrderReference(value *string)() {
    if m != nil {
        m.customerPurchaseOrderReference = value
    }
}
// SetDiscountAmount sets the discountAmount property value. 
func (m *SalesInvoice) SetDiscountAmount(value *float64)() {
    if m != nil {
        m.discountAmount = value
    }
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. 
func (m *SalesInvoice) SetDiscountAppliedBeforeTax(value *bool)() {
    if m != nil {
        m.discountAppliedBeforeTax = value
    }
}
// SetDueDate sets the dueDate property value. 
func (m *SalesInvoice) SetDueDate(value *string)() {
    if m != nil {
        m.dueDate = value
    }
}
// SetEmail sets the email property value. 
func (m *SalesInvoice) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetExternalDocumentNumber sets the externalDocumentNumber property value. 
func (m *SalesInvoice) SetExternalDocumentNumber(value *string)() {
    if m != nil {
        m.externalDocumentNumber = value
    }
}
// SetInvoiceDate sets the invoiceDate property value. 
func (m *SalesInvoice) SetInvoiceDate(value *string)() {
    if m != nil {
        m.invoiceDate = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *SalesInvoice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNumber sets the number property value. 
func (m *SalesInvoice) SetNumber(value *string)() {
    if m != nil {
        m.number = value
    }
}
// SetOrderId sets the orderId property value. 
func (m *SalesInvoice) SetOrderId(value *string)() {
    if m != nil {
        m.orderId = value
    }
}
// SetOrderNumber sets the orderNumber property value. 
func (m *SalesInvoice) SetOrderNumber(value *string)() {
    if m != nil {
        m.orderNumber = value
    }
}
// SetPaymentTerm sets the paymentTerm property value. 
func (m *SalesInvoice) SetPaymentTerm(value *PaymentTerm)() {
    if m != nil {
        m.paymentTerm = value
    }
}
// SetPaymentTermsId sets the paymentTermsId property value. 
func (m *SalesInvoice) SetPaymentTermsId(value *string)() {
    if m != nil {
        m.paymentTermsId = value
    }
}
// SetPhoneNumber sets the phoneNumber property value. 
func (m *SalesInvoice) SetPhoneNumber(value *string)() {
    if m != nil {
        m.phoneNumber = value
    }
}
// SetPricesIncludeTax sets the pricesIncludeTax property value. 
func (m *SalesInvoice) SetPricesIncludeTax(value *bool)() {
    if m != nil {
        m.pricesIncludeTax = value
    }
}
// SetSalesInvoiceLines sets the salesInvoiceLines property value. 
func (m *SalesInvoice) SetSalesInvoiceLines(value []SalesInvoiceLine)() {
    if m != nil {
        m.salesInvoiceLines = value
    }
}
// SetSalesperson sets the salesperson property value. 
func (m *SalesInvoice) SetSalesperson(value *string)() {
    if m != nil {
        m.salesperson = value
    }
}
// SetSellingPostalAddress sets the sellingPostalAddress property value. 
func (m *SalesInvoice) SetSellingPostalAddress(value *PostalAddressType)() {
    if m != nil {
        m.sellingPostalAddress = value
    }
}
// SetShipmentMethod sets the shipmentMethod property value. 
func (m *SalesInvoice) SetShipmentMethod(value *ShipmentMethod)() {
    if m != nil {
        m.shipmentMethod = value
    }
}
// SetShipmentMethodId sets the shipmentMethodId property value. 
func (m *SalesInvoice) SetShipmentMethodId(value *string)() {
    if m != nil {
        m.shipmentMethodId = value
    }
}
// SetShippingPostalAddress sets the shippingPostalAddress property value. 
func (m *SalesInvoice) SetShippingPostalAddress(value *PostalAddressType)() {
    if m != nil {
        m.shippingPostalAddress = value
    }
}
// SetShipToContact sets the shipToContact property value. 
func (m *SalesInvoice) SetShipToContact(value *string)() {
    if m != nil {
        m.shipToContact = value
    }
}
// SetShipToName sets the shipToName property value. 
func (m *SalesInvoice) SetShipToName(value *string)() {
    if m != nil {
        m.shipToName = value
    }
}
// SetStatus sets the status property value. 
func (m *SalesInvoice) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetTotalAmountExcludingTax sets the totalAmountExcludingTax property value. 
func (m *SalesInvoice) SetTotalAmountExcludingTax(value *float64)() {
    if m != nil {
        m.totalAmountExcludingTax = value
    }
}
// SetTotalAmountIncludingTax sets the totalAmountIncludingTax property value. 
func (m *SalesInvoice) SetTotalAmountIncludingTax(value *float64)() {
    if m != nil {
        m.totalAmountIncludingTax = value
    }
}
// SetTotalTaxAmount sets the totalTaxAmount property value. 
func (m *SalesInvoice) SetTotalTaxAmount(value *float64)() {
    if m != nil {
        m.totalTaxAmount = value
    }
}

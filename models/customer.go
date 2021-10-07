package models

import (
	"log"

	"github.com/creasty/defaults"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	NIK          string       `json:"nik" gorm:"column:nik"`
	FirstName    string       `json:"first_name" gorm:"column:first_name"`
	LastName     string       `json:"last_name" gorm:"column:last_name"`
	Address      *Address     `json:"address,omitempty" gorm:"embedded"`
	Company      *Company     `json:"company,omitempty" gorm:"embedded"`
	Contact      interface{}  `json:"contact,omitempty" gorm:"-"`
	ContactEmail ContactEmail `json:"-" gorm:"embedded"`
	ContactPhone ContactPhone `json:"-" gorm:"embedded"`

	Account []Account `json:"account_details,omitempty"`

	NickName    string `json:"nick_name" gorm:"-"`
	Citizenship string `default:"Unknown" json:"citizenship" gorm:"column:nationality"`

	ArrToDelete []string `json:"arr_to_delete,omitempty"`
}

type Address struct {
	Street  string `json:"street,omitempty" gorm:"column:address_street"`
	Zipcode string `json:"zip_code,omitempty" gorm:"column:address_zipcode"`
	City    string `json:"city,omitempty" gorm:"column:address_city"`
	Country string `json:"country,omitempty" gorm:"column:address_country"`
}

type Company struct {
	CompanyIndustry string `json:"company_industry" gorm:"column:company_industry"`
	CompanyName     string `json:"company_name" gorm:"column:company_name"`
	Partner         string `json:"partner" default:"unknown"`
}

type Contact struct {
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
}

type ContactEmail struct {
	ContactTypeEmail  string `gorm:"column:contact_type_email"`
	ContactValueEmail string `gorm:"column:contact_value_email"`
}

type ContactPhone struct {
	ContactTypePhone  string `gorm:"column:contact_type_phone"`
	ContactValuePhone string `gorm:"column:contact_value_phone"`
}

func (cus *Customer) MappingContact() {
	var contactArray []Contact

	if cus.ContactEmail.ContactValueEmail != "" {
		email := Contact{
			ContactType:  cus.ContactEmail.ContactTypeEmail,
			ContactValue: cus.ContactEmail.ContactValueEmail,
		}
		log.Println(email.ContactValue)
		contactArray = append(contactArray, email)
	}
	if cus.ContactPhone.ContactValuePhone != "" {
		phone := Contact{
			ContactType:  cus.ContactPhone.ContactTypePhone,
			ContactValue: cus.ContactPhone.ContactValuePhone,
		}
		log.Println(phone.ContactValue)
		contactArray = append(contactArray, phone)
	}

	if len(contactArray) == 1 {
		cus.Contact = contactArray[0]
	} else if len(contactArray) == 0 {
		cus.Contact = nil
	} else {
		cus.Contact = contactArray
	}
}

func (cus *Customer) DefineRequestDetails(details *RequestDetails) {
	if details.Address <= 0 {
		cus.Address = nil
	}
	if details.Company <= 0 {
		cus.Company = nil
	}
	if details.Contact <= 0 {
		cus.Contact = nil
		cus.ArrToDelete = nil
	}
}

func (c *Customer) AfterFind(tx *gorm.DB) (err error) {

	var contactArr []Contact

	if c.ContactEmail.ContactValueEmail != "" {
		contactArr = append(contactArr, Contact{
			ContactType:  c.ContactEmail.ContactTypeEmail,
			ContactValue: c.ContactEmail.ContactValueEmail,
		})
	}

	if c.ContactPhone.ContactValuePhone != "" {
		contactArr = append(contactArr, Contact{
			ContactType:  c.ContactPhone.ContactTypePhone,
			ContactValue: c.ContactPhone.ContactValuePhone,
		})
	}

	if len(contactArr) > 1 {
		c.Contact = contactArr
	} else if len(contactArr) == 1 {
		c.Contact = contactArr[0]
	} else {
		c.Contact = nil
	}

	c.NickName = c.FirstName

	if err := defaults.Set(c); err != nil {
		log.Println(err)
	}

	c.ArrToDelete = []string{"dummy1", "dummy2", "dummy3"}

	return nil
}

// GORM
type Tabler interface {
	TableName() string
}

func (Customer) TableName() string {
	return "customer_details"
}

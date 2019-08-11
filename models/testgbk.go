package models

type Testgbk struct {
	X string `orm:"column(x);size(5000);null"`
}

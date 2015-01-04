package models

import (
	_ "errors"
	"fmt"
	_ "reflect"
	_ "strings"

	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Country struct {
	Id             string  `orm:"column(Code);pk"`
	Name           string  `orm:"column(Name);size(52)"`
	Continent      string  `orm:"column(Continent)"`
	Region         string  `orm:"column(Region);size(26)"`
	SurfaceArea    float32 `orm:"column(SurfaceArea)"`
	IndepYear      int16   `orm:"column(IndepYear);null"`
	Population     int     `orm:"column(Population)"`
	LifeExpectancy float32 `orm:"column(LifeExpectancy);null"`
	GNP            float32 `orm:"column(GNP);null"`
	GNPOld         float32 `orm:"column(GNPOld);null"`
	LocalName      string  `orm:"column(LocalName);size(45)"`
	GovernmentForm string  `orm:"column(GovernmentForm);size(45)"`
	HeadOfState    string  `orm:"column(HeadOfState);size(60);null"`
	Capital        int     `orm:"column(Capital);null"`
	Code2          string  `orm:"column(Code2);size(2)"`
}

func (c *Country) TableName() string {
	return "Country"
}

func init() {
	orm.RegisterModel(new(Country))
}

func DeleteCountry(c string) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Country{Id: c}); err == nil {
		return num, nil
	} else {
		return 0, err
	}
}

func SaveCountry(new_c Country, insert bool) (int64, error) {
	o := orm.NewOrm()
	var num int64
	var err error
	if insert {
		if num, err = o.Insert(&new_c); err == nil {
			return num, nil
		}
	} else {
		if num, err = o.Update(&new_c); err == nil && num > 0 {
			return num, nil
		}
	}
	return 0, err
}

// GetAllCountry retrieves all Country matches certain condition. Returns empty list if
// no records exist
func GetAllCountry(search string) ([]*Country, error) {
	o := orm.NewOrm()
	o.Using("default")

	// var count int
	// o.Raw("select count(*) as Count from Country").QueryRow(&count)
	// beego.Info(count)

	// err := o.QueryTable("Country").All(&country)

	//working code
	var country []*Country
	// num, err := o.QueryTable("Country").All(&country)

	// fmt.Printf("Returned Rows Num: %s, %s", num, err)

	// beego.Info(qs)
	c := new(Country)
	num, err := o.QueryTable(c).Filter("Name__icontains", search).All(&country) // return a QuerySetter
	//qs.Filter("name__contains", "slene")
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	// fmt.Printf("qs %s", qs)
	// beego.Info(qs)
	// beego.Info(err)
	// beego.Info(num)
	// if err != orm.ErrNoRows && num > 0 {
	// 	return country, ""
	// }
	return country, err
	// qs := o.QueryTable(new(Country))
	// m := make(map[int]interface{})
	// m[0] = "india"
	// return
}

func GetCountry(countryId string) (c *Country, err error) {
	o := orm.NewOrm()
	o.Using("default")
	// c := new(Country)
	c = &Country{Id: countryId}
	fmt.Printf("GetCountry %s", countryId)
	// err := o.QueryTable(c).One(&c) // return a QuerySetter
	// return c, err
	err = o.Read(c)
	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	}
	return c, err

}

// // AddCountry insert a new Country into database and returns
// // last inserted Id on success.
// func AddCountry(m *Country) (id int64, err error) {
// 	o := orm.NewOrm()
// 	id, err = o.Insert(m)
// 	return
// }

// // GetCountryById retrieves Country by Id. Returns error if
// // Id doesn't exist
// func GetCountryById(id int) (v *Country, err error) {
// 	o := orm.NewOrm()
// 	v = &Country{Id: id}
// 	if err = o.Read(v); err == nil {
// 		return v, nil
// 	}
// 	return nil, err
// }

// // GetAllCountry retrieves all Country matches certain condition. Returns empty list if
// // no records exist
// func GetAllCountry(query map[string]string, fields []string, sortby []string, order []string,
// 	offset int64, limit int64) (ml []interface{}, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable(new(Country))
// 	// query k=v
// 	for k, v := range query {
// 		// rewrite dot-notation to Object__Attribute
// 		k = strings.Replace(k, ".", "__", -1)
// 		qs = qs.Filter(k, v)
// 	}
// 	// order by:
// 	var sortFields []string
// 	if len(sortby) != 0 {
// 		if len(sortby) == len(order) {
// 			// 1) for each sort field, there is an associated order
// 			for i, v := range sortby {
// 				orderby := ""
// 				if order[i] == "desc" {
// 					orderby = "-" + v
// 				} else if order[i] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 			qs = qs.OrderBy(sortFields...)
// 		} else if len(sortby) != len(order) && len(order) == 1 {
// 			// 2) there is exactly one order, all the sorted fields will be sorted by this order
// 			for _, v := range sortby {
// 				orderby := ""
// 				if order[0] == "desc" {
// 					orderby = "-" + v
// 				} else if order[0] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 		} else if len(sortby) != len(order) && len(order) != 1 {
// 			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
// 		}
// 	} else {
// 		if len(order) != 0 {
// 			return nil, errors.New("Error: unused 'order' fields")
// 		}
// 	}

// 	var l []Country
// 	qs = qs.OrderBy(sortFields...)
// 	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
// 		if len(fields) == 0 {
// 			for _, v := range l {
// 				ml = append(ml, v)
// 			}
// 		} else {
// 			// trim unused fields
// 			for _, v := range l {
// 				m := make(map[string]interface{})
// 				val := reflect.ValueOf(v)
// 				for _, fname := range fields {
// 					m[fname] = val.FieldByName(fname).Interface()
// 				}
// 				ml = append(ml, m)
// 			}
// 		}
// 		return ml, nil
// 	}
// 	return nil, err
// }

// // UpdateCountry updates Country by Id and returns error if
// // the record to be updated doesn't exist
// func UpdateCountryById(m *Country) (err error) {
// 	o := orm.NewOrm()
// 	v := Country{Id: m.Id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Update(m); err == nil {
// 			fmt.Println("Number of records updated in database:", num)
// 		}
// 	}
// 	return
// }

// // DeleteCountry deletes Country by Id and returns error if
// // the record to be deleted doesn't exist
// func DeleteCountry(id int) (err error) {
// 	o := orm.NewOrm()
// 	v := Country{Id: id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Delete(&Country{Id: id}); err == nil {
// 			fmt.Println("Number of records deleted in database:", num)
// 		}
// 	}
// 	return
// }

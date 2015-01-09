package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/rtanwar/bee_test/utils"
	"reflect"
	"strings"
)

type Users struct {
	Id                    int    `orm:"column(id);auto"`
	IpAddress             string `orm:"column(ip_address);size(15)"`
	Username              string `orm:"column(username);size(100)"`
	Password              string `orm:"column(password);size(255)"`
	Salt                  string `orm:"column(salt);size(255);null"`
	Email                 string `orm:"column(email);size(100)"`
	ActivationCode        string `orm:"column(activation_code);size(40);null"`
	ForgottenPasswordCode string `orm:"column(forgotten_password_code);size(40);null"`
	ForgottenPasswordTime uint   `orm:"column(forgotten_password_time);null"`
	RememberCode          string `orm:"column(remember_code);size(40);null"`
	CreatedOn             uint   `orm:"column(created_on)"`
	LastLogin             uint   `orm:"column(last_login);null"`
	Active                uint8  `orm:"column(active);null"`
	FirstName             string `orm:"column(first_name);size(50);null"`
	LastName              string `orm:"column(last_name);size(50);null"`
	Company               string `orm:"column(company);size(100);null"`
	Phone                 string `orm:"column(phone);size(20);null"`
}

func init() {
	orm.RegisterModel(new(Users))
}

// AddUsers insert a new Users into database and returns
// last inserted Id on success.
func AddUsers(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func CheckUser(identity string, password string) bool {
	fmt.Print("Inside Checking user")
	var u *Users
	var err error
	o := orm.NewOrm()
	u = &Users{Username: identity}
	err = o.Read(u)

	salt := u.Salt //utils.GetRandomString(10)
	encodedPwd := salt + "$" + utils.EncodePassword(password, salt)
	fmt.Printf("\nsalt: %s\npassword: %s", salt, encodedPwd)

	if err = o.Read(u, "Username"); err == nil {
		if (u.Username == identity) && (u.Password == encodedPwd) {
			return true
		}
	}
	return false
}

func GetUsersById(id int) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsers retrieves all Users matches certain condition. Returns empty list if
// no records exist
func GetAllUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Users
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUsers updates Users by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersById(m *Users) (err error) {
	o := orm.NewOrm()
	v := Users{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsers deletes Users by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsers(id int) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

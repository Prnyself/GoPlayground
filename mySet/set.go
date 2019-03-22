package mySet

type Set map[interface{}]bool

// whether the elem is existed in the set
func (this Set) IsExist(elem interface{}) bool {
	return this[elem]
}

// add an elem to the set
// if the elem is already existed, return false.
// otherwise return true for success adding
func (this Set) Add(elem interface{}) bool {
	if this.IsExist(elem) {
		return false
	}
	this[elem] = true
	return true
}

// return the length of the set, which means how many different elem in the set
func (this Set) Len() int {
	return len(this)
}

// delete an elem from the set
// return false if the elem is not existed
func (this Set) Delete(elem interface{}) bool {
	if !this.IsExist(elem) {
		return false
	}
	delete(this, elem)
	return true
}

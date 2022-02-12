package form

// Define a new type errors, which we will use to hold the validation errors
// message from form. The name of the form field will be used as the key in
// this map.
type errors map[string][]string

// Implement an Add method to add error messages for a given field to the map
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Implement a Get() method to retrieve the first error message for a give
// field from the map
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}

package GoVector

type VectorError struct {
	Error string
}

func CreateVector (coordinates ...float32) (Vector, VectorError) {

	if l := len(coordinates); l == 2 {
		return Vector{
			coordinates[0],
			coordinates[1],
			0,
		},
			nil
	}	else if l == 3 {
		return Vector{
			coordinates[0],
			coordinates[1],
			coordinates[2],
		},
			nil
	}	else {
		return nil,
			VectorError{
				"needs to be two or three variables",
			}
	}

	return nil,
		VectorError{
			"Unexpected VectorError",
		}
}

type Vector struct {
	x float32
	y float32
	z float32
}


func (v *Vector) set(coordinates ...float32) VectorError {

	if l := len(coordinates); l == 2 {
		v.x = coordinates[0]
		v.y = coordinates[1]
		return nil

	}	else if l == 3 {
		v.x = coordinates[0]
		v.y = coordinates[1]
		v.z = coordinates[2]
		return nil

	}	else {
		return VectorError{
			"needs to be two or three variables",
		}
	}

	return VectorError{
		"Unexpected VectorError",
	}
}

func (v Vector)  add(v2 Vector) Vector{

	return Vector{
		v.x + v2.x,
		v.y + v2.y,
		v.z + v2.z,
	}

}
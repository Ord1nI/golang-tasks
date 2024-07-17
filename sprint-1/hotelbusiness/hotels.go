//go:build !solution

package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	var minDate int = 100000
	var maxDate int = 0


	m := make(map[int]int)

	for _, i := range guests {
		if i.CheckInDate < minDate {
			minDate = i.CheckInDate
		}
		if i.CheckOutDate > maxDate {
			maxDate = i.CheckOutDate
		}
		for q := i.CheckInDate; q < i.CheckOutDate; q++ {
			m[q] += 1
		}
	}

	var loads []Load

	for q, tmpPersons := minDate, 0; q <= maxDate; q++ {
		if tmp := m[q]; tmp != tmpPersons {
			loads = append(loads, Load{q, tmp})
			tmpPersons = tmp
		}
	}

	return loads
}

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
	var min_date int = 100000
	var max_date int = 0


	m := make(map[int]int)

	for _, i := range guests {
		if i.CheckInDate < min_date {
			min_date = i.CheckInDate
		}
		if i.CheckOutDate > max_date {
			max_date = i.CheckOutDate
		}
		for q := i.CheckInDate; q < i.CheckOutDate; q++ {
			m[q] += 1
		}
	}

	var loads []Load

	for q, tmp_persons := min_date, 0; q <= max_date; q++ {
		if tmp := m[q]; tmp != tmp_persons {
			loads = append(loads, Load{q, tmp})
			tmp_persons = tmp
		}
	}

	return loads
}

package p

func notPrintfFuncAtAll1() {
	o := orm{}
	o.Update()
	o.Where()
}

type orm struct {
}

func (o orm) Update() orm {
	return o
}

func (o orm) Delete() orm {
	return o
}

func (o orm) Save() orm {
	return o
}

func (o orm) Where() orm {
	return o
}

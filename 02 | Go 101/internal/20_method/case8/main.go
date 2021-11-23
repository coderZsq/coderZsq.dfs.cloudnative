package main

type StringSet map[string]struct{}
func (ss StringSet) Has(key string) bool {
	_, present := ss[key] // 永不会产生恐慌，即使ss为nil。
	return present
}

type Age int
func (age *Age) IsNil() bool {
	return age == nil
}
func (age *Age) Increase() {
	*age++ // 如果age是一个空指针，则此行将产生一个恐慌。
}

func main() {
	_ = (StringSet(nil)).Has   // 不会产生恐慌
	_ = ((*Age)(nil)).IsNil    // 不会产生恐慌
	_ = ((*Age)(nil)).Increase // 不会产生恐慌

	_ = (StringSet(nil)).Has("key") // 不会产生恐慌
	_ = ((*Age)(nil)).IsNil()       // 不会产生恐慌

	// 下面这行将产生一个恐慌，但是此恐慌不是在调用方法的时
	// 候产生的，而是在此方法体内解引用空指针的时候产生的。
	((*Age)(nil)).Increase()
}


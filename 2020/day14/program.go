package day14

// Program is a represantation of the task for day14
type Program struct {
	andMask      int
	orMask       int
	floatMask    []int
	noChangeMask int
	operators    []map[string]int
}

// Run runs the function
func (p *Program) Run(cache map[int]int) map[int]int {
	for _, val := range p.operators {
		memLocation := val["memLocation"]
		value := val["value"]
		and := value & p.andMask
		result := and | p.orMask
		cache[memLocation] = result
	}

	return cache
}

// RunPart2 does exactly that
func (p *Program) RunPart2(cache map[int]int) map[int]int {
	for i := 0; i < len(p.operators); i++ {
		val := p.operators[i]
		memLocation := val["memLocation"]
		value := val["value"]
		for _, floatMask := range p.floatMask {
			memLocation = p.orMask | (p.noChangeMask & memLocation) | floatMask
			cache[memLocation] = value
		}
	}

	return cache
}

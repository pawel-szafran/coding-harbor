package buildorder

func FindBuildOrder(modules []Module) ([]ModuleName, error) {
	var (
		toVisit       = make([]ModuleName, len(modules))
		deps          = map[ModuleName][]ModuleName{}
		visited       = map[ModuleName]bool{}
		beeingVisited = map[ModuleName]bool{}
		order         = make([]ModuleName, 0, len(modules))
	)
	for i, m := range modules {
		toVisit[i] = m.Name
		deps[m.Name] = m.Deps
	}

	var visitDepsFirst func([]ModuleName) error
	visitDepsFirst = func(toVisit []ModuleName) error {
		for _, m := range toVisit {
			if !visited[m] {
				if beeingVisited[m] {
					return ErrDepCycle
				}
				beeingVisited[m] = true
				err := visitDepsFirst(deps[m])
				if err != nil {
					return err
				}
				beeingVisited[m] = false
				visited[m] = true
				order = append(order, m)
			}
		}
		return nil
	}

	err := visitDepsFirst(toVisit)
	return order, err
}

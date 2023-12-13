package day13

type Sparse struct {
	mat [][]int
}

func NewSparse(size int) *Sparse {
	mat := make([][]int, size)
	return &Sparse{
		mat: mat,
	}
}

func (s *Sparse) Set(idx int, val []int) {
	s.mat[idx] = val
}

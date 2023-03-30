package problem566

func matrixReshape(mat [][]int, r int, c int) [][]int {
    var reshaped [][]int = make([][]int, r)
    for i := 0; i < len(reshaped); i++ {
        reshaped[i] = make([]int, c)
    }
    
    var elems []int = make([]int, len(mat) * len(mat[0]))
    
    if len(mat) * len(mat[0]) == r * c {
        k := 0
        for i := 0; i < len(mat); i++ {
            for j := 0; j < len(mat[i]); j++ {
                elems[k] = mat[i][j]
                k++
            }
        }
        
        k = 0
        for i := 0; i < len(reshaped); i++ {
            for j := 0; j < len(reshaped[0]); j++ {
                reshaped[i][j] = elems[k]
                k++
            }
        }
        return reshaped
    } else {
        return mat
    }
}

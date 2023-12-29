package main

import (
	"math"
	"strconv"
	"strings"
)

func solveDay24Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ps, ds := make([][3]int, len(lines)), make([][3]int, len(lines))
	for i, s := range lines {
		s1, s2, _ := strings.Cut(s, "@")
		for j, p := range strings.Split(s1, ",") {
			v, _ := strconv.Atoi(strings.TrimSpace(p))
			ps[i][j] = v
		}
		for j, d := range strings.Split(s2, ",") {
			v, _ := strconv.Atoi(strings.TrimSpace(d))
			ds[i][j] = v
		}
	}
	// lo, hi := 7.0, 27.0 // for sample input
	var lo, hi float64 = 200_000_000_000_000, 400_000_000_000_000
	count := 0
	for i := 0; i < len(ps)-1; i++ {
		pix, piy := float64(ps[i][0]), float64(ps[i][1])
		dix, diy := float64(ds[i][0]), float64(ds[i][1])
		for j := i + 1; j < len(ps); j++ {
			pjx, pjy := float64(ps[j][0]), float64(ps[j][1])
			djx, djy := float64(ds[j][0]), float64(ds[j][1])
			// ax + by + c = 0 => y = sx + c => sx - y + c = 0 => a = s, b = -1; c = y - sx
			si, sj := diy/dix, djy/djx        // slopes
			ci, cj := piy-si*pix, pjy-sj*pjx  // constant offsets
			cx := (ci - cj) / (sj - si)       // intersection x
			cy := (sj*ci - si*cj) / (sj - si) // intersection y
			if dix > 0 && cx < pix || dix < 0 && cx > pix || diy > 0 && cy < piy || diy < 0 && cy > piy {
				continue // past intersection for i
			}
			if djx > 0 && cx < pjx || djx < 0 && cx > pjx || djy > 0 && cy < pjy || djy < 0 && cy > pjy {
				continue // past intersection for j
			}
			if cx < lo || cx > hi || cy < lo || cy > hi {
				continue // intersection outside the test area
			}
			count++
		}
	}
	return count
}

func solveDay24Part2(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	hs := make([][6]float64, len(lines))
	for i, s := range lines {
		s1, s2, _ := strings.Cut(s, "@")
		for j, ss := range append(strings.Split(s1, ","), strings.Split(s2, ",")...) {
			v, _ := strconv.Atoi(strings.TrimSpace(ss))
			hs[i][j] = float64(v)
		}
	}
	// Solve linear equations for rock's px, dx, py, dy
	mat := make([][]float64, 4)
	res := make([]float64, 4)
	for i := 0; i < 4; i++ {
		mat[i] = make([]float64, 4)
		p1x, p1y, p2x, p2y := hs[i][0], hs[i][1], hs[i+1][0], hs[i+1][1]
		d1x, d1y, d2x, d2y := hs[i][3], hs[i][4], hs[i+1][3], hs[i+1][4]
		mat[i][0] = d2y - d1y
		mat[i][1] = d1x - d2x
		mat[i][2] = p1y - p2y
		mat[i][3] = p2x - p1x
		res[i] = p1y*d1x - p1x*d1y + p2x*d2y - p2y*d2x
	}
	gaussEliminate(res, mat)
	rpx := math.Round(res[0])
	rpy := math.Round(res[1])
	rdx := math.Round(res[2]) // rdy not used
	// Solve for rock's pz, dz
	mat = make([][]float64, 2)
	res = make([]float64, 2)
	for i := 0; i < 2; i++ {
		mat[i] = make([]float64, 2)
		p1x, p1z, p2x, p2z := hs[i][0], hs[i][2], hs[i+1][0], hs[i+1][2]
		d1x, d1z, d2x, d2z := hs[i][3], hs[i][5], hs[i+1][3], hs[i+1][5]
		mat[i][0] = d1x - d2x
		mat[i][1] = p2x - p1x
		res[i] = p1z*d1x - p1x*d1z + p2x*d2z - p2z*d2x - d2z*rpx + d1z*rpx - p1z*rdx + p2z*rdx
	}
	gaussEliminate(res, mat)
	rpz := math.Round(res[0]) // rdz not used
	return int64(rpx) + int64(rpy) + int64(rpz)
}

func gaussEliminate(res []float64, mat [][]float64) {
	n := len(mat)
	for i := 0; i < n; i++ {
		v := mat[i][i]
		res[i] /= v
		for j := 0; j < n; j++ {
			mat[i][j] /= v
		}
		for ii := 0; ii < n; ii++ {
			if ii != i {
				v := mat[ii][i]
				res[ii] -= v * res[i]
				for j := 0; j < n; j++ {
					mat[ii][j] -= v * mat[i][j]
				}
			}
		}
	}
}

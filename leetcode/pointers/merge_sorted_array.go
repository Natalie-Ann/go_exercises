func andrewMerge(nums1 []int, m int, nums2 []int, n int) {
  idx1 := m - 1
  idx2 := n - 1
  for n > 0 {
    if idx1 < 0 || nums1[idx1] < nums2[idx2] {
      nums1[idx1+n] = nums2[idx2]
      idx2--
      n--
    } else {
      nums1[idx1+n] = nums1[idx1]
      idx1--
    }
  }
}

function juliaMerge(nums1, m, nums2, n) {
  // could potentially put in an early return here for when n === 0
  let [idxM, idxN, writeIdx] = [m - 1, n - 1, m + n - 1];

  while (writeIdx >= 0) {
    // could potentially put in an early return here for when idxN < 0
    if (idxN < 0 || nums1[idxM] > nums2[idxN]) {
      nums1[writeIdx] = nums1[idxM];
      idxM -= 1;
    } else {
      nums1[writeIdx] = nums2[idxN];
      idxN -= 1;
    }

    writeIdx -= 1;
  }
}

algorithms
    keep a pointer at the end
    keep a pointer at the end of n
    keep a pointer at the end of m
    if the val at nums1[n] > nums2[m]
        write nums1[n] at the write idx
        n + 1
    otherwise
        write nums2[n] at the write idx
        m + 1
        
    increment the write idx
*/

func nykaelaMerge(nums1 []int, m int, nums2 []int, n int)  {
	idxM := m - 1
	idxN := n - 1
	writeIdx := m + n - 1
	
	for writeIdx > 0 {
			if idxM < 0 || nums1[idxM] < nums2[idxN] {
					nums1[writeIdx] = nums2[idxN]
					idxN--
			} else {
					nums1[writeIdx] = nums1[idxM]
					idxM--
			}
	}
	
	writeIdx--
}